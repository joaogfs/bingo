package buffer

import (
	"bytes"
	"fmt"
)

// Immutable Buffer structure. Once initialized, buf remains the same
type Buffer struct {
	buf         []byte
	index       int
	length      int
	initialized bool
}

// Initializes the buffer. Can't be changed in any way after initialized
func (b *Buffer) Init(data []byte) error {
	var err error = nil
	if !b.initialized {
		b.buf = data[:]
		b.index = 0
		b.length = len(data)
		b.initialized = true
	} else {
		err = fmt.Errorf("buffer already initialized")
	}

	return err
}

// returns all the data. index is not changed
func (b *Buffer) All() []byte {
	new_b := b.buf[:]

	return new_b
}

// Returns the next n bytes. Sets index accordingly
func (b *Buffer) Next(n int) []byte {
	var new_b []byte

	//pp.Printf("Entering Next: %+v\nArg: %v\n", b, n)

	if b.index+n >= b.length { // make sure it's not out of bounds
		new_b = b.buf[b.index:b.length]
		b.index = b.length
	} else {
		new_b = b.buf[b.index:(b.index + n)]
		b.index += n
	}

	return new_b
}

// Seek n bytes
func (b *Buffer) Seek(n int) {
	if b.index+n >= b.length { // make sure it's not out of bounds
		b.index = b.length
	} else if b.index+n <= 0 {
		b.index = 0
	} else {
		b.index += n
	}
}

// Jump to offset n
func (b *Buffer) Jump(n int) error {
	if n < 0 || n >= b.length {
		return fmt.Errorf("offset %v is out of bounds", n)
	}

	b.index = n
	return nil
}

func (b *Buffer) Position() int {
	return b.index
}

// Read until delim. Returns either the entire rest of the buffer or an empty slice if delim is not present
func (b *Buffer) ReadUntil(delim []byte, returnEmpty bool) []byte {
	delim_index := bytes.Index(b.buf, delim)
	var new_b []byte
	if delim_index != -1 {
		new_b = b.buf[:delim_index]
	} else if !returnEmpty {
		new_b = b.buf[:]
	}
	return new_b
}
