package binary

import "testing"

var putbuf = []byte{8: 0}

var (
	boolean bool
	u8      uint8
	u16     uint16
	u32     uint32
	u64     uint64
)

var big = []byte{
	0x01,
	0x01,
	0x01, 0x02,
	0x01, 0x02, 0x03, 0x04,
	0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f,
}

var little = []byte{
	0x01,
	0x01,
	0x02, 0x01,
	0x04, 0x03, 0x02, 0x01,
	0x0f, 0x0e, 0x0d, 0x0c, 0x0b, 0x0a, 0x09, 0x08,
}

func BenchmarkPutUint8(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PutUint8(putbuf[:1], uint8(i))
	}
}

func BenchmarkPutBool(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PutBool(putbuf[:1], true)
	}
}

func BenchmarkUint8(b *testing.B) {
	for i := 0; i < b.N; i++ {
		u8 = Uint8(little[:1])
	}
}

func BenchmarkBool(b *testing.B) {
	for i := 0; i < b.N; i++ {
		boolean = Bool(little[:1])
	}
}
