package main

import (
	"aoc2023/pkg/math"
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	g, err := createGraph("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	n := problem1(g)
	fmt.Println(n)

	n = problem2(g)
	fmt.Println(n)
}

func problem1(graph *Graph) int {

	count := 0
	nodes := graph.Maps["AAA"]
	node := ""
	for node != "ZZZ" {
		for _, dir := range graph.Direction {
			if dir == rune('L') {
				node = nodes[0]
			} else {
				node = nodes[1]
			}
			count++
			if node == "ZZZ" {
				return count
			}
			nodes = graph.Maps[node]
		}
	}

	return count
}

func problem2(graph *Graph) int {
	nodes := graph.Maps

	var curs []string
	for name := range nodes {
		if name[2] == 'A' {
			curs = append(curs, name)
		}
	}

	distance := 0
	distances := make(map[int]int)
	for i := 0; ; i = (i + 1) % len(graph.Direction) {
		if len(distances) == len(curs) {
			break
		}

		dir := graph.Direction[i]
		for curIdx, name := range curs {
			if dir == 'L' {
				curs[curIdx] = nodes[name][0]
			} else {
				curs[curIdx] = nodes[name][1]
			}
		}
		distance++

		for idx, cur := range curs {
			if cur[2] == 'Z' {
				if _, exists := distances[idx]; exists {
					continue
				}
				distances[idx] = distance
			}
		}
	}

	numbers := make([]int, 0, len(curs))
	for _, d := range distances {
		numbers = append(numbers, d)
	}
	return math.LeastCommonMultipleList(numbers)
}

func createGraph(input string) (*Graph, error) {
	payload, err := os.ReadFile(input)
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(bytes.NewReader(payload))
	i := 0

	var (
		maps       = make(map[string][]string)
		keys       []string
		directions string
	)

	for scanner.Scan() {
		line := scanner.Text()
		if i == 0 {
			directions = line
			i++
		} else if line == "" {
			continue
		} else {
			items := strings.Split(line, "=")
			key := strings.Trim(items[0], " ")
			keys = append(keys, key)
			vals := strings.Trim(items[1], " (")
			vals = strings.Trim(vals, ")")
			values := strings.Split(vals, ", ")
			maps[key] = []string{values[0], values[1]}
		}
	}
	return &Graph{
		Direction: directions,
		Maps:      maps,
		Keys:      keys,
	}, nil
}

type Graph struct {
	Direction string
	Keys      []string
	Maps      map[string][]string
}
