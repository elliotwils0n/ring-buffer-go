package ringbuffer

import (
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	// given - when
	rb := New[int]()

	// then
	if rb.size != 0 {
		t.Fatalf(`should initialize with size: 0, has: %d`, rb.size)
	}
	if rb.capacity != 32 {
		t.Fatalf(`should initialize with capacity: 32, has: %d`, rb.capacity)
	}
	if rb.head != 0 {
		t.Fatalf(`should initialize with head: 0, has: %d`, rb.head)
	}
	if rb.tail != 0 {
		t.Fatalf(`should initialize with tail: 0, has: %d`, rb.tail)
	}
	if !reflect.DeepEqual(rb.buffer, make([]int, 32, 32)) {
		t.Fatalf(`should initialize with buffer of size 32 with zero values`)
	}
}

func TestNewWithCapacity(t *testing.T) {
	// given - when
	rb := NewWithCapacity[int](128)

	// then
	if rb.size != 0 {
		t.Fatalf(`should initialize with size: 0, has: %d`, rb.size)
	}
	if rb.capacity != 128 {
		t.Fatalf(`should initialize with capacity: 128, has: %d`, rb.capacity)
	}
	if rb.head != 0 {
		t.Fatalf(`should initialize with head: 0, has: %d`, rb.head)
	}
	if rb.tail != 0 {
		t.Fatalf(`should initialize with tail: 0, has: %d`, rb.tail)
	}
	if !reflect.DeepEqual(rb.buffer, make([]int, 128, 128)) {
		t.Fatalf(`should initialize with buffer of size 128 with zero values`)
	}
}

func TestPushBack(t *testing.T) {
	// given
	rb := NewWithCapacity[int](2)
	expectedBuffer := []int{1, 2}

	// when
	rb.PushBack(1)
	rb.PushBack(2)

	// then
	if rb.size != 2 {
		t.Fatalf(`should have size: 2, has: %d`, rb.size)
	}
	if rb.capacity != 2 {
		t.Fatalf(`should have capacity: 2, has: %d`, rb.capacity)
	}
	if rb.head != 0 {
		t.Fatalf(`should have head: 0, has: %d`, rb.head)
	}
	if rb.tail != 1 {
		t.Fatalf(`should have tail: 1, has: %d`, rb.tail)
	}
	if !reflect.DeepEqual(rb.buffer, expectedBuffer) {
		t.Fatalf(`should have buffer: %v, has: %v`, expectedBuffer, rb.buffer)
	}
}

func TestPushFront(t *testing.T) {
	// given
	rb := NewWithCapacity[int](2)
	expectedBuffer := []int{1, 2}

	// when
	rb.PushFront(1)
	rb.PushFront(2)

	// then
	if rb.size != 2 {
		t.Fatalf(`should have size: 2, has: %d`, rb.size)
	}
	if rb.capacity != 2 {
		t.Fatalf(`should have capacity: 2, has: %d`, rb.capacity)
	}
	if rb.head != 1 {
		t.Fatalf(`should have head: 1, has: %d`, rb.head)
	}
	if rb.tail != 0 {
		t.Fatalf(`should have tail: 0, has: %d`, rb.tail)
	}
	if !reflect.DeepEqual(rb.buffer, expectedBuffer) {
		t.Fatalf(`When PushFront() called then should have buffer: %v, has: %v`, expectedBuffer, rb.buffer)
	}
}

