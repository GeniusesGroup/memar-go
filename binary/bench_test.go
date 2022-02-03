package binary

import "testing"

var putbuf  = []byte{8: 0}

var (
	boolean bool
	u8      uint8
	u16     uint16
	u32     uint32
	u64     uint64
)

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

func BenchmarkBigEndian_PutUint16(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BigEndian.PutUint16(putbuf[:2], uint16(i))
	}
}

func BenchmarkLittleEndian_PutUint16(b *testing.B) {
	for i := 0; i < b.N; i++ {
		LittleEndian.PutUint16(putbuf[:2], uint16(i))
	}
}

func BenchmarkBigEndian_PutUint32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BigEndian.PutUint32(putbuf[:4], uint32(i))
	}
}

func BenchmarkLittleEndian_PutUint32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		LittleEndian.PutUint32(putbuf[:4], uint32(i))
	}
}

func BenchmarkBigEndian_PutUint64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BigEndian.PutUint64(putbuf, uint64(i))
	}
}

func BenchmarkLittleEndian_PutUint64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		LittleEndian.PutUint64(putbuf, uint64(i))
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

func BenchmarkBigEndian_Uint16(b *testing.B) {
	for i := 0; i < b.N; i++ {
		u16 = BigEndian.Uint16(big[2:4])
	}
}

func BenchmarkLittleEndian_Uint16(b *testing.B) {
	for i := 0; i < b.N; i++ {
		u16 = LittleEndian.Uint16(little[2:4])
	}
}

func BenchmarkBigEndian_Uint32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		u32 = BigEndian.Uint32(big[4:8])
	}
}

func BenchmarkLittleEndian_Uint32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		u32 = LittleEndian.Uint32(little[4:8])
	}
}

func BenchmarkBigEndian_Uint64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		u64 = BigEndian.Uint64(big[8:])
	}
}

func BenchmarkLittleEndian_Uint64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		u64 = LittleEndian.Uint64(little[8:])
	}
}
