// Package bitmap provides access to a fast, low-memory bitmap with protection against invalid offsets.
package bitmap

// constants used by various funcs
const shift int = 6
const ss int = 0b111111
const mask uint64 = 0xFFFFFFFFFFFFFFFF
const m1 uint64 = 0x5555555555555555  //binary: 0101...
const m2 uint64 = 0x3333333333333333  //binary: 00110011..
const m4 uint64 = 0x0f0f0f0f0f0f0f0f  //binary:  4 zeros,  4 ones ...
const h01 uint64 = 0x0101010101010101 //the sum of 256 to the power of 0,1,2,3..

type Bitmap struct {
	max  int
	bits []uint64
}

// New creates a bitmap with len bits.
func New(len int) Bitmap {
	b := Bitmap{}
	b.max = len
	s := len >> shift
	if len&ss > 0 {
		s++
	}
	b.bits = make([]uint64, s)
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
	b.bits[offset>>shift] |= uint64(1 << (offset & ss))
	return true
}

// Clear the bit at the given offset.
func (b *Bitmap) Clear(offset int) bool {
	if !b.Valid(offset) {
		return false
	}
	b.bits[offset>>shift] &= mask ^ uint64(1<<(offset&ss))
	return true
}

// Flip/negate the bit at the given offset.
func (b *Bitmap) Flip(offset int) bool {
	if !b.Valid(offset) {
		return false
	}
	b.bits[offset>>shift] ^= uint64(1 << (offset & ss))
	return true
}

// Test the bit at the given offset. This method will return false if the offset is also invalid for the current
// bitmap. If you are unsure about whether the given offset is valid, call Valid on the offset first.
func (b *Bitmap) Test(offset int) bool {
	if !b.Valid(offset) {
		return false
	}
	return b.bits[offset>>shift]&uint64(1<<(offset&ss)) > 0
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
	x := ss - (n & ss)
	n >>= shift
	for i := 0; i < n; i++ {
		b.bits[i] &= a.bits[i]
	}
	b.bits[n] &= a.bits[n] & (mask >> x)
}

// Or does a bitwise-or with another bitmap
func (b *Bitmap) Or(a Bitmap) {
	m, n := b.max-1, a.max-1
	if m < n {
		n = m
	}
	x := ss - (n & ss)
	n >>= shift
	for i := 0; i < n; i++ {
		b.bits[i] |= a.bits[i]
	}
	b.bits[n] |= a.bits[n] & (mask >> x)
}

// Xor does a bitwise-xor with another bitmap
func (b *Bitmap) Xor(a Bitmap) {
	m, n := b.max-1, a.max-1
	if m < n {
		n = m
	}
	x := ss - (n & ss)
	n >>= shift
	for i := 0; i < n; i++ {
		b.bits[i] ^= a.bits[i]
	}
	b.bits[n] ^= a.bits[n] & (mask >> x)
}

// Flip/negate all bits in the bitmap
func (b *Bitmap) FlipAll() {
	n := b.max - 1
	x := ss - (n & ss)
	n >>= shift
	for i := 0; i < n; i++ {
		b.bits[i] ^= mask
	}
	b.bits[n] ^= mask >> x
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
	x := ss - (n & ss)
	n >>= shift
	for i, n := 0, len(b.bits); i < n; i++ {
		b.bits[i] = mask
	}
	b.bits[n] &= mask >> x
}

// Sum returns the hamming weight of the bitmap, that is, the number of bits that are set.
//
// https://en.wikipedia.org/wiki/Hamming_weight
func (b *Bitmap) Sum() int {
	c := 0
	for i, n := 0, len(b.bits); i < n; i++ {
		x := b.bits[i]
		x -= (x >> 1) & m1
		x = (x & m2) + ((x >> 2) & m2)
		x = (x + (x >> 4)) & m4
		c += int((x * h01) >> 56)
	}
	return c
}
