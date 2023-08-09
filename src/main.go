package main

import (
	//"fmt"
	"github.com/Valentine112/sudoku/pkg"
)

// A test sample
// It is a 9-dimesional sudoku problem
// Zero means empty, so spaces with 0 are blank and needs to be filled with the right number
var  testSudoku = [][][][]int{
	{
		{{0, 3, 4}, {0, 0, 7}, {0, 0, 8}},
		{{0, 8, 0}, {0, 6, 5}, {0, 0, 0}},
		{{0, 0, 0}, {3, 0, 0}, {0, 7, 0}}},
	{
		{{2, 0, 0}, {0, 0, 0}, {7, 0, 0}},
		{{7, 1, 0}, {0, 4, 0}, {0, 9, 6}},
		{{0, 0, 5}, {0, 0, 0}, {0, 0, 1}}},
	{
		{{0, 5, 0}, {0, 0, 2}, {0, 0, 0}},
		{{0, 0, 0}, {1, 7, 0}, {0, 6, 0}},
		{{6, 0, 0}, {9, 0, 0}, {4, 3, 0}}},
}

// Entry point
func main() {
	// Every other func would be called from within this
	pkg.Process(testSudoku)
}