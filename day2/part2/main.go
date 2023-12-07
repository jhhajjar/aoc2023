package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func min_power_set(input string) int {
	init_split := strings.Split(input, ":")
	game_str := init_split[1]

	games := strings.Split(game_str, ";")
	max_red, max_blue, max_green := 0, 0, 0
	for _, game := range games {
		cubes := strings.Split(game, ",")
		for _, cube := range cubes {
			number, err := strconv.Atoi(strings.Split(cube, " ")[1])
			handlerr(err)
			color := strings.Split(cube, " ")[2]
			if color == "red" && number > max_red {
				max_red = number
			} else if color == "blue" && number > max_blue {
				max_blue = number
			} else if color == "green" && number > max_green {
				max_green = number
			}
		}
	}

	return max_green * max_red * max_blue
}

func handlerr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	// read in file
	file, err := os.Open("input.txt")
	handlerr(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	sum := 0
	for scanner.Scan() {
		sum += min_power_set(scanner.Text())
	}
	fmt.Println(sum)
}
