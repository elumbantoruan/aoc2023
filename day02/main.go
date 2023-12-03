package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// https://adventofcode.com/2023/day/2

func main() {
	findGames("input01.txt")
}

func findGames(inputFile string) error {
	payload, err := os.ReadFile(inputFile)
	if err != nil {
		return err
	}

	var (
		sumID1 int
		sumID2 int
	)

	scanner := bufio.NewScanner(bytes.NewReader(payload))
	for scanner.Scan() {
		line := scanner.Text()
		gameCubes := strings.Split(line, ": ")
		game := strings.Split(gameCubes[0], " ")
		gID, _ := strconv.Atoi(game[1])
		maxR, maxG, maxB := 0, 0, 0

		turns := strings.Split(gameCubes[1], "; ")
		for _, turn := range turns {
			cubes := strings.Split(turn, ", ")
			for _, cube := range cubes {
				items := strings.Split(cube, " ")
				count, _ := strconv.Atoi(items[0])
				color := items[1]

				if color == "red" {
					maxR = max(maxR, count)
				} else if color == "green" {
					maxG = max(maxG, count)
				} else {
					maxB = max(maxB, count)
				}
			}
		}

		if maxR <= 12 && maxG <= 13 && maxB <= 14 {
			sumID1 += gID
		}
		sumID2 += maxR * maxG * maxB
	}

	fmt.Println(sumID1)
	fmt.Println(sumID2)
	return nil
}
