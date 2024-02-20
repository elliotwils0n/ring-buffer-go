package ring_buffer

import (
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	// given - when
	rb := New[int]()

	// then
	if rb.size != 0 {
		t.Fatalf(`New() should initialize with size: 0, has: %d`, rb.size)
	}
	if rb.capacity != 32 {
		t.Fatalf(`New() should initialize with capacity: 32, has: %d`, rb.capacity)
	}
	if rb.front != 0 {
		t.Fatalf(`New() should initialize with front: 0, has: %d`, rb.front)
	}
	if rb.tail != -1 {
		t.Fatalf(`New() should initialize with front: 0, has: %d`, rb.tail)
	}
	if !reflect.DeepEqual(rb.buffer, make([]int, 32, 32)) {
		t.Fatalf(`New() should initialize with buffer of size 32 with nil values`)
	}
}

func TestNewWithCapacity(t *testing.T) {
	// given - when
	rb := NewWithCapacity[int](128)

	// then
	if rb.size != 0 {
		t.Fatalf(`NewWithCapacity(128) should initialize with size: 0, has: %d`, rb.size)
	}
	if rb.capacity != 128 {
		t.Fatalf(`NewWithCapacity(128) should initialize with capacity: 128, has: %d`, rb.capacity)
	}
	if rb.front != 0 {
		t.Fatalf(`NewWithCapacity(128) should initialize with front: 0, has: %d`, rb.front)
	}
	if rb.tail != -1 {
		t.Fatalf(`NewWithCapacity(128) should initialize with front: 0, has: %d`, rb.tail)
	}
	if !reflect.DeepEqual(rb.buffer, make([]int, 128, 128)) {
		t.Fatalf(`NewWithCapacity(128) should initialize with buffer of size 128 with nil values`)
	}
}

func TestPushBack(t *testing.T) {
	// given
	rb := NewWithCapacity[int](3)
	expectedBuffer := []int{8008135, 0, 0}

	// when
	rb.PushBack(8008135)

	// then
	if rb.size != 1 {
		t.Fatalf(`When PushBack() called then should have size: 1, has: %d`, rb.size)
	}
	if rb.capacity != 3 {
		t.Fatalf(`When PushBack() called then should have original capacity: 3, has: %d`, rb.capacity)
	}
	if rb.front != 0 {
		t.Fatalf(`When PushBack() called then should have front: 0, has: %d`, rb.front)
	}
	if rb.tail != 0 {
		t.Fatalf(`When PushBack() called then should have tail: 0, has: %d`, rb.tail)
	}
	if !reflect.DeepEqual(rb.buffer, expectedBuffer) {
		t.Fatalf(`When PushBack() called then should have buffer: %v, has: %v`, expectedBuffer, rb.buffer)
	}
}

func TestPushBackExtendingCapacity(t *testing.T) {
	// given
	rb := NewWithCapacity[int](3)
	expectedBuffer := []int{1, 2, 3, 4, 0, 0}

	// when
	rb.PushBack(1)
	rb.PushBack(2)
	rb.PushBack(3)
	rb.PushBack(4)

	// then
	if rb.size != 4 {
		t.Fatalf(`When 4 * PushBack() called then should have size: 4, has: %d`, rb.size)
	}
	if rb.capacity != 6 {
		t.Fatalf(`When 4 * PushBack() called then should have extended capacity: 6, has: %d`, rb.capacity)
	}
	if rb.front != 0 {
		t.Fatalf(`When 4 * PushBack() called then should have front: 0, has: %d`, rb.front)
	}
	if rb.tail != 3 {
		t.Fatalf(`When 4 * PushBack() called then should have tail: 3, has: %d`, rb.tail)
	}
	if !reflect.DeepEqual(rb.buffer, expectedBuffer) {
		t.Fatalf(`When 4 * PushBack() called then should have buffer: %v, has: %v`, expectedBuffer, rb.buffer)
	}
}

