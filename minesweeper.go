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
columns
# +----+----+----+
# | X  |    |    |
# +--------------+
# |    |  X | X  |
# +----+----+----+   rows
*/

/*
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

	fmt.Printf("Grid generated: %+v\n", g)
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
		fmt.Printf("Checking coordinates (%v, %v)\n", x, y)

		// if this random point isn't a mine, make it a mine and decrement
		// our mine count
		if g.points[y][x].mine == false {
			g.points[y][x].mine = true
			i--
		}

		fmt.Printf("%v mines left to place\n", i)

	}

	return nil

}

// getPoint gets a point based on the x, y coordinates supplied
func (g *grid) getPoint(x, y int) (point, error) {

	var p point

	// check our inputs
	if x > g.columns || x < 0 {
		return p, fmt.Errorf("x is outside of grid %v x %v", g.columns, g.rows)
	}
	if y > g.rows || y < 0 {
		return p, fmt.Errorf("y is outside of grid %v x %v", g.columns, g.rows)
	}

	// get our point off the grid
	p = g.points[y][x]

	return p, nil

}

// findAdjacent walks the grid and sets the adjacent value for each point
func (g *grid) findAdjacent() error {

	// walk all points
	for _, row := range g.points {
		for _, p := range row {

			var adj int

			// check all 8 adjacent points
			// this feels really suboptimal, but I can't think of a better
			// solution right now

			// only process subtractions from x if it is greater than 0
			// the boundry of our grid
			if p.x > 0 {
				if p.y > 0 {
					if g.points[p.y-1][p.x-1].mine == true {
						adj++
					}
				}
				if p.y < g.rows {
					if g.points[p.y+1][p.x-1].mine == true {
						adj++
					}
				}
				if g.points[p.y][p.x-1].mine == true {
					adj++
				}
			}

			// only process additions to x if it is less than our boundry
			if p.x < g.columns {

				if p.y > 0 {
					if g.points[p.y-1][p.x+1].mine == true {
						adj++
					}
				}
				if p.y < g.rows {
					if g.points[p.y+1][p.x+1].mine == true {
						adj++
					}
				}
				if g.points[p.y][p.x+1].mine == true {
					adj++
				}

			}

			// check our boundry when adding to y
			if p.y < g.rows {
				if g.points[p.y+1][p.x].mine == true {
					adj++
				}
			}

			// check our boundry when subtracting from y
			if p.y > 0 {
				if g.points[p.y-1][p.x].mine == true {
					adj++
				}
			}

			// set the adjacent count
			g.points[p.y][p.x].adjacent = adj

		}
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

	// populate adjacent counts
	err = g.findAdjacent()
	if err != nil {
		return g, fmt.Errorf("failed to set adjacent counts on points: %v", err)
	}

	return g, nil
}

func main() {

	log.Println("Generating grid...")
	g, err := generate(3, 3, 2)
	if err != nil {
		log.Fatalf("failed to generate the grid: %v", err)
	}

	fmt.Printf("%+v", g)

}
