package binary

import "testing"

func BenchmarkLittleEndian_PutUint16(b *testing.B) {
	for i := 0; i < b.N; i++ {
		LittleEndian.PutUint16(putbuf[:2], uint16(i))
	}
}

func BenchmarkLittleEndian_PutUint32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		LittleEndian.PutUint32(putbuf[:4], uint32(i))
	}
}

func BenchmarkLittleEndian_PutUint64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		LittleEndian.PutUint64(putbuf, uint64(i))
	}
}

func BenchmarkLittleEndian_Uint16(b *testing.B) {
	for i := 0; i < b.N; i++ {
		u16 = LittleEndian.Uint16(little[2:4])
	}
}

func BenchmarkLittleEndian_Uint32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		u32 = LittleEndian.Uint32(little[4:8])
	}
}

func BenchmarkLittleEndian_Uint64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		u64 = LittleEndian.Uint64(little[8:])
	}
}
