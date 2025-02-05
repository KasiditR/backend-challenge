package main

import (
	"errors"
	"fmt"
	"strings"
)

func main() {
	var input string
	fmt.Print("Type L R = : ")
	fmt.Scanf("%v", &input)
	err := validateInput(input)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(encoded(input))
}

func encoded(input string) []int {
	n := len(input)
	numbers := make([]int, n+1)

	for i, char := range input {
		switch char {
		case 'L':
			if numbers[i] <= numbers[i+1] {
				numbers[i] = numbers[i+1] + 1
				if i > 0 && input[i-1] == 'L' && numbers[i] == numbers[i-1] {
					numbers[i-1]++
				}
			}
		case 'R':
			if numbers[i+1] <= numbers[i] {
				numbers[i+1] = numbers[i] + 1
			}
		}
	}

	for i, char := range input {
		if char == '=' {
			if i+1 == len(input) {
				numbers[i+1] = numbers[i]
			} else {
				numbers[i] = numbers[i+1]
			}
		}
	}

	return numbers
}

func validateInput(input string) error {
	countL := strings.Count(input, "L")
	countR := strings.Count(input, "R")
	countEqual := strings.Count(input, "=")

	if (countL + countR + countEqual) != len(input) {
		return errors.New("invalid input: must contain only L, R, and =")
	}

	return nil
}
