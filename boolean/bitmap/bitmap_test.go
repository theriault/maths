package bitmap

import "testing"

func TestValid(t *testing.T) {
	b := New(1)
	if b.Valid(-1) {
		t.Error("bit -1 should not be valid")
	}
	if !b.Valid(0) {
		t.Error("bit 0 should be valid")
	}
	if b.Valid(1) {
		t.Error("bit 1 should not be valid")
	}
	if b.Test(1) {
		t.Error("bit 1 should not be testable")
	}
	if b.Set(1) {
		t.Error("bit 1 should not be settable")
	}
	if b.Clear(1) {
		t.Error("bit 1 should not be clearable")
	}
	if b.Flip(1) {
		t.Error("bit 1 should not be flippable")
	}
}

func TestTest(t *testing.T) {
	b := New(1)
	if b.Test(0) {
		t.Error("bit 0 should not be set")
	}
	b.Set(0)
	if !b.Test(0) {
		t.Error("bit 0 should be set")
	}
}

func TestSet(t *testing.T) {
	b := New(1)
	b.Set(0)
	if !b.Test(0) {
		t.Error("bit 0 should be set")
	}
}

func TestFlip(t *testing.T) {
	b := New(1)
	b.Flip(0)
	if !b.Test(0) {
		t.Error("bit 0 should be set")
	}
	b.Flip(0)
	if b.Test(0) {
		t.Error("bit 0 should not be set")
	}
}

func TestClear(t *testing.T) {
	b := New(1)
	b.Set(0)
	b.Clear(0)
	if b.Test(0) {
		t.Error("bit 0 should not be set")
	}
}

func TestLen(t *testing.T) {
	b := New(999)
	if b.Len() != 999 {
		t.Error("len should be 999")
	}
	b = New(8)
	if b.Len() != 8 {
		t.Error("len should be 8")
	}
}

func TestAnd(t *testing.T) {
	a := New(65)
	b := New(66)
	a.Set(0)
	b.Set(0)
	a.Set(1)
	b.Set(2)
	a.And(b)

	if !a.Test(0) || a.Test(1) || a.Test(2) || a.Test(3) || a.Test(4) {
		t.Errorf("got %b, expected 0b00001", a.bits)
	}
}

func TestOr(t *testing.T) {
	a := New(65)
	b := New(66)
	a.Set(0)
	b.Set(0)
	a.Set(1)
	b.Set(2)
	b.Set(9)
	a.Or(b)

	if !a.Test(0) || !a.Test(1) || !a.Test(2) || a.Test(3) || a.Test(8) {
		t.Errorf("got %b, expected 0b0111", a.bits)
	}
}

func TestXor(t *testing.T) {
	a := New(65)
	b := New(66)
	a.Set(0)
	b.Set(0)
	a.Set(1)
	b.Set(2)
	a.Xor(b)

	if a.Test(0) || !a.Test(1) || !a.Test(2) || a.Test(3) {
		t.Errorf("got %b, expected 0b00001", a.bits)
	}
}

func TestSetAll(t *testing.T) {
	a := New(65)
	a.SetAll()

	if !a.Test(0) || !a.Test(1) || !a.Test(2) || !a.Test(3) || !a.Test(8) {
		t.Errorf("got %b, expected 0b1_1111_1111", a.bits)
	}
}

func TestFlipAll(t *testing.T) {
	a := New(65)
	a.Set(0)
	a.Set(2)
	a.FlipAll()

	if a.Test(0) || !a.Test(1) || a.Test(2) || !a.Test(8) {
		t.Errorf("got %b, expected 0b1_1111_1010", a.bits)
	}
}

func TestClearAll(t *testing.T) {
	a := New(65)
	a.Set(0)
	a.Set(1)
	a.Set(2)
	a.Set(3)
	a.Set(8)
	a.ClearAll()

	if a.Test(0) || a.Test(1) || a.Test(2) || a.Test(3) || a.Test(8) {
		t.Errorf("got %b, expected 0b0_0000_0000", a.bits)
	}
}

func TestSum(t *testing.T) {
	a := New(65)
	a.Set(0)
	a.Set(1)
	a.Set(2)
	a.Set(3)

	if c := a.Sum(); c != 4 {
		t.Errorf("got %d, expected 4", c)
	}
}

func BenchmarkBitmapFlip(b *testing.B) {
	m := New(10000000)
	for i := 0; i < b.N; i++ {
		m.Flip(79999999)
	}
}

func BenchmarkBitmapFlipAll(b *testing.B) {
	m := New(100000)
	for i := 0; i < b.N; i++ {
		m.FlipAll()
	}
}

func BenchmarkBitmapXor(b *testing.B) {
	a := New(500)
	c := New(500)
	c.FlipAll()
	for i := 0; i < b.N; i++ {
		a.Xor(c)
	}
}

func BenchmarkBitmapSum(b *testing.B) {
	m := New(80000)
	m.Set(0)
	m.Set(8)
	m.Set(16)
	m.Set(24)
	m.Set(32)
	m.Set(40)
	for i := 0; i < b.N; i++ {
		_ = m.Sum()
	}
}
