package main

import "testing"

func TestFavourableValue(t *testing.T) {
	//Test 1: Single Ace adjustment (favour blackjack)
	hand := Hand{}
	card1 := NewCard(SUIT_SPADES, ACE)
	card2 := NewCard(SUIT_HEARTS, ACE)
	card3 := NewCard(SUIT_HEARTS, NINE)
	hand.AddCard(&card1)
	hand.AddCard(&card2)
	hand.AddCard(&card3)
	handValue := hand.FavourableValue()
	expectedValue := 21
	if handValue != expectedValue {
		t.Errorf("Hand value not adjusted! Expected %d but got %d", expectedValue, handValue)
	}

	//Test 2: No adjustment (bust)
	hand = Hand{}
	card1 = NewCard(SUIT_SPADES, NINE)
	card2 = NewCard(SUIT_HEARTS, TEN)
	card3 = NewCard(SUIT_HEARTS, JACK)
	hand.AddCard(&card1)
	hand.AddCard(&card2)
	hand.AddCard(&card3)
	handValue = hand.FavourableValue()
	expectedValue = 29
	if handValue != expectedValue {
		t.Errorf("Adjustment not expected. Got %d instead of %d", handValue, expectedValue)
	}

	//Test 3: No adjustment (blackjack)
	hand = Hand{}
	card1 = NewCard(SUIT_SPADES, ACE)
	card2 = NewCard(SUIT_HEARTS, TEN)
	hand.AddCard(&card1)
	hand.AddCard(&card2)
	handValue = hand.FavourableValue()
	expectedValue = 21
	if handValue != expectedValue {
		t.Errorf("Adjustment not expected. Got %d instead of %d", handValue, expectedValue)
	}

}
