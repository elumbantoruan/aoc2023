package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	count := problem1("input.txt")
	fmt.Println(count)
	count = problem2("input.txt")
	fmt.Println(count)
}

func problem1(input string) int {
	payload, err := os.ReadFile(input)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(bytes.NewReader(payload))
	var (
		times     []int
		distances []int
	)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "Time:") {
			timeItems := strings.Split(strings.Trim(strings.Split(line, ":")[1], " "), " ")
			for _, ti := range timeItems {
				if len(ti) == 0 {
					continue
				}
				num, _ := strconv.Atoi(ti)
				times = append(times, num)
			}
		} else if strings.HasPrefix(line, "Distance:") {
			distanceItems := strings.Split(strings.Trim(strings.Split(line, ":")[1], " "), " ")
			for _, di := range distanceItems {
				if len(di) == 0 {
					continue
				}
				dit, _ := strconv.Atoi(di)
				distances = append(distances, dit)
			}
		}
	}

	n := len(times)
	counts := []int{}
	count := 0
	product := 1
	for i := 0; i < n; i++ {
		count = 0
		for t := 1; t <= times[i]; t++ {
			dist := (times[i] - t) * t
			if dist > distances[i] {
				count++
			}
		}
		if count > 0 {
			counts = append(counts, count)
		}
	}
	for i := 0; i < n; i++ {
		product *= counts[i]
	}
	return product
}

func problem2(input string) int {
	payload, err := os.ReadFile(input)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(bytes.NewReader(payload))
	var (
		times     int
		distances int
	)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "Time:") {
			timeItems := strings.Split(strings.Trim(strings.Split(line, ":")[1], " "), " ")
			timeBuffer := ""
			for _, ti := range timeItems {
				if len(ti) == 0 {
					continue
				}
				timeBuffer += ti
			}
			times, _ = strconv.Atoi(timeBuffer)
		} else if strings.HasPrefix(line, "Distance:") {
			distanceItems := strings.Split(strings.Trim(strings.Split(line, ":")[1], " "), " ")
			distanceBuffer := ""
			for _, di := range distanceItems {
				if len(di) == 0 {
					continue
				}
				distanceBuffer += di
			}
			distances, _ = strconv.Atoi(distanceBuffer)
		}
	}

	count := 0

	for t := 1; t <= times; t++ {
		dist := (times - t) * t
		if dist > distances {
			count++
		}
	}

	return count
}
