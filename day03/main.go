package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)

func main() {
	sum, totalGears := sumPartNumbers("input01.txt")
	fmt.Println(sum, totalGears)
}

var offsets = [][]int{
	{0, 1},
	{0, -1},
	{1, 0},
	{1, 1},
	{1, -1},
	{-1, 0},
	{-1, -1},
	{-1, 1},
}

func sumPartNumbers(inputFile string) (int, int) {
	payload, err := os.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(bytes.NewReader(payload))
	var (
		board   [][]byte
		sum     int
		numbers []int
		gears   = make(map[string][]int)
	)

	for scanner.Scan() {
		line := scanner.Text()
		board = append(board, []byte(line))
	}

	for row := 0; row < len(board); row++ {
		foundSymbol := false
		beginNumber := false
		left := 0
		gearRatio := ""
		for col := 0; col < len(board[0]); col++ {
			if unicode.IsNumber(rune(board[row][col])) {
				if !beginNumber {
					left = col
					beginNumber = true
				}
				if !foundSymbol {
					foundSymbol, gearRatio = scanSymbol(board, row, col)
					if gearRatio != "" {
						if _, ok := gears[gearRatio]; !ok {
							gears[gearRatio] = nil
						}
					}
				}
				if col == len(board[0])-1 {
					if foundSymbol {
						part := string(board[row][left : col+1])
						number, _ := strconv.Atoi(part)
						sum += number
						numbers = append(numbers, number)

						if _, ok := gears[gearRatio]; ok {
							gears[gearRatio] = append(gears[gearRatio], number)
						}
						foundSymbol = false
					}
				}
			} else {
				if beginNumber {
					beginNumber = false
					part := string(board[row][left:col])
					number, _ := strconv.Atoi(part)
					if foundSymbol {
						sum += number
						numbers = append(numbers, number)

						if _, ok := gears[gearRatio]; ok {
							gears[gearRatio] = append(gears[gearRatio], number)
						}
						foundSymbol = false
					}
				}
			}
		}
	}

	totalGears := 0
	for _, v := range gears {
		if len(v) == 2 {
			totalGears += v[0] * v[1]
		}
	}

	return sum, totalGears
}

func findGearRatios(engineSchematic [][]byte) int {
	rows := len(engineSchematic)
	cols := len(engineSchematic[0])
	totalGearRatio := 0

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if engineSchematic[i][j] == '*' {
				partCount := 0
				gearRatio := 1

				// Check adjacent numbers (including diagonals)
				for x := -1; x <= 1; x++ {
					for y := -1; y <= 1; y++ {
						newI, newJ := i+x, j+y
						if newI >= 0 && newI < rows && newJ >= 0 && newJ < cols && isPartNumber(engineSchematic[newI][newJ]) {
							partCount++
							gearRatio *= int(engineSchematic[newI][newJ] - '0')
						}
					}
				}

				// If exactly two adjacent part numbers, add the gear ratio to the total
				if partCount == 2 {
					totalGearRatio += gearRatio
				}
			}
		}
	}

	return totalGearRatio
}

func isPartNumber(char byte) bool {
	return char >= '0' && char <= '9'
}

type gearPos struct {
	row int
	col int
}

func sumGearRatios(inputFile string) int {
	return -1
}

func scanSymbol(board [][]byte, row, col int) (bool, string) {

	for _, offset := range offsets {
		rowF := row + offset[0]
		colF := col + offset[1]

		if rowF < 0 || rowF >= len(board) || colF < 0 || colF >= len(board[0]) {
			continue
		}

		if board[rowF][colF] != '.' && !unicode.IsNumber(rune(board[rowF][colF])) {
			if board[rowF][colF] == '*' {
				return true, fmt.Sprintf("%d,%d", rowF, colF)
			} else {
				return true, ""
			}
		}
	}
	return false, ""
}

func scanGears(board [][]byte, row, col int) []int {
	var gear []int

	if board[row][col] == '*' {
		for _, offset := range offsets {
			rowF := row + offset[0]
			colF := col + offset[1]

			if rowF < 0 || rowF >= len(board) || colF < 0 || colF >= len(board[0]) {
				continue
			}

		}

	}

	return gear
}
