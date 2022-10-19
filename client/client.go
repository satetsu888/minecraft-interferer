package client

import (
	"context"
	"fmt"

	"golang.org/x/sync/errgroup"

	"github.com/itchyny/maze"
	"github.com/satetsu888/minecraft-rcon-builder/client/command"
	"github.com/satetsu888/minecraft-rcon-builder/client/query"
	"github.com/satetsu888/minecraft-rcon-builder/model"
	"github.com/willroberts/minecraft-client"
)

type Client struct {
	Client *minecraft.Client
}

func NewClient(hostport string, rconPassword string) (*Client, error) {
	client, err := minecraft.NewClient(minecraft.ClientOptions{
		Hostport: hostport,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create client: %w", err)
	}
	if err := client.Authenticate(rconPassword); err != nil {
		return nil, err
	}
	return &Client{Client: client}, nil
}

func (c Client) FetchPlayerList() (count int, maxCount int, playerList []string, err error) {
	return query.FetchPlayerList(c.Client)
}

func (c Client) FetchPlayer(playerName string) (model.Player, error) {
	rawPosition, err := query.FetchPlayerRawPosition(c.Client, playerName)
	if err != nil {
		return model.Player{}, err
	}

	rotation, err := query.FetchPlayerRotation(c.Client, playerName)
	if err != nil {
		return model.Player{}, err
	}

	dimension, err := query.FetchPlayerDimention(c.Client, playerName)
	if err != nil {
		return model.Player{}, err
	}
	rawPosition.Dimension = dimension

	return model.Player{
		Name:        playerName,
		RawPosition: rawPosition,
		Rotation:    rotation,
	}, nil
}

func (c Client) SendChat(message string) error {
	return command.SendChat(c.Client, message)
}

func (c Client) BuildBlocks(pos model.Position, facing model.Direction, structure model.Structure) error {
	blocks := structure.Blocks
	eg, ctx := errgroup.WithContext(context.Background())
	for i := 0; i < len(blocks); i++ {
		i := i
		eg.Go(func() error {
			for j := 0; j < len(blocks[i]); j++ {
				for k := 0; k < len(blocks[i][j]); k++ {
					if !blocks[i][j][k].IsNull() {
						relatevePos := pos.GetRelative(i-structure.BasePoint.X, j-structure.BasePoint.Y, k-structure.BasePoint.Z, facing)
						err := command.SetBlock(c.Client, relatevePos.X, relatevePos.Y, relatevePos.Z, blocks[i][j][k].GetRelativeString(facing))
						if err != nil {
							ctx.Err()
							return err
						}
					}
				}
			}
			ctx.Done()
			return nil
		})
	}
	if err := eg.Wait(); err != nil {
		return err
	}
	return nil
}

func (c Client) FillBlocks(x1, y1, z1, x2, y2, z2 int, blockName string) error {
	return command.FillBlock(c.Client, x1, y1, z1, x2, y2, z2, blockName)
}

func (c Client) BuildMaze(pos model.Position, blockX, blockZ, height, roadWidth int, blockName string) error {
	x := pos.X
	y := pos.Y
	z := pos.Z
	wallWidth := 1
	sizeX := blockX*(roadWidth+wallWidth) + wallWidth
	sizeZ := blockZ*(roadWidth+wallWidth) + wallWidth
	blocks := make([][][]model.Block, sizeX)

	// fill bloocks
	for i := 0; i < sizeX; i++ {
		blocks[i] = make([][]model.Block, height)
		for j := 0; j < height; j++ {
			blocks[i][j] = make([]model.Block, sizeZ)
			for k := 0; k < sizeZ; k++ {
				blocks[i][j][k] = model.Block{BlockName: blockName}
			}
		}
	}

	maze := maze.NewMaze(blockX, blockZ)
	maze.Generate()

	// convert maze position to minecraft block position and mark roads as false
	for indexX, line := range maze.Directions {
		for indexY, direction := range line {
			up := bool(direction&0b0001 != 0)    // means -Z
			down := bool(direction&0b0010 != 0)  // means +Z
			left := bool(direction&0b0100 != 0)  // means -X
			right := bool(direction&0b1000 != 0) // means +X

			X := (roadWidth+wallWidth)*indexX + 1
			Y := (roadWidth+wallWidth)*indexY + 1

			for h := 0; h < height; h++ {
				// center of maze blocks
				for lx := 0; lx < roadWidth; lx++ {
					for ly := 0; ly < roadWidth; ly++ {
						blocks[X+lx][h][Y+ly] = model.Block{}
					}
				}
				// right and left walls
				for lx := 0; lx < roadWidth; lx++ {
					if right {
						blocks[X+lx][h][Y+roadWidth] = model.Block{}
					}
					if left {
						blocks[X+lx][h][Y-wallWidth] = model.Block{}
					}
				}
				// up and down walls
				for ly := 0; ly < roadWidth; ly++ {
					if up {
						blocks[X-wallWidth][h][Y+ly] = model.Block{}
					}
					if down {
						blocks[X+roadWidth][h][Y+ly] = model.Block{}
					}
				}
			}

		}

	}

	c.FillBlocks(x, y, z, x+sizeX-1, y+height-1, z+sizeZ-1, "minecraft:air")
	c.BuildBlocks(pos, model.Direction("south"), model.Structure{BasePoint: model.Vec3{X: 0, Y: 0, Z: 0}, Blocks: blocks})
	return nil
}
