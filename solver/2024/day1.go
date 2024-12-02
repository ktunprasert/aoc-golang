package year_2024

import (
	"slices"
	"strconv"
	"strings"
)

type Day1 struct{}

func (d Day1) Part1(input string) (int, error) {
	l, r, _ := d.Parse(input)

	slices.Sort(l)
	slices.Sort(r)

	var sum int
	for i := 0; i < len(l); i++ {
		diff := l[i] - r[i]
		if diff < 0 {
			diff = -diff
		}
		sum += diff
	}

	return sum, nil
}

func (d Day1) Part2(input string) (int, error) {
	l, r, _ := d.Parse(input)

	counter := make(map[int]int)
	for _, v := range r {
		counter[v]++
	}

	var sum int
	for _, v := range l {
		if count, ok := counter[v]; ok {
			sum += v * count
		}
	}

	return sum, nil
}

func (d Day1) Parse(input string) ([]int, []int, error) {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	left := make([]int, len(lines))
	right := make([]int, len(lines))
	for i, line := range lines {
		items := strings.Split(line, " ")
		n_items := make([]int, 0)

		for _, item := range items {
			if item == "" {
				continue
			}

			n, err := strconv.Atoi(item)
			if err != nil {
				return nil, nil, err
			}

			n_items = append(n_items, n)
		}

		left[i] = n_items[0]
		right[i] = n_items[1]
	}

	return left, right, nil
}
