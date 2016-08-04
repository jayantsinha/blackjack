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

//Test for deck size after multiple draws
func TestDraw(t *testing.T) {
  deck := deck
  deck = CreateNewDeck()
  initialSizeOfDeck := len(deck)
  _, deck = deck.Draw()
  newSizeOfDeck := len(deck)
  expectednewSize := initialSizeOfDeck - 1
  if newSizeOfDeck != expectednewSize {
    t.Errorf("Expected deck is %d, found %d", expectednewSize, newSizeOfDeck)
  }

}
