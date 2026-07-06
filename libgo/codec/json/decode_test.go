/* For license and copyright information please see the LEGAL file in the code repository */

package json

import (
	"bytes"
	"encoding/json"
	"fmt"
	"testing"

	error_p "memar/process/error/protocol"
)

type test1 struct {
	CaptchaID [16]byte // `json:",string"`
	Image     []byte   // `json:",string"`
}

var unMarshaledTest1 = test1{
	CaptchaID: [16]byte{167, 7, 56, 140, 146, 65, 70, 25, 183, 113, 230, 83, 166, 148, 108, 210},
}
var marshaledTest1 []byte
var marshaledTest1Memar []byte
var marshaledTest1Easy []byte

func init() {
	const sliceLen = 2400
	unMarshaledTest1.Image = make([]byte, sliceLen)
	var j uint8
	for i := 0; i < sliceLen; i++ {
		unMarshaledTest1.Image[i] = j
		j++
	}

	marshaledTest1, _ = json.Marshal(&unMarshaledTest1)
	marshaledTest1Memar = unMarshaledTest1.memarEncoder()
	// marshaledTest1 = unMarshaledTest1.memarEncoder()
	// fmt.Print("Syllab test initialized!!", "\n")
}

/*
	Decode Benchmark
*/

func Benchmark1MemarDecode(b *testing.B) {
	for n := 0; n < b.N; n++ {
		var t test1
		t.memarDecoder(marshaledTest1Memar)
	}
}

/*
	Decode Tests
*/

func Test1MemarDecode(b *testing.T) {
	var t test1
	var err = t.memarDecoder(marshaledTest1Memar)
	if err != nil {
		fmt.Print(err, "\n")
		b.Fail()
	} else if t.CaptchaID != unMarshaledTest1.CaptchaID {
		fmt.Print("Decoded CaptchaID not same\n")
		b.Fail()
	} else if !bytes.Equal(t.Image, unMarshaledTest1.Image) {
		fmt.Print("Decoded Image not same\n")
		b.Fail()
	}
}

/*
	Encode Benchmarks
*/

func Benchmark1MemarEncode(b *testing.B) {
	for n := 0; n < b.N; n++ {
		unMarshaledTest1.memarEncoder()
	}
}

/*
	Encode Tests
*/

func Test1MemarEncode(t *testing.T) {
	var buf []byte = unMarshaledTest1.memarEncoder()
	if !bytes.Equal(buf, marshaledTest1Memar) {
		t.Error("Encoded unMarshaledTest1 not same\n")
		t.Error("len--cap of test: ", len(buf), "--", cap(buf), "\n")
		t.Error("len--cap of base: ", len(marshaledTest1Memar), "--", cap(marshaledTest1Memar), "\n")
		t.Error(string(buf), "\n")
		t.Error(string(marshaledTest1Memar), "\n")
		t.Fail()
	}
	// fmt.Print(string(buf))
}

// TODO:::

/*
	memar Encoder and decoder (this package)
*/

func (t *test1) memarDecoder(buf []byte) (err error_p.Error) {
	var decoder DecoderUnsafeMinified
	decoder.Init(buf)

	for err == nil {
		var keyName = decoder.DecodeKey()
		switch keyName {
		case "":
			return
		case "CaptchaID":
			err = decoder.DecodeByteArrayAsBase64(t.CaptchaID[:])
		case "Image":
			t.Image, err = decoder.DecodeByteSliceAsBase64()
		default:
			err = decoder.NotFoundKeyStrict()
		}

		if decoder.End() {
			return
		}
	}
	return
}

func (t *test1) memarEncoder() []byte {
	var encoder Encoder
	encoder.Init(make([]byte, 0, t.JSON_Length()))
	encoder.EncodeString(`{"CaptchaID":"`)
	encoder.EncodeByteSliceAsBase64(t.CaptchaID[:])
	encoder.EncodeString(`","Image":"`)
	encoder.EncodeByteSliceAsBase64(t.Image)
	encoder.EncodeString(`"}`)
	return encoder.Buf()
}

func (t *test1) JSON_Length() (ln int) {
	ln = ((len(t.Image)*8 + 5) / 6)
	ln += 49
	return
}
