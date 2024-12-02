package main

import (
	"os"

	year_2024 "aoc/solver/2024"
)

func main() {
	input, err := os.ReadFile("input/2024-1.txt")
	if err != nil {
		panic(err)
	}

	year_2024.Day1{}.Solve(string(input))
}
