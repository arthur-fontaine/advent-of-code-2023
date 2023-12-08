package puzzles

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func (h HandType) parse_hand_with_joker(hand_str string) error {
	character_counts := map[string]int{}

	for _, character := range hand_str {
		character_counts[string(character)] = strings.Count(hand_str, string(character))
	}

	joker_number := character_counts["J"]
	if joker_number < 5 {
		delete(character_counts, "J")
	} else {
		joker_number = 0
	}

	must_character_counts := make([]int, len(h.character_counts))
	copy(must_character_counts, h.character_counts)

	for _, character_count := range character_counts {
		must_character_count_index := -1
		for i, v := range must_character_counts {
			if v == character_count {
				must_character_count_index = i
				break
			}
		}

		// Try with joker
		if must_character_count_index == -1 {
			for i, v := range must_character_counts {
				if v == character_count+joker_number {
					must_character_count_index = i
					joker_number = 0
					break
				}
			}
		}

		if must_character_count_index == -1 {
			return fmt.Errorf("hand not type of %v", h.name)
		}

		must_character_counts = append(must_character_counts[:must_character_count_index], must_character_counts[must_character_count_index+1:]...)
	}

	if len(must_character_counts) == 0 {
		return nil
	}

	return fmt.Errorf("hand not type of %v", h.name)
}

func parse_hand_type_with_joker(hand_str string) (HandType, error) {
	for _, hand_type := range hand_types {
		if hand_type.parse_hand_with_joker(hand_str) == nil {
			return hand_type, nil
		}
	}

	return HandType{}, fmt.Errorf("cannot satisfy any hand type")
}

var cards_with_joker = []CamelCard{
	{value: "A", rank: 0},
	{value: "K", rank: 1},
	{value: "Q", rank: 2},
	{value: "T", rank: 4},
	{value: "9", rank: 5},
	{value: "8", rank: 6},
	{value: "7", rank: 7},
	{value: "6", rank: 8},
	{value: "5", rank: 9},
	{value: "4", rank: 10},
	{value: "3", rank: 11},
	{value: "2", rank: 12},
	{value: "J", rank: 13},
}

func parse_cards_with_joker(hand_str string) ([]CamelCard, error) {
	found_cards := []CamelCard{}

	for _, character := range hand_str {
		found_card := false

		for _, card := range cards_with_joker {
			if card.value == string(character) {
				found_cards = append(found_cards, card)
				found_card = true
				break
			}
		}

		if !found_card {
			return found_cards, fmt.Errorf("cannot find card satisfies %v", string(character))
		}
	}

	return found_cards, nil
}

func parse_hand_with_joker_with_bid(row string) Hand {
	hand_with_bid := Hand{}

	parts := strings.Split(row, " ")

	bid_str := parts[1]
	bid, err := strconv.Atoi(bid_str)
	if err != nil {
		panic(err)
	}

	hand_str := parts[0]

	hand_type, err := parse_hand_type_with_joker(hand_str)
	if err != nil {
		panic(err)
	}

	cards, err := parse_cards_with_joker(hand_str)
	if err != nil {
		panic(err)
	}

	hand_with_bid.bid = bid
	hand_with_bid.cards = cards
	hand_with_bid.hand_type = hand_type

	return hand_with_bid
}

func parse_hands_with_joker_with_bid(input string) []Hand {
	rows := strings.Split(input, "\n")

	hands_with_bid := []Hand{}

	for _, row := range rows {
		hands_with_bid = append(hands_with_bid, parse_hand_with_joker_with_bid(row))
	}

	return hands_with_bid
}

func day7_part2() any {
	input, err := os.ReadFile("resources/7/input.txt")
	if err != nil {
		panic(err)
	}

	hands := parse_hands_with_joker_with_bid(string(input))
	sorted_hands := sort_hands(hands)

	return calculate_total_winnings(sorted_hands)
}

func init() {
	RegisterPuzzle(7, 2, day7_part2)
}
