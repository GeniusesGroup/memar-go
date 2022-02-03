package binary

import (
	"fmt"
	"reflect"
	"testing"
)

type ByteOrder interface {
	Uint16([]byte) uint16
	Uint32([]byte) uint32
	Uint64([]byte) uint64
	PutUint16([]byte, uint16)
	PutUint32([]byte, uint32)
	PutUint64([]byte, uint64)
	fmt.Stringer
}

type Struct struct {
	Bool   bool
	Uint8  uint8
	Uint16 uint16
	Uint32 uint32
	Uint64 uint64
}

var data = Struct{
	Bool:   true,
	Uint8:  0x01,
	Uint16: 0x0102,
	Uint32: 0x01020304,
	Uint64: 0x08090a0b0c0d0e0f,
}

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

func check(t *testing.T, f string, order ByteOrder, have, want interface{}) {
	if !reflect.DeepEqual(have, want) {
		t.Errorf("%v %v:\n\thave %+v\n\twant %+v", f, order, have, want)
	}
}

func testRead(t *testing.T, order ByteOrder, b []byte, s1 Struct) {
	var s2 = Struct{
		Bool:   Bool(b[:1]),
		Uint8:  Uint8(b[1:2]),
		Uint16: order.Uint16(b[2:4]),
		Uint32: order.Uint32(b[4:8]),
		Uint64: order.Uint64(b[8:16]),
	}
	check(t, "Read", order, s2, s1)
}

func testWrite(t *testing.T, order ByteOrder, b []byte, s1 Struct) {
	buf := make([]byte, 16)
	PutBool(buf[:1], s1.Bool)
	PutUint8(buf[1:2], s1.Uint8)
	order.PutUint16(buf[2:4], s1.Uint16)
	order.PutUint32(buf[4:8], s1.Uint32)
	order.PutUint64(buf[8:16], s1.Uint64)
	check(t, "Read", order, buf, b)
}

func TestReadBigEndian(t *testing.T) {
	testRead(t, BigEndian, big, data)
}

func TestWriteBigEndian(t *testing.T) {
	testWrite(t, BigEndian, big, data)
}

func TestReadLittleEndian(t *testing.T) {
	testRead(t, LittleEndian, little, data)
}

func TestWriteLittleEndian(t *testing.T) {
	testWrite(t, LittleEndian, little, data)
}

func (b bigEndian) String() string {
	return "BigEndian"
}

func (b littleEndian) String() string {
	return "LittleEndian"
}
