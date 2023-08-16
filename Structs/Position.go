package Structs

import "math"

// Position describes position of a tile. Simply a three-dimensional vector
type Position struct {
	X int `json:"x"`
	Y int `json:"y"`
	Z int `json:"z"`
}

func (pos Position) Sub(other Position) Position {
	return Position{
		X: pos.X - other.X,
		Y: pos.Y - other.Y,
		Z: pos.Z - other.Z,
	}
}

func (pos Position) Length() float64 {
	return math.Sqrt(float64(pos.X*pos.X + pos.Y*pos.Y + pos.Z*pos.Z))
}

func (pos *Position) DistanceTo(other Position) float64 {
	return pos.Sub(other).Length()
}
