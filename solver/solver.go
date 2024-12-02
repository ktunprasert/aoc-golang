package solver

type Solver interface {
	part1(input string) (string, error)
	part2(input string) (string, error)
}
