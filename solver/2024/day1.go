package year_2024

import (
	"math"
	"slices"
	"strconv"
	"strings"
)

type Day1 struct{}

func (d Day1) part1(input string) (int, error) {
	l, r, _ := d.parse(input)

	slices.Sort(l)
	slices.Sort(r)

	var sum int
	for i := 0; i < len(l); i++ {
		sum += int(math.Abs(float64(l[i] - r[i])))
	}

	return sum, nil
}

func (d Day1) parse(input string) ([]int, []int, error) {
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
