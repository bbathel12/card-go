package card

import (
	"fmt"
	"sort"
	"strconv"
)

type Hand []card

func (h Hand) Len() int {
	return len(h)
}

func (h Hand) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h Hand) Less(i, j int) bool {
	return h[i].Value < h[j].Value
}

func (h *Hand) Sort() {
	sort.Sort(h)
}

func (h Hand) show() {
	for _, v := range h {
		var value string
		switch v.Value {
		case 14:
			value = "A"
		case 13:
			value = "K"
		case 12:
			value = "Q"
		case 11:
			value = "J"
		default:
			value = strconv.Itoa(v.Value)
		}

		fmt.Printf("%2s%1c ", value, v.Suit)
	}
	fmt.Println()
}

func (h Hand) Occurences() map[int]int {
	var occurences map[int]int = make(map[int]int, 15)

	for _, card := range h {
		occurences[card.Value]++
	}

	return occurences
}

func (h Hand) highCard() card {
	return h[len(h)-1]
}
