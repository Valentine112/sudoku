package pkg

import (
	"github.com/Valentine112/sudoku/pkg/types"
	"reflect"
)

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

//Process, this func would call every other func in the pkg space
func Process(data [][][][]int) []types.Position {
	return getMissingPosition(data)
}