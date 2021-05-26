// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	type item struct {
		id    int
		name  string
		price int
	}

	type game struct {
		item
		genre string
	}

	games := []game{
		{
			item:  item{id: 1, name: "god of war", price: 50},
			genre: "action adventure",
		},
		{item: item{id: 2, name: "x-com 2", price: 40}, genre: "strategy"},
		{item: item{id: 3, name: "minecraft", price: 20}, genre: "sandbox"},
	}

	gamesMapById := make(map[int]game, len(games))
	for _, gr := range games {
		gamesMapById[gr.id] = gr
	}

	inputScan := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("list: list of all games")
		fmt.Println("quit: closing")
		fmt.Println("id N: get a game by id")
		fmt.Println("save: exports json and quit")

		if !inputScan.Scan() {
			break
		}

		scanInput := strings.Fields(inputScan.Text())
		if len(scanInput) == 0 {
			continue
		}

		switch scanInput[0] {
		case "quit":
			fmt.Println("bye!")
			return

		case "list":
			for _, gameRecord := range games {
				fmt.Println(gameRecord.id, gameRecord.name, gameRecord.genre, gameRecord.price)
			}

		case "id":
			if len(scanInput) == 2 {
				idValue, err := strconv.Atoi(scanInput[1])
				if err != nil {
					fmt.Println("Wrong Data")
				} else if gameValue, ok := gamesMapById[idValue]; ok {
					fmt.Println("Found Data", gameValue.name)
				} else {
					fmt.Println("No Record avilable ", idValue)
				}
			} else {
				fmt.Println("Wrong data")
			}

		case "save":
			type jsonGame struct {
				ID    int    `json:"id"`
				Name  string `json:"name"`
				Genre string `json:"genre"`
				Price int    `json:"price"`
			}

			var encodableGame []jsonGame
			for _, g := range games {
				encodableGame = append(encodableGame,
					jsonGame{g.id, g.name, g.genre, g.price})
			}

			out, err := json.MarshalIndent(encodableGame, "", "\t")
			if err != nil {
				fmt.Println("error:", err)
				continue
			}

			fmt.Println(string(out))
			return
		}
	}
}
