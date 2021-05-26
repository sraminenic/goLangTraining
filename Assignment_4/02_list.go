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

	in := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("list: list of all games")
		fmt.Println("quit: closing")

		if !in.Scan() {
			break
		}

		switch in.Text() {
		case "quit":
			fmt.Println("bye!")
			return

		case "list":
			for _, gameRecord := range games {
				fmt.Println(gameRecord.id, gameRecord.name, gameRecord.genre, gameRecord.price)
			}
		}
	}
}
