package minecraft_interferer

import (
	"log"
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
		return nil, err
	}
	if err := client.Authenticate(rconPassword); err != nil {
		log.Fatal(err)
	}
	return &Client{Client: client}, nil
}

func (c *Client) FetchPlayer(playerName string) (Player, error) {
	baseReg := regexp.MustCompile(`(?P<playerName>\S+) has the following entity data: (?P<data>.+)`)

	// get position
	res, err := c.Client.SendCommand("data get entity @p[name=" + playerName + "] Pos")
	if err != nil {
		return Player{}, err
	}
	result := baseReg.FindAllStringSubmatch(res.Body, -1)
	name := result[0][1]

	positionReg := regexp.MustCompile(`\[(?P<x>\S+)d, (?P<y>\S+)d, (?P<z>\S+)d\]`)
	result = positionReg.FindAllStringSubmatch(res.Body, -1)
	x, err := strconv.ParseFloat(result[0][1], 64)
	if err != nil {
		return Player{}, err
	}
	y, err := strconv.ParseFloat(result[0][2], 64)
	if err != nil {
		return Player{}, err
	}
	z, err := strconv.ParseFloat(result[0][3], 64)
	if err != nil {
		return Player{}, err
	}

	// get direction
	res, err = c.Client.SendCommand("data get entity @p[name=" + playerName + "] Rotation")
	if err != nil {
		return Player{}, err
	}
	result = baseReg.FindAllStringSubmatch(res.Body, -1)
	data := result[0][2]
	rotationReg := regexp.MustCompile(`\[(?P<pitch>\S+)f, (?P<yow>\S+)f\]`)
	result = rotationReg.FindAllStringSubmatch(data, -1)
	yow, err := strconv.ParseFloat(result[0][1], 64)
	if err != nil {
		return Player{}, err
	}
	pitch, err := strconv.ParseFloat(result[0][2], 64)
	if err != nil {
		return Player{}, err
	}

	// get dimension
	res, err = c.Client.SendCommand("data get entity @p[name=" + playerName + "] Dimension")
	if err != nil {
		return Player{}, err
	}
	result = baseReg.FindAllStringSubmatch(res.Body, -1)
	dimension := result[0][2]

	return Player{
		Name: name,
		Position: Position{
			X: x,
			Y: y,
			Z: z,
		},
		Rotation: Rotation{
			Yaw:   yow,
			Pitch: pitch,
		},
		Dimension: Dimension(dimension),
	}, nil
}
