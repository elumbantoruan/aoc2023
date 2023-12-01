package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	num, err := produceNumber("input01.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(num)

	num, err = produceNumber2("input01.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(num)
}

func produceNumber(inputFile string) (int, error) {
	payload, err := os.ReadFile(inputFile)
	if err != nil {
		return 0, err
	}
	var total = 0
	scanner := bufio.NewScanner(bytes.NewReader(payload))
	for scanner.Scan() {
		text := scanner.Text()
		left := 0
		right := len(text) - 1

		nl := byte(' ')
		nr := byte(' ')

		for left <= right {
			if nl == ' ' {
				if unicode.IsNumber(rune(text[left])) {
					nl = text[left]
				} else {
					left++
				}
			}

			if nr == ' ' {
				if unicode.IsNumber(rune(text[right])) {
					nr = text[right]
				} else {
					right--
				}
			}
			if nl != ' ' && nr != ' ' {
				break
			}
		}

		numB := string([]byte{nl, nr})
		num, _ := strconv.Atoi(numB)
		total += num
	}
	return total, nil
}

func produceNumber2(inputFile string) (int, error) {

	payload, err := os.ReadFile(inputFile)
	if err != nil {
		return -1, err
	}
	var total int
	scanner := bufio.NewScanner(bytes.NewReader(payload))
	for scanner.Scan() {
		text := scanner.Text()

		// left & right index
		left := 0
		right := len(text) - 1
		// left & right segment index
		ls := left
		rs := right

		nl := byte(' ')
		nr := byte(' ')

		for left <= right {
			if nl == ' ' {
				if unicode.IsNumber(rune(text[left])) {
					nl = text[left]
				} else {
					leftSegment := text[ls : left+1]
					if b, ok := findNumber(leftSegment); ok {
						nl = b
					} else {
						left++
					}
				}
			}

			if nr == ' ' {
				if unicode.IsNumber(rune(text[right])) {
					nr = text[right]
				} else {
					rightSegment := text[right : rs+1]
					if b, ok := findNumber(rightSegment); ok {
						nr = b
					} else {
						right--
					}
				}
			}

			if nl != ' ' && nr != ' ' {
				break
			}
		}

		numB := string([]byte{nl, nr})
		num, _ := strconv.Atoi(numB)
		total += num
	}

	return total, nil
}

func findNumber(s string) (byte, bool) {
	textNumByte := map[string]byte{
		"one":   '1',
		"two":   '2',
		"three": '3',
		"four":  '4',
		"five":  '5',
		"six":   '6',
		"seven": '7',
		"eight": '8',
		"nine":  '9',
	}

	for k, v := range textNumByte {
		if strings.Contains(s, k) {
			return v, true
		}
	}
	return ' ', false
}
