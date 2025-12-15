package day9

import "testing"

func TestSolve(t *testing.T) {
	input := "7,1\r\n11,1\r\n11,7\r\n9,7\r\n9,5\r\n2,5\r\n2,3\r\n7,3"
	expected := 50

	got := Solve(input)

	if expected != got {
		t.Errorf("Solve: expected: %d, got: %d", expected, got)
	}
}
