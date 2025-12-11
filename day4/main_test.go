package day4

import "testing"

func TestSolve(t *testing.T) {
	input := "..@@.@@@@.\r\n@@@.@.@.@@\r\n@@@@@.@.@@\r\n@.@@@@..@.\r\n@@.@@@@.@@\r\n.@@@@@@@.@\r\n.@.@.@.@@@\r\n@.@@@.@@@@\r\n.@@@@@@@@.\r\n@.@.@@@.@."
	expected := 13

	got := Solve(input)

	if expected != got {
		t.Errorf("Solve: expected: %d, got: %d", expected, got)
	}
}
