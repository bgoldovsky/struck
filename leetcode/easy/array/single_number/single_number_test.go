package single_number

import "testing"

func TestSingleNumberNaive(t *testing.T) {
	tests := []struct {
		nums []int
		exp  int
	}{
		{
			nums: []int{2, 2, 1},
			exp:  1,
		},
		{
			nums: []int{4, 1, 2, 1, 2},
			exp:  4,
		},
		{
			nums: []int{1},
			exp:  1,
		},
	}

	for _, tt := range tests {
		act := singleNumberNaive(tt.nums)
		if tt.exp != act {
			t.Errorf("exp: %v, act: %v", tt.exp, act)
		}
	}
}

func TestSingleNumber(t *testing.T) {
	tests := []struct {
		nums []int
		exp  int
	}{
		{
			nums: []int{2, 2, 1},
			exp:  1,
		},
		{
			nums: []int{4, 1, 2, 1, 2},
			exp:  4,
		},
		{
			nums: []int{1},
			exp:  1,
		},
	}

	for _, tt := range tests {
		act := singleNumber(tt.nums)
		if tt.exp != act {
			t.Errorf("exp: %v, act: %v", tt.exp, act)
		}
	}
}
