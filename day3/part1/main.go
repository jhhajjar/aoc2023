package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func main() {
	file, err := os.Open("../input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var engine []string
	for scanner.Scan() {
		r := scanner.Text()
		engine = append(engine, r)
	}
	sum := analyze(engine)
	fmt.Println(sum)
}

func analyze(engine []string) int {
	sum := 0
	var currNum []rune
	relevant := false
	for i, line := range engine {
		for j, ch := range line {
			if unicode.IsDigit(ch) {
				currNum = append(currNum, ch)
				relevant = look_at_neighbors(engine, i, j) || relevant
			} else if relevant && len(currNum) >= 0 {
				number, err := strconv.Atoi(string(currNum))
				handlerr(err)
				sum += number
				currNum = nil
				relevant = false
			} else if !relevant && len(currNum) >= 0 {
				currNum = nil
			}
		}
		if relevant && len(currNum) >= 0 {
			number, err := strconv.Atoi(string(currNum))
			handlerr(err)
			sum += number
			currNum = nil
			relevant = false
		}
	}
	return sum
}

func handlerr(err error) {
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}

func look_at_neighbors(engine []string, i, j int) bool {
	offsets := []int{-1, 0, 1}
	has_sign := false
	for _, i_off := range offsets {
		for _, j_off := range offsets {
			valid_i, valid_j := i+i_off >= 0 && i+i_off < len(engine), j+j_off >= 0 && j+j_off < len(engine[0])
			if valid_i && valid_j {
				neighbor := engine[i+i_off][j+j_off]
				curr_has_sign := !(neighbor == '.' || (neighbor >= '0' && neighbor <= '9'))
				has_sign = curr_has_sign || has_sign
			}
		}
	}
	return has_sign
}
