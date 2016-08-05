package main

//Configs
const (
	CARDS_IN_SUIT        = 13
	BUST                 = 21
)

//Game Result Constants
const (
	GAME_LOST      string = "Game Lost"
	GAME_WON              = "Game Won"
	GAME_BLACKJACK        = "Blackjack!"
	GAME_BUST             = "Bust!"
)

type Player struct {
	hand       Hand
	playerName string
}

type Human Player

type Dealer Player

type Round struct {
  num int
  deck Deck
  player Human
  dealer Dealer
  outcome string
}

func NewGame() Game {
  game := Game{}
  d := CreateNewDeck()
  d = d.Shuffle()
	game.deck = d
	game.player = Human{
		hand:       Hand{},
		playerName: "Player",
	}

	game.dealer = Dealer{
		hand:       Hand{},
		playerName: "Dealer",
	}

  //deal 2 cards to both player and dealer


	return game

}

func DealCard(player Player, deck Deck) Player {
  card, deck =
}
