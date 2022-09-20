package command

import (
	"fmt"

	"github.com/willroberts/minecraft-client"
)

func FillBlock(c *minecraft.Client, x1, y1, z1, x2, y2, z2 int, blockName string) error {
	_, err := c.SendCommand(fmt.Sprintf("fill %d %d %d %d %d %d %s", x1, y1, z1, x2, y2, z2, blockName))
	return err
}

func SetBlock(c *minecraft.Client, x, y, z int, blockName string) error {
	_, err := c.SendCommand(fmt.Sprintf("setblock %d %d %d %s", x, y, z, blockName))
	return err
}
