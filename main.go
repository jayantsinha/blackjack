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
		if round.outcome == GAME_BUST || round.outcome == GAME_BLACKJACK || round.outcome == GAME_LOST || round.outcome == GAME_WON || round.outcome == "ND" {
			round.dealer = Dealer{}
			round.player = Human{}
			for i := 0; i < 2; i++ {
				card, round.deck = round.deck.Draw()
				round.player.hand.AddCard(&card)
				card, round.deck = round.deck.Draw()
				round.dealer.hand.AddCard(&card)
			}

			round.num += 1
			round.outcome = ""
			fmt.Printf("\n========================================================\n")
			fmt.Printf("Dealer's Card: %v, [hidden card]\n", round.dealer.hand[0])
			fmt.Printf("Your Cards: %v\n", round.player.hand)
		}
		pHandVal = player.hand.FavourableValue()
		if pHandVal >= BUST {
			round.outcome = GAME_BUST
			fmt.Println(GAME_BUST)
			fmt.Printf("Dealer's Hand: %v\n", round.dealer.hand)
		} else if pHandVal == BLACKJACK {
			round.outcome = GAME_BLACKJACK
			fmt.Println(GAME_BLACKJACK)
			fmt.Printf("Dealer's Hand: %v\n", round.dealer.hand)
		} else {
			for round.outcome == "" {
				fmt.Println("Enter your action: Hit or Stand or Surrender [H/S/X]: ")
				reader.Scan()
				text = reader.Text()
				switch text {
				case "H", "h":
					card, round.deck = round.deck.Draw()
					round.player.DealCard(card)
					fmt.Println("Your card: ", card)
					if round.player.hand.FavourableValue() == BLACKJACK {
						if round.dealer.hand.FavourableValue() == BLACKJACK {
							round.outcome = GAME_DRAW
							fmt.Printf("Dealer's Hand: %v\n", round.dealer.hand)
							fmt.Println(GAME_DRAW)
						} else {
							round.outcome = GAME_BLACKJACK
							fmt.Printf("Dealer's Hand: %v\n", round.dealer.hand)
							fmt.Println(GAME_BLACKJACK)
						}
					} else if round.player.hand.FavourableValue() >= BUST {
						round.outcome = GAME_BUST
						fmt.Printf("Dealer's Hand: %v\n", round.dealer.hand)
						fmt.Println(GAME_BUST)
					}
				case "S", "s":
					round = round.DealersTurn()
					fmt.Printf("Dealer's Hand: %v\n", round.dealer.hand)
					fmt.Println(round.outcome)
				case "X", "x":
					round.outcome = GAME_LOST
					fmt.Printf("Dealer's Hand: %v\n", round.dealer.hand)
					fmt.Println(GAME_LOST)
				}
			}
		}

	}
}
