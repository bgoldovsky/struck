package deque

import (
	"testing"
)

func TestDeque_Back_Success(t *testing.T) {
	tests := []struct {
		push int64
		pop  int64
	}{
		{
			push: 1,
			pop:  5,
		},
		{
			push: 2,
			pop:  4,
		},
		{
			push: 3,
			pop:  3,
		},
		{
			push: 4,
			pop:  2,
		},
		{
			push: 5,
			pop:  1,
		},
	}

	d := New(5)

	for _, tt := range tests {
		_ = d.PushBack(tt.push)
	}

	for _, tt := range tests {
		act, _ := d.PopBack()
		if act != tt.pop {
			t.Errorf("expected: %v, act: %v", tt.pop, act)
		}
	}
}

func TestDeque_Back_Error(t *testing.T) {
	tests := []int64{100, 200}

	d := New(2)

	for _, tt := range tests {
		_ = d.PushBack(tt)
	}

	err := d.PushBack(300)
	if err == nil {
		t.Error("push back expected error, got nil")
	}

	for range tests {
		_, _ = d.PopBack()
	}

	_, err = d.PopBack()
	if err == nil {
		t.Error("exp back expected error, got nil")
	}
}

func TestDeque_Front_Success(t *testing.T) {
	tests := []struct {
		push int64
		pop  int64
	}{
		{
			push: 1,
			pop:  5,
		},
		{
			push: 2,
			pop:  4,
		},
		{
			push: 3,
			pop:  3,
		},
		{
			push: 4,
			pop:  2,
		},
		{
			push: 5,
			pop:  1,
		},
	}

	d := New(5)

	for _, tt := range tests {
		_ = d.PushFront(tt.push)
	}

	for _, tt := range tests {
		act, _ := d.PopFront()
		if act != tt.pop {
			t.Errorf("expected: %v, act: %v", tt.pop, act)
		}
	}
}

func TestDeque_Front_Error(t *testing.T) {
	tests := []int64{100, 200}

	d := New(2)

	for _, tt := range tests {
		_ = d.PushFront(tt)
	}

	err := d.PushFront(300)
	if err == nil {
		t.Error("push front expected error, got nil")
	}

	for range tests {
		_, _ = d.PopFront()
	}

	_, err = d.PopFront()
	if err == nil {
		t.Error("exp front expected error, got nil")
	}
}

/*
	push_front -479
	pop_back
	push_back -119
	pop_front
	pop_front
	push_back 906
*/
func TestDeque_Success(t *testing.T) {
	d := New(3)

	err := d.PushFront(-479)
	if err != nil {
		t.Errorf("expected: nil, act: %v", err)
	}

	act, err := d.PopBack()
	if err != nil {
		t.Errorf("expected: nil, act: %v", err)
	}
	if act != -479 {
		t.Errorf("expected: %v, act: %v", -419, act)
	}

	err = d.PushBack(-119)
	if err != nil {
		t.Errorf("expected: nil, act: %v", err)
	}

	act, err = d.PopFront()
	if err != nil {
		t.Errorf("expected: nil, act: %v", err)
	}
	if act != -119 {
		t.Errorf("expected: %v, act: %v", -419, act)
	}

	act, err = d.PopFront()
	if err == nil {
		t.Errorf("expected: err, act: %v", err)
	}
	if act != 0 {
		t.Errorf("expected: %v, act: %v", 0, act)
	}

	err = d.PushBack(906)
	if err != nil {
		t.Errorf("expected: nil, act: %v", err)
	}
}

func TestDeque_Mix_Success(t *testing.T) {
	tests := []struct {
		push int64
		exp  int64
	}{
		{
			push: 1,
			exp:  5,
		},
		{
			push: 2,
			exp:  4,
		},
		{
			push: 3,
			exp:  3,
		},
		{
			push: 4,
			exp:  2,
		},
		{
			push: 5,
			exp:  1,
		},
	}

	d := New(10)

	for _, tt := range tests {
		_ = d.PushFront(tt.push)
	}

	for _, tt := range tests {
		_ = d.PushBack(tt.push * 10)
	}

	for _, tt := range tests {
		act, _ := d.PopFront()
		if act != tt.exp {
			t.Errorf("expected: %v, act: %v", tt.exp, act)
		}
	}

	for _, tt := range tests {
		act, _ := d.PopBack()
		if act != tt.exp*10 {
			t.Errorf("expected: %v, act: %v", tt.exp*10, act)
		}
	}
}
