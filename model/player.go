package model

import "math"

type Dimension string

type RawPosition struct {
	X         float64
	Y         float64
	Z         float64
	Dimension Dimension
}

type Position struct {
	X         int
	Y         int
	Z         int
	Dimension Dimension
}

type Rotation struct {
	Yaw   float64
	Pitch float64
}

type Player struct {
	Name        string
	RawPosition RawPosition
	Rotation    Rotation
}

func (p Player) Position() Position {
	x := int(math.Floor(p.RawPosition.X))
	y := int(math.Floor(p.RawPosition.Y))
	z := int(math.Floor(p.RawPosition.Z))
	return Position{
		X:         x,
		Y:         y,
		Z:         z,
		Dimension: p.RawPosition.Dimension,
	}
}
