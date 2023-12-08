package puzzles

import (
	"fmt"
	"os"
	"strings"
)

func find_prime_factors(number int) []int {
	// This function surely does not work for every case, but it works for day8part2 of AoC2023

	prime_factors := []int{}

	for i := 1; ; i++ {
		for j := 1; j < i; j++ {
			if j != 1 && j != i && number%i == 0 {
				prime_factors = append(prime_factors, i)
				number /= i
				break
			}
		}

		if number == 1 {
			return prime_factors
		}
	}
}

func find_lcm(numbers []int) int {
	all_prime_factors := map[int]bool{}

	for _, number := range numbers {
		prime_factors := find_prime_factors(number)
		for _, prime_factor := range prime_factors {
			all_prime_factors[prime_factor] = true
		}
	}

	lcm := 1
	for v := range all_prime_factors {
		lcm *= v
	}

	return lcm
}

func get_path_from_to_xxz(nodes []Node, sequence string, node_name_a string) []Node {
	paths := []Node{}

	node_a, err := find_node_by_name(nodes, node_name_a)
	if err != nil {
		panic(err)
	}

	current_node := node_a
	for i := 0; ; i++ {
		direction := string(sequence[i])

		next_node_name := ""
		if direction == "L" {
			next_node_name = current_node.left
		} else if direction == "R" {
			next_node_name = current_node.right
		} else {
			panic(fmt.Errorf("unknown direction %s", direction))
		}

		next_node, err := find_node_by_name(nodes, next_node_name)
		if err != nil {
			panic(err)
		}

		paths = append(paths, next_node)

		if string(next_node.name[len(next_node.name)-1]) == "Z" {
			return paths
		}

		current_node = next_node

		if i+1 == len(sequence) {
			i = -1
		}
	}
}

func get_simultaneously_step_number(nodes []Node, sequence string) int {
	start_nodes := []Node{}
	for _, node := range nodes {
		if string(node.name[len(node.name)-1]) == "A" {
			start_nodes = append(start_nodes, node)
		}
	}

	numbers_to_find_lcm := []int{}
	for _, node := range start_nodes {
		numbers_to_find_lcm = append(numbers_to_find_lcm, len(get_path_from_to_xxz(nodes, sequence, node.name)))
	}

	return find_lcm(numbers_to_find_lcm)
}

func day8_part2() any {
	input, err := os.ReadFile("resources/8/input.txt")
	if err != nil {
		panic(err)
	}

	input = []byte(strings.TrimSpace(string(input)))

	nodes := parse_nodes(get_nodes_str(string(input)))
	sequence := get_sequence(string(input))

	return get_simultaneously_step_number(nodes, sequence)
}

func init() {
	RegisterPuzzle(8, 2, day8_part2)
}
