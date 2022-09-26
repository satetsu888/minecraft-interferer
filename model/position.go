package model

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

func (p Position) GetRelative(x, y, z int, direction Direction) Position {
	switch direction {
	case North:
		return Position{
			X:         p.X - x,
			Y:         p.Y + y,
			Z:         p.Z - z,
			Dimension: p.Dimension,
		}
	case South:
		return Position{
			X:         p.X + x,
			Y:         p.Y + y,
			Z:         p.Z + z,
			Dimension: p.Dimension,
		}
	case East:
		return Position{
			X:         p.X + z,
			Y:         p.Y + y,
			Z:         p.Z - x,
			Dimension: p.Dimension,
		}
	case West:
		return Position{
			X:         p.X - z,
			Y:         p.Y + y,
			Z:         p.Z + x,
			Dimension: p.Dimension,
		}
	default:
		return p
	}
}
