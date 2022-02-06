package binary

import "testing"

func TestUint8(t *testing.T) {
	var expected uint8 = 0x01
	if got := Uint8([]byte{0x01}); got != expected {
		t.Errorf("Uint8(): got %x, want %x", got, expected)
	}
}

func TestBool(t *testing.T) {
	var expected bool = true
	if got := Bool([]byte{0x01}); got != expected {
		t.Errorf("Bool(): got %t, want %t", got, expected)
	}
}

func TestPutUint8(t *testing.T) {
	var got = [1]byte{0x00}
	var expected = [1]byte{0x10}
	if PutUint8(got[:], 0x10); got != expected {
		t.Errorf("Uint8(): got %v, want %v", got, expected)
	}
}

func TestPutBool(t *testing.T) {
	var got = [1]byte{0x00}
	var expected = [1]byte{0x01}
	if PutBool(got[:], true); got != expected {
		t.Errorf("PutBool(): got %v, want %v", got, expected)
	}
}
