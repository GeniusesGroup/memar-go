/* For license and copyright information please see the LEGAL file in the code repository */

package json

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"testing"

	buffer_p "memar/computer/buffer/protocol"
	error_p "memar/process/error/protocol"
)

type test2 struct {
	SimpleObject  map[string]uint64
	StringObject  map[string]string
	ComplexObject map[string]*innerTest2
	Array         [16]byte
	ArraySlice    []uint64
	ArrayBase64   []byte
	ComplexArray  []*innerTest2
	String        string
	IntegerNumber uint64
	FloatNumber   float64
	Boolean       bool
}

type innerTest2 struct {
	String string
	Number uint64
}

var unMarshaledTest2 = test2{
	SimpleObject: map[string]uint64{
		"SimpleObject1": 2727753243454,
		"SimpleObject2": 5468687,
		"SimpleObject3": 568788978978964537,
	},
	StringObject: map[string]string{
		"StringObject1": "test2727753243454test",
		"StringObject2": "test5468687test",
		"StringObject3": "test568788978978964537test",
	},
	ComplexObject: map[string]*innerTest2{
		"ComplexObject1": {String: "ComplexArray1", Number: 46548464},
		"ComplexObject2": {String: "ComplexArray2", Number: 546786453413},
		"ComplexObject3": {String: "ComplexArray3", Number: 645678678678677},
	},
	Array:      [16]byte{167, 7, 56, 140, 146, 65, 70, 25, 183, 113, 230, 83, 166, 148, 108, 210},
	ArraySlice: []uint64{568788978978964537, 574673574, 879894765, 687654654},
	ComplexArray: []*innerTest2{
		{String: "ComplexArray1", Number: 46548464},
		{String: "ComplexArray2", Number: 546786453413},
		{String: "ComplexArray3", Number: 645678678678677},
	},
	String:        "TestTest46548464TestTest46548464TestTest46548464",
	IntegerNumber: 5687889789789645371,
	FloatNumber:   151875.8564,
	Boolean:       true,
}
var marshaledTest2 []byte
var marshaledTest2Easy []byte

func init() {
	const sliceLen = 2400
	unMarshaledTest2.ArrayBase64 = make([]byte, sliceLen)
	var j uint8
	for i := 0; i < sliceLen; i++ {
		unMarshaledTest2.ArrayBase64[i] = j
		j++
	}

	marshaledTest2, _ = json.Marshal(&unMarshaledTest2)
	// fmt.Print(len(marshaledTest2), "--", cap(marshaledTest2), "\n")  // >> len,cap = 4099--4864
	// marshaledTest2 = unMarshaledTest2.memarEncoder()
	// fmt.Print(len(marshaledTest2), "--", cap(marshaledTest2), "\n")  // >> len,cap = 4099--4152
	// fmt.Print("Syllab test initialized!!", "\n")
}

/*
	Decode Benchmark
*/

func Benchmark2MemarDecode(b *testing.B) {
	for n := 0; n < b.N; n++ {
		var t test2
		var decoder DecoderUnsafeMinified[marshaledTest2]
		decoder.Init()
		t.memarDecoder(decoder)
	}
}

/*
	Encode Benchmarks
*/

func Benchmark2MemarEncode(b *testing.B) {
	for n := 0; n < b.N; n++ {
		unMarshaledTest2.memarEncoder()
	}
}

/*
	Tests
*/

func Test2MemarDecode(t *testing.T) {
	var t2 test2
	var err = t2.memarDecoder(marshaledTest2)
	if err != nil {
		fmt.Print("Decoded face error:", err, "\n")
		t.Fail()
	} else if t2.IntegerNumber != unMarshaledTest2.IntegerNumber {
		fmt.Print("Decoded IntegerNumber not same\n")
		t.Fail()
	} else if !bytes.Equal(t2.ArrayBase64, unMarshaledTest2.ArrayBase64) {
		fmt.Print("Decoded ArrayBase64 not same\n")
		t.Fail()
	} else if t2.Boolean != unMarshaledTest2.Boolean {
		fmt.Print("Decoded Boolean not same\n")
		t.Fail()
	}
}

/*
	memar Encoder and decoder (this package)
*/

