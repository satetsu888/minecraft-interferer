package player

type Dimension string

type Position struct {
	X float64
	Y float64
	Z float64
}

type Rotation struct {
	Yaw   float64
	Pitch float64
}

type Player struct {
	Name      string
	Position  Position
	Rotation  Rotation
	Dimension Dimension
}
