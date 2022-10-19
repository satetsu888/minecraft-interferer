package model

import (
	"strings"
)

type Vec3 struct {
	X int
	Y int
	Z int
}

type Block struct {
	BlockName string
	State     BlockState
}
type BlockState struct {
	Axis   string // x, y, z
	Facing string // north, south, west, east, up, down
	Type   string // for half block: top, bottom, double
	Half   string // for trap door: top, bottom
}

func (b Block) IsNull() bool {
	return b.BlockName == ""
}

func (b Block) GetRelativeString(facing Direction) string {
	states := []string{}

	if b.State.Axis != "" {
		relativeAxis := relativeAxis(b.State.Axis, facing)
		states = append(states, "axis="+relativeAxis)
	}

	if b.State.Facing != "" {
		relativeFacing := relativeFacing(b.State.Facing, facing)
		states = append(states, "facing="+relativeFacing)
	}

	if b.State.Type != "" {
		states = append(states, "type="+b.State.Type)
	}

	if b.State.Half != "" {
		states = append(states, "half="+b.State.Half)
	}

	if len(states) > 0 {
		return b.BlockName + "[" + strings.Join(states, ",") + "]"
	} else {
		return b.BlockName
	}
}

func relativeAxis(originalAxis string, facing Direction) string {
	if facing == North || facing == South {
		return originalAxis
	}

	if originalAxis == "x" {
		return "z"
	} else if originalAxis == "z" {
		return "x"
	} else {
		return originalAxis
	}
}

func relativeFacing(originalFacing string, facing Direction) string {
	if originalFacing == "up" || originalFacing == "down" {
		return originalFacing
	}

	if facing == South {
		return originalFacing
	}

	if facing == North {
		if originalFacing == "north" {
			return "south"
		} else if originalFacing == "south" {
			return "north"
		} else if originalFacing == "east" {
			return "west"
		} else if originalFacing == "west" {
			return "east"
		}
	}

	if facing == East {
		if originalFacing == "north" {
			return "west"
		} else if originalFacing == "south" {
			return "east"
		} else if originalFacing == "east" {
			return "north"
		} else if originalFacing == "west" {
			return "south"
		}
	}

	if facing == West {
		if originalFacing == "north" {
			return "east"
		} else if originalFacing == "south" {
			return "west"
		} else if originalFacing == "east" {
			return "south"
		} else if originalFacing == "west" {
			return "north"
		}
	}

	return originalFacing
}

type Structure struct {
	BasePoint Vec3
	Blocks    [][][]Block
}
