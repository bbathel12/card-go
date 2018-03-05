package card

import (
	"math/rand"
	"time"
)

type Deck struct {
	Cards []Card
}

func NewDeck() (d *Deck) {
	d = new(Deck)

	suits := [4]rune{hearts, diamonds, spades, clubs}
	values := [13]int{two, three, four, five, six, seven, eight, nine,
		ten, jack, queen, king, ace}
	for _, suit := range suits {
		for _, value := range values {
			d.Cards = append(d.Cards, Card{suit, value})
		}
	}
	return
}

func (d *Deck) Shuffle() {
	for i := 0; i < len(d.Cards); i++ {

		rand.Seed(time.Now().UTC().UnixNano())

		r := rand.Intn(len(d.Cards))

		d.Cards[r], d.Cards[i] = d.Cards[i], d.Cards[r]
	}
}

func (d *Deck) Pop() (popped Card, ok bool) {

	if len(d.Cards) >= 1 {
		popped, d.Cards = d.Cards[0], d.Cards[1:]
		ok = true
	} else {
		popped = Card{}
		ok = false
	}

	return

}
