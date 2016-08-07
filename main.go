package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewScanner(os.Stdin)
	fmt.Print("Press enter to Start: ")
	reader.Scan()
	text := reader.Text()
	if text == "" {
		play()
	}
	if err := reader.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}

func (r *Round) init() {
	round := *r
	deck := CreateNewDeck()
	deck.Shuffle()
	round.deck = deck

	round.player = Human{
		hand:       Hand{},
		playerName: "Player",
	}

	round.dealer = Dealer{
		hand:       Hand{},
		playerName: "Dealer",
	}

	*r = round
}

func play() {
	reader := bufio.NewScanner(os.Stdin)
	text := reader.Text()
	round := Round{}
	round.init()
	card := Card{}
	player := round.player
	//dealer := round.dealer
	round.outcome = "ND"
	pHandVal := 0

	for round.deck.HasMoreCards() {
		if round.outcome == "ND" {
			fmt.Println("Inside if")
			for i := 0; i < 2; i++ {
				card, round.deck = round.deck.Draw()
				round.player.hand.AddCard(&card)
				card, round.deck = round.deck.Draw()

				round.dealer.hand.AddCard(&card)
			}
			fmt.Println(round.dealer)
			round.num += 1
			round.outcome = "ND"
			fmt.Printf("Dealer's Card: %v, [hidden card]\n", round.dealer.hand[0])
			fmt.Printf("Your Cards: %v\n", round.player.hand)
		}
		pHandVal = player.hand.FavourableValue()
		if pHandVal >= BUST {
			fmt.Println(GAME_BUST)
		} else if pHandVal == BLACKJACK {
			fmt.Println(GAME_BLACKJACK)
		} else {
			for round.outcome != "ND" {
				fmt.Println("Enter your action: Hit or Stand or Surrender [H/S/X]: ")
				reader.Scan()
				text = reader.Text()
				switch text {
				case "H", "h":
					card, round.deck = round.deck.Draw()
					round.player.DealCard(card)
					fmt.Println("Your card: ", card)
					if round.player.hand.FavourableValue() == BLACKJACK {
						round.outcome = GAME_BLACKJACK
						fmt.Println(GAME_BLACKJACK)
					} else if round.player.hand.FavourableValue() >= BUST {
						round.outcome = GAME_BUST
						fmt.Println(GAME_BUST)
					}
				case "S", "s":
					round = round.DealersTurn()
					fmt.Println(round.outcome)
				case "X", "x":
					round.outcome = GAME_LOST
					fmt.Println(GAME_LOST)
				}
			}
		}

	}
}
