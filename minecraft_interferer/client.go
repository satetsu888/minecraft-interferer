package minecraft_interferer

import (
	"fmt"
	"regexp"
	"strconv"

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

func (c Client) FetchPlayerPosition(playerName string) (Position, error) {
	baseReg := regexp.MustCompile(`(?P<playerName>\S+) has the following entity data: (?P<data>.+)`)

	res, err := c.Client.SendCommand("data get entity @p[name=" + playerName + "] Pos")
	if err != nil {
		return Position{}, fmt.Errorf("failed to fetch player Pos: %w", err)
	}
	result := baseReg.FindAllStringSubmatch(res.Body, -1)
	data := result[0][2]

	positionReg := regexp.MustCompile(`\[(?P<x>\S+)d, (?P<y>\S+)d, (?P<z>\S+)d\]`)
	result = positionReg.FindAllStringSubmatch(data, -1)
	x, err := strconv.ParseFloat(result[0][1], 64)
	if err != nil {
		return Position{}, err
	}
	y, err := strconv.ParseFloat(result[0][2], 64)
	if err != nil {
		return Position{}, err
	}
	z, err := strconv.ParseFloat(result[0][3], 64)
	if err != nil {
		return Position{}, err
	}
	return Position{
		X: x,
		Y: y,
		Z: z,
	}, nil
}

func (c Client) FetchPlayerRotation(playerName string) (Rotation, error) {
	baseReg := regexp.MustCompile(`(?P<playerName>\S+) has the following entity data: (?P<data>.+)`)

	res, err := c.Client.SendCommand("data get entity @p[name=" + playerName + "] Rotation")
	if err != nil {
		return Rotation{}, fmt.Errorf("failed to fetch player Rotation: %w", err)
	}
	result := baseReg.FindAllStringSubmatch(res.Body, -1)
	data := result[0][2]

	rotationReg := regexp.MustCompile(`\[(?P<yaw>\S+)f, (?P<pitch>\S+)f\]`)
	result = rotationReg.FindAllStringSubmatch(data, -1)
	yaw, err := strconv.ParseFloat(result[0][1], 64)
	if err != nil {
		return Rotation{}, err
	}
	pitch, err := strconv.ParseFloat(result[0][2], 64)
	if err != nil {
		return Rotation{}, err
	}
	return Rotation{
		Yaw:   yaw,
		Pitch: pitch,
	}, nil
}

func (c Client) FetchPlayerDimention(playerName string) (Dimension, error) {
	baseReg := regexp.MustCompile(`(?P<playerName>\S+) has the following entity data: (?P<data>.+)`)

	res, err := c.Client.SendCommand("data get entity @p[name=" + playerName + "] Dimension")
	if err != nil {
		return Dimension(""), fmt.Errorf("failed to fetch player Dimension: %w", err)
	}
	result := baseReg.FindAllStringSubmatch(res.Body, -1)
	data := result[0][2]

	dimensionReg := regexp.MustCompile(`(?P<dimension>\S+)`)
	result = dimensionReg.FindAllStringSubmatch(data, -1)
	return Dimension(result[0][1]), nil
}

func (c Client) FetchPlayer(playerName string) (Player, error) {
	position, err := c.FetchPlayerPosition(playerName)
	if err != nil {
		return Player{}, err
	}

	rotation, err := c.FetchPlayerRotation(playerName)
	if err != nil {
		return Player{}, err
	}

	dimension, err := c.FetchPlayerDimention(playerName)
	if err != nil {
		return Player{}, err
	}

	return Player{
		Name:      playerName,
		Position:  position,
		Rotation:  rotation,
		Dimension: dimension,
	}, nil
}
