package main

import "testing"

func TestDealCard(t *testing.T) {
	human := Human{}
	dealer := Dealer{}
	card := NewCard(SUIT_SPADES, JACK)
	human.DealCard(card)
	dealer.DealCard(card)

	if len(human.hand) == 1 {
		t.Errorf("Expected player's hand length to be 1, %d found", len(human.hand))
	}
	if len(dealer.hand) == 1 {
		t.Errorf("Expected dealer's hand length to be 1, %d found", len(dealer.hand))
	}
}
