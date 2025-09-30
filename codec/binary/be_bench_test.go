package binary

import "testing"

func BenchmarkBigEndian_PutUint16(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BigEndian.PutUint16(putbuf[:2], uint16(i))
	}
}

func BenchmarkBigEndian_PutUint32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BigEndian.PutUint32(putbuf[:4], uint32(i))
	}
}

func BenchmarkBigEndian_PutUint64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BigEndian.PutUint64(putbuf, uint64(i))
	}
}

func BenchmarkBigEndian_Uint16(b *testing.B) {
	for i := 0; i < b.N; i++ {
		u16 = BigEndian.Uint16(big[2:4])
	}
}

func BenchmarkBigEndian_Uint32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		u32 = BigEndian.Uint32(big[4:8])
	}
}

func BenchmarkBigEndian_Uint64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		u64 = BigEndian.Uint64(big[8:])
	}
}
