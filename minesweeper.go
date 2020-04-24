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
	"math/rand"
)

// grid is a slice of slices that contain a point
/*
[
  [point, point, point],
  [point, point, point],
]
*/
type grid struct {
	rows     int
	columns  int
	capacity int
	points   [][]point
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
		return fmt.Errorf("can't build a grid of negative size, got %v x %v", r, c)
	}

	// set our attributes
	g.rows = r
	g.columns = c
	g.capacity = r * c

	// build our grid
	for i := 0; i < r; i++ {
		row := []point{}
		for j := 0; j < c; j++ {
			p := point{
				mine:     false,
				adjacent: 0,
				x:        i,
				y:        j,
			}
			row = append(row, p)
		}
		g = append(g, row)
	}

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
	for i := m; i > 0; {

		x := rand.Intn(g.rows)
		y := rand.Intn(g.columns)

		if g[x][y].mine == false {
			g[x][y].mine = true
			m--
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
	for x, r := range grid {
		for y, c := range r {

			adj := 0

			// TODO add additional checks
			if grid[x+1][y+1].mine == true {
				adj++
			}

		}

	}

	return nil
}

func main() {

	var g grid

}
