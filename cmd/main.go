package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/satetsu888/minecraft-interferer/minecraft_interferer"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	rconHostPort := os.Getenv("RCON_HOSTPORT")
	rconPassowrd := os.Getenv("RCON_PASSWORD")
	playerName := os.Getenv("PLAYER_NAME")

	client, err := minecraft_interferer.NewClient(rconHostPort, rconPassowrd)
	if err != nil {
		panic(err)
	}

	/*
		resp, err := client.Client.SendCommand("data get entity @p[name=satetsu888]")
		if err != nil {
			log.Fatal("error: ", err)
		}
		log.Println(resp.Body)
	*/

	player, err := client.FetchPlayer(playerName)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v", player)
}
