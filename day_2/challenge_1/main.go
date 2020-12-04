package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
)

func main() {
	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	passwordRegex := regexp.MustCompile(`(?m)(\d+)-(\d+)\s([A-Za-z]):\s(.*)$`)

	for scanner.Scan() {
		passwordRules := passwordRegex.FindStringSubmatch(scanner.Text())

		if len(passwordRules) < 5 {
			fmt.Println("Password is incorrect")
			continue
		}

		//passwordTestReg := regexp.Compile(`(?m)`)
	}
}
