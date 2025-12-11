package day6

import "testing"

func TestSolve(t *testing.T) {
	input := "123 328  51 64 \r\n 45 64  387 23 \r\n  6 98  215 314\r\n*   +   *   +  "
	expected := 3263827

	got := Solve(input)

	if expected != got {
		t.Errorf("Solve: expected: %d, got: %d", expected, got)
	}
}
