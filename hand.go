package main

type Hand []Card

//Returns total value of hand (actual and alternate)
func (hand Hand) TotalValue() (actual, alternate int) {
	for _, card := range hand {
		actual += card.value
		alternate += card.alternateValue
	}
	return actual, alternate
}

func (hand *Hand) AddCard(card *Card) {
	h := append(*hand, *card)
	*hand = h
}

//Adjust the aces in hand in favour of the player
//If; for example; there are 2 aces in hand,
//then the total value should be 12 rather than 22
//for 2 aces and a nine will total to 21
func (hand Hand) FavourableValue() int {
	var acesInHand []Card
	var nonAceCard []Card
	var nonAceCardSum = 0
	var retHandValue = 0
	for _, card := range hand {
		if card.value != card.alternateValue {
			acesInHand = append(acesInHand, card)
		} else {
			nonAceCardSum += card.value
			nonAceCard = append(nonAceCard, card)
		}
	}

	minAcesValue := len(acesInHand)
	diff := BUST - nonAceCardSum

	switch minAcesValue {
	case 1:
		if diff < 11 {
			retHandValue += 1
		} else if diff >= 11 {
			retHandValue += 11
		}
	case 2, 3, 4:
		if diff < 11 {
			retHandValue += minAcesValue
		} else if diff >= 11 {
			retHandValue += minAcesValue + 10
		}
	}

	return retHandValue + nonAceCardSum
}

/*
func swapCardValue(card Card) Card {
	c := card
	c.value = c.alternateValue
	return c
}
*/
