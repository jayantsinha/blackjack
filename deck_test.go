package main

import "testing"

var deck Deck

//Check deck capacity after init
func TestCreateNewDeck(t *testing.T) {
  d := deck
  d = CreateNewDeck()
  actualDeckSize := len(d)
  expectedDeckSize := 52
  if actualDeckSize != expectedDeckSize {
    t.Errorf("Expected deck size %d but got %d", expectedDeckSize, actualDeckSize)
  }
}

//Test shuffling
func TestShuffle(t *testing.T) {
  deckNormal := deck
  deckShuffled := deck
  deckNormal = CreateNewDeck()
  deckShuffled = CreateNewDeck()
  deckShuffled = deckShuffled.Shuffle()

  if len(deckNormal) == len(deckShuffled) {
    //Compare both decks
    equality := 1
    for i := 0; i<len(deckNormal); i++ {
      if deckNormal[i] == deckShuffled[i] {
        equality &= 1
      } else {
        equality &= 0
      }
    }
    if equality == 1 {
      t.Error("Expected the deck to be Shuffled!")
    }
  } else {
    t.Errorf("Expected length of deck after shuffle to be %d but got %d", len(deckNormal), len(deckShuffled))
  }

}

 //TestDraw verifies the following:
 //1. Deck size should reduce after drawing a card
 //2. Drawn card should not exist in the deck
func TestDraw(t *testing.T) {
  card := Card{}
  deck := deck
  deck = CreateNewDeck()
  initialSizeOfDeck := len(deck)
  card, deck = deck.Draw()
  newSizeOfDeck := len(deck)
  expectednewSize := initialSizeOfDeck - 1
  if newSizeOfDeck != expectednewSize {
    t.Errorf("Expected deck is %d, found %d", expectednewSize, newSizeOfDeck)
  }

  for r := range deck {
    if deck[r] == card {
      t.Errorf("Card %v should not exists in the deck", card)
    }
  }

}
