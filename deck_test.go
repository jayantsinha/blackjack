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
