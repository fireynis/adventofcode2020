package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	passwordRegex := regexp.MustCompile(`(?m)(\d+)-(\d+)\s([A-Za-z]):\s(.*)$`)

	correctPasswords := 0
	for scanner.Scan() {
		passwordRules := passwordRegex.FindStringSubmatch(scanner.Text())

		if len(passwordRules) < 5 {
			fmt.Println("Password is incorrect")
			continue
		}

		min, err := strconv.Atoi(passwordRules[1])
		if err != nil {
			log.Fatal(err)
		}
		max, err := strconv.Atoi(passwordRules[2])
		if err != nil {
			log.Fatal(err)
		}

		count := strings.Count(passwordRules[4], passwordRules[3])

		if count <= max && count >= min {
			correctPasswords++
		}
	}

	fmt.Printf("Correct passwords count: %d\n", correctPasswords)
}
