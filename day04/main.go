package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"math"
	"os"
	"strings"
)

// https://adventofcode.com/2023/day/4
func main() {
	totalScore, totalCards := calculateWorthPoint("input01.txt")
	fmt.Println(totalScore, totalCards)
}

func calculateWorthPoint(input string) (int, int) {
	payload, err := os.ReadFile(input)
	if err != nil {
		log.Fatal(err)
	}
	var (
		totalScore int
		totalCards int
		count      int
		matches    []int
	)
	bscanner := bufio.NewScanner(bytes.NewReader(payload))
	for bscanner.Scan() {
		count = 0
		line := bscanner.Text()
		cardItems := strings.Split(line, ": ")
		listNumbers := strings.Split(cardItems[1], "|")
		winningNumbers := strings.Split(listNumbers[0], " ")
		numbers := strings.Split(listNumbers[1], " ")

		var keys = make(map[string]interface{})

		for _, win := range winningNumbers {
			if win != "" {
				keys[win] = nil
			}
		}

		for _, num := range numbers {
			if _, ok := keys[num]; ok {
				count++
			}
		}

		matches = append(matches, count)

		totalScore += int(math.Pow(float64(2), float64(count-1)))

	}

	copies := make([]int, len(matches))
	for i := 0; i < len(copies); i++ {
		copies[i] = 1
	}

	for i, n := range matches {
		start := i + 1
		end := i + n + 1
		for j := start; j < end; j++ {
			copies[j] += copies[i]
		}
	}

	for _, v := range copies {
		totalCards += v
	}

	return totalScore, totalCards
}
