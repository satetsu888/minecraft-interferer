package query

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/satetsu888/minecraft-interferer/player"
	"github.com/willroberts/minecraft-client"
)

func FetchPlayerPosition(c *minecraft.Client, playerName string) (player.Position, error) {
	baseReg := regexp.MustCompile(`(?P<playerName>\S+) has the following entity data: (?P<data>.+)`)

	res, err := c.SendCommand("data get entity @p[name=" + playerName + "] Pos")
	if err != nil {
		return player.Position{}, fmt.Errorf("failed to fetch player Pos: %w", err)
	}
	result := baseReg.FindAllStringSubmatch(res.Body, -1)
	data := result[0][2]

	positionReg := regexp.MustCompile(`\[(?P<x>\S+)d, (?P<y>\S+)d, (?P<z>\S+)d\]`)
	result = positionReg.FindAllStringSubmatch(data, -1)
	x, err := strconv.ParseFloat(result[0][1], 64)
	if err != nil {
		return player.Position{}, err
	}
	y, err := strconv.ParseFloat(result[0][2], 64)
	if err != nil {
		return player.Position{}, err
	}
	z, err := strconv.ParseFloat(result[0][3], 64)
	if err != nil {
		return player.Position{}, err
	}
	return player.Position{
		X: x,
		Y: y,
		Z: z,
	}, nil
}

func FetchPlayerRotation(c *minecraft.Client, playerName string) (player.Rotation, error) {
	baseReg := regexp.MustCompile(`(?P<playerName>\S+) has the following entity data: (?P<data>.+)`)

	res, err := c.SendCommand("data get entity @p[name=" + playerName + "] Rotation")
	if err != nil {
		return player.Rotation{}, fmt.Errorf("failed to fetch player Rotation: %w", err)
	}
	result := baseReg.FindAllStringSubmatch(res.Body, -1)
	data := result[0][2]

	rotationReg := regexp.MustCompile(`\[(?P<yaw>\S+)f, (?P<pitch>\S+)f\]`)
	result = rotationReg.FindAllStringSubmatch(data, -1)
	yaw, err := strconv.ParseFloat(result[0][1], 64)
	if err != nil {
		return player.Rotation{}, err
	}
	pitch, err := strconv.ParseFloat(result[0][2], 64)
	if err != nil {
		return player.Rotation{}, err
	}
	return player.Rotation{
		Yaw:   yaw,
		Pitch: pitch,
	}, nil
}

func FetchPlayerDimention(c *minecraft.Client, playerName string) (player.Dimension, error) {
	baseReg := regexp.MustCompile(`(?P<playerName>\S+) has the following entity data: (?P<data>.+)`)

	res, err := c.SendCommand("data get entity @p[name=" + playerName + "] Dimension")
	if err != nil {
		return player.Dimension(""), fmt.Errorf("failed to fetch player Dimension: %w", err)
	}
	result := baseReg.FindAllStringSubmatch(res.Body, -1)
	data := result[0][2]

	dimensionReg := regexp.MustCompile(`(?P<dimension>\S+)`)
	result = dimensionReg.FindAllStringSubmatch(data, -1)
	return player.Dimension(result[0][1]), nil
}
