package pkg

import (
	"github.com/Valentine112/sudoku/pkg/types"
	"reflect"
	"fmt"
)

/**
* Layering starts from the top core
* That means layer is the top core
* Layer1 is the core after the top core and so on
*/

func rowColumnBox(vPosi []types.Position, data [][][][]int) {
	var rowType types.Row
	var rowValues []types.Row
	//var columnValues []int

	missingLen := reflect.ValueOf(vPosi).Len()
	// Loop through each unknown value to get their contracdictory row, col & box value
	for i := 0; i < missingLen; i++ {
		// Get the possible row values
		elem := vPosi[i]
		row := data[elem.Layer][elem.Layer1]

		for _, r := range row {
			rowType.Element = i + 1
			rowType.RowValues = append(rowType.RowValues, r[0], r[1], r[2])

			rowValues = append(rowValues, rowType)
		}
	}
	
	fmt.Println(rowValues)

}

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

func findPossibleValues(unknownValues []types.Position, data [][][][]int) {
	// Fetch all the possible values to for each empty container
	/*for _, val := range unknownValues {
		/*layer3 := val.Layer3
		layer2 := val.Layer2
		layer1 := val.Layer1
		layer := val.Layer

		// Get the mising value, then fetch the values in it row, column and box
		var missingValue int = data[layer][layer1][layer2][layer3]*/

	//}

	rowColumnBox(unknownValues, data)
}

//Process, this func would call every other func in the pkg space
func Process(data [][][][]int) {
	missingPositions := getMissingPosition(data)

	findPossibleValues(missingPositions, data)
}