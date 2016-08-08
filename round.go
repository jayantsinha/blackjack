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
	GAME_DRAW             = "Game Draw"
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
	*human = p
}

func (dealer *Dealer) DealCard(card Card) {
	p := *dealer
	p.hand.AddCard(&card)
	*dealer = p
}

//Called when player opts to Stand
func (r *Round) DealersTurn() Round {
	round := *r
	//dealer := round.dealer
	//player := round.player
	card := Card{}
	dealerHandVal := round.dealer.hand.FavourableValue()
	playerHandVal := round.player.hand.FavourableValue()
	//Check Natural
	if dealerHandVal == BLACKJACK {
		round.outcome = GAME_LOST
		return round
	} else if dealerHandVal == BUST {
		round.outcome = GAME_WON
		return round
	} else {
		for round.dealer.hand.FavourableValue() <= 17 {
			card, round.deck = round.deck.Draw()
			round.dealer.DealCard(card)
			dealerHandVal = round.dealer.hand.FavourableValue()
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
		} else if dealerHandVal == playerHandVal {
			round.outcome = GAME_DRAW
			return round
		}
	}
	return round
}
