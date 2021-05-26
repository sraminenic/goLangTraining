package main

import "fmt"

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

	for _, gameRecord := range games {
		fmt.Println(gameRecord.id, gameRecord.name, gameRecord.price, gameRecord.genre)
	}

}
