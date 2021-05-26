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

	var games [3]game
	games[0] = game{item: item{id: 1, name: "god of war", price: 50}, genre: "action adventure"}
	games[1] = game{item: item{id: 2, name: "x-com 2r", price: 30}, genre: "strategy"}
	games[2] = game{item: item{id: 3, name: "god of war", price: 20}, genre: "sandbox"}

	gamesMapById := make(map[int]game, len(games))
	for _, gr := range games {
		gamesMapById[gr.id] = gr
	}

	inputScan := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("list: list of all games")
		fmt.Println("quit: closing")
		fmt.Println("id N: get a game by id")

		if !inputScan.Scan() {
			break
		}

		scanInput := strings.Fields(inputScan.Text())
		if len(scanInput) > 0 {
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
			}

		}
	}
}
