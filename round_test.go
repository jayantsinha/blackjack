package main

import "testing"

func TestNewGame(t *testing.T) {
  game := NewGame()
  player := game.player
  dealer := game.dealer
  deck := game.deck
  expectedDeckSize := 48
  expectedPlayerHandSize, expectedDealerHandSize := 2, 2

  if len(deck) != expectedDeckSize {
    t.Errorf("Expected deck size to be %d, but found size to be %d",expectedDeckSize, len(deck))
  }

  if len(player.hand) != expectedPlayerHandSize {
    t.Errorf("Expected player to have %d cards, %d found", expectedPlayerHandSize, len(player.hand))
  }

  if len(dealer.hand) != expectedDealerHandSize {
    t.Errorf("Expected dealer to have %d cards, %d found", expectedDealerHandSize, len(dealer.hand))
  }
}
