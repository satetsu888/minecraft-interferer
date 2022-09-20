package query

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/satetsu888/minecraft-interferer/model"
	"github.com/willroberts/minecraft-client"
)

func FetchPlayerRawPosition(c *minecraft.Client, playerName string) (model.RawPosition, error) {
	reg := regexp.MustCompile(`(?P<playerName>\S+) has the following entity data: \[(?P<x>\S+)d, (?P<y>\S+)d, (?P<z>\S+)d\]`)

	res, err := c.SendCommand("data get entity @p[name=" + playerName + "] Pos")
	if err != nil {
		return model.RawPosition{}, fmt.Errorf("failed to fetch player Pos: %w", err)
	}

	result := reg.FindAllStringSubmatch(res.Body, -1)
	x, err := strconv.ParseFloat(result[0][2], 64)
	if err != nil {
		return model.RawPosition{}, err
	}
	y, err := strconv.ParseFloat(result[0][3], 64)
	if err != nil {
		return model.RawPosition{}, err
	}
	z, err := strconv.ParseFloat(result[0][4], 64)
	if err != nil {
		return model.RawPosition{}, err
	}

	return model.RawPosition{
		X: x,
		Y: y,
		Z: z,
	}, nil
}

func FetchPlayerRotation(c *minecraft.Client, playerName string) (model.Rotation, error) {
	reg := regexp.MustCompile(`(?P<playerName>\S+) has the following entity data: \[(?P<yaw>\S+)f, (?P<pitch>\S+)f\]`)

	res, err := c.SendCommand("data get entity @p[name=" + playerName + "] Rotation")
	if err != nil {
		return model.Rotation{}, fmt.Errorf("failed to fetch player Rotation: %w", err)
	}
	result := reg.FindAllStringSubmatch(res.Body, -1)
	yaw, err := strconv.ParseFloat(result[0][2], 64)
	if err != nil {
		return model.Rotation{}, err
	}
	pitch, err := strconv.ParseFloat(result[0][3], 64)
	if err != nil {
		return model.Rotation{}, err
	}
	return model.Rotation{
		Yaw:   yaw,
		Pitch: pitch,
	}, nil
}

func FetchPlayerDimention(c *minecraft.Client, playerName string) (model.Dimension, error) {
	reg := regexp.MustCompile(`(?P<playerName>\S+) has the following entity data: (?P<dimension>\S+)`)

	res, err := c.SendCommand("data get entity @p[name=" + playerName + "] Dimension")
	if err != nil {
		return model.Dimension(""), fmt.Errorf("failed to fetch player Dimension: %w", err)
	}
	result := reg.FindAllStringSubmatch(res.Body, -1)
	return model.Dimension(result[0][2]), nil
}
