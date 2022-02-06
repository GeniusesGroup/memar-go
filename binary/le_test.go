package binary

import "testing"

func TestLittleEndian_Uint16(t *testing.T) {
	var expected uint16 = 0x0b0a
	if got := LittleEndian.Uint16([]byte{0x0a, 0x0b}); got != expected {
		t.Errorf("LittleEndian.Uint16(): got %x, want %x", got, expected)
	}
}

func TestLittleEndian_Uint32(t *testing.T) {
	var expected uint32 = 0x0d0c0b0a
	if got := LittleEndian.Uint32([]byte{0x0a, 0x0b, 0x0c, 0x0d}); got != expected {
		t.Errorf("LittleEndian.Uint32(): got %x, want %x", got, expected)
	}
}

func TestLittleEndian_Uint64(t *testing.T) {
	var expected uint64 = 0x0f0e0d0c0b0a0908
	if got := LittleEndian.Uint64([]byte{0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f}); got != expected {
		t.Errorf("LittleEndian.Uint64(): got %x, want %x", got, expected)
	}
}

func TestLittleEndian_PutUint16(t *testing.T) {
	var got [2]byte
	var expected = [2]byte{0x0b, 0x0a}
	if LittleEndian.PutUint16(got[:], 0x0a0b); got != expected {
		t.Errorf("LittleEndian.PutUint16(): got %x, want %x", got, expected)
	}
}

func TestLittleEndian_PutUint32(t *testing.T) {
	var got [4]byte
	var expected = [4]byte{0x0d, 0x0c, 0x0b, 0x0a}
	if LittleEndian.PutUint32(got[:], 0x0a0b0c0d); got != expected {
		t.Errorf("LittleEndian.PutUint32(): got %x, want %x", got, expected)
	}
}

func TestLittleEndian_PutUint64(t *testing.T) {
	var got [8]byte
	var expected = [8]byte{0x0f, 0x0e, 0x0d, 0x0c, 0x0b, 0x0a, 0x09, 0x08}
	if LittleEndian.PutUint64(got[:], 0x08090a0b0c0d0e0f); got != expected {
		t.Errorf("LittleEndian.PutUint64(): got %x, want %x", got, expected)
	}
}
