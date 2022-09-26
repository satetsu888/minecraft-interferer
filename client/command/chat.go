package command

import (
	"fmt"

	"github.com/willroberts/minecraft-client"
)

func SendChat(c *minecraft.Client, message string) error {
	_, err := c.SendCommand(fmt.Sprintf("say %s", message))
	return err
}
