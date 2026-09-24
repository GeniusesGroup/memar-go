/* For license and copyright information please see the LEGAL file in the code repository */

package syllab

import (
	"bytes"
	"fmt"
	"testing"

	syllab_errs "memar/codec/syllab/errors"
	error_p "memar/process/error/protocol"
)

/*
	Test data
*/

type test1 struct {
	CaptchaID uint64
	Image     []byte
}

var unMarshaledTest1 = test1{
	CaptchaID: 1824074400375950161,
}
var syllabMarshaledTest1 []byte
var protoBufMarshaledTest1 []byte

func init() {
	const sliceLen = 2400
	unMarshaledTest1.Image = make([]byte, sliceLen)
	var j uint8
	for i := 0; i < sliceLen; i++ {
		unMarshaledTest1.Image[i] = j
		j++
	}

	syllabMarshaledTest1 = make([]byte, unMarshaledTest1.Syllab_Length())
	unMarshaledTest1.ToSyllab(syllabMarshaledTest1, 0, unMarshaledTest1.Syllab_StackLength())

	// fmt.Print("Syllab test initialized!!", "\n")
}

/*
	Decode && Encode Benchmarks
*/

func Benchmark1SyllabDecode(b *testing.B) {
	for n := 0; n < b.N; n++ {
		var t test1
		t.FromSyllab(syllabMarshaledTest1, 0)
	}
}

func Benchmark1SyllabDecodeSafely(b *testing.B) {
	for n := 0; n < b.N; n++ {
		var t test1
		t.FromSyllabSafely(syllabMarshaledTest1, 0)
	}
}

func Benchmark1SyllabEncode(b *testing.B) {
	for n := 0; n < b.N; n++ {
		var buf = make([]byte, unMarshaledTest1.Syllab_Length())
		unMarshaledTest1.ToSyllab(buf, 0, unMarshaledTest1.Syllab_StackLength())
	}
}

/*
	Decode && Encode Tests
*/

func Test1SyllabUnsafeDecode(b *testing.T) {
	var t test1
	var err = t.CheckSyllab(syllabMarshaledTest1)
	if err != nil {
		fmt.Print(err, "\n")
		b.Fail()
		return
	}

	t.FromSyllabSafely(syllabMarshaledTest1, 0)
	if t.CaptchaID != unMarshaledTest1.CaptchaID {
		fmt.Print("Decoded CaptchaID not same\n")
		b.Fail()
	} else if !bytes.Equal(t.Image, unMarshaledTest1.Image) {
		fmt.Print("Decoded Image not same\n")
		b.Fail()
	}
}

func Test1SyllabSafeDecode(b *testing.T) {
	var t test1
	var err = t.CheckSyllab(syllabMarshaledTest1)
	if err != nil {
		fmt.Print(err, "\n")
		b.Fail()
		return
	}

	t.FromSyllab(syllabMarshaledTest1, 0)
	if t.CaptchaID != unMarshaledTest1.CaptchaID {
		fmt.Print("Decoded CaptchaID not same\n")
		b.Fail()
	} else if !bytes.Equal(t.Image, unMarshaledTest1.Image) {
		fmt.Print("Decoded Image not same\n")
		b.Fail()
	}
}

func Test1SyllabUnSafeEncode(b *testing.T) {
	var buf = make([]byte, unMarshaledTest1.Syllab_Length())
	unMarshaledTest1.ToSyllab(buf, 0, unMarshaledTest1.Syllab_StackLength())
	if len(buf) != len(syllabMarshaledTest1) {
		b.Fail()
	}
}

/*
	Syllab Encoder and Decoder
*/

func (t *test1) CheckSyllab(buf []byte) (err error_p.Error) {
	if len(buf) < int(t.Syllab_StackLength()) {
		err = &syllab_errs.ShortArrayDecode
	}
	return
}

func (t *test1) FromSyllab(buf []byte, stackIndex uint32) {
	t.CaptchaID = GetUInt64(buf, 0)
	t.Image = GetByteArray(buf, 8)
}

func (t *test1) FromSyllabSafely(buf []byte, stackIndex uint32) {
	t.CaptchaID = GetUInt64(buf, 0)
	t.Image = GetByteArraySafely(buf, 8)
}

func (t *test1) ToSyllab(buf []byte, stackIndex, heapIndex uint32) (freeHeapIndex uint32) {
	SetUInt64(buf, 0, t.CaptchaID)
	freeHeapIndex = SetByteArray(buf, t.Image, 8, heapIndex)
	return
}

func (t *test1) Syllab_StackLength() uint32 {
	return 16 // 16 = 8+(1*8)
}

func (t *test1) Syllab_HeapLength() (ln uint32) {
	ln += uint32(len(t.Image))
	return
}

func (t *test1) Syllab_Length() uint64 {
	return uint64(t.Syllab_StackLength() + t.Syllab_HeapLength())
}
