package ring_buffer

import (
	"errors"
)

type RingBuffer struct {
	buffer   []any
	size     int
	capacity int
	front    int
	tail     int
}

func New() *RingBuffer {
	return NewWithCapacity(32)
}

func NewWithCapacity(capacity int) *RingBuffer {
	return &RingBuffer{
		buffer:   make([]any, capacity, capacity),
		size:     0,
		capacity: capacity,
		front:    0,
		tail:     -1,
	}
}

func (rb *RingBuffer) PushBack(element any) {
	if rb.full() {
		rb.extendCapacity()
	}
	rb.tail++
	rb.size++
	rb.buffer[rb.tail%rb.capacity] = element
}

func (rb *RingBuffer) PopFront() (any, error) {
	if rb.empty() {
		return nil, errors.New("Pop on an empty RingBuffer.")
	}
	value := rb.buffer[rb.front%rb.capacity]
	rb.front++
	rb.size--
	return value, nil
}

func (rb *RingBuffer) PeekFront() (any, error) {
	if rb.empty() {
		return nil, errors.New("PeekFront on an empty RingBuffer.")
	}
	return rb.buffer[rb.front%rb.capacity], nil
}

func (rb *RingBuffer) PeekTail() (any, error) {
	if rb.empty() {
		return nil, errors.New("PeekTail on an empty RingBuffer.")
	}
	return rb.buffer[rb.tail%rb.capacity], nil
}

func (rb *RingBuffer) empty() bool {
	return rb.size == 0
}

func (rb *RingBuffer) full() bool {
	return rb.size == rb.capacity
}

func (rb *RingBuffer) extendCapacity() {
	new_capacity := rb.capacity * 2
	new_buffer := make([]any, new_capacity, new_capacity)
	for offset := range rb.size {
		new_buffer[offset] = rb.buffer[(rb.front+offset)%rb.capacity]
	}
	rb.buffer = new_buffer
	rb.capacity = new_capacity
	rb.front = 0
	rb.tail = rb.size - 1
}
