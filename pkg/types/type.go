package types

//Position, this stores the empty missing values for each positions 
type Position struct {
	Layer int
	Layer1 int
	Layer2 int
	Layer3 int
}

// Row
type Row struct {
	Element int
	RowValues []int
}

type MissingProps struct {
	Position Position
	Possible []int
	Box int
}
