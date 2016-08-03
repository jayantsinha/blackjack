package main

type Hand []Card

//Returns total value of hand (actual and alternate)
func (hand Hand) Value() actual, alternate int {
  for card := range hand {
    actual += card.value
    alternate += card.alternateValue
  }
  return actual alternate;
}

func (hand Hand) AddCard(card Card) Hand {
  return append(hand, card)
}
