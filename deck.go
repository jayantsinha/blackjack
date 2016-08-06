package main

import (
	"math/rand"
	"time"
)

type Deck []Card

//Create a new Deck
func CreateNewDeck() Deck {
	deck := Deck{}
	deck = createSuit(SUIT_HEARTS, deck)
	deck = createSuit(SUIT_SPADES, deck)
	deck = createSuit(SUIT_BLACKS, deck)
	deck = createSuit(SUIT_DIAMONDS, deck)
	return deck
}

//Add given suit type to the deck
func createSuit(suit string, deck Deck) Deck {
	for i := 0; i < CARDS_IN_SUIT; i++ {
		deck = append(deck, NewCard(suit, i))
	}
	return deck
}

//Knuth or Fisher-Yates shuffle
func (d *Deck) Shuffle() {
	rand.Seed(time.Now().UnixNano() / int64(time.Millisecond))
	deck := *d
	for i := range deck {
		j := rand.Intn(i + 1)
		deck[i], deck[j] = deck[j], deck[i]
	}
	*d = deck
}

//Draw a card from the Deck
func (deck Deck) Draw() (Card, Deck) {
	return deck[0], deck[1:len(deck)]
}

func (deck Deck) HasMoreCards() bool {
	ret := false
	if len(deck) != 0 {
		ret = true
	} else {
		ret = false
	}
	return ret
}
