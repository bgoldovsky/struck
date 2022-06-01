package main

import "testing"

func Test_IsPalindrome(t *testing.T) {
	type args struct {
		head *ListNode
	}

	tests := []struct {
		args args
		exp  bool
	}{
		{
			args: args{
				head: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val: 2,
							Next: &ListNode{
								Val:  1,
								Next: nil,
							},
						},
					},
				},
			},
			exp: true,
		},
		{
			args: args{
				head: nil,
			},
			exp: true,
		},
		{
			args: args{
				head: &ListNode{
					Val:  1,
					Next: nil,
				},
			},
			exp: true,
		},
	}
	for _, tt := range tests {
		if act := IsPalindrome(tt.args.head); act != tt.exp {
			t.Errorf("act %v, exp %v", act, tt.exp)
		}
	}
}
