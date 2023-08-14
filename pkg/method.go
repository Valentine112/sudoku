package pkg

import (
	"github.com/Valentine112/sudoku/pkg/types"
	"reflect"
	"fmt"
	"sort"
)

/**
* Layering starts from the top core
* That means layer is the top core
* Layer1 is the core after the top core and so on
*/

// This funct returns the missing elements along with their position on each axis
func getMissingPosition(sudoku [][][][]int) []types.Position {
	var result []types.Position
	var posi types.Position
	//var position types.Position
	// For a 9 dimensional array, we would be looping thrice
	for ind, layer := range sudoku {
		// Confirm if the next value is an array
		if reflect.ValueOf(layer).Kind() == reflect.Slice {
			// Loop again for the second time
			for ind1, layer1 := range layer {

				if reflect.ValueOf(layer1).Kind() == reflect.Slice {

					for ind2, layer2 := range layer1 {

						if(reflect.ValueOf(layer2).Kind() == reflect.Slice) {

							// Supposed to be the last loop
							for ind3, layer3 := range layer2 {
								// Verify that the value is an Integer and not array
								if reflect.ValueOf(layer3).Kind() == reflect.Int {
									if layer3 == 0 {
										posi.Layer = ind
										posi.Layer1 = ind1
										posi.Layer2 = ind2
										posi.Layer3 = ind3

										result = append(result, posi)
									}
								}
							}
						}
					}
				}
			}
		}
	}

	return result
}

func findPossibleValues(unknownValues []types.Position, data [][][][]int, dimen int) []types.MissingProps {
	var missingData []types.MissingProps
	var missingProps types.MissingProps

	for _, val := range unknownValues {
		var result []int
		
		for i := 1; i <= 9; i++ {
			found := 0
			// Check for the row
			layer := data[val.Layer][val.Layer1]
			// The 3rd row
			for _, val1 := range layer {
				// Fourth and final row with the values
				for _, val2 := range val1 {
					if i == val2 {
						found = 1
						break
					}
				}

				if found == 1 {
					break
				}
			}

			// Check for the column if the value is still a prospect
			if found == 0 {
				// Loop through the highest layer fist
				for _, val1 := range data {
					// Loop through the second layer next
					for _, val2 := range val1 {
						if i == val2[val.Layer2][val.Layer3] {
							found = 1
							break
						}
					}

					if found == 1 {
						break
					}
				}
			}

			// Check for the box if the value is still a prospect
			if found == 0 {
				// Loop through the missing value first layer
				for _, val1 := range data[val.Layer] {
					for _, val2 := range val1[val.Layer2] {
						if i == val2 {
							found = 1
							break
						}
					}
					if found == 1 {
						break
					}
				}

			}

			if found == 0 {
				result = append(result, i)
			}
		}

		missingProps.Position = val
		missingProps.Possible = result

		// Fetching the box
		// Getting the exact box would help in making the elimination process easier
		switch val.Layer {
			case 0:
				missingProps.Box = (val.Layer2 + 1)
				break

			case 1:
				missingProps.Box = 3 + (val.Layer2 + 1)
				break

			case 2:
				missingProps.Box = 6 + (val.Layer2 + 1)
				break
		}

		missingData = append(missingData, missingProps)
	}

	return missingData
}

func elimination(missingData []types.MissingProps) {
	for ind, val := range missingData {
		// Start by eliminating every other poss that goes a single poss in a block
		if len(val.Possible) == 1 {
			// Loop through the whole data again
			// Then eliminate this value from every other possible value that has this value in it, but must be in the same row, column or box

			for ind1, val1 := range missingData {
				// Most likely be wraping this in a function

				// Declare the value details that we are working with
				position := val.Position
				//box := val.Box
				//possible := val.Possible

				// Declare the current value details
				position1 := val1.Position
				//box1 := val1.Box
				//possible1 := val1.Possible

				// Elminate from row
				if position.Layer == position1.Layer && position.Layer1 == position1.Layer1 {
					// Check that its not the same value
					if ind != ind1 {
						// Get the position of the value
						// Then save as array since the delete only accepts array indexes
						var posi []int
						posi = append(posi, sort.SearchInts(possible1, possible[0]))
						// Delete the values here
						break
					}
					// Get the position of the value in the list
					//posi := sort.SearchInts(possible1, possible[0])
					// Access the main data
					//missingData.

				}
			}
		}
	}
}

//Process, this func would call every other func in the pkg space
func Process(data [][][][]int, dimen int) {
	missingPositions := getMissingPosition(data)

	possibleValues := findPossibleValues(missingPositions, data, dimen)

	elimination(possibleValues)
}