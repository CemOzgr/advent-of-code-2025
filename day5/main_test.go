package day5

import "testing"

func TestSolve(t *testing.T) {
	input := "3-5\r\n10-14\r\n16-20\r\n12-18\r\n\r\n1\r\n5\r\n8\r\n11\r\n17\r\n32"
	expected := 14

	got := Solve(input)

	if expected != got {
		t.Errorf("Solve: excepted: %d, got: %d", expected, got)
	}
}
