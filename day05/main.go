package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	g := createGarden("input.txt")
	l := problem1(g)
	fmt.Println(l)
	l = problem2(g)
	fmt.Println(l)
}

func problem1(garden Garden) int {

	var lowestLocation = math.MaxInt
	for _, seed := range garden.Seeds {
		target := seed
		for _, p := range garden.Phases {

			for _, r := range p.Ranges {
				if target >= r.SourceStart && target < r.SourceStart+r.Length {
					target = r.DestinationStart + (target - r.SourceStart)
					break
				}
			}
			if p.Name == "humidity-to-location" {
				lowestLocation = min(lowestLocation, target)
			}
		}

	}
	return lowestLocation
}

func problem2(garden Garden) int {

	var lowestLocation = math.MaxInt
	var seeds []int

	for i := 1; i < len(garden.Seeds); i = i + 2 {
		seed := garden.Seeds[i-1]
		for j := seed; j < seed+garden.Seeds[i]; j++ {
			seeds = append(seeds, j)
		}
	}

	for _, seed := range seeds {
		target := seed
		for _, p := range garden.Phases {

			for _, r := range p.Ranges {
				if target >= r.SourceStart && target < r.SourceStart+r.Length {
					target = r.DestinationStart + (target - r.SourceStart)
					break
				}
			}
			if p.Name == "humidity-to-location" {
				lowestLocation = min(lowestLocation, target)
			}
		}

	}
	return lowestLocation
}

func createGarden(input string) Garden {
	payload, err := os.ReadFile(input)
	if err != nil {
		log.Fatal(err)
	}
	var (
		seeds  []int
		phases []Phase
	)
	scanner := bufio.NewScanner(bytes.NewReader(payload))
	var currPhase Phase
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "seeds:") {
			line = strings.TrimLeft(line, "seeds: ")
			for _, item := range strings.Split(line, " ") {
				num, _ := strconv.Atoi(item)
				seeds = append(seeds, num)
			}
		} else {
			if strings.HasSuffix(line, " map:") {
				currPhase = Phase{}
				currPhase.Name = strings.TrimSuffix(line, " map:")
			} else if len(line) == 0 {
				if len(currPhase.Ranges) > 0 {
					// in between phases
					phases = append(phases, currPhase)
				}
			} else if unicode.IsNumber(rune(line[0])) {
				items := strings.Split(line, " ")
				phaseRange := PhaseRange{}
				phaseRange.DestinationStart, _ = strconv.Atoi(items[0])
				phaseRange.SourceStart, _ = strconv.Atoi(items[1])
				phaseRange.Length, _ = strconv.Atoi(items[2])
				currPhase.Ranges = append(currPhase.Ranges, phaseRange)
			}
		}
	}
	// EOF
	phases = append(phases, currPhase)
	return Garden{
		Seeds:  seeds,
		Phases: phases,
	}
}

type Garden struct {
	Seeds  []int
	Phases []Phase
}

type PhaseRange struct {
	DestinationStart int
	SourceStart      int
	Length           int
}

type Phase struct {
	Name   string
	Ranges []PhaseRange
}
