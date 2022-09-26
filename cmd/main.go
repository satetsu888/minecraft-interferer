package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/satetsu888/minecraft-rcon-builder/client"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	rconHostPort := os.Getenv("RCON_HOSTPORT")
	rconPassowrd := os.Getenv("RCON_PASSWORD")
	playerName := os.Getenv("PLAYER_NAME")

	client, err := client.NewClient(rconHostPort, rconPassowrd)
	if err != nil {
		panic(err)
	}

	/*
		resp, err := client.Client.SendCommand("data get entity @p[name=satetsu888]")
		if err != nil {
			fmt.Printf("%+v", err)
		}
		log.Println(resp.Body)
	*/

	player, err := client.FetchPlayer(playerName)
	if err != nil {
		fmt.Printf("%+v", err)
		panic(err)
	}
	fmt.Printf("%+v", player)
	fmt.Println()

	// pos := player.Position()

	/*
		err = client.BuildMaze(pos.X+5, pos.Y, pos.Z, 5, 5, 2, 4, "minecraft:stone")
		if err != nil {
			fmt.Printf("%+v", err)
			panic(err)
		}
	*/
	/*
		err = client.FillBlocks(pos.X, pos.Y, pos.Z, pos.X+30, pos.Y+20, pos.Z-30, "minecraft:air")
		if err != nil {
			fmt.Printf("%+v", err)
			panic(err)
		}
	*/
	/*
		count, maxCouont, playerList, err := client.FetchPlayerList()
		if err != nil {
			panic(err)
		}
		fmt.Printf("%+v", count)
		fmt.Println()
		fmt.Printf("%+v", maxCouont)
		fmt.Println()
		fmt.Printf("%+v", playerList)
	*/

	err = client.SendChat("Hello, world!")
	if err != nil {
		panic(err)
	}
}
