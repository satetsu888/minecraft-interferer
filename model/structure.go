package model

type Vec3 struct {
	X int
	Y int
	Z int
}

type Block struct {
	BlockName string
}

func (b Block) IsNull() bool {
	return b.BlockName == ""
}

type Structure struct {
	BasePoint Vec3
	Blocks    [][][]Block
}
