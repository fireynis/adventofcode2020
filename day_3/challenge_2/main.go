package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var slopes = []slope{{1, 1}, {3, 1}, {5, 1}, {7, 1}, {1, 2}}

type slope struct {
	x int
	y int
}

type point struct {
	x int
	y int
}

func main() {
	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var hillMap [][]string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		text := scanner.Text()

		row := strings.Split(text, "")
		hillMap = append(hillMap, row)
	}

	if len(hillMap) < 1 {
		log.Fatal(fmt.Errorf("no lines in index file"))
	}

	//goland:noinspection GoNilness
	maxX := len(hillMap[0]) - 1
	floor := len(hillMap) - 1
	var treeCounts []int

	for _, curSlope := range slopes {
		count := 0
		var locPoint point
		for true {
			if //goland:noinspection GoNilness
			hillMap[locPoint.y][locPoint.x] == "#" {
				count++
			}

			if locPoint.y == floor {
				break
			}

			if locPoint.x+curSlope.x > maxX {
				locPoint.x = (locPoint.x + curSlope.x) - maxX - 1
			} else {
				locPoint.x = locPoint.x + curSlope.x
			}

			if locPoint.y+curSlope.y > floor {
				locPoint.y = floor
			} else {
				locPoint.y = locPoint.y + curSlope.y
			}
		}
		treeCounts = append(treeCounts, count)
	}

	totalMulti := 1
	for _, treeCount := range treeCounts {
		totalMulti = totalMulti * treeCount
	}

	log.Printf("Trees enroute: %d", totalMulti)
}
