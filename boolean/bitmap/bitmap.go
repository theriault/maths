// Package bitmap provides access to a fast, low-memory bitmap with protection against invalid offsets.
package bitmap

type Bitmap struct {
	max  int
	bits []byte
}

// New creates a bitmap with len bits.
func New(len int) Bitmap {
	b := Bitmap{}
	b.max = len
	s := len >> 3
	if len&7 > 0 {
		s++
	}
	b.bits = make([]byte, s)
	return b
}

// Valid returns whether or not the offset is valid within the bitmap.
func (b *Bitmap) Valid(offset int) bool {
	return offset >= 0 && offset < b.max
}

// Set the bit at the given offset.
func (b *Bitmap) Set(offset int) bool {
	if !b.Valid(offset) {
		return false
	}
	b.bits[offset>>3] |= byte(1 << (offset & 7))
	return true
}

// Clear the bit at the given offset.
func (b *Bitmap) Clear(offset int) bool {
	if !b.Valid(offset) {
		return false
	}
	b.bits[offset>>3] &= 255 ^ byte(1<<(offset&7))
	return true
}

// Flip/negate the bit at the given offset.
func (b *Bitmap) Flip(offset int) bool {
	if !b.Valid(offset) {
		return false
	}
	b.bits[offset>>3] ^= byte(1 << (offset & 7))
	return true
}

// Test the bit at the given offset. This method will return false if the offset is also invalid for the current
// bitmap. If you are unsure about whether the given offset is valid, call Valid on the offset first.
func (b *Bitmap) Test(offset int) bool {
	if !b.Valid(offset) {
		return false
	}
	return b.bits[offset>>3]&byte(1<<(offset&7)) > 0
}

// Len returns the number of bits available in the bitmap.
func (b *Bitmap) Len() int {
	return b.max
}

// And does a bitwise-and with another bitmap
func (b *Bitmap) And(a Bitmap) {
	m, n := b.max-1, a.max-1
	if m < n {
		n = m
	}
	x := 7 - (n & 7)
	n >>= 3
	for i := 0; i < n; i++ {
		b.bits[i] &= a.bits[i]
	}
	b.bits[n] &= a.bits[n] & (255 >> x)
}

// Or does a bitwise-or with another bitmap
func (b *Bitmap) Or(a Bitmap) {
	m, n := b.max-1, a.max-1
	if m < n {
		n = m
	}
	x := 7 - (n & 7)
	n >>= 3
	for i := 0; i < n; i++ {
		b.bits[i] |= a.bits[i]
	}
	b.bits[n] |= a.bits[n] & (255 >> x)
}

// And does a bitwise-xor with another bitmap
func (b *Bitmap) Xor(a Bitmap) {
	m, n := b.max-1, a.max-1
	if m < n {
		n = m
	}
	x := 7 - (n & 7)
	n >>= 3
	for i := 0; i < n; i++ {
		b.bits[i] ^= a.bits[i]
	}
	b.bits[n] ^= a.bits[n] & (255 >> x)
}

// Flip/negate all bits in the bitmap
func (b *Bitmap) FlipAll() {
	n := b.max - 1
	x := 7 - (n & 7)
	n >>= 3
	for i := 0; i < n; i++ {
		b.bits[i] ^= 255
	}
	b.bits[n] ^= 255 >> x
}

// Clear all bits in the bitmap
func (b *Bitmap) ClearAll() {
	for i, n := 0, len(b.bits); i < n; i++ {
		b.bits[i] = 0
	}
}

// Set all bits in the bitmap
func (b *Bitmap) SetAll() {
	n := b.max - 1
	x := 7 - (n & 7)
	n >>= 3
	for i, n := 0, len(b.bits); i < n; i++ {
		b.bits[i] = 255
	}
	b.bits[n] &= 255 >> x
}
