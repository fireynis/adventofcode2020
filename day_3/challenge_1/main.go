package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

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
	var locPoint point
	treeCount := 0
	for true {
		if //goland:noinspection GoNilness
		hillMap[locPoint.y][locPoint.x] == "#" {
			treeCount++
		}

		if locPoint.y == floor {
			break
		}

		if locPoint.x+3 > maxX {
			locPoint.x = (locPoint.x + 3) - maxX - 1
		} else {
			locPoint.x = locPoint.x + 3
		}

		locPoint.y++
	}

	log.Printf("Trees enroute: %d", treeCount)
}
