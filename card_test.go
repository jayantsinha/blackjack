package main

import "testing"

func TestNewCard(t *testing.T) {
	card := Card{}
	card = NewCard(SUIT_HEARTS, ACE)
	aceCard := Card{
		suit:           SUIT_HEARTS,
		face:           "A",
		value:          11,
		alternateValue: 1,
	}

	if card.face != aceCard.face {
		t.Errorf("Card value should be %v, found %v", aceCard.face, card.face)
	}
	if card.value != aceCard.value {
		t.Errorf("Card value should be %d, found %d", aceCard.value, card.value)
	}
	if card.alternateValue != aceCard.alternateValue {
		t.Errorf("Card value should be %d, found %d", aceCard.alternateValue, card.alternateValue)
	}
	if card.suit != aceCard.suit {
		t.Errorf("Card value should be %v, found %v", aceCard.suit, card.suit)
	}
}