func TestPopFront(t *testing.T) {
	// given
	rb := NewWithCapacity[int](6)
	expectedBuffer := []int{1, 2, 3, 4, 0, 0}

	// when
	rb.PushBack(1)
	rb.PushBack(2)
	rb.PushBack(3)
	rb.PushBack(4)
	pop1, err1 := rb.PopFront()
	pop2, err2 := rb.PopFront()
	pop3, err3 := rb.PopFront()
	pop4, err4 := rb.PopFront()
	pop5, err5 := rb.PopFront()

	// then
	if pop1 != 1 {
		t.Fatalf(`When (1) PopFront() called then should return 1, has: %d`, pop1)
	}
	if err1 != nil {
		t.Fatalf(`When (1) PopFront() called then no error expected, has: %v`, err1)
	}
	if pop2 != 2 {
		t.Fatalf(`When (2) PopFront() called then should return 2, has: %d`, pop2)
	}
	if err2 != nil {
		t.Fatalf(`When (2) PopFront() called then no error expected, has: %v`, err2)
	}
	if pop3 != 3 {
		t.Fatalf(`When (3) PopFront() called then should return 3, has: %d`, pop3)
	}
	if err3 != nil {
		t.Fatalf(`When (3) PopFront() called then no error expected, has: %v`, err3)
	}
	if pop4 != 4 {
		t.Fatalf(`When (4) PopFront() called then should return 4, has: %d`, pop4)
	}
	if err4 != nil {
		t.Fatalf(`When (4) PopFront() called then no error expected, has: %v`, err4)
	}
	if pop5 != 0 {
		t.Fatalf(`When (5) PopFront() called then should return zero value, has: %d`, pop2)
	}
	if err5 == nil {
		t.Fatalf(`When (5) PopFront() called then error expected`)
	}

	if rb.size != 0 {
		t.Fatalf(`When few pushed and all popped then should have size: 0, has: %d`, rb.size)
	}
	if rb.capacity != 6 {
		t.Fatalf(`When few pushed and all popped then should have extended capacity: 6, has: %d`, rb.capacity)
	}
	// it is expected to front be at index 4 and tail at index 3
	// another PushBack would increment tail, which would make front and tail point to the same index
	if rb.front != 4 {
		t.Fatalf(`When new pushed and all popped then should have front: 0, has: %d`, rb.front)
	}
	if rb.tail != 3 {
		t.Fatalf(`When few pushed and all popped then should have tail: 3, has: %d`, rb.tail)
	}
	if !reflect.DeepEqual(rb.buffer, expectedBuffer) {
		t.Fatalf(`When few pushed and all popped then should have buffer: %v, has: %v`, expectedBuffer, rb.buffer)
	}
}

func TestRingBehaviour(t *testing.T) {
	// given
	rb := NewWithCapacity[int](6)
	expectedBuffer := []int{3, 4, 5, 6, 1, 2}

	// when
	rb.PushBack(1)
	rb.PushBack(2)
	rb.PushBack(3)
	rb.PushBack(4)
	_, _ = rb.PopFront()
	_, _ = rb.PopFront()
	_, _ = rb.PopFront()
	_, _ = rb.PopFront()
	rb.PushBack(1)
	rb.PushBack(2)
	rb.PushBack(3)
	rb.PushBack(4)
	rb.PushBack(5)
	rb.PushBack(6)

	// then
	if rb.size != 6 {
		t.Fatalf(`Should have size: 6, has: %d`, rb.size)
	}
	if rb.capacity != 6 {
		t.Fatalf(`Should have extended capacity: 6, has: %d`, rb.capacity)
	}
	if rb.front != 4 {
		t.Fatalf(`Should have front: 4, has: %d`, rb.front)
	}
	if rb.tail != 9 {
		t.Fatalf(`Should have tail: 9, has: %d`, rb.tail)
	}
	if !reflect.DeepEqual(rb.buffer, expectedBuffer) {
		t.Fatalf(`Should have buffer: %v, has: %v`, expectedBuffer, rb.buffer)
	}
}

