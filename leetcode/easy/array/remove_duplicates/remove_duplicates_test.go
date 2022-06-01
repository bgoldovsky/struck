package remove_dublicates

import (
	"reflect"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	tests := []struct {
		nums    []int
		expNums []int
		exp     int
	}{
		{
			nums:    []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			expNums: []int{0, 1, 2, 3, 4},
			exp:     5,
		},
		{
			nums:    []int{1, 1, 1},
			expNums: []int{1},
			exp:     1,
		},
		{
			nums:    []int{1, 2, 3},
			expNums: []int{1, 2, 3},
			exp:     3,
		},
		{
			nums:    nil,
			expNums: nil,
			exp:     0,
		},
	}

	for _, tt := range tests {
		act := removeDuplicates(tt.nums)
		if act != tt.exp {
			t.Errorf("exp: %v, act: %v", tt.exp, act)
		}

		if !reflect.DeepEqual(tt.expNums, tt.nums[:act]) {
			t.Errorf("exp: %v, act: %v", tt.expNums, tt.nums)
		}
	}
}