func (t *test2) memarDecoder(decoder *DecoderUnsafeMinified[buffer_p.Buffer]) (err error_p.Error) {
	for err == nil {
		var keyName = decoder.DecodeKey()
		fmt.Println(keyName)
		switch keyName {
		case "ArraySlice":
			t.ArraySlice = make([]uint64, 0, 8)
			var value uint64
			for !decoder.CheckToken(']') {
				value, err = decoder.Decode_Integer_U64()
				t.ArraySlice = append(t.ArraySlice, value)
				decoder.Offset(1)
			}
		case "ArrayBase64":
			t.ArrayBase64, err = decoder.DecodeByteSliceAsBase64()
		case "Array":
			err = decoder.DecodeByteArrayAsNumber(t.Array[:])
		case "Boolean": // Boolean":false, || Boolean":true,
			t.Boolean, err = decoder.DecodeBool()
		case "ComplexObject":
			t.ComplexObject = make(map[string]*innerTest2, 16)
			for decoder.buf[0] != '}' {
				var key string
				var value innerTest2
				decoder.Offset(2) // remove		'{"'	||	 ',"'
				key = decoder.DecodeKey()
				for !decoder.CheckToken('}') {
					decoder.Offset(2) // remove		':"'	||	 ',"'
					switch decoder.buf[0] {
					case 'S':
						// String":"",
						decoder.Offset(9)
						value.String, err = decoder.DecodeString()
					case 'N':
						// Number":0}
						decoder.Offset(8)
						value.Number, err = decoder.Decode_Integer_U64()
						if err != nil {
							return
						}
						decoder.Offset(1)
					}
				}
				t.ComplexObject[key] = &value
			}
		case "ComplexArray":
			t.ComplexArray = make([]*innerTest2, 16)
			for decoder.buf[0] != ']' {
				var value innerTest2
				decoder.Offset(1) // remove		'['	||	 ','
				for !decoder.CheckToken('}') {
					decoder.Offset(2) // remove		'{"'
					switch decoder.buf[0] {
					case 'S':
						// String":"",
						decoder.Offset(9)
						value.String, err = decoder.DecodeString()
					case 'N':
						// Number":0}
						decoder.Offset(8)
						value.Number, err = decoder.Decode_Integer_U64()
						if err != nil {
							return
						}
					}
				}
				t.ComplexArray = append(t.ComplexArray, &value)
			}
		case "FloatNumber": // FloatNumber":12.3,
			t.FloatNumber, err = decoder.DecodeFloat64AsNumber()
		case "IntegerNumber": // IntegerNumber":0000,"
			t.IntegerNumber, err = decoder.Decode_Integer_U64()
		case "SimpleObject": // SimpleObject":{"":0,"":0}	StringObject":{"":"","":""}		String":""
			t.SimpleObject = make(map[string]uint64, 16) // TODO::: make efficient enough?
			for !decoder.CheckToken('}') {
				var key string
				var value uint64
				decoder.Offset(2) // remove		'{"'	||	 ',"'
				key = decoder.DecodeKey()
				value, err = decoder.Decode_Integer_U64()
				if err != nil {
					return
				}
				t.SimpleObject[key] = value
			}
			decoder.Offset(1)
		case "StringObject":
			var key string
			var value string
			t.StringObject = make(map[string]string, 16) // TODO::: make efficient enough?
			for decoder.buf[0] != '}' {
				decoder.Offset(2)
				key = decoder.DecodeKey()
				decoder.Offset(1)
				value, err = decoder.DecodeString()
				t.StringObject[key] = value
			}
			decoder.Offset(1)
		case "String":
			t.String, err = decoder.DecodeString()
		default:
			err = decoder.NotFoundKeyStrict()
		}

		if decoder.End() {
			// Reach last item!
			return
		}
	}
	return
}

