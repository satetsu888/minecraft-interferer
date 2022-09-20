package main

import (
	"fmt"
	"log"
	"os"

	"github.com/itchyny/maze"
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

	pos := player.Position()

	mazeSize := 30
	lineWidth := 4
	wallWidth := 1
	size := mazeSize*(lineWidth+wallWidth) + wallWidth
	height := 2
	blocks := make([][][]bool, size)
	for i := 0; i < size; i++ {
		blocks[i] = make([][]bool, size)
		for j := 0; j < height; j++ {
			blocks[i][j] = make([]bool, size)
			for k := 0; k < size; k++ {
				blocks[i][j][k] = true
			}
		}
	}

	maze := maze.NewMaze(mazeSize, mazeSize)
	maze.Generate()

	for y, line := range maze.Directions {
		for x, direction := range line {
			up := bool(direction&0b0001 != 0)
			down := bool(direction&0b0010 != 0)
			left := bool(direction&0b0100 != 0)
			right := bool(direction&0b1000 != 0)
			// fmt.Printf("%d, %d  up:%v, down: %v, left: %v, right: %v", x, y, up, down, left, right)
			// fmt.Println()
			X := (lineWidth+wallWidth)*x + 1
			Y := (lineWidth+wallWidth)*y + 1
			for h := 0; h < height; h++ {
				for lx := 0; lx < lineWidth; lx++ {
					for ly := 0; ly < lineWidth; ly++ {
						blocks[X+lx][h][Y+ly] = false
					}
				}
				for ly := 0; ly < lineWidth; ly++ {
					if right {
						blocks[X+lineWidth][h][Y+ly] = false
					}
					if left {
						blocks[X-wallWidth][h][Y+ly] = false
					}
				}
				for lx := 0; lx < lineWidth; lx++ {
					if up {
						blocks[X+lx][h][Y-wallWidth] = false
					}
					if down {
						blocks[X+lx][h][Y+lineWidth] = false
					}
				}
			}

		}

	}

	mazeX := pos.X + 5
	mazeY := pos.Y
	mazeZ := pos.Z - size/2

	client.FillBlocks(mazeX, mazeY, mazeZ, mazeX+size-1, mazeY+height-1, mazeZ+size-1, "minecraft:air")
	client.BuildBlocks(mazeX, mazeY, mazeZ, blocks, "minecraft:stone")
}
