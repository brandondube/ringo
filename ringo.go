// Package ringo provides fast ring buffers that support getting continuous slices, even when they roll over the end of the underlying slice.
package ringo

// CircleF64 is a ring buffer of f64 values.  It is not concurrent safe.
type CircleF64 struct {
	buf    []float64
	cursor int
	filled bool
}

// Append adds a value to the buffer
func (c *CircleF64) Append(f float64) {
	c.buf[c.cursor] = f
	if c.cursor == cap(c.buf) {
		c.cursor = 0
		c.filled = true
	} else {
		c.cursor++
	}
}

// Head gets the most recent addition
func (c *CircleF64) Head() float64 {
	return c.buf[c.cursor]
}

// Tail gets the least recent addition.  It returns zero if the buffer is empty
func (c *CircleF64) Tail() float64 {
	if c.cursor == 0 {
		return 0
	}
	return c.buf[c.cursor-1]
}

// Contiguous gets a slice of the values in the buffer from least to most recent
func (c *CircleF64) Contiguous() []float64 {
	if c.filled {
		chunk1 := c.buf[c.cursor:]
		chunk2 := c.buf[:c.cursor]
		out := append(chunk1, chunk2...)
		return out
	}
	return c.buf[:c.cursor]
}

// Initialize creates a new slice of zeros and resets the internal state of the buffer.
// It may be called multiple times.
func (c *CircleF64) Initialize(size int) {
	c.buf = make([]float64, size)
	c.filled = false
	c.cursor = 0
}
