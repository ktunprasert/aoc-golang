package main

import (
	"fmt"
	"os"

	year_2024 "aoc/solver/2024"
)

func main() {
	input, err := os.ReadFile("input/2024-1.txt")
	if err != nil {
		panic(err)
	}

	answer, _ := year_2024.Day1{}.Part1(string(input))
	answer2, _ := year_2024.Day1{}.Part2(string(input))
	fmt.Println(answer, answer2)
}
