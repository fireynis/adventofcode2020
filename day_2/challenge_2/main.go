package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
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

		pos1, err := strconv.Atoi(passwordRules[1])
		if err != nil {
			log.Fatal(err)
		}
		pos1 = pos1 - 1
		pos2, err := strconv.Atoi(passwordRules[2])
		if err != nil {
			log.Fatal(err)
		}
		pos2 = pos2 - 1

		password := passwordRules[4]
		keyRune := passwordRules[3]
		//var firstPosition,secondPosition bool
		//if pos1 < len(password) {
		//	firstPosition = string([]rune(password)[min-1]) == keyRune
		//}
		firstPosition := string([]rune(password)[pos1]) == keyRune
		secondPosition := string([]rune(password)[pos2]) == keyRune
		//if max-1 < len(password)-1 {
		//	secondPosition = string([]rune(password)[max-1]) == keyRune
		//}

		if (firstPosition || secondPosition) && !(firstPosition && secondPosition) {
			correctPasswords++
		}
	}

	fmt.Printf("Correct passwords count: %d\n", correctPasswords)
}
