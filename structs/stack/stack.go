package main

/*
	Стек

	Работает по принципу LIFO (Last in first out)
	==============================
	I. Time: O(1)
	II. Space: O(1)
*/

type Stack interface {
	Push(item int64)
	Pop() int64
	IsEmpty() bool
}

type stack struct {
	Store []interface{}
}

func New(size int) *stack {
	return &stack{
		Store: make([]interface{}, 0, size),
	}
}

// Push Добавляет элемент в стен
// Time O(1)/O(n) - если количество элементов больше заданного размера стека
// Space O(1)/O(n) - если количество элементов больше заданного размера стека
func (s *stack) Push(item int64) {
	s.Store = append(s.Store, item)
}

// Pop Извлекает элемент из стека
// Time O(1)
// Space O(1)
func (s *stack) Pop() interface{} {
	if s.IsEmpty() {
		return -1
	}

	lastIdx := len(s.Store) - 1

	item := s.Store[lastIdx]
	s.Store = s.Store[:lastIdx]

	return item
}

// IsEmpty Проверяет пуст ли стек
// Time O(1)
// Space O(1)
func (s *stack) IsEmpty() bool {
	return len(s.Store) == 0
}