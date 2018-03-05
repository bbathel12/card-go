package card

import (
	"math/rand"
	"time"
)

type Deck struct {
	cards []card
}

func NewDeck() (d *Deck) {
	d = new(Deck)

	suits := [4]rune{hearts, diamonds, spades, clubs}
	values := [13]int{two, three, four, five, six, seven, eight, nine,
		ten, jack, queen, king, ace}
	for _, suit := range suits {
		for _, value := range values {
			d.cards = append(d.cards, card{suit, value})
		}
	}
	return
}

func (d *Deck) Shuffle() {
	for i := 0; i < len(d.cards); i++ {

		rand.Seed(time.Now().UTC().UnixNano())

		r := rand.Intn(len(d.cards))

		d.cards[r], d.cards[i] = d.cards[i], d.cards[r]
	}
}

func (d *Deck) Pop() (popped card, ok bool) {

	if len(d.cards) >= 1 {
		popped, d.cards = d.cards[0], d.cards[1:]
		ok = true
	} else {
		popped = card{}
		ok = false
	}

	return

}
