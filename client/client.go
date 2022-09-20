package client

import (
	"fmt"
	"sync"

	"github.com/satetsu888/minecraft-interferer/client/command"
	"github.com/satetsu888/minecraft-interferer/client/query"
	"github.com/satetsu888/minecraft-interferer/model"
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

func (c Client) BuildBlocks(x, y, z int, blocks [][][]bool, blockName string) error {
	wg := new(sync.WaitGroup)
	for i := 0; i < len(blocks); i++ {
		wg.Add(1)
		go func(i int) error {
			for j := 0; j < len(blocks[i]); j++ {
				for k := 0; k < len(blocks[i][j]); k++ {
					if blocks[i][j][k] {
						err := command.SetBlock(c.Client, x+i, y+j, z+k, blockName)
						if err != nil {
							return err
						}
					}
				}
			}
			wg.Done()
			return nil
		}(i)
	}
	wg.Wait()
	return nil
}

func (c Client) FillBlocks(x1, y1, z1, x2, y2, z2 int, blockName string) error {
	return command.FillBlock(c.Client, x1, y1, z1, x2, y2, z2, blockName)
}
