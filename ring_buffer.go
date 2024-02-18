package ring_buffer

import (
	"errors"
)

type RingBuffer[T any] struct {
	buffer   []T
	size     int
	capacity int
	front    int
	tail     int
}

func New[T any]() *RingBuffer[T] {
	return NewWithCapacity[T](32)
}

func NewWithCapacity[T any](capacity int) *RingBuffer[T] {
	return &RingBuffer[T]{
		buffer:   make([]T, capacity, capacity),
		size:     0,
		capacity: capacity,
		front:    0,
		tail:     -1,
	}
}

func (rb *RingBuffer[T]) PushBack(element T) {
	if rb.full() {
		rb.extendCapacity()
	}
	rb.tail++
	rb.size++
	rb.buffer[rb.tail%rb.capacity] = element
}

func (rb *RingBuffer[T]) PopFront() (T, error) {
	if rb.empty() {
		var zero T
		return zero, errors.New("Pop on an empty RingBuffer.")
	}
	value := rb.buffer[rb.front%rb.capacity]
	rb.front++
	rb.size--
	return value, nil
}

func (rb *RingBuffer[T]) PeekFront() (T, error) {
	if rb.empty() {
		var zero T
		return zero, errors.New("PeekFront on an empty RingBuffer.")
	}
	return rb.buffer[rb.front%rb.capacity], nil
}

func (rb *RingBuffer[T]) PeekTail() (T, error) {
	if rb.empty() {
		var zero T
		return zero, errors.New("PeekTail on an empty RingBuffer.")
	}
	return rb.buffer[rb.tail%rb.capacity], nil
}

func (rb *RingBuffer[T]) empty() bool {
	return rb.size == 0
}

func (rb *RingBuffer[T]) full() bool {
	return rb.size == rb.capacity
}

func (rb *RingBuffer[T]) extendCapacity() {
	new_capacity := rb.capacity * 2
	new_buffer := make([]T, new_capacity, new_capacity)
	for offset := range rb.size {
		new_buffer[offset] = rb.buffer[(rb.front+offset)%rb.capacity]
	}
	rb.buffer = new_buffer
	rb.capacity = new_capacity
	rb.front = 0
	rb.tail = rb.size - 1
}
