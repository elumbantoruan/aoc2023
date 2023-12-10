package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	numbers, err := readInput("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	n, _ := problem1(numbers)
	fmt.Println(n)

	n, _ = problem2(numbers)
	fmt.Println(n)
}

func readInput(input string) ([][]int, error) {
	reader, err := os.Open(input)
	if err != nil {
		return nil, err
	}
	var numberList [][]int

	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := scanner.Text()
		nums := strings.Split(line, " ")
		var numbers []int
		for _, num := range nums {
			v, _ := strconv.Atoi(num)
			numbers = append(numbers, v)
		}
		numberList = append(numberList, numbers)
	}
	return numberList, nil
}

func problem1(numberList [][]int) (int, error) {
	var (
		sequenceList [][]int
		total        = 0
	)

	for _, numbers := range numberList {
		var (
			sequences []int
		)

		sequenceList = append(sequenceList, numbers)
		sumNum := 0

		for j := 1; j < len(sequenceList[len(sequenceList)-1]); j++ {
			curr := sequenceList[len(sequenceList)-1][j]
			prev := sequenceList[len(sequenceList)-1][j-1]
			num := curr - prev
			sumNum += num

			sequences = append(sequences, num)
			if j == len(sequenceList[len(sequenceList)-1])-1 {
				sequenceList = append(sequenceList, sequences)
				if sumNum == 0 {
					break
				}
				sequences = nil
				j = 0
				sumNum = 0
			}
		}

		for i := len(sequenceList) - 1; i >= 0; i-- {
			if i == 0 {
				final := sequenceList[i][len(sequenceList[i])-1]
				total += final

				sequenceList = nil

				break
			}
			var sum = 0
			curr := sequenceList[i]
			if i == len(sequenceList)-1 {
				sequenceList[len(sequenceList)-1] = append(sequenceList[len(sequenceList)-1], 0)
			}
			above := sequenceList[i-1]
			sum = curr[len(curr)-1] + above[len(above)-1]
			sequenceList[i-1] = append(sequenceList[i-1], sum)
		}
	}

	return total, nil
}

func problem2(numberList [][]int) (int, error) {
	var (
		sequenceList [][]int
		total        = 0
	)

	for _, numbers := range numberList {
		var (
			sequences []int
		)

		sequenceList = append(sequenceList, numbers)
		sumNum := 0

		for j := 1; j < len(sequenceList[len(sequenceList)-1]); j++ {
			curr := sequenceList[len(sequenceList)-1][j]
			prev := sequenceList[len(sequenceList)-1][j-1]
			num := curr - prev
			sumNum += num

			sequences = append(sequences, num)
			if j == len(sequenceList[len(sequenceList)-1])-1 {
				sequenceList = append(sequenceList, sequences)
				if sumNum == 0 {
					break
				}
				sequences = nil
				j = 0
				sumNum = 0
			}
		}

		for i := len(sequenceList) - 1; i >= 0; i-- {
			if i == 0 {
				final := sequenceList[i][0]
				total += final

				sequenceList = nil

				break
			}
			var sum = 0
			curr := sequenceList[i]
			if i == len(sequenceList)-1 {
				sequenceList[len(sequenceList)-1] = append([]int{0}, sequenceList[len(sequenceList)-1]...)
			}
			above := sequenceList[i-1]
			sum = above[0] - curr[0]
			sequenceList[i-1] = append([]int{sum}, sequenceList[i-1]...)
		}
	}

	return total, nil
}
