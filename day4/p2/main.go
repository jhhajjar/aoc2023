package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func handlerr(err error) {
	if err != nil {
		panic(err)
	}
}

func dictGet(dict map[int]int, key int) int {
	if value, ok := dict[key]; ok {
		return value
	}
	return 1
}

func main() {
	file, err := os.Open("../input.txt")
	handlerr(err)
	defer file.Close()

	counter := make(map[int]int)
	scanner := bufio.NewScanner(file)
	line := 1
	counter[line] = 1
	for scanner.Scan() {
		number_matches := game(scanner.Text())
		for i := 1; i <= number_matches; i++ {
			counter[line+i] = dictGet(counter, line+i) + dictGet(counter, line)
		}
		line += 1
	}
	total := 0
	for i := 1; i < line; i++ {
		total += dictGet(counter, i)
	}
	fmt.Println(total)
}

func game(line string) int {
	var game string = strings.TrimSpace(strings.Split(line, ":")[1])
	var winning_string string = strings.TrimSpace(strings.Split(game, "|")[0])
	var our_string string = strings.TrimSpace(strings.Split(game, "|")[1])

	winning_nums := make(map[int]bool)
	for _, num := range strings.Split(winning_string, " ") {
		number, err := strconv.Atoi(num)
		if err != nil {
			continue
		}
		winning_nums[number] = true
	}
	num_matches := 0
	for _, num := range strings.Split(our_string, " ") {
		number, err := strconv.Atoi(num)
		if err != nil {
			continue
		}
		if winning_nums[number] {
			num_matches += 1
		}
	}
	return num_matches
}
