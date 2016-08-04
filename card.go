package main

import "strconv"

//Faces
const (
	TWO int = iota
	THREE
	FOUR
	FIVE
	SIX
	SEVEN
	EIGHT
	NINE
	TEN
	JACK
	QUEEN
	KING
	ACE
)

//Suits
const (
	SUIT_HEARTS   = "H"
	SUIT_SPADES   = "S"
	SUIT_DIAMONDS = "D"
	SUIT_BLACKS   = "B"
)

type Card struct {
	suit           string
	face           string
	value          int
	alternateValue int //Same as value except for Aces
}

//Returns a new new card
func NewCard(s string, val int) Card {
	f, v, altVal := getCardValues(val)
	card := Card{
		suit:           s,
		face:           f,
		value:          v,
		alternateValue: altVal,
	}
	return card
}

//Get card values
func getCardValues(val int) (face string, value, alternateValue int) {
	switch val {
	case JACK:
		return "J", 10, 10
	case QUEEN:
		return "Q", 10, 10
	case KING:
		return "K", 10, 10
	case ACE:
		return "A", 11, 1
	}

	return strconv.Itoa(val + 2), val + 2, val + 2
}
