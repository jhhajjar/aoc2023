package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func handlerr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	file, err := os.Open("../input.txt")
	handlerr(err)
	defer file.Close()

	sum := 0.0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		sum += game(scanner.Text())
	}
	fmt.Println(sum)
}

func game(line string) float64 {
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
	if num_matches > 0 {
		return math.Pow(2, float64(num_matches-1))
	} else {
		return 0
	}
}
