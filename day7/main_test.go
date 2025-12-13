package day7

import "testing"

func TestSolve(t *testing.T) {
	input := ".......S.......\r\n...............\r\n.......^.......\r\n...............\r\n......^.^......\r\n...............\r\n.....^.^.^.....\r\n...............\r\n....^.^...^....\r\n...............\r\n...^.^...^.^...\r\n...............\r\n..^...^.....^..\r\n...............\r\n.^.^.^.^.^...^.\r\n..............."
	expected := 21

	got := Solve(input)

	if expected != got {
		t.Errorf("Solve: expected: %d, got: %d", expected, got)
	}
}
