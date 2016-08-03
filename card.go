package main

//Faces
const (
  TWO = iota
  THREE
  FOUR
  FIVE
  SIX
  SEVEN
  EIGHT
  NINE
  TEN
  JACK
  QUEEN
  KING
  ACE
)

//Suits
const(
  SUIT_HEARTS = "H"
  SUIT_SPADES= "S"
  SUIT_DIAMONDS = "D"
  SUIT_BLACKS = "B"
)

type Card struct {
  suit string
  value int
  alternateValue int  //Same as value except for Aces
}

//Returns a new new card
func NewCard(suit string, value int) Card {
  val, altVal := getCardValues(value)
  card := Card {
    suit = suit,
    value = val,
    alternateValue = altVal,
  }
  return card
}

//Get card values
func getCardValues(val int) (value, alternateValue int) {
  switch {
  case JACK, QUEEN, KING:
    return 10, 10
  case ACE:
    return 11, 1
  }
  return val+2, val+2
}
