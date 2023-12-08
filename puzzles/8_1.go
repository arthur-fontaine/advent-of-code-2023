package puzzles

import (
	"fmt"
	"os"
	"strings"
)

type Node struct {
	name  string
	left  string
	right string
}

func find_node_by_name(nodes []Node, name string) (Node, error) {
	for _, node := range nodes {
		if node.name == name {
			return node, nil
		}
	}

	return Node{}, fmt.Errorf("node not found")
}

func get_path_from_to(nodes []Node, sequence string, node_name_a string, node_name_b string) []Node {
	paths := []Node{}

	if node_name_a == node_name_b {
		return paths
	}

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

		if next_node.name == node_name_b {
			return paths
		}

		current_node = next_node

		if i+1 == len(sequence) {
			i = -1
		}
	}
}

func get_sequence(input string) string {
	return strings.Split(input, "\n")[0]
}

func get_nodes_str(input string) []string {
	return strings.Split(input, "\n")[2:]
}

func parse_nodes(nodes_str []string) []Node {
	nodes := []Node{}

	for _, node_str := range nodes_str {
		node := Node{}

		node_parts := strings.Split(node_str, " = ")

		node.name = node_parts[0]
		node_lrs := strings.Split(strings.Trim(node_parts[1], "()"), ", ")
		node.left = node_lrs[0]
		node.right = node_lrs[1]

		nodes = append(nodes, node)
	}

	return nodes
}

func day8_part1() any {
	input, err := os.ReadFile("resources/8/input.txt")
	if err != nil {
		panic(err)
	}

	nodes := parse_nodes(get_nodes_str(string(input)))
	sequence := get_sequence(string(input))

	path := get_path_from_to(nodes, sequence, "AAA", "ZZZ")

	return len(path)
}

func init() {
	RegisterPuzzle(8, 1, day8_part1)
}
