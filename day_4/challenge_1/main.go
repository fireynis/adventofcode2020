package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	fields := make(map[string]string)
	validPassports := 0
	var line string
	for scanner.Scan() {
		line = scanner.Text()
		if line == "" {
			if _, ok := fields["cid"]; ok {
				if len(fields) == 8 {
					validPassports++
				}
			} else {
				if len(fields) == 7 {
					validPassports++
				}
			}
			fields = make(map[string]string)
			continue
		}
		passportFields := strings.Split(line, " ")

		for _, passFields := range passportFields {
			data := strings.Split(passFields, ":")
			fields[data[0]] = data[1]
		}
	}

	if _, ok := fields["cid"]; ok {
		if len(fields) == 8 {
			validPassports++
		}
	} else {
		if len(fields) == 7 {
			validPassports++
		}
	}
	fmt.Printf("Valid passports: %d", validPassports)
}
