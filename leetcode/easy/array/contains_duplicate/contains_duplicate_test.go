package contains_duplicate

import "testing"

func TestContainsDuplicate(t *testing.T) {
	tests := []struct {
		nums []int
		exp  bool
	}{
		{
			nums: []int{1, 2, 3, 1},
			exp:  true,
		},
		{
			nums: []int{1, 2, 3, 4},
			exp:  false,
		},
		{
			nums: []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2},
			exp:  true,
		},
	}

	for _, tt := range tests {
		act := containsDuplicate(tt.nums)
		if tt.exp != act {
			t.Errorf("exp: %v, act: %v", tt.exp, act)
		}
	}
}
