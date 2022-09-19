package client

import (
	"fmt"

	"github.com/satetsu888/minecraft-interferer/client/query"
	"github.com/satetsu888/minecraft-interferer/player"
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

func (c Client) FetchPlayer(playerName string) (player.Player, error) {
	position, err := query.FetchPlayerPosition(c.Client, playerName)
	if err != nil {
		return player.Player{}, err
	}

	rotation, err := query.FetchPlayerRotation(c.Client, playerName)
	if err != nil {
		return player.Player{}, err
	}

	dimension, err := query.FetchPlayerDimention(c.Client, playerName)
	if err != nil {
		return player.Player{}, err
	}
	position.Dimension = dimension

	return player.Player{
		Name:     playerName,
		Position: position,
		Rotation: rotation,
	}, nil
}
