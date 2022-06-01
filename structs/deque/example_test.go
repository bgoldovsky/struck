package deque

import (
	"fmt"
)

func Example_back() {
	tests := []int64{1, 2}

	d := New(2)

	for _, tt := range tests {
		_ = d.PushBack(tt)
	}

	// Output:
	// 2
	// 1

	act, _ := d.PopBack()
	fmt.Println(act)

	act, _ = d.PopBack()
	fmt.Println(act)
}
