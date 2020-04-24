/*
# In the game of Minesweeper you are given a Two Dimensional grid with a variable number of
# mines randomly located within the grid.
#
# Write a function that generates a random game grid. The function should take arguments for
# Length, Width and Number of Mines.
#
# The function should return the constructed 2D grid.  Place at most one mine per cell
#
# Your function gen(L, W, M), given L of 2, W of 3, M of 3 could return the following:
#
# +----+----+----+
# | X  |    |    |
# +--------------+
# |    |  X | X  |
# +----+----+----+

# How would you change your code if you also had to generate the tile numbers (the number on a non-mine tile should represent the number of mines in adjacent tiles)

Example:
# Your function gen(L, W, M), given L of 3, W of 3, M of 2 could return the following:
# +--+--+--+
# | X | 2 | 1 |
# +--+--+--+
# | 1 | 2 | X |
# +--+--+--+
# | 0 | 1 | 1 |
# +--+--+--+
*/
package main

import (
	"fmt"
	"log"
	"math/rand"
	"strings"
	"time"
)

// grid is a slice of slices that contain a point
type grid struct {
	rows     int
	columns  int
	capacity int
	/*
		[
		  [point, point, point],
		  [point, point, point],
		]
	*/
	points [][]point
}

// point contains whether or not a mine exists, a count of adjacent mines,
// and it's own x,y coordinates in the grid
type point struct {
	mine     bool
	adjacent int
	x        int
	y        int
}

// display prints the grid using ASCII art
func (g *grid) display() {

	for i := g.rows; i > 0; i-- {
		fmt.Printf(strings.Repeat("+---", g.columns))
		fmt.Printf("+\n")
		for _, p := range g.points[i-1] {
			symbol := " "
			if p.adjacent != 0 {
				symbol = fmt.Sprintf("%v", p.adjacent)
			}
			if p.mine == true {
				symbol = "X"
			}

			fmt.Printf("| %v ", symbol)
		}
		fmt.Printf("|\n")
	}

	fmt.Printf(strings.Repeat("+---", g.columns))
	fmt.Printf("+\n")
}

// gen builds a grid based on the provided rows and columns
func (g *grid) gen(r, c int) error {

	// check for negative inputs
	if r < 1 || c < 1 {
		return fmt.Errorf("can't build a grid of negative or zero size, got %v x %v", r, c)
	}

	// set our attributes
	g.rows = r
	g.columns = c
	g.capacity = r * c

	// build our grid
	g.points = make([][]point, g.rows)
	for y := range g.points {
		g.points[y] = make([]point, g.columns)
		for x := range g.points[y] {
			g.points[y][x].x = x
			g.points[y][x].y = y
		}
	}

	//fmt.Printf("Grid generated: %+v\n", g)
	return nil
}

// populate adds the given number of mines to the grid
func (g *grid) populate(m int) error {

	// check input
	if g.capacity < m {
		return fmt.Errorf("can't fit %v mines into grid sized %v x %v", m, g.rows, g.columns)
	}
	if m < 1 {
		return fmt.Errorf("please specifiy at least one mine, got %v", m)
	}

	// populate with mines
	fmt.Printf("Populating grid with %v mines...\n", m)
	for i := m; i > 0; {

		// get random coordinates
		x := rand.Intn(g.columns)
		y := rand.Intn(g.rows)
		//fmt.Printf("Checking coordinates (%v, %v)\n", x, y)

		// if this random point isn't a mine, make it a mine
		if g.points[y][x].mine == false {

			// now it's a mine
			g.points[y][x].mine = true

			// decrement our mine count
			i--

			// increment adjacent points mine counters
			// build a list of points to increment
			pts := [][]int{
				[]int{y + 1, x},
				[]int{y - 1, x},
				[]int{y, x + 1},
				[]int{y, x - 1},
				[]int{y + 1, x + 1},
				[]int{y - 1, x - 1},
				[]int{y + 1, x - 1},
				[]int{y - 1, x + 1},
			}

			// process the adjacent points
			for _, pt := range pts {

				// friendly variable names
				adjY := pt[0]
				adjX := pt[1]

				// handle out of range by checking bounds
				if adjY < 0 || adjX < 0 || adjY > g.rows-1 || adjX > g.columns-1 {
					// the point we are trying to increment is outside the grid,
					// so skip this one.
					continue
				}

				// increment the adjacent point
				g.points[adjY][adjX].adjacent++

			}

		}

		fmt.Printf("%v mines left to place\n", i)

	}

	return nil

}

// generate takes the number of rows, columns, and mines you'd like to use
// to generate the grid.
func generate(r, c, m int) (grid, error) {

	var g grid

	// generate our grid
	err := g.gen(r, c)
	if err != nil {
		return g, fmt.Errorf("failed to generate grid: %v", err)
	}

	// populate our grid with mines
	err = g.populate(m)
	if err != nil {
		return g, fmt.Errorf("failed to populate mines: %v", err)
	}

	return g, nil
}

func main() {

	// generate a random seed
	rand.Seed(time.Now().UnixNano())

	rows := 10
	columns := 10
	mines := 10

	log.Printf("Generating %v by %v grid...\n", rows, columns)
	g, err := generate(rows, columns, mines)
	if err != nil {
		log.Fatalf("failed to generate the grid: %v", err)
	}

	g.display()

}
