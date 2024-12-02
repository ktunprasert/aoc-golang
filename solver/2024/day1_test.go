package year_2024

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var example = `
3   4
4   3
2   5
1   3
3   9
3   3
`

func TestDay1_P1_Example(t *testing.T) {
	solver := Day1{}
	ans, err := solver.part1(example)
	require.NoError(t, err)
	require.Equal(t, 11, ans)
}
