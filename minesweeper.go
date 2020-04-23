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

type grid [][]point

type point struct {
  mine bool
  adjacent int
  x int
  y int
}

/*[
  [bool, bool, bool],
  [bool, bool, bool],
]
*/

//map[string]bool

//XY = bool


func gen(r, c, m int) (grid, error) {
  
  // check input
  if r * c < m {
    return nil, fmt.Errorf("can't fit %v mines into grid sized %v x %v", m, r, c)
  }
  
  if m < 1 {
    return nil,fmt.Errorf("please specifiy at least one mine, got %v", m)
  }
  
  if r < 1 || c < 1 {
    return nil, fmt.Errorf("can't build a grid of negative size, got %v x %v", r, c)
  }
  
  g := grid
  mines := m
  
  // build our grid
  for i := 0; i < r; i++ {
    row := []point
    for j :=0; j < c; j++ {
      p := point{
        mine: false,
        adjacent: 0,
        x: i,
        y: j,
      }
      row = append(row, point)
    }
    grid = append(grid, row)
  }
  
  // populate with mines
  for i := mines; i > 0 {
  
    x := random.Int(0, r)
    y := random.Int(0, c)
    
    if grid[x][y].mine == false {
      grid[x][y].mine = true
      mines--
    }   
  }
  
  // populate adjacent counts
  for x, r := range grid {
    for y, c := range row {
      
      adj := 0
      
      if grid[x+1][y+1].mine == true {
        adj++
      }
      ...
      
      
    }
  
  
  return grid
}

func (p *point) adjacent() {
  
  adj := 0
  
  if 
  
}

func main() {
  
}
