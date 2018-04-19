package main

import (
	"fmt"
	"strings"

	"./deck"
)

type Hand []deck.Card

func (h Hand) String() string {
	strs := make([]string, len(h))
	for i := range h {
		strs[i] = h[i].String()
	}
	return strings.Join(strs, ", ")
}

func main() {
	cards := deck.New(deck.Deck(1), deck.Shuffle)
	var card deck.Card
	var hero, villain Hand

	// Draw 2 cards each
	for i := 0; i < 2; i++ {
		for _, hand := range []*Hand{&hero, &villain} {
			// cards is the draw pile
			card, cards = draw(cards)
			*hand = append(*hand, card)
		}
	}

	fmt.Println("Hero:", hero)
	fmt.Println("Villain:", villain)
}

func draw(cards []deck.Card) (deck.Card, []deck.Card) {
	return cards[0], cards[1:]
}
