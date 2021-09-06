package main

import (
	"testing"
)

func TestRingBuffer_Write(t *testing.T) {
	rb := New(5)
	for i := int64(0); i < 5; i++ {
		err := rb.Write(i)
		if err != nil {
			t.Errorf("expected: no errors, act: %v", err)
		}
	}

	for i := int64(0); i < 5; i++ {
		act := rb.buf[i]
		if act != i {
			t.Errorf("expected: %v, act: %v", i, act)
		}
	}
}

func TestRingBuffer_Write_Error(t *testing.T) {
	rb := New(5)
	for i := int64(0); i < 5; i++ {
		_ = rb.Write(i)
	}

	err := rb.Write(10)
	if err == nil {
		t.Error("expected error, act nil")
	}
}

func TestRingBuffer_Read(t *testing.T) {
	rb := New(5)
	for i := int64(0); i < 5; i++ {
		_ = rb.Write(i)
	}

	for i := int64(0); i < 5; i++ {
		act, err := rb.Read()
		if err != nil {
			t.Errorf("expected: no errors, act: %v", err)
		}

		if act != i {
			t.Errorf("expected: %v, act: %v", i, act)
		}
	}
}

func TestRingBuffer_Read_Error(t *testing.T) {
	rb := New(5)

	act, err := rb.Read()
	if err == nil {
		t.Error("expected error, act nil")
	}

	if act != 0 {
		t.Errorf("expected: 0 act: %v", act)
	}
}