func TestPopFront(t *testing.T) {
	// given
	rb := NewWithCapacity[int](2)
	expectedBuffer := []int{1, 2}

	// when
	rb.PushBack(1)
	rb.PushBack(2)
	pop1, err1 := rb.PopFront()
	pop2, err2 := rb.PopFront()
	pop3, err3 := rb.PopFront()

	// then
	if pop1 != 1 || err1 != nil {
		t.Fatalf(`pop1 should be 1, is: %d`, pop1)
		t.Fatalf(`err1 should be nil, is: %v`, err1)
	}
	if pop2 != 2 || err2 != nil {
		t.Fatalf(`pop2 should be 2, is: %d`, pop2)
		t.Fatalf(`err2 should be nil, is: %v`, err2)
	}
	if pop3 != 0 || err3 == nil {
		t.Fatalf(`pop3 should be zero value, is: %d`, pop3)
		t.Fatalf(`err3 should not be non-nil, is: %v`, err3)
	}

	if rb.size != 0 {
		t.Fatalf(`should have size: 0, has: %d`, rb.size)
	}
	if rb.capacity != 2 {
		t.Fatalf(`should have extended capacity: 2, has: %d`, rb.capacity)
	}
	if rb.head != 0 {
		t.Fatalf(`should have head: 0, has: %d`, rb.head)
	}
	// next push would correct tail index
	if rb.tail != 1 {
		t.Fatalf(`should have tail: 1, has: %d`, rb.tail)
	}
	if !reflect.DeepEqual(rb.buffer, expectedBuffer) {
		t.Fatalf(`should have buffer: %v, has: %v`, expectedBuffer, rb.buffer)
	}
}

func TestPopBack(t *testing.T) {
	// given
	rb := NewWithCapacity[int](2)
	expectedBuffer := []int{1, 2}

	// when
	rb.PushFront(1)
	rb.PushFront(2)
	pop1, err1 := rb.PopBack()
	pop2, err2 := rb.PopBack()
	pop3, err3 := rb.PopBack()

	// then
	if pop1 != 1 || err1 != nil {
		t.Fatalf(`pop1 should be 1, is: %d`, pop1)
		t.Fatalf(`err1 should be nil, is: %v`, err1)
	}
	if pop2 != 2 || err2 != nil {
		t.Fatalf(`pop1 should be 2, is: %d`, pop1)
		t.Fatalf(`err1 should be nil, is: %v`, err1)
	}
	if pop3 != 0 || err3 == nil {
		t.Fatalf(`pop3 should be zero value, is: %d`, pop3)
		t.Fatalf(`err1 should be non-nil, is: %v`, err1)
	}

	if rb.size != 0 {
		t.Fatalf(`should have size: 0, has: %d`, rb.size)
	}
	if rb.capacity != 2 {
		t.Fatalf(`should have extended capacity: 6, has: %d`, rb.capacity)
	}
	// next push would correct head index
	if rb.head != 1 {
		t.Fatalf(`should have head: 1, has: %d`, rb.head)
	}
	if rb.tail != 0 {
		t.Fatalf(`should have tail: 0, has: %d`, rb.tail)
	}
	if !reflect.DeepEqual(rb.buffer, expectedBuffer) {
		t.Fatalf(`should have buffer: %v, has: %v`, expectedBuffer, rb.buffer)
	}
}

func TestRollRight(t *testing.T) {
	// given
	rb := NewWithCapacity[int](3)
	expectedBuffer := []int{3, 1, 2}

	// when
	rb.PushBack(0)
	rb.PushBack(1)
	_, _ = rb.PopFront()
	rb.PushBack(2)
	rb.PushBack(3)

	// then
	if rb.size != 3 {
		t.Fatalf(`should have size: 3, has: %d`, rb.size)
	}
	if rb.capacity != 3 {
		t.Fatalf(`should have capacity: 3, has: %d`, rb.capacity)
	}
	if rb.head != 1 {
		t.Fatalf(`should have head: 1, has: %d`, rb.head)
	}
	if rb.tail != 0 {
		t.Fatalf(`should have tail: 0, has: %d`, rb.tail)
	}
	if !reflect.DeepEqual(rb.buffer, expectedBuffer) {
		t.Fatalf(`should have buffer: %v, has: %v`, expectedBuffer, rb.buffer)
	}
}

func TestRollLeft(t *testing.T) {
	// given
	rb := NewWithCapacity[int](3)
	expectedBuffer := []int{1, 3, 2}

	// when
	rb.PushBack(1)
	rb.PushFront(2)
	rb.PushFront(3)

	// then
	if rb.size != 3 {
		t.Fatalf(`should have size: 3, has: %d`, rb.size)
	}
	if rb.capacity != 3 {
		t.Fatalf(`should have capacity: 3, has: %d`, rb.capacity)
	}
	if rb.head != 1 {
		t.Fatalf(`should have head: 1, has: %d`, rb.head)
	}
	if rb.tail != 0 {
		t.Fatalf(`should have tail: 0, has: %d`, rb.tail)
	}
	if !reflect.DeepEqual(rb.buffer, expectedBuffer) {
		t.Fatalf(`should have buffer: %v, has: %v`, expectedBuffer, rb.buffer)
	}
}

