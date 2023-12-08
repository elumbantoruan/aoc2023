package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {

	n, err := problem1("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(n)

	n, err = problem2("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(n)

}

func problem1(input string) (int, error) {
	payload, err := os.ReadFile(input)
	if err != nil {
		return -1, err
	}
	var (
		handsMap = make(map[Hand][]HandBid)
		count    = 0
	)
	var cardsStrength = map[byte]int{
		'2': 0,
		'3': 1,
		'4': 2,
		'5': 3,
		'6': 4,
		'7': 5,
		'8': 6,
		'9': 7,
		'T': 8,
		'J': 9,
		'Q': 10,
		'K': 11,
		'A': 12,
	}

	scanner := bufio.NewScanner(bytes.NewReader(payload))
	for scanner.Scan() {
		line := scanner.Text()
		items := strings.Split(line, " ")
		labels := items[0]
		bid, _ := strconv.Atoi(items[1])
		handType, _ := getHandType(items[0], false)

		handsMap[handType] = append(handsMap[handType], HandBid{Hand: handType, Labels: labels, Bid: bid})
		count++
	}

	var ranks = make([]int, count)
	var rankIndex = count - 1
	var handTypesCount = 6

	// start from the biggest and assign to the latest rank
	for i := handTypesCount; i >= 0; i-- {
		handType := Hand(i)
		if vals, ok := handsMap[handType]; ok {
			if len(vals) == 1 {
				ranks[rankIndex] = vals[0].Bid
				rankIndex--
			} else {
				sort.SliceStable(vals, func(x, y int) bool {

					lbx := vals[x].Labels
					lby := vals[y].Labels
					for b := 0; b < len(lbx); b++ {
						if lbx[b] != lby[b] {
							return cardsStrength[lbx[b]] < cardsStrength[lby[b]]
						}
					}
					return lbx < lby
				})
				for n := len(vals) - 1; n >= 0; n-- {
					ranks[rankIndex] = vals[n].Bid
					rankIndex--
				}
			}
		}
	}

	var sum int
	for i := 0; i < len(ranks); i++ {
		sum += (i + 1) * ranks[i]
	}

	return sum, nil

}

func problem2(input string) (int, error) {
	payload, err := os.ReadFile(input)
	if err != nil {
		return -1, err
	}
	var (
		handsMap = make(map[Hand][]HandBid)
		count    = 0
	)
	var cardsStrength = map[byte]int{
		'J': 0,
		'2': 1,
		'3': 2,
		'4': 3,
		'5': 4,
		'6': 5,
		'7': 6,
		'8': 7,
		'9': 8,
		'T': 9,
		'Q': 10,
		'K': 11,
		'A': 12,
	}

	scanner := bufio.NewScanner(bytes.NewReader(payload))
	for scanner.Scan() {
		line := scanner.Text()
		items := strings.Split(line, " ")
		labels := items[0]
		bid, _ := strconv.Atoi(items[1])
		handType, jokerLabel := getHandType(items[0], true)

		handsMap[handType] = append(handsMap[handType], HandBid{Hand: handType, Labels: labels, Bid: bid, Joker: jokerLabel})
		count++
	}

	var ranks = make([]int, count)
	var rankIndex = count - 1
	var handTypesCount = 6

	// start from the biggest and assign to the latest rank
	for i := handTypesCount; i >= 0; i-- {
		handType := Hand(i)
		if vals, ok := handsMap[handType]; ok {
			if len(vals) == 1 {
				ranks[rankIndex] = vals[0].Bid
				rankIndex--
			} else {
				sort.SliceStable(vals, func(x, y int) bool {

					lbx := vals[x].Labels
					lby := vals[y].Labels
					for b := 0; b < len(lbx); b++ {
						if lbx[b] != lby[b] {
							return cardsStrength[lbx[b]] < cardsStrength[lby[b]]
						}
					}
					return lbx < lby
				})
				for n := len(vals) - 1; n >= 0; n-- {
					ranks[rankIndex] = vals[n].Bid
					rankIndex--
				}
			}
		}
	}

	var sum int
	for i := 0; i < len(ranks); i++ {
		sum += (i + 1) * ranks[i]
	}

	return sum, nil

}

type HandBid struct {
	Hand   Hand
	Labels string
	Joker  string
	Bid    int
}

type Hand int

const (
	HandHighCard Hand = iota
	HandOnePair
	HandTwoPair
	HandThreeKind
	HandFullHouse
	HandFourKind
	HandFiveKind
)

func getHandType(s string, jokerMode bool) (Hand, string) {
	var jokerLabel string
	if jokerMode {
		sbytes := []byte(s)
		jc := bytes.Count(sbytes, []byte{'J'})
		if jc > 0 {
			countType := make(map[byte]int)
			for i := 0; i < len(sbytes); i++ {
				if sbytes[i] == 'J' {
					continue
				}
				countType[sbytes[i]]++
			}
			maxV := 0
			maxK := byte(' ')
			for k, v := range countType {
				if v > maxV {
					maxV = v
					maxK = k
				}
			}
			for i := 0; i < len(sbytes); i++ {
				if sbytes[i] == 'J' && jc > 0 {
					sbytes[i] = maxK
					jc--
				}
			}
			jokerLabel = string(sbytes)
		}
	}
	handMap := make(map[byte]int)
	if jokerMode && jokerLabel != "" {
		s = jokerLabel
	}
	for i := 0; i < len(s); i++ {
		handMap[s[i]]++
	}
	kinds := len(handMap)
	switch kinds {
	case 1:
		return HandFiveKind, jokerLabel // all the same
	case 2: // two different kind of hands
		maxVal := 0
		for _, v := range handMap {
			maxVal = max(maxVal, v)
		}
		if maxVal == 4 {
			return HandFourKind, jokerLabel // four are the same and one different, i.e: AA8AA
		} else {
			return HandFullHouse, jokerLabel // three different and two different of types, i.e: 23332
		}
	case 3: // three different kind of hands
		maxVal := 0
		for _, v := range handMap {
			maxVal = max(maxVal, v)
		}
		if maxVal == 3 { // three are the same and remaining two cards are each different, i.e: TTT98
			return HandThreeKind, jokerLabel
		} else if maxVal == 2 {
			return HandTwoPair, jokerLabel // two cards share one label, two other cards share a second label, and one has a third label, i.e: 23432
		}
	case 4: // four different kind of hands, two cards share one label, and other three cards have a different label, i.e: A23A4
		return HandOnePair, jokerLabel
	case 5:
		return HandHighCard, jokerLabel
	}
	return HandHighCard, jokerLabel
}
