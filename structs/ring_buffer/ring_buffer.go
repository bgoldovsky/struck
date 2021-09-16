package main

import (
	"errors"
)

/*
	Ring-buffer

	Кольцевой буфер, или циклический буфер (англ. ring-buffer)
	Структура данных, использующая единственный буфер фиксированного размера,
	как будто бы после последнего элемента сразу же снова идет первый.
	Такая структура легко предоставляет возможность буферизации потоков данных.
	==============================
	I. Time: O(1)
	II. Space: O(1)
*/

type ringBuffer struct {
	buf      []interface{}
	readPos  int // Начинается с 0
	writePos int // Начинается с -1
}

func New(size int) *ringBuffer {
	return &ringBuffer{
		buf:      make([]interface{}, size),
		writePos: -1,
	}
}

// Read Читает элемент из буфера
// Извлекая элемент мы не удаляем его, а увеличиваем маркер чтения
// Time O(1)
// Space O(1)
func (r *ringBuffer) Read() (interface{}, error) {
	if r.isEmpty() {
		return 0, errors.New("ring buffer is empty")
	}

	idx := r.nextReadIdx()
	val := r.buf[idx]
	r.readPos++

	return val, nil
}

// Write добавляет элемент в буфер
// Time O(1)
// Space O(1)
func (r *ringBuffer) Write(item interface{}) error {
	if r.isFull() {
		return errors.New("ring buffer is full")
	}

	r.buf[r.nextWriteIdx()] = item
	r.writePos++

	return nil
}

func (r *ringBuffer) isFull() bool {
	unread := (r.writePos - r.readPos) + 1
	return unread == len(r.buf)
}

func (r *ringBuffer) isEmpty() bool {
	return r.writePos < r.readPos
}

func (r *ringBuffer) nextWriteIdx() int {
	// Для записи используем инкремент текущей позиции
	pos := r.writePos
	pos++

	return pos % len(r.buf)
}

func (r *ringBuffer) nextReadIdx() int {
	return r.readPos % len(r.buf)
}
