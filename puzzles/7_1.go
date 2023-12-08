package puzzles

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type HandType struct {
	name             string
	rank             int
	character_counts []int
}

func (h HandType) parse_hand(hand_str string) error {
	character_counts := map[string]int{}

	for _, character := range hand_str {
		character_counts[string(character)] = strings.Count(hand_str, string(character))
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

var hand_types = []HandType{
	{name: "Five of a kind", rank: 1, character_counts: []int{5}},
	{name: "Four of a kind", rank: 2, character_counts: []int{4, 1}},
	{name: "Full house", rank: 3, character_counts: []int{3, 2}},
	{name: "Three of a kind", rank: 4, character_counts: []int{3, 1, 1}},
	{name: "Two pair", rank: 5, character_counts: []int{2, 2, 1}},
	{name: "One pair", rank: 6, character_counts: []int{2, 1, 1, 1}},
	{name: "High card", rank: 7, character_counts: []int{1, 1, 1, 1, 1}},
}

func parse_hand_type(hand_str string) (HandType, error) {
	for _, hand_type := range hand_types {
		if hand_type.parse_hand(hand_str) == nil {
			return hand_type, nil
		}
	}

	return HandType{}, fmt.Errorf("cannot satisfy any hand type")
}

type CamelCard struct {
	value string
	rank  int
}

var cards = []CamelCard{
	{value: "A", rank: 0},
	{value: "K", rank: 1},
	{value: "Q", rank: 2},
	{value: "J", rank: 3},
	{value: "T", rank: 4},
	{value: "9", rank: 5},
	{value: "8", rank: 6},
	{value: "7", rank: 7},
	{value: "6", rank: 8},
	{value: "5", rank: 9},
	{value: "4", rank: 10},
	{value: "3", rank: 11},
	{value: "2", rank: 12},
}

func parse_cards(hand_str string) ([]CamelCard, error) {
	found_cards := []CamelCard{}

	for _, character := range hand_str {
		found_card := false

		for _, card := range cards {
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

type Hand struct {
	cards     []CamelCard
	hand_type HandType
	bid       int
}

func parse_hand_with_bid(row string) Hand {
	hand_with_bid := Hand{}

	parts := strings.Split(row, " ")

	bid_str := parts[1]
	bid, err := strconv.Atoi(bid_str)
	if err != nil {
		panic(err)
	}

	hand_str := parts[0]

	hand_type, err := parse_hand_type(hand_str)
	if err != nil {
		panic(err)
	}

	cards, err := parse_cards(hand_str)
	if err != nil {
		panic(err)
	}

	hand_with_bid.bid = bid
	hand_with_bid.cards = cards
	hand_with_bid.hand_type = hand_type

	return hand_with_bid
}

func parse_hands_with_bid(input string) []Hand {
	rows := strings.Split(input, "\n")

	hands_with_bid := []Hand{}

	for _, row := range rows {
		hands_with_bid = append(hands_with_bid, parse_hand_with_bid(row))
	}

	return hands_with_bid
}

func sort_hands(hands []Hand) []Hand {
	sorted_hands := make([]Hand, len(hands))
	copy(sorted_hands, hands)

	for i := 0; i < len(sorted_hands)-1; i++ {
		hand_1 := sorted_hands[i]
		hand_2 := sorted_hands[i+1]

		if hand_2.hand_type.rank > hand_1.hand_type.rank {
			sorted_hands[i], sorted_hands[i+1] = sorted_hands[i+1], sorted_hands[i]
			i = -1
		} else if hand_2.hand_type.rank == hand_1.hand_type.rank {
			for card_index := 0; card_index < len(hand_1.cards); card_index++ {
				hand_1_card_rank := hand_1.cards[card_index].rank
				hand_2_card_rank := hand_2.cards[card_index].rank

				if hand_2_card_rank > hand_1_card_rank {
					sorted_hands[i], sorted_hands[i+1] = sorted_hands[i+1], sorted_hands[i]
					i = -1
					break
				} else if hand_2_card_rank < hand_1_card_rank {
					break
				}
			}
		}
	}

	return sorted_hands
}

func calculate_total_winnings(hands []Hand) int {
	total_winnings := 0

	for i, hand := range hands {
		total_winnings += hand.bid * (i + 1)
	}

	return total_winnings
}

func day7_part1() any {
	input, err := os.ReadFile("resources/7/input.txt")
	if err != nil {
		panic(err)
	}

	hands := parse_hands_with_bid(string(input))
	sorted_hands := sort_hands(hands)

	return calculate_total_winnings(sorted_hands)
}

func init() {
	RegisterPuzzle(7, 1, day7_part1)
}
