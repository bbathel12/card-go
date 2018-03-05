package card

import (
	_ "fmt"
	"testing"
)

func Test_Equal(t *testing.T) {
	h10 := card{hearts, 10}
	d10 := card{diamonds, 10}
	//	s10 := card{spades, 10}
	c1 := card{clubs, 1}
	c1c := card{clubs, 1}

	if h10.Equal(d10) {
		t.Error("h10 shouldn't Equal d10")
	}

	if d10.Equal(c1) {
		t.Error("d10 shouldn't Equal c1")
	}

	if !c1.Equal(c1c) {
		t.Error("c1 should Equal c1c")
	}
}

func Test_Shuffle(t *testing.T) {
	d := NewDeck()
	copyofcards := make([]card, len(d.cards))
	copy(copyofcards, d.cards)
	d.Shuffle()

	var matches int = 0
	for k, card := range d.cards {
		if card.Equal(copyofcards[k]) {
			matches++
		}
	}
	//	fmt.Println(matches)
	if matches > len(d.cards)/2 {
		t.Error("more than 50% match")
	}

}

func Test_pop(t *testing.T) {
	d := NewDeck()
	d.Shuffle()

	for _, card := range d.cards {
		oldLength := len(d.cards)
		popped, ok := d.pop()
		if ok && !popped.Equal(card) {
			t.Error(popped, "does not Equal", card)
		}
		if oldLength == len(d.cards)-1 {
			t.Error(oldLength, "didn't shrink", len(d.cards))
		}
	}
}

func Benchmark_Shuffle(b *testing.B) {
	d := NewDeck()

	for i := 0; i <= b.N; i++ {
		d.Shuffle()
	}

}
