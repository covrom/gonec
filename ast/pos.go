package ast

// Position provides interface to store code locations.
type Position struct {
	Line   int
	Column int
}

// Pos interface provies two functions to get/set the position for expression or statement.
type Pos interface {
	Position() Position
	SetPosition(Position)
}

// PosImpl provies commonly implementations for Pos.
type PosImpl struct {
	Pos Position
}

// Position return the position of the expression or statement.
func (x *PosImpl) Position() Position {
	return x.Pos
}

// SetPosition is a function to specify position of the expression or statement.
func (x *PosImpl) SetPosition(pos Position) {
	x.Pos = pos
}
