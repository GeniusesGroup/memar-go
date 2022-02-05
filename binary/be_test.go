package binary

import "testing"

func TestBigEndian_Uint16(t *testing.T) {
	var expected uint16 = 0x0a0b
	if got := BigEndian.Uint16([]byte{0x0a, 0x0b}); got != expected {
		t.Errorf("BigEndian.Uint16(): got %x, want %x", got, expected)
	}
}

func TestBigEndian_Uint32(t *testing.T) {
	var expected uint32 = 0x0a0b0c0d
	if got := BigEndian.Uint32([]byte{0x0a, 0x0b, 0x0c, 0x0d}); got != expected {
		t.Errorf("BigEndian.Uint32(): got %x, want %x", got, expected)
	}
}

func TestBigEndian_Uint64(t *testing.T) {
	var expected uint64 = 0x08090a0b0c0d0e0f
	if got := BigEndian.Uint64([]byte{0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f}); got != expected {
		t.Errorf("BigEndian.Uint64(): got %x, want %x", got, expected)
	}
}

func TestBigEndian_PutUint16(t *testing.T) {
	var got [2]byte
	var expected = [2]byte{0x0a, 0x0b}
	if BigEndian.PutUint16(got[:], 0x0a0b); got != expected {
		t.Errorf("BigEndian.PutUint16(): got %x, want %x", got, expected)
	}
}

func TestBigEndian_PutUint32(t *testing.T) {
	var got [4]byte
	var expected = [4]byte{0x0a, 0x0b, 0x0c, 0x0d}
	if BigEndian.PutUint32(got[:], 0x0a0b0c0d); got != expected {
		t.Errorf("BigEndian.PutUint32(): got %x, want %x", got, expected)
	}
}

func TestBigEndian_PutUint64(t *testing.T) {
	var got [8]byte
	var expected = [8]byte{0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f}
	if BigEndian.PutUint64(got[:], 0x08090a0b0c0d0e0f); got != expected {
		t.Errorf("BigEndian.PutUint64(): got %x, want %x", got, expected)
	}
}
