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

var yearRegex = regexp.MustCompile(`^\d{4}$`)
var pidRegex = regexp.MustCompile(`^\d{9}$`)
var hclRegex = regexp.MustCompile(`^#[0-9a-f]{6}$`)

var eyeColors = []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}

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
				if len(fields) == 8 && validData(fields) {
					validPassports++
				}
			} else {
				if len(fields) == 7 && validData(fields) {
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
		if len(fields) == 8 && validData(fields) {
			validPassports++
		}
	} else {
		if len(fields) == 7 && validData(fields) {
			validPassports++
		}
	}
	fmt.Printf("Valid passports: %d", validPassports)
}

func validData(fields map[string]string) bool {
	for name, data := range fields {
		switch name {
		case "byr":
			if !(validateYear(data, 1920, 2002)) {
				return false
			}
			continue
		case "iyr":
			if !(validateYear(data, 2010, 2020)) {
				return false
			}
			continue
		case "eyr":
			if !(validateYear(data, 2020, 2030)) {
				return false
			}
			continue
		case "hcl":
			if !(hclRegex.MatchString(data)) {
				return false
			}
			continue
		case "pid":
			if !(pidRegex.MatchString(data)) {
				return false
			}
			continue
		case "ecl":
			var valid bool
			for _, color := range eyeColors {
				if color == data {
					valid = true
					break
				}
			}
			if !valid {
				return false
			}
			continue
		case "hgt":
			min := 59
			max := 76
			if strings.Contains(data, "cm") {
				min = 150
				max = 193
			}
			data = strings.Replace(strings.Replace(data, "cm", "", -1), "in", "", -1)
			height, err := strconv.Atoi(data)
			if err != nil {
				return false
			}

			if !(height >= min && height <= max) {
				return false
			}
		}
	}
	return true
}

func validateYear(year string, min, max int) bool {
	if !(yearRegex.MatchString(year)) {
		return false
	}

	yearInt, err := strconv.Atoi(year)

	if err != nil {
		return false
	}

	return yearInt >= min && yearInt <= max
}
