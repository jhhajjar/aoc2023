package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func determine_valid(input string) int {
	init_split := strings.Split(input, ":")
	id_str, game_str := init_split[0], init_split[1]
	num := strings.Split(id_str, " ")[1]
	id, err := strconv.Atoi(num)
	handlerr(err)

	games := strings.Split(game_str, ";")
	valid := true
	for _, game := range games {
		cubes := strings.Split(game, ",")
		for _, cube := range cubes {
			cube := strings.TrimSpace(cube)
			number, err := strconv.Atoi(strings.Split(cube, " ")[0])
			handlerr(err)
			color := strings.Split(cube, " ")[1]
			game_valid := (color == "red" && number <= 12) || (color == "blue" && number <= 14) || (color == "green" && number <= 13)
			valid = valid && game_valid
		}
	}

	if valid {
		return id
	}

	return 0
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
		sum += determine_valid(scanner.Text())
	}
	fmt.Println(sum)
}
