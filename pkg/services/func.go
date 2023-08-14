package services

func DeleteElements(slice []int, index []int) {
	// Sort index in descending order
	for _, val := range index {
		slice = append(slice[:val], slice[val+1:]...)
	}
}