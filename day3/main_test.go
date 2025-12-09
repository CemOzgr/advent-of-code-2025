package day3

import "testing"

func TestSolve(t *testing.T) {
	tests := []struct {
		bank   string
		expect int
	}{
		{"987654321111111", 987654321111},
		{"811111111111119", 811111111119},
		{"234234234234278", 434234234278},
		{"818181911112111", 888911112111},
	}

	for _, tc := range tests {
		got := Solve(tc.bank)
		if got != tc.expect {
			t.Errorf("Solve(%s): expected :%d got %d", tc.bank, tc.expect, got)
		}
	}
}
