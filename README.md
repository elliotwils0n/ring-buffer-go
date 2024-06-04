# ring-buffer-go

## Description
Ring Buffer implementation based on Go's slices.

## Import module
```shell
go get -u github.com/elliotwils0n/ring-buffer-go
```

## Usage

#### Import package
```go
import (
    "github.com/elliotwils0n/ring-buffer-go"
)
```

#### Init Ring Buffer
With or without initial capacity (defaults to 32)
```go
rb := ringbuffer.New[int]()
```
```go
rb := ringbuffer.NewWithCapacity[int](10);
```

#### Push, pop and peek elements with Ring Buffer
Push back
```go
rb.PushBack(123)
```
Push front
```go
rb.PushFront(123)
```

Pop front, error returned on empty Ring Buffer
```go
element, err := rb.PopFront()
```
Pop back, error returned on empty Ring Buffer
```go
element, err := rb.PopBack()
```

Peek front/back, error returned on empty Ring Buffer
```go
front_element, err := rb.PeekFront()
tail_element, err := rb.PeekTail()
```
