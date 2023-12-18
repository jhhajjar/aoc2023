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
	var curr_num []rune
	var curr_star []int
	star_map := make(map[string][]int)
	for i, line := range engine {
		for j, ch := range line {
			if unicode.IsDigit(ch) {
				curr_num = append(curr_num, ch)
				if curr_star == nil {
					curr_star = look_at_neighbors(engine, i, j)
				}
			} else if curr_star != nil && len(curr_num) >= 0 {
				number, err := strconv.Atoi(string(curr_num))
				handlerr(err)
				key := fmt.Sprintf("%d,%d", curr_star[0], curr_star[1])
				star_map[key] = append(star_map[key], number)
				curr_num = nil
				curr_star = nil
			} else if curr_star == nil && len(curr_num) >= 0 {
				curr_num = nil
			}
		}
		if curr_star != nil && len(curr_num) >= 0 {
			number, err := strconv.Atoi(string(curr_num))
			handlerr(err)
			key := fmt.Sprintf("%d,%d", curr_star[0], curr_star[1])
			star_map[key] = append(star_map[key], number)
			curr_num = nil
			curr_star = nil
		} else if curr_star == nil && len(curr_num) >= 0 {
			curr_num = nil
		}
	}
	sum := 0
	for key := range star_map {
		// fmt.Println(key, star_map[key])
		if len(star_map[key]) != 2 {
			continue
		}
		sum += star_map[key][0] * star_map[key][1]
	}
	return sum
}

func handlerr(err error) {
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}

func look_at_neighbors(engine []string, i, j int) []int {
	offsets := []int{-1, 0, 1}
	for _, i_off := range offsets {
		for _, j_off := range offsets {
			valid_i, valid_j := i+i_off >= 0 && i+i_off < len(engine), j+j_off >= 0 && j+j_off < len(engine[0])
			if valid_i && valid_j {
				neighbor := engine[i+i_off][j+j_off]
				found_star := neighbor == '*'
				if found_star {
					return []int{i + i_off, j + j_off}
				}
			}
		}
	}
	return nil
}
