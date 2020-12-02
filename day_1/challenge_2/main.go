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

	maxLength := len(numbers) - 1
	for numIndex, number := range numbers {
		for i := numIndex; i < maxLength; i++ {
			sumOne := number + numbers[i]
			for j := 0; j < maxLength; j++ {
				if j == i || j == numIndex {
					continue
				}
				if sumOne+numbers[j] == 2020 {
					fmt.Printf("%d\n", number*numbers[i]*numbers[j])
					return
				}
			}
		}
	}
}
