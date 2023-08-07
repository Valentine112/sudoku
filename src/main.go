package main

import "fmt"

var  testSudoku = [][][][]int{
	{
		{{6, 3, 0}, {0, 2, 0}, {0, 0, 0}},
		{{0, 0, 0}, {0, 0, 3}, {0, 1, 7}},
		{{0, 8, 1}, {0, 0, 0}, {4, 3, 0}}},
	{
		{{0, 9, 6}, {0, 0, 0}, {0, 8, 0}},
		{{4, 0, 0}, {7, 6, 2}, {0, 0, 0}},
		{{5, 7, 0}, {0, 0, 0}, {6, 0, 0}}},
	{
		{{0, 6, 0}, {3, 0, 9}, {0, 0, 0}},
		{{0, 2, 0}, {0, 0, 0}, {0, 0, 0}},
		{{0, 0, 0}, {0, 6, 0}, {0, 0, 9}}},
}

// Entry point
func main() {
	// Every other func would be called from within this
	fmt.Println(testSudoku)
}