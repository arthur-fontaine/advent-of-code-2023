package puzzles

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Card struct {
	id              int
	winning_numbers []int
	own_numbers     []int
	scratched       bool
}

func (c Card) get_matching_numbers() int {
	matching_numbers := 0
	for _, own_number := range c.own_numbers {
		for _, winning_number := range c.winning_numbers {
			if own_number == winning_number {
				matching_numbers++
				break
			}
		}
	}

	return matching_numbers
}

func create_card(card_row string) Card {
	card := Card{scratched: false}

	card_description := strings.Split(card_row, ":")

	card_id, err := strconv.Atoi(strings.Trim(strings.Split(card_description[0], "Card ")[1], " "))
	if err != nil {
		panic(err)
	}
	card.id = card_id

	separated_numbers := strings.Split(card_description[1], "|")
	winning_numbers := strings.Split(strings.Trim(separated_numbers[0], " "), " ")
	own_numbers := strings.Split(strings.Trim(separated_numbers[1], " "), " ")

	for _, own_number := range own_numbers {
		if len(own_number) > 0 {
			number, err := strconv.Atoi(own_number)
			if err != nil {
				panic(err)
			}
			card.own_numbers = append(card.own_numbers, number)
		}
	}

	for _, winning_number := range winning_numbers {
		if len(winning_number) > 0 {
			number, err := strconv.Atoi(winning_number)
			if err != nil {
				panic(err)
			}
			card.winning_numbers = append(card.winning_numbers, number)
		}
	}

	return card
}

func find_card_by_id(cards []Card, card_index int) (Card, error) {
	for _, card := range cards {
		if card.id == card_index {
			return card, nil
		}
	}

	return Card{}, fmt.Errorf("card not found")
}

var cached_cards = map[int][]Card{}

func play_scratch_round(cards *[]Card, card_index int) {
	if card_index >= len(*cards) {
		return
	}

	card := (*cards)[card_index]

	card.scratched = true

	added_cards := []Card{}

	if cached_cards[card.id] != nil {
		added_cards = cached_cards[card.id]
	} else {
		for i := card.id + 1; i <= card.id+card.get_matching_numbers(); i++ {
			card_copy, err := find_card_by_id(*cards, i)
			if err != nil {
				continue
			}
			card_copy.scratched = false
			added_cards = append(added_cards, card_copy)
		}

		cached_cards[card.id] = added_cards
	}

	*cards = append(*cards, added_cards...)

	(*cards)[card_index] = card
}

func day4_part2() any {
	input, err := os.ReadFile("resources/4/input.txt")
	if err != nil {
		panic(err)
	}

	cards := []Card{}
	for _, card_line := range strings.Split(string(input), "\n") {
		cards = append(cards, create_card(card_line))
	}

	for card_index := 0; !cards[len(cards)-1].scratched; card_index++ {
		play_scratch_round(&cards, card_index)
	}

	return len(cards)
}

func init() {
	RegisterPuzzle(4, 2, day4_part2)
}