func (t *test2) memarEncoder() (buf []byte) {
	var encoder Encoder[buffer_p.Buffer]
	encoder.Init(make([]byte, 0, t.JSON_Length()))

	encoder.EncodeString(`{"SimpleObject":{`)
	if t.SimpleObject != nil {
		for key, value := range t.SimpleObject {
			encoder.EncodeKey(key)
			encoder.Encode_Integer_U64(value)
			encoder.EncodeByte(',')
		}
		encoder.RemoveTrailingComma()
	}

	encoder.EncodeString(`},"StringObject":{`)
	if t.StringObject != nil {
		for key, value := range t.StringObject {
			encoder.EncodeKey(key)
			encoder.EncodeStringValue(value)
		}
		encoder.RemoveTrailingComma()
	}

	encoder.EncodeString(`},"ComplexObject":{`)
	if t.ComplexObject != nil {
		for key, value := range t.ComplexObject {
			if value != nil {
				encoder.EncodeKey(key)
				encoder.EncodeString(`{"String":"`)
				encoder.EncodeString(value.String)
				encoder.EncodeString(`","Number":`)
				encoder.Encode_Integer_U64(value.Number)
				encoder.EncodeString("},")
			}
		}
		encoder.RemoveTrailingComma()
	}

	encoder.EncodeString(`},"Array":[`)
	encoder.EncodeByteSliceAsNumber(t.Array[:])

	encoder.EncodeString(`],"ArraySlice":[`)
	encoder.Encode_Integer_U64SliceAsNumber(t.ArraySlice)

	encoder.EncodeString(`],"ArrayBase64":"`)
	encoder.EncodeByteSliceAsBase64(t.ArrayBase64)

	encoder.EncodeString(`","ComplexArray":[`)
	if t.ComplexObject != nil {
		for _, value := range t.ComplexArray {
			if value != nil {
				encoder.EncodeString(`{"String":"`)
				encoder.EncodeString(value.String)
				encoder.EncodeString(`","Number":`)
				encoder.Encode_Integer_U64(value.Number)
				encoder.EncodeString("},")
			}
		}
		encoder.RemoveTrailingComma()
	}

	encoder.EncodeString(`],"String":"`)
	encoder.EncodeString(t.String)

	encoder.EncodeString(`","IntegerNumber":`)
	encoder.Encode_Integer_U64(t.IntegerNumber)

	encoder.EncodeString(`,"FloatNumber":`)
	encoder.EncodeFloat64(t.FloatNumber)

	encoder.EncodeString(`,"Boolean":`)
	encoder.EncodeBoolean(t.Boolean)

	encoder.EncodeByte('}')
	return encoder.buf
}

func (t *test2) JSON_Length() (ln int) {
	ln = 185 // len(`{"SimpleObject":{},"StringObject":{},"ComplexObject":{},"Array":[],"ArraySlice":[],"ArrayBase64":"","ComplexArray":[],"String":"","IntegerNumber":,"FloatNumber":,"Boolean":}`)
	if t.SimpleObject != nil {
		ln += len(t.SimpleObject) * 20 // TODO::: IS it worth to calculate value size??
		for key := range t.SimpleObject {
			ln += len(key) + 3 // 3 = len('"":')
			// ln += 20 // TODO::: IS it worth to calculate value size??
		}
	}
	if t.StringObject != nil {
		for key, value := range t.StringObject {
			ln += len(key)
			ln += len(value)
		}
	}
	if t.ComplexObject != nil {
		ln += len(t.ComplexObject) * 24 // 24=len(`{"String":"","Number":},`)
		for key, value := range t.ComplexObject {
			if value != nil {
				ln += len(key)
				ln += len(value.String)
				ln += 20 // TODO::: Is it worth to calculate value size??
			}
		}
	}
	ln += 16 * 4                 // >> len(t.Array)
	ln += len(t.ArraySlice) * 20 // TODO::: Is it worth to calculate size??
	ln += base64.StdEncoding.EncodedLen(len(t.ArrayBase64))
	if t.ComplexArray != nil {
		for _, value := range t.ComplexArray {
			if value != nil {
				ln += 24 // 24=len(`{"String":"","Number":},`)
				ln += len(value.String)
				ln += 20 // TODO::: Is it worth to calculate value size??
			}
		}
	}
	ln += len(t.String)
	ln += 20 // >> len(t.IntegerNumber) TODO::: Is it worth to calculate integer size??
	ln += 20 // >> len(t.FloatNumber) TODO::: Is it worth to calculate integer size??
	ln += 5  // >> len("false") as bigger one!!
	return
}
