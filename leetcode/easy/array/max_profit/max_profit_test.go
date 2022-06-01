package max_profit

import (
	"testing"
)

func TestMaxProfit(t *testing.T) {
	tests := []struct {
		nums []int
		exp  int
	}{
		{
			nums: []int{7, 1, 5, 3, 6, 4},
			exp:  7,
		},
		{
			nums: []int{1, 2, 3, 4, 5},
			exp:  4,
		},
		{
			nums: []int{7, 6, 4, 3, 1},
			exp:  0,
		},
	}

	for _, tt := range tests {
		act := maxProfit(tt.nums)
		if act != tt.exp {
			t.Errorf("exp: %v, act: %v", tt.exp, act)
		}
	}
}
