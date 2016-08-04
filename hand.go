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

func (hand Hand) AddCard(card Card) Hand {
  return append(hand, card)
}
