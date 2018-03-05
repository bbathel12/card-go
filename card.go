package card

import (
	_ "fmt"
)

var BlankCard = Card{}

const (
	spades   = '♠'
	hearts   = '♥'
	diamonds = '♦'
	clubs    = '♣'
)

const (
	_ = iota
	_
	two
	three
	four
	five
	six
	seven
	eight
	nine
	ten
	jack
	queen
	king
	ace
)

type corner struct {
	X, Y int
}

type pixel struct {
	X, Y   int
	Symbol rune
}

type Card struct {
	Suit  rune
	Value int
}

func (c Card) equal(c2 Card) bool {
	if c.Suit == c2.Suit && c.Value == c.Value {
		return true
	}
	return false
}
