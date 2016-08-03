package main

//Configs
const (
  MAX_SUITS_IN_DECK = 4
  CARDS_IN_SUIT = 13
)



//Game Result Constants
const(
  GAME_LOST = iota
  GAME_WON
  GAME_BLACKJACK
  GAME_BUST
)


//Game consists of Player, Dealer and Deck
type Game struct {
  deck Deck
  player Player
  dealer Dealer
}

type Player struct {
  hand Hand
}

type Dealer struct {
  hand Hand
}