func TestIndexesReset(t *testing.T) {
	// given
	rb := NewWithCapacity[int](3)
	expectedBuffer := []int{4, 5, 3}

	// when
	rb.PushBack(1)
	rb.PushBack(2)
	rb.PushBack(3)
	_, _ = rb.PopFront()
	_, _ = rb.PopFront()
	rb.PushBack(4)
	rb.PushBack(5)
	_, _ = rb.PopFront()

	// then
	if rb.size != 2 {
		t.Fatalf(`Should have size: 2, has: %d`, rb.size)
	}
	if rb.capacity != 3 {
		t.Fatalf(`Should have original capacity: 3, has: %d`, rb.capacity)
	}
	if rb.front != 0 {
		t.Fatalf(`Should have front: 0, has: %d`, rb.front)
	}
	if rb.tail != 1 {
		t.Fatalf(`Should have tail: 1, has: %d`, rb.tail)
	}
	if !reflect.DeepEqual(rb.buffer, expectedBuffer) {
		t.Fatalf(`Should have buffer: %v, has: %v`, expectedBuffer, rb.buffer)
	}
}

func TestPeekFront(t *testing.T) {
	// given
	rb := NewWithCapacity[int](13)

	// when
	pf1, err1 := rb.PeekFront()
	rb.PushBack(12)
	pf2, err2 := rb.PeekFront()
	rb.PushBack(21)
	pf3, err3 := rb.PeekFront()
	rb.PopFront()
	pf4, err4 := rb.PeekFront()
	rb.PopFront()
	pf5, err5 := rb.PeekFront()

	// then
	if pf1 != 0 {
		t.Fatalf(`(1) PeekFront() should return zero value when empty RingBuffer, has: %d`, pf1)
	}
	if err1 == nil {
		t.Fatalf(`(1) PeekFront() should return error when empty RingBuffer, but it didn't`)
	}
	if pf2 != 12 {
		t.Fatalf(`(2) PeekFront() should return 12 but returned: %d`, pf2)
	}
	if err2 != nil {
		t.Fatalf(`(2) PeekFront() should not return error when non empty RingBuffer`)
	}
	if pf3 != 12 {
		t.Fatalf(`(3) PeekFront() should return 12 but returned: %d`, pf3)
	}
	if err3 != nil {
		t.Fatalf(`(3) PeekFront() should not return error when non empty RingBuffer`)
	}
	if pf4 != 21 {
		t.Fatalf(`(4) PeekFront() should return 21 but returned: %d`, pf4)
	}
	if err4 != nil {
		t.Fatalf(`(4) PeekFront() should not return error when non empty RingBuffer`)
	}
	if pf5 != 0 {
		t.Fatalf(`(5) PeekFront() should return zero value when empty RingBuffer, has: %d`, pf5)
	}
	if err5 == nil {
		t.Fatalf(`(5) PeekFront() should return error when empty RingBuffer, but it didn't`)
	}
}

func TestPeekTail(t *testing.T) {
	// given
	rb := NewWithCapacity[int](5)

	// when
	pf1, err1 := rb.PeekTail()
	rb.PushBack(11)
	pf2, err2 := rb.PeekTail()
	rb.PushBack(22)
	pf3, err3 := rb.PeekTail()
	rb.PushBack(33)
	pf4, err4 := rb.PeekTail()
	rb.PushBack(44)
	pf5, err5 := rb.PeekTail()

	// then
	if pf1 != 0 {
		t.Fatalf(`(1) PeekTail() should return zero value when empty RingBuffer, has: %d`, pf1)
	}
	if err1 == nil {
		t.Fatalf(`(1) PeekTail() should return error when empty RingBuffer, but it didn't`)
	}
	if pf2 != 11 {
		t.Fatalf(`(2) PeekTail() should return 11 but returned: %d`, pf2)
	}
	if err2 != nil {
		t.Fatalf(`(2) PeekTail() should not return error when non empty RingBuffer`)
	}
	if pf3 != 22 {
		t.Fatalf(`(3) PeekTail() should return 22 but returned: %d`, pf3)
	}
	if err3 != nil {
		t.Fatalf(`(3) PeekTail() should not return error when non empty RingBuffer`)
	}
	if pf4 != 33 {
		t.Fatalf(`(4) PeekTail() should return 33 but returned: %d`, pf4)
	}
	if err4 != nil {
		t.Fatalf(`(4) PeekTaiil() should not return error when non empty RingBuffer`)
	}
	if pf5 != 44 {
		t.Fatalf(`(5) PeekTail() should return 44 when empty RingBuffer, has: %d`, pf5)
	}
	if err5 != nil {
		t.Fatalf(`(5) PeekTaiil() should not return error when non empty RingBuffer`)
	}
}
