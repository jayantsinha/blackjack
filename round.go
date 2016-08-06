package main

//Configs
const (
	CARDS_IN_SUIT = 13
	BUST          = 22
	BLACKJACK     = 21
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
	num     int
	deck    Deck
	player  Human
	dealer  Dealer
	outcome string
}

func (human *Human) DealCard(card Card) {
	p := *human
	p.hand.AddCard(&card)
}

func (dealer *Dealer) DealCard(card Card) {
	p := *dealer
	p.hand.AddCard(&card)
}

//Called when player opts to Stand
func (round Round) DealersTurn() Round {
	dealer := round.dealer
	player := round.player
	card := Card{}
	dealerHandVal := dealer.hand.FavourableValue()
	playerHandVal := player.hand.FavourableValue()
	//Check Natural
	if dealerHandVal == BLACKJACK {
		round.outcome = GAME_LOST
		return round
	} else if dealerHandVal == BUST {
		round.outcome = GAME_WON
		return round
	} else {
		for dealer.hand.FavourableValue() <= 17 {
			card, round.deck = round.deck.Draw()
			dealer.DealCard(card)
			dealerHandVal = dealer.hand.FavourableValue()
			if (dealerHandVal > playerHandVal) && (dealerHandVal <= BLACKJACK) {
				round.outcome = GAME_LOST
				return round
			}
			if dealerHandVal >= BUST {
				round.outcome = GAME_WON
				return round
			}
		}
		if dealerHandVal < playerHandVal {
			round.outcome = GAME_WON
			return round
		}
	}
	return round
}
