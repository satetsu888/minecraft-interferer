package model

import "math"

type Dimension string

type Direction string

const (
	North Direction = "north"
	South Direction = "south"
	East  Direction = "east"
	West  Direction = "west"
)

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

func (p Player) Direction() Direction {
	if p.Rotation.Yaw >= -45 && p.Rotation.Yaw < 45 {
		return South
	} else if p.Rotation.Yaw >= 45 && p.Rotation.Yaw < 135 {
		return West
	} else if p.Rotation.Yaw >= 135 || p.Rotation.Yaw < -135 {
		return North
	} else {
		return East
	}
}
