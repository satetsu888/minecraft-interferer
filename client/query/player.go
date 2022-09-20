package query

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/satetsu888/minecraft-interferer/player"
	"github.com/willroberts/minecraft-client"
)

func FetchPlayerPosition(c *minecraft.Client, playerName string) (player.Position, error) {
	reg := regexp.MustCompile(`(?P<playerName>\S+) has the following entity data: \[(?P<x>\S+)d, (?P<y>\S+)d, (?P<z>\S+)d\]`)

	res, err := c.SendCommand("data get entity @p[name=" + playerName + "] Pos")
	if err != nil {
		return player.Position{}, fmt.Errorf("failed to fetch player Pos: %w", err)
	}

	result := reg.FindAllStringSubmatch(res.Body, -1)
	x, err := strconv.ParseFloat(result[0][2], 64)
	if err != nil {
		return player.Position{}, err
	}
	y, err := strconv.ParseFloat(result[0][3], 64)
	if err != nil {
		return player.Position{}, err
	}
	z, err := strconv.ParseFloat(result[0][4], 64)
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
	reg := regexp.MustCompile(`(?P<playerName>\S+) has the following entity data: \[(?P<yaw>\S+)f, (?P<pitch>\S+)f\]`)

	res, err := c.SendCommand("data get entity @p[name=" + playerName + "] Rotation")
	if err != nil {
		return player.Rotation{}, fmt.Errorf("failed to fetch player Rotation: %w", err)
	}
	result := reg.FindAllStringSubmatch(res.Body, -1)
	yaw, err := strconv.ParseFloat(result[0][2], 64)
	if err != nil {
		return player.Rotation{}, err
	}
	pitch, err := strconv.ParseFloat(result[0][3], 64)
	if err != nil {
		return player.Rotation{}, err
	}
	return player.Rotation{
		Yaw:   yaw,
		Pitch: pitch,
	}, nil
}

func FetchPlayerDimention(c *minecraft.Client, playerName string) (player.Dimension, error) {
	reg := regexp.MustCompile(`(?P<playerName>\S+) has the following entity data: (?P<dimension>\S+)`)

	res, err := c.SendCommand("data get entity @p[name=" + playerName + "] Dimension")
	if err != nil {
		return player.Dimension(""), fmt.Errorf("failed to fetch player Dimension: %w", err)
	}
	result := reg.FindAllStringSubmatch(res.Body, -1)
	return player.Dimension(result[0][2]), nil
}
