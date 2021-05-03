package main

import (	
	"github.com/kyokomi/emoji/v2"
	"fmt"
	"time"
	"math/rand"	
)


func main() {   	
	printRandomThreeEmoji();	    
}

func printRandomThreeEmoji() {
    var emojiLen int = len(emoji.CodeMap())
	var emojiPrintCount int = 0
	rand.Seed(time.Now().UTC().UnixNano())
	emojiKey := generateEmojiKey();
	for ; emojiPrintCount < 3; {
		fmt.Println(emoji.CodeMap()[emojiKey[rand.Intn(emojiLen)]])	
		emojiPrintCount++
	}
}

func generateEmojiKey() []string {
	var emojiesKey []string = make([]string, len(emoji.CodeMap()))
	var i int = 0
	for k := range emoji.CodeMap() {
		emojiesKey[i] = k
		i++
	}
	return emojiesKey
}