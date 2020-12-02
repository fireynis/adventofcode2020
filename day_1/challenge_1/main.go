package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var numbers []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		number, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		numbers = append(numbers, number)
	}

	for numIndex, number := range numbers {
		for i := numIndex; i < len(numbers)-1; i++ {
			if number+numbers[i] == 2020 {
				fmt.Printf("%d", number*numbers[i])
				return
			}
		}
	}
}