func TestExtendingCapacity(t *testing.T) {
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
		t.Fatalf(`should have size: 4, has: %d`, rb.size)
	}
	if rb.capacity != 6 {
		t.Fatalf(`should have extended capacity: 6, has: %d`, rb.capacity)
	}
	if rb.head != 0 {
		t.Fatalf(`should have head: 0, has: %d`, rb.head)
	}
	if rb.tail != 3 {
		t.Fatalf(`should have tail: 3, has: %d`, rb.tail)
	}
	if !reflect.DeepEqual(rb.buffer, expectedBuffer) {
		t.Fatalf(`should have buffer: %v, has: %v`, expectedBuffer, rb.buffer)
	}
}

func TestExtendingCapacityOfRightRolled(t *testing.T) {
	// given
	rb := NewWithCapacity[int](2)
	expectedBuffer := []int{1, 2, 3, 0}

	// when
	rb.PushBack(0)
	rb.PushBack(1)
	_, _ = rb.PopFront()
	rb.PushBack(2)
	rb.PushBack(3)

	// then
	if rb.size != 3 {
		t.Fatalf(`should have size: 3, has: %d`, rb.size)
	}
	if rb.capacity != 4 {
		t.Fatalf(`should have extended capacity: 4, has: %d`, rb.capacity)
	}
	if rb.head != 0 {
		t.Fatalf(`should have head: 0, has: %d`, rb.head)
	}
	if rb.tail != 2 {
		t.Fatalf(`should have tail: 2, has: %d`, rb.tail)
	}
	if !reflect.DeepEqual(rb.buffer, expectedBuffer) {
		t.Fatalf(`should have buffer: %v, has: %v`, expectedBuffer, rb.buffer)
	}
}

func TestExtendingCapacityOfLeftRolled(t *testing.T) {
	// given
	rb := NewWithCapacity[int](2)
	expectedBuffer := []int{1, 2, 0, 3}

	// when
	rb.PushBack(1)
	rb.PushBack(2)
	rb.PushFront(3)

	// then
	if rb.size != 3 {
		t.Fatalf(`should have size: 3, has: %d`, rb.size)
	}
	if rb.capacity != 4 {
		t.Fatalf(`should have extended capacity: 4, has: %d`, rb.capacity)
	}
	if rb.head != 3 {
		t.Fatalf(`should have head: 3, has: %d`, rb.head)
	}
	if rb.tail != 1 {
		t.Fatalf(`should have tail: 1, has: %d`, rb.tail)
	}
	if !reflect.DeepEqual(rb.buffer, expectedBuffer) {
		t.Fatalf(`should have buffer: %v, has: %v`, expectedBuffer, rb.buffer)
	}
}

func TestPeekFront(t *testing.T) {
	// given
	rb := NewWithCapacity[int](2)

	// when
	pf1, err1 := rb.PeekFront()
	rb.PushBack(12)
	pf2, err2 := rb.PeekFront()

	// then
	if pf1 != 0 || err1 == nil {
		t.Fatalf(`should be zero value when empty RingBuffer, is: %d`, pf1)
		t.Fatalf(`should return error when empty RingBuffer, but it didn't`)
	}
	if pf2 != 12 || err2 != nil {
		t.Fatalf(`should return 12 but returned: %d`, pf2)
		t.Fatalf(`should not return error when non empty RingBuffer`)
	}
}

func TestPeekBack(t *testing.T) {
	// given
	rb := NewWithCapacity[int](2)

	// when
	pf1, err1 := rb.PeekBack()
	rb.PushBack(11)
	pf2, err2 := rb.PeekBack()

	// then
	if pf1 != 0 || err1 == nil {
		t.Fatalf(`should be zero value when empty RingBuffer, is: %d`, pf1)
		t.Fatalf(`should return error when empty RingBuffer, but it didn't`)
	}
	if pf2 != 11 || err2 != nil {
		t.Fatalf(`should return 11 but returned: %d`, pf2)
		t.Fatalf(`should not return error when non empty RingBuffer`)
	}
}
