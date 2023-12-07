package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func find_number(input string) int {
	var digits []rune
	for pos, char := range input {
		if unicode.IsDigit(char) {
			digits = append(digits, char)
		} else if strings.HasPrefix(input[pos:], "one") {
			digits = append(digits, '1')
		} else if strings.HasPrefix(input[pos:], "two") {
			digits = append(digits, '2')
		} else if strings.HasPrefix(input[pos:], "three") {
			digits = append(digits, '3')
		} else if strings.HasPrefix(input[pos:], "four") {
			digits = append(digits, '4')
		} else if strings.HasPrefix(input[pos:], "five") {
			digits = append(digits, '5')
		} else if strings.HasPrefix(input[pos:], "six") {
			digits = append(digits, '6')
		} else if strings.HasPrefix(input[pos:], "seven") {
			digits = append(digits, '7')
		} else if strings.HasPrefix(input[pos:], "eight") {
			digits = append(digits, '8')
		} else if strings.HasPrefix(input[pos:], "nine") {
			digits = append(digits, '9')
		}
	}
	number := fmt.Sprintf("%c%c", digits[0], digits[len(digits)-1])
	final, err := strconv.Atoi(number)
	if err != nil {
		panic(err)
	}
	return final
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	sum := 0
	for scanner.Scan() {
		sum += find_number(scanner.Text())
	}
	fmt.Println(sum)
}
