/* For license and copyright information please see LEGAL file in repository */

package json

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"runtime"
	"testing"

	"../protocol"

	jsoniter "github.com/json-iterator/go"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
	fflib "github.com/pquerna/ffjson/fflib/v1"
)

/*
Benchmark2JsonDecode-8       	   16280	     72752 ns/op	    5352 B/op	      91 allocs/op
Benchmark2LibgoDecode-8      	  108826	     11077 ns/op	    6416 B/op	      16 allocs/op
Benchmark2EasyDecode-8       	   92332	     12984 ns/op	    4256 B/op	      35 allocs/op
Benchmark2FFDecode-8         	   47458	     25140 ns/op	    8773 B/op	      45 allocs/op
Benchmark2JsoniterDecode-8   	   41660	     28993 ns/op	   11899 B/op	     103 allocs/op

Benchmark2JsonEncode-8       	   51574	     23350 ns/op	    8301 B/op	      41 allocs/op
Benchmark2LibgoEncode-8      	  155941	      7809 ns/op	    4864 B/op	       1 allocs/op
Benchmark2EasyEncode-8       	   76269	     14945 ns/op	   17294 B/op	      11 allocs/op
Benchmark2FFEncode-8         	   42363	     28328 ns/op	   19953 B/op	      62 allocs/op
Benchmark2JsoniterEncode-8   	   85363	     13688 ns/op	    9227 B/op	      15 allocs/op

note1: This benchmark is not apple to apple due to EasyJson encode||decode array as base64 string!
note2: Libgo decoder performance better by: 6.5X from standard GO, 1.2X from EasyJson, 2.2X from FFJson and 2.5X from Jsoniter
note3: Libgo encoder performance better by: 4.8X from standard GO, 1.9X from EasyJson, 3.6X from FFJson and 1.7X from Jsoniter
*/

/*
	Test data
*/

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
		"ComplexObject1": &innerTest2{String: "ComplexArray1", Number: 46548464},
		"ComplexObject2": &innerTest2{String: "ComplexArray2", Number: 546786453413},
		"ComplexObject3": &innerTest2{String: "ComplexArray3", Number: 645678678678677},
	},
	Array:      [16]byte{167, 7, 56, 140, 146, 65, 70, 25, 183, 113, 230, 83, 166, 148, 108, 210},
	ArraySlice: []uint64{568788978978964537, 574673574, 879894765, 687654654},
	ComplexArray: []*innerTest2{
		&innerTest2{String: "ComplexArray1", Number: 46548464},
		&innerTest2{String: "ComplexArray2", Number: 546786453413},
		&innerTest2{String: "ComplexArray3", Number: 645678678678677},
	},
	String:        "TestTest46548464TestTest46548464TestTest46548464",
	IntegerNumber: 5687889789789645371,
	FloatNumber:   151875.8564,
	Boolean:       true,
}
var marshaledTest2 []byte
var marshaledTest2Easy []byte

func init() {
	fmt.Print("Number of CPU used:", runtime.NumCPU(), "\n")

	const sliceLen = 2400
	unMarshaledTest2.ArrayBase64 = make([]byte, sliceLen)
	var j uint8
	for i := 0; i < sliceLen; i++ {
		unMarshaledTest2.ArrayBase64[i] = j
		j++
	}

	marshaledTest2, _ = json.Marshal(&unMarshaledTest2)
	marshaledTest2Easy, _ = unMarshaledTest2.easyEncoder()
	// fmt.Print(len(marshaledTest2), "--", cap(marshaledTest2), "\n")  // >> len,cap = 4099--4864
	// marshaledTest2 = unMarshaledTest2.libgoEncoder()
	// fmt.Print(len(marshaledTest2), "--", cap(marshaledTest2), "\n")  // >> len,cap = 4099--4152
	// fmt.Print("Syllab test initialized!!", "\n")
}

/*
	Decode Benchmark
*/

func Benchmark2JsonDecode(b *testing.B) {
	for n := 0; n < b.N; n++ {
		var t test2
		json.Unmarshal(marshaledTest2, &t)
	}
}

func Benchmark2LibgoDecode(b *testing.B) {
	for n := 0; n < b.N; n++ {
		var t test2
		t.libgoDecoder(marshaledTest2)
	}
}

func Benchmark2EasyDecode(b *testing.B) {
	for n := 0; n < b.N; n++ {
		var t test2
		t.easyDecoder(marshaledTest2Easy)
	}
}

func Benchmark2FFDecode(b *testing.B) {
	for n := 0; n < b.N; n++ {
		var t test2
		t.ffDecoder(marshaledTest2)
	}
}

func Benchmark2JsoniterDecode(b *testing.B) {
	for n := 0; n < b.N; n++ {
		var t test2
		jsoniter.Unmarshal(marshaledTest2, &t)
	}
}

/*
	Encode Benchmarks
*/

func Benchmark2JsonEncode(b *testing.B) {
	for n := 0; n < b.N; n++ {
		json.Marshal(&unMarshaledTest2)
	}
}

func Benchmark2LibgoEncode(b *testing.B) {
	for n := 0; n < b.N; n++ {
		unMarshaledTest2.libgoEncoder()
	}
}

func Benchmark2EasyEncode(b *testing.B) {
	for n := 0; n < b.N; n++ {
		unMarshaledTest2.easyEncoder()
	}
}

func Benchmark2FFEncode(b *testing.B) {
	for n := 0; n < b.N; n++ {
		unMarshaledTest2.ffEncoder()
	}
}

func Benchmark2JsoniterEncode(b *testing.B) {
	for n := 0; n < b.N; n++ {
		jsoniter.Marshal(&unMarshaledTest2)
	}
}

/*
	Tests
*/

func Test2LibgoDecode(t *testing.T) {
	var t2 test2
	var err = t2.libgoDecoder(marshaledTest2)
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

func Test2EasyDecode(t *testing.T) {
	var t2 test2
	var err = t2.easyDecoder(marshaledTest2)
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
	libgo Encoder and decoder (this package)
*/

func (t *test2) libgoDecoder(buf []byte) (err protocol.Error) {
	var decoder = DecoderUnsafeMinifed{
		Buf: buf,
	}
	for err == nil {
		var keyName = decoder.DecodeKey()
		fmt.Println(keyName)
		switch keyName {
		case "ArraySlice":
			t.ArraySlice = make([]uint64, 0, 8)
			var value uint64
			for !decoder.CheckToken(']') {
				value, err = decoder.DecodeUInt64()
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
			for decoder.Buf[0] != '}' {
				var key string
				var value innerTest2
				decoder.Offset(2) // remove		'{"'	||	 ',"'
				key = decoder.DecodeKey()
				for !decoder.CheckToken('}') {
					decoder.Offset(2) // remove		':"'	||	 ',"'
					switch decoder.Buf[0] {
					case 'S':
						// String":"",
						decoder.Offset(9)
						value.String, err = decoder.DecodeString()
					case 'N':
						// Number":0}
						decoder.Offset(8)
						value.Number, err = decoder.DecodeUInt64()
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
			for decoder.Buf[0] != ']' {
				var value innerTest2
				decoder.Offset(1) // remove		'['	||	 ','
				for !decoder.CheckToken('}') {
					decoder.Offset(2) // remove		'{"'
					switch decoder.Buf[0] {
					case 'S':
						// String":"",
						decoder.Offset(9)
						value.String, err = decoder.DecodeString()
					case 'N':
						// Number":0}
						decoder.Offset(8)
						value.Number, err = decoder.DecodeUInt64()
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
			t.IntegerNumber, err = decoder.DecodeUInt64()
		case "SimpleObject": // SimpleObject":{"":0,"":0}	StringObject":{"":"","":""}		String":""
			t.SimpleObject = make(map[string]uint64, 16) // TODO::: make efficient enough?
			for !decoder.CheckToken('}') {
				var key string
				var value uint64
				decoder.Offset(2) // remove		'{"'	||	 ',"'
				key = decoder.DecodeKey()
				value, err = decoder.DecodeUInt64()
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
			for decoder.Buf[0] != '}' {
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

		if len(decoder.Buf) < 3 {
			// Reach last item!
			return
		}
	}
	return
}

func (t *test2) libgoEncoder() (buf []byte) {
	var encoder = Encoder{
		Buf: make([]byte, 0, t.LenAsJSON()),
	}

	encoder.EncodeString(`{"SimpleObject":{`)
	if t.SimpleObject != nil {
		for key, value := range t.SimpleObject {
			encoder.EncodeKey(key)
			encoder.EncodeUInt64(value)
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
				encoder.EncodeUInt64(value.Number)
				encoder.EncodeString("},")
			}
		}
		encoder.RemoveTrailingComma()
	}

	encoder.EncodeString(`},"Array":[`)
	encoder.EncodeByteSliceAsNumber(t.Array[:])

	encoder.EncodeString(`],"ArraySlice":[`)
	encoder.EncodeUInt64SliceAsNumber(t.ArraySlice)

	encoder.EncodeString(`],"ArrayBase64":"`)
	encoder.EncodeByteSliceAsBase64(t.ArrayBase64)

	encoder.EncodeString(`","ComplexArray":[`)
	if t.ComplexObject != nil {
		for _, value := range t.ComplexArray {
			if value != nil {
				encoder.EncodeString(`{"String":"`)
				encoder.EncodeString(value.String)
				encoder.EncodeString(`","Number":`)
				encoder.EncodeUInt64(value.Number)
				encoder.EncodeString("},")
			}
		}
		encoder.RemoveTrailingComma()
	}

	encoder.EncodeString(`],"String":"`)
	encoder.EncodeString(t.String)

	encoder.EncodeString(`","IntegerNumber":`)
	encoder.EncodeUInt64(t.IntegerNumber)

	encoder.EncodeString(`,"FloatNumber":`)
	encoder.EncodeFloat64(t.FloatNumber)

	encoder.EncodeString(`,"Boolean":`)
	encoder.EncodeBoolean(t.Boolean)

	encoder.EncodeByte('}')
	return encoder.Buf
}

func (t *test2) LenAsJSON() (ln int) {
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

/*
	EasyJson Encoder and decoder
	https://github.com/mailru/easyjson
	>> easyjson -all ./benchmark2_test.go
*/

// easyEncoder supports json.Marshaler interface
func (t test2) easyEncoder() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonA6c3493fEncodeBench(&w, t)
	return w.Buffer.BuildBytes(), w.Error
}

// easyDecoder supports json.Unmarshaler interface
func (t *test2) easyDecoder(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonA6c3493fDecodeBench(&r, t)
	return r.Error()
}

func easyjsonA6c3493fDecodeBench(in *jlexer.Lexer, out *test2) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "SimpleObject":
			if in.IsNull() {
				in.Skip()
			} else {
				in.Delim('{')
				out.SimpleObject = make(map[string]uint64)
				for !in.IsDelim('}') {
					key := string(in.String())
					in.WantColon()
					var v1 uint64 = uint64(in.Uint64())
					(out.SimpleObject)[key] = v1
					in.WantComma()
				}
				in.Delim('}')
			}
		case "StringObject":
			if in.IsNull() {
				in.Skip()
			} else {
				in.Delim('{')
				out.StringObject = make(map[string]string)
				for !in.IsDelim('}') {
					key := string(in.String())
					in.WantColon()
					var v2 string = string(in.String())
					(out.StringObject)[key] = v2
					in.WantComma()
				}
				in.Delim('}')
			}
		case "ComplexObject":
			if in.IsNull() {
				in.Skip()
			} else {
				in.Delim('{')
				out.ComplexObject = make(map[string]*innerTest2)
				for !in.IsDelim('}') {
					key := string(in.String())
					in.WantColon()
					var v3 *innerTest2
					if in.IsNull() {
						in.Skip()
						v3 = nil
					} else {
						if v3 == nil {
							v3 = new(innerTest2)
						}
						(*v3).UnmarshalEasyJSON(in)
					}
					(out.ComplexObject)[key] = v3
					in.WantComma()
				}
				in.Delim('}')
			}
		case "Array":
			if in.IsNull() {
				in.Skip()
			} else {
				copy(out.Array[:], in.Bytes())
			}
		case "ArraySlice":
			if in.IsNull() {
				in.Skip()
				out.ArraySlice = nil
			} else {
				in.Delim('[')
				if out.ArraySlice == nil {
					if !in.IsDelim(']') {
						out.ArraySlice = make([]uint64, 0, 8)
					} else {
						out.ArraySlice = []uint64{}
					}
				} else {
					out.ArraySlice = (out.ArraySlice)[:0]
				}
				for !in.IsDelim(']') {
					var v5 uint64 = uint64(in.Uint64())
					out.ArraySlice = append(out.ArraySlice, v5)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "ArrayBase64":
			if in.IsNull() {
				in.Skip()
				out.ArrayBase64 = nil
			} else {
				out.ArrayBase64 = in.Bytes()
			}
		case "ComplexArray":
			if in.IsNull() {
				in.Skip()
				out.ComplexArray = nil
			} else {
				in.Delim('[')
				if out.ComplexArray == nil {
					if !in.IsDelim(']') {
						out.ComplexArray = make([]*innerTest2, 0, 8)
					} else {
						out.ComplexArray = []*innerTest2{}
					}
				} else {
					out.ComplexArray = (out.ComplexArray)[:0]
				}
				for !in.IsDelim(']') {
					var v7 *innerTest2
					if in.IsNull() {
						in.Skip()
						v7 = nil
					} else {
						if v7 == nil {
							v7 = new(innerTest2)
						}
						(*v7).UnmarshalEasyJSON(in)
					}
					out.ComplexArray = append(out.ComplexArray, v7)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "String":
			out.String = string(in.String())
		case "IntegerNumber":
			out.IntegerNumber = uint64(in.Uint64())
		case "FloatNumber":
			out.FloatNumber = float64(in.Float64())
		case "Boolean":
			out.Boolean = bool(in.Bool())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonA6c3493fEncodeBench(out *jwriter.Writer, in test2) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"SimpleObject\":"
		out.RawString(prefix[1:])
		if in.SimpleObject == nil && (out.Flags&jwriter.NilMapAsEmpty) == 0 {
			out.RawString(`null`)
		} else {
			out.RawByte('{')
			v8First := true
			for v8Name, v8Value := range in.SimpleObject {
				if v8First {
					v8First = false
				} else {
					out.RawByte(',')
				}
				out.String(string(v8Name))
				out.RawByte(':')
				out.Uint64(uint64(v8Value))
			}
			out.RawByte('}')
		}
	}
	{
		const prefix string = ",\"StringObject\":"
		out.RawString(prefix)
		if in.StringObject == nil && (out.Flags&jwriter.NilMapAsEmpty) == 0 {
			out.RawString(`null`)
		} else {
			out.RawByte('{')
			v9First := true
			for v9Name, v9Value := range in.StringObject {
				if v9First {
					v9First = false
				} else {
					out.RawByte(',')
				}
				out.String(string(v9Name))
				out.RawByte(':')
				out.String(string(v9Value))
			}
			out.RawByte('}')
		}
	}
	{
		const prefix string = ",\"ComplexObject\":"
		out.RawString(prefix)
		if in.ComplexObject == nil && (out.Flags&jwriter.NilMapAsEmpty) == 0 {
			out.RawString(`null`)
		} else {
			out.RawByte('{')
			v10First := true
			for v10Name, v10Value := range in.ComplexObject {
				if v10First {
					v10First = false
				} else {
					out.RawByte(',')
				}
				out.String(string(v10Name))
				out.RawByte(':')
				if v10Value == nil {
					out.RawString("null")
				} else {
					(*v10Value).MarshalEasyJSON(out)
				}
			}
			out.RawByte('}')
		}
	}
	{
		const prefix string = ",\"Array\":"
		out.RawString(prefix)
		out.Base64Bytes(in.Array[:])
	}
	{
		const prefix string = ",\"ArraySlice\":"
		out.RawString(prefix)
		if in.ArraySlice == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v12, v13 := range in.ArraySlice {
				if v12 > 0 {
					out.RawByte(',')
				}
				out.Uint64(uint64(v13))
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"ArrayBase64\":"
		out.RawString(prefix)
		out.Base64Bytes(in.ArrayBase64)
	}
	{
		const prefix string = ",\"ComplexArray\":"
		out.RawString(prefix)
		if in.ComplexArray == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v16, v17 := range in.ComplexArray {
				if v16 > 0 {
					out.RawByte(',')
				}
				if v17 == nil {
					out.RawString("null")
				} else {
					(*v17).MarshalEasyJSON(out)
				}
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"String\":"
		out.RawString(prefix)
		out.String(string(in.String))
	}
	{
		const prefix string = ",\"IntegerNumber\":"
		out.RawString(prefix)
		out.Uint64(uint64(in.IntegerNumber))
	}
	{
		const prefix string = ",\"FloatNumber\":"
		out.RawString(prefix)
		out.Float64(float64(in.FloatNumber))
	}
	{
		const prefix string = ",\"Boolean\":"
		out.RawString(prefix)
		out.Bool(bool(in.Boolean))
	}
	out.RawByte('}')
}

func easyjsonA6c3493fDecodeBench1(in *jlexer.Lexer, out *innerTest2) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "String":
			out.String = string(in.String())
		case "Number":
			out.Number = uint64(in.Uint64())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonA6c3493fEncodeBench1(out *jwriter.Writer, in innerTest2) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"String\":"
		out.RawString(prefix[1:])
		out.String(string(in.String))
	}
	{
		const prefix string = ",\"Number\":"
		out.RawString(prefix)
		out.Uint64(uint64(in.Number))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (it innerTest2) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonA6c3493fEncodeBench1(&w, it)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (it innerTest2) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonA6c3493fEncodeBench1(w, it)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (it *innerTest2) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonA6c3493fDecodeBench1(&r, it)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (it *innerTest2) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonA6c3493fDecodeBench1(l, it)
}

/*
	ffJson Encoder and decoder
	https://github.com/pquerna/ffjson
	>> ffjson ./benchmark2_test.go
*/

// ffEncoder marshal bytes to json - template
func (it *innerTest2) ffEncoder() ([]byte, error) {
	var buf fflib.Buffer
	if it == nil {
		buf.WriteString("null")
		return buf.Bytes(), nil
	}
	err := it.MarshalJSONBuf(&buf)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// MarshalJSONBuf marshal buff to json - template
func (it *innerTest2) MarshalJSONBuf(buf fflib.EncodingBuffer) error {
	if it == nil {
		buf.WriteString("null")
		return nil
	}
	var err error
	var obj []byte
	_ = obj
	_ = err
	buf.WriteString(`{"String":`)
	fflib.WriteJsonString(buf, string(it.String))
	buf.WriteString(`,"Number":`)
	fflib.FormatBits2(buf, uint64(it.Number), 10, false)
	buf.WriteByte('}')
	return nil
}

const (
	ffjtinnerTest2base = iota
	ffjtinnerTest2nosuchkey

	ffjtinnerTest2String

	ffjtinnerTest2Number
)

var ffjKeyinnerTest2String = []byte("String")

var ffjKeyinnerTest2Number = []byte("Number")

// ffDecoder umarshall json - template of ffjson
func (j *innerTest2) ffDecoder(input []byte) error {
	fs := fflib.NewFFLexer(input)
	return j.UnmarshalJSONFFLexer(fs, fflib.FFParse_map_start)
}

// UnmarshalJSONFFLexer fast json unmarshall - template ffjson
func (j *innerTest2) UnmarshalJSONFFLexer(fs *fflib.FFLexer, state fflib.FFParseState) error {
	var err error
	currentKey := ffjtinnerTest2base
	_ = currentKey
	tok := fflib.FFTok_init
	wantedTok := fflib.FFTok_init

mainparse:
	for {
		tok = fs.Scan()
		//	println(fmt.Sprintf("debug: tok: %v  state: %v", tok, state))
		if tok == fflib.FFTok_error {
			goto tokerror
		}

		switch state {

		case fflib.FFParse_map_start:
			if tok != fflib.FFTok_left_bracket {
				wantedTok = fflib.FFTok_left_bracket
				goto wrongtokenerror
			}
			state = fflib.FFParse_want_key
			continue

		case fflib.FFParse_after_value:
			if tok == fflib.FFTok_comma {
				state = fflib.FFParse_want_key
			} else if tok == fflib.FFTok_right_bracket {
				goto done
			} else {
				wantedTok = fflib.FFTok_comma
				goto wrongtokenerror
			}

		case fflib.FFParse_want_key:
			// json {} ended. goto exit. woo.
			if tok == fflib.FFTok_right_bracket {
				goto done
			}
			if tok != fflib.FFTok_string {
				wantedTok = fflib.FFTok_string
				goto wrongtokenerror
			}

			kn := fs.Output.Bytes()
			if len(kn) <= 0 {
				// "" case. hrm.
				currentKey = ffjtinnerTest2nosuchkey
				state = fflib.FFParse_want_colon
				goto mainparse
			} else {
				switch kn[0] {

				case 'N':

					if bytes.Equal(ffjKeyinnerTest2Number, kn) {
						currentKey = ffjtinnerTest2Number
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'S':

					if bytes.Equal(ffjKeyinnerTest2String, kn) {
						currentKey = ffjtinnerTest2String
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				}

				if fflib.SimpleLetterEqualFold(ffjKeyinnerTest2Number, kn) {
					currentKey = ffjtinnerTest2Number
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyinnerTest2String, kn) {
					currentKey = ffjtinnerTest2String
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				currentKey = ffjtinnerTest2nosuchkey
				state = fflib.FFParse_want_colon
				goto mainparse
			}

		case fflib.FFParse_want_colon:
			if tok != fflib.FFTok_colon {
				wantedTok = fflib.FFTok_colon
				goto wrongtokenerror
			}
			state = fflib.FFParse_want_value
			continue
		case fflib.FFParse_want_value:

			if tok == fflib.FFTok_left_brace || tok == fflib.FFTok_left_bracket || tok == fflib.FFTok_integer || tok == fflib.FFTok_double || tok == fflib.FFTok_string || tok == fflib.FFTok_bool || tok == fflib.FFTok_null {
				switch currentKey {

				case ffjtinnerTest2String:
					goto handle_String

				case ffjtinnerTest2Number:
					goto handle_Number

				case ffjtinnerTest2nosuchkey:
					err = fs.SkipField(tok)
					if err != nil {
						return fs.WrapErr(err)
					}
					state = fflib.FFParse_after_value
					goto mainparse
				}
			} else {
				goto wantedvalue
			}
		}
	}

handle_String:

	/* handler: j.String type=string kind=string quoted=false*/

	{

		{
			if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
			}
		}

		if tok == fflib.FFTok_null {

		} else {

			outBuf := fs.Output.Bytes()

			j.String = string(string(outBuf))

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Number:

	/* handler: j.Number type=uint64 kind=uint64 quoted=false*/

	{
		if tok != fflib.FFTok_integer && tok != fflib.FFTok_null {
			return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for uint64", tok))
		}
	}

	{

		if tok == fflib.FFTok_null {

		} else {

			tval, err := fflib.ParseUint(fs.Output.Bytes(), 10, 64)

			if err != nil {
				return fs.WrapErr(err)
			}

			j.Number = uint64(tval)

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

wantedvalue:
	return fs.WrapErr(fmt.Errorf("wanted value token, but got token: %v", tok))
wrongtokenerror:
	return fs.WrapErr(fmt.Errorf("ffjson: wanted token: %v, but got token: %v output=%s", wantedTok, tok, fs.Output.String()))
tokerror:
	if fs.BigError != nil {
		return fs.WrapErr(fs.BigError)
	}
	err = fs.Error.ToError()
	if err != nil {
		return fs.WrapErr(err)
	}
	panic("ffjson-generated: unreachable, please report bug.")
done:

	return nil
}

// ffEncoder marshal bytes to json - template
func (j *test2) ffEncoder() ([]byte, error) {
	var buf fflib.Buffer
	if j == nil {
		buf.WriteString("null")
		return buf.Bytes(), nil
	}
	err := j.MarshalJSONBuf(&buf)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// MarshalJSONBuf marshal buff to json - template
func (j *test2) MarshalJSONBuf(buf fflib.EncodingBuffer) error {
	if j == nil {
		buf.WriteString("null")
		return nil
	}
	var err error
	var obj []byte
	_ = obj
	_ = err
	if j.SimpleObject == nil {
		buf.WriteString(`{"SimpleObject":null`)
	} else {
		buf.WriteString(`{"SimpleObject":{ `)
		for key, value := range j.SimpleObject {
			fflib.WriteJsonString(buf, key)
			buf.WriteString(`:`)
			fflib.FormatBits2(buf, uint64(value), 10, false)
			buf.WriteByte(',')
		}
		buf.Rewind(1)
		buf.WriteByte('}')
	}
	if j.StringObject == nil {
		buf.WriteString(`,"StringObject":null`)
	} else {
		buf.WriteString(`,"StringObject":{ `)
		for key, value := range j.StringObject {
			fflib.WriteJsonString(buf, key)
			buf.WriteString(`:`)
			fflib.WriteJsonString(buf, string(value))
			buf.WriteByte(',')
		}
		buf.Rewind(1)
		buf.WriteByte('}')
	}
	buf.WriteString(`,"ComplexObject":`)
	/* Falling back. type=map[string]*json.innerTest2 kind=map */
	err = buf.Encode(j.ComplexObject)
	if err != nil {
		return err
	}
	buf.WriteString(`,"Array":`)
	buf.WriteString(`[`)
	for i, v := range j.Array {
		if i != 0 {
			buf.WriteString(`,`)
		}
		fflib.FormatBits2(buf, uint64(v), 10, false)
	}
	buf.WriteString(`]`)
	buf.WriteString(`,"ArraySlice":`)
	if j.ArraySlice != nil {
		buf.WriteString(`[`)
		for i, v := range j.ArraySlice {
			if i != 0 {
				buf.WriteString(`,`)
			}
			fflib.FormatBits2(buf, uint64(v), 10, false)
		}
		buf.WriteString(`]`)
	} else {
		buf.WriteString(`null`)
	}
	buf.WriteString(`,"ArrayBase64":`)
	if j.ArrayBase64 != nil {
		buf.WriteString(`"`)
		{
			enc := base64.NewEncoder(base64.StdEncoding, buf)
			enc.Write(reflect.Indirect(reflect.ValueOf(j.ArrayBase64)).Bytes())
			enc.Close()
		}
		buf.WriteString(`"`)
	} else {
		buf.WriteString(`null`)
	}
	buf.WriteString(`,"ComplexArray":`)
	if j.ComplexArray != nil {
		buf.WriteString(`[`)
		for i, v := range j.ComplexArray {
			if i != 0 {
				buf.WriteString(`,`)
			}

			{

				if v == nil {
					buf.WriteString("null")
				} else {

					err = v.MarshalJSONBuf(buf)
					if err != nil {
						return err
					}

				}

			}
		}
		buf.WriteString(`]`)
	} else {
		buf.WriteString(`null`)
	}
	buf.WriteString(`,"String":`)
	fflib.WriteJsonString(buf, string(j.String))
	buf.WriteString(`,"IntegerNumber":`)
	fflib.FormatBits2(buf, uint64(j.IntegerNumber), 10, false)
	buf.WriteString(`,"FloatNumber":`)
	fflib.AppendFloat(buf, float64(j.FloatNumber), 'g', -1, 64)
	if j.Boolean {
		buf.WriteString(`,"Boolean":true`)
	} else {
		buf.WriteString(`,"Boolean":false`)
	}
	buf.WriteByte('}')
	return nil
}

const (
	ffjttest2base = iota
	ffjttest2nosuchkey

	ffjttest2SimpleObject

	ffjttest2StringObject

	ffjttest2ComplexObject

	ffjttest2Array

	ffjttest2ArraySlice

	ffjttest2ArrayBase64

	ffjttest2ComplexArray

	ffjttest2String

	ffjttest2IntegerNumber

	ffjttest2FloatNumber

	ffjttest2Boolean
)

var ffjKeytest2SimpleObject = []byte("SimpleObject")

var ffjKeytest2StringObject = []byte("StringObject")

var ffjKeytest2ComplexObject = []byte("ComplexObject")

var ffjKeytest2Array = []byte("Array")

var ffjKeytest2ArraySlice = []byte("ArraySlice")

var ffjKeytest2ArrayBase64 = []byte("ArrayBase64")

var ffjKeytest2ComplexArray = []byte("ComplexArray")

var ffjKeytest2String = []byte("String")

var ffjKeytest2IntegerNumber = []byte("IntegerNumber")

var ffjKeytest2FloatNumber = []byte("FloatNumber")

var ffjKeytest2Boolean = []byte("Boolean")

// ffDecoder umarshall json - template of ffjson
func (j *test2) ffDecoder(input []byte) error {
	fs := fflib.NewFFLexer(input)
	return j.UnmarshalJSONFFLexer(fs, fflib.FFParse_map_start)
}

// UnmarshalJSONFFLexer fast json unmarshall - template ffjson
func (j *test2) UnmarshalJSONFFLexer(fs *fflib.FFLexer, state fflib.FFParseState) error {
	var err error
	currentKey := ffjttest2base
	_ = currentKey
	tok := fflib.FFTok_init
	wantedTok := fflib.FFTok_init

mainparse:
	for {
		tok = fs.Scan()
		//	println(fmt.Sprintf("debug: tok: %v  state: %v", tok, state))
		if tok == fflib.FFTok_error {
			goto tokerror
		}

		switch state {

		case fflib.FFParse_map_start:
			if tok != fflib.FFTok_left_bracket {
				wantedTok = fflib.FFTok_left_bracket
				goto wrongtokenerror
			}
			state = fflib.FFParse_want_key
			continue

		case fflib.FFParse_after_value:
			if tok == fflib.FFTok_comma {
				state = fflib.FFParse_want_key
			} else if tok == fflib.FFTok_right_bracket {
				goto done
			} else {
				wantedTok = fflib.FFTok_comma
				goto wrongtokenerror
			}

		case fflib.FFParse_want_key:
			// json {} ended. goto exit. woo.
			if tok == fflib.FFTok_right_bracket {
				goto done
			}
			if tok != fflib.FFTok_string {
				wantedTok = fflib.FFTok_string
				goto wrongtokenerror
			}

			kn := fs.Output.Bytes()
			if len(kn) <= 0 {
				// "" case. hrm.
				currentKey = ffjttest2nosuchkey
				state = fflib.FFParse_want_colon
				goto mainparse
			} else {
				switch kn[0] {

				case 'A':

					if bytes.Equal(ffjKeytest2Array, kn) {
						currentKey = ffjttest2Array
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffjKeytest2ArraySlice, kn) {
						currentKey = ffjttest2ArraySlice
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffjKeytest2ArrayBase64, kn) {
						currentKey = ffjttest2ArrayBase64
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'B':

					if bytes.Equal(ffjKeytest2Boolean, kn) {
						currentKey = ffjttest2Boolean
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'C':

					if bytes.Equal(ffjKeytest2ComplexObject, kn) {
						currentKey = ffjttest2ComplexObject
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffjKeytest2ComplexArray, kn) {
						currentKey = ffjttest2ComplexArray
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'F':

					if bytes.Equal(ffjKeytest2FloatNumber, kn) {
						currentKey = ffjttest2FloatNumber
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'I':

					if bytes.Equal(ffjKeytest2IntegerNumber, kn) {
						currentKey = ffjttest2IntegerNumber
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'S':

					if bytes.Equal(ffjKeytest2SimpleObject, kn) {
						currentKey = ffjttest2SimpleObject
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffjKeytest2StringObject, kn) {
						currentKey = ffjttest2StringObject
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffjKeytest2String, kn) {
						currentKey = ffjttest2String
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				}

				if fflib.SimpleLetterEqualFold(ffjKeytest2Boolean, kn) {
					currentKey = ffjttest2Boolean
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffjKeytest2FloatNumber, kn) {
					currentKey = ffjttest2FloatNumber
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffjKeytest2IntegerNumber, kn) {
					currentKey = ffjttest2IntegerNumber
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeytest2String, kn) {
					currentKey = ffjttest2String
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffjKeytest2ComplexArray, kn) {
					currentKey = ffjttest2ComplexArray
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeytest2ArrayBase64, kn) {
					currentKey = ffjttest2ArrayBase64
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeytest2ArraySlice, kn) {
					currentKey = ffjttest2ArraySlice
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffjKeytest2Array, kn) {
					currentKey = ffjttest2Array
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffjKeytest2ComplexObject, kn) {
					currentKey = ffjttest2ComplexObject
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeytest2StringObject, kn) {
					currentKey = ffjttest2StringObject
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeytest2SimpleObject, kn) {
					currentKey = ffjttest2SimpleObject
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				currentKey = ffjttest2nosuchkey
				state = fflib.FFParse_want_colon
				goto mainparse
			}

		case fflib.FFParse_want_colon:
			if tok != fflib.FFTok_colon {
				wantedTok = fflib.FFTok_colon
				goto wrongtokenerror
			}
			state = fflib.FFParse_want_value
			continue
		case fflib.FFParse_want_value:

			if tok == fflib.FFTok_left_brace || tok == fflib.FFTok_left_bracket || tok == fflib.FFTok_integer || tok == fflib.FFTok_double || tok == fflib.FFTok_string || tok == fflib.FFTok_bool || tok == fflib.FFTok_null {
				switch currentKey {

				case ffjttest2SimpleObject:
					goto handle_SimpleObject

				case ffjttest2StringObject:
					goto handle_StringObject

				case ffjttest2ComplexObject:
					goto handle_ComplexObject

				case ffjttest2Array:
					goto handle_Array

				case ffjttest2ArraySlice:
					goto handle_ArraySlice

				case ffjttest2ArrayBase64:
					goto handle_ArrayBase64

				case ffjttest2ComplexArray:
					goto handle_ComplexArray

				case ffjttest2String:
					goto handle_String

				case ffjttest2IntegerNumber:
					goto handle_IntegerNumber

				case ffjttest2FloatNumber:
					goto handle_FloatNumber

				case ffjttest2Boolean:
					goto handle_Boolean

				case ffjttest2nosuchkey:
					err = fs.SkipField(tok)
					if err != nil {
						return fs.WrapErr(err)
					}
					state = fflib.FFParse_after_value
					goto mainparse
				}
			} else {
				goto wantedvalue
			}
		}
	}

handle_SimpleObject:

	/* handler: j.SimpleObject type=map[string]uint64 kind=map quoted=false*/

	{

		{
			if tok != fflib.FFTok_left_bracket && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for ", tok))
			}
		}

		if tok == fflib.FFTok_null {
			j.SimpleObject = nil
		} else {

			j.SimpleObject = make(map[string]uint64, 0)

			wantVal := true

			for {

				var k string

				var tmpJSimpleObject uint64

				tok = fs.Scan()
				if tok == fflib.FFTok_error {
					goto tokerror
				}
				if tok == fflib.FFTok_right_bracket {
					break
				}

				if tok == fflib.FFTok_comma {
					if wantVal == true {
						// TODO(pquerna): this isn't an ideal error message, this handles
						// things like [,,,] as an array value.
						return fs.WrapErr(fmt.Errorf("wanted value token, but got token: %v", tok))
					}
					continue
				} else {
					wantVal = true
				}

				/* handler: k type=string kind=string quoted=false*/

				{

					{
						if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
							return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
						}
					}

					if tok == fflib.FFTok_null {

					} else {

						outBuf := fs.Output.Bytes()

						k = string(string(outBuf))

					}
				}

				// Expect ':' after key
				tok = fs.Scan()
				if tok != fflib.FFTok_colon {
					return fs.WrapErr(fmt.Errorf("wanted colon token, but got token: %v", tok))
				}

				tok = fs.Scan()
				/* handler: tmpJSimpleObject type=uint64 kind=uint64 quoted=false*/

				{
					if tok != fflib.FFTok_integer && tok != fflib.FFTok_null {
						return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for uint64", tok))
					}
				}

				{

					if tok == fflib.FFTok_null {

					} else {

						tval, err := fflib.ParseUint(fs.Output.Bytes(), 10, 64)

						if err != nil {
							return fs.WrapErr(err)
						}

						tmpJSimpleObject = uint64(tval)

					}
				}

				j.SimpleObject[k] = tmpJSimpleObject

				wantVal = false
			}

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_StringObject:

	/* handler: j.StringObject type=map[string]string kind=map quoted=false*/

	{

		{
			if tok != fflib.FFTok_left_bracket && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for ", tok))
			}
		}

		if tok == fflib.FFTok_null {
			j.StringObject = nil
		} else {

			j.StringObject = make(map[string]string, 0)

			wantVal := true

			for {

				var k string

				var tmpJStringObject string

				tok = fs.Scan()
				if tok == fflib.FFTok_error {
					goto tokerror
				}
				if tok == fflib.FFTok_right_bracket {
					break
				}

				if tok == fflib.FFTok_comma {
					if wantVal == true {
						// TODO(pquerna): this isn't an ideal error message, this handles
						// things like [,,,] as an array value.
						return fs.WrapErr(fmt.Errorf("wanted value token, but got token: %v", tok))
					}
					continue
				} else {
					wantVal = true
				}

				/* handler: k type=string kind=string quoted=false*/

				{

					{
						if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
							return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
						}
					}

					if tok == fflib.FFTok_null {

					} else {

						outBuf := fs.Output.Bytes()

						k = string(string(outBuf))

					}
				}

				// Expect ':' after key
				tok = fs.Scan()
				if tok != fflib.FFTok_colon {
					return fs.WrapErr(fmt.Errorf("wanted colon token, but got token: %v", tok))
				}

				tok = fs.Scan()
				/* handler: tmpJStringObject type=string kind=string quoted=false*/

				{

					{
						if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
							return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
						}
					}

					if tok == fflib.FFTok_null {

					} else {

						outBuf := fs.Output.Bytes()

						tmpJStringObject = string(string(outBuf))

					}
				}

				j.StringObject[k] = tmpJStringObject

				wantVal = false
			}

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_ComplexObject:

	/* handler: j.ComplexObject type=map[string]*json.innerTest2 kind=map quoted=false*/

	{

		{
			if tok != fflib.FFTok_left_bracket && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for ", tok))
			}
		}

		if tok == fflib.FFTok_null {
			j.ComplexObject = nil
		} else {

			j.ComplexObject = make(map[string]*innerTest2, 0)

			wantVal := true

			for {

				var k string

				var tmpJComplexObject *innerTest2

				tok = fs.Scan()
				if tok == fflib.FFTok_error {
					goto tokerror
				}
				if tok == fflib.FFTok_right_bracket {
					break
				}

				if tok == fflib.FFTok_comma {
					if wantVal == true {
						// TODO(pquerna): this isn't an ideal error message, this handles
						// things like [,,,] as an array value.
						return fs.WrapErr(fmt.Errorf("wanted value token, but got token: %v", tok))
					}
					continue
				} else {
					wantVal = true
				}

				/* handler: k type=string kind=string quoted=false*/

				{

					{
						if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
							return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
						}
					}

					if tok == fflib.FFTok_null {

					} else {

						outBuf := fs.Output.Bytes()

						k = string(string(outBuf))

					}
				}

				// Expect ':' after key
				tok = fs.Scan()
				if tok != fflib.FFTok_colon {
					return fs.WrapErr(fmt.Errorf("wanted colon token, but got token: %v", tok))
				}

				tok = fs.Scan()
				/* handler: tmpJComplexObject type=*json.innerTest2 kind=ptr quoted=false*/

				{
					if tok == fflib.FFTok_null {

						tmpJComplexObject = nil

					} else {

						if tmpJComplexObject == nil {
							tmpJComplexObject = new(innerTest2)
						}

						err = tmpJComplexObject.UnmarshalJSONFFLexer(fs, fflib.FFParse_want_key)
						if err != nil {
							return err
						}
					}
					state = fflib.FFParse_after_value
				}

				j.ComplexObject[k] = tmpJComplexObject

				wantVal = false
			}

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Array:

	/* handler: j.Array type=[16]uint8 kind=array quoted=false*/

	{

		{
			if tok != fflib.FFTok_left_brace && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for ", tok))
			}
		}

		j.Array = [16]uint8{}

		if tok != fflib.FFTok_null {
			wantVal := true

			idx := 0
			for {

				var tmpJArray uint8

				tok = fs.Scan()
				if tok == fflib.FFTok_error {
					goto tokerror
				}
				if tok == fflib.FFTok_right_brace {
					break
				}

				if tok == fflib.FFTok_comma {
					if wantVal == true {
						// TODO(pquerna): this isn't an ideal error message, this handles
						// things like [,,,] as an array value.
						return fs.WrapErr(fmt.Errorf("wanted value token, but got token: %v", tok))
					}
					continue
				} else {
					wantVal = true
				}

				/* handler: tmpJArray type=uint8 kind=uint8 quoted=false*/

				{
					if tok != fflib.FFTok_integer && tok != fflib.FFTok_null {
						return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for uint8", tok))
					}
				}

				{

					if tok == fflib.FFTok_null {

					} else {

						tval, err := fflib.ParseUint(fs.Output.Bytes(), 10, 8)

						if err != nil {
							return fs.WrapErr(err)
						}

						tmpJArray = uint8(tval)

					}
				}

				// Standard json.Unmarshal ignores elements out of array bounds,
				// that what we do as well.
				if idx < 16 {
					j.Array[idx] = tmpJArray
					idx++
				}

				wantVal = false
			}
		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_ArraySlice:

	/* handler: j.ArraySlice type=[]uint64 kind=slice quoted=false*/

	{

		{
			if tok != fflib.FFTok_left_brace && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for ", tok))
			}
		}

		if tok == fflib.FFTok_null {
			j.ArraySlice = nil
		} else {

			j.ArraySlice = []uint64{}

			wantVal := true

			for {

				var tmpJArraySlice uint64

				tok = fs.Scan()
				if tok == fflib.FFTok_error {
					goto tokerror
				}
				if tok == fflib.FFTok_right_brace {
					break
				}

				if tok == fflib.FFTok_comma {
					if wantVal == true {
						// TODO(pquerna): this isn't an ideal error message, this handles
						// things like [,,,] as an array value.
						return fs.WrapErr(fmt.Errorf("wanted value token, but got token: %v", tok))
					}
					continue
				} else {
					wantVal = true
				}

				/* handler: tmpJArraySlice type=uint64 kind=uint64 quoted=false*/

				{
					if tok != fflib.FFTok_integer && tok != fflib.FFTok_null {
						return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for uint64", tok))
					}
				}

				{

					if tok == fflib.FFTok_null {

					} else {

						tval, err := fflib.ParseUint(fs.Output.Bytes(), 10, 64)

						if err != nil {
							return fs.WrapErr(err)
						}

						tmpJArraySlice = uint64(tval)

					}
				}

				j.ArraySlice = append(j.ArraySlice, tmpJArraySlice)

				wantVal = false
			}
		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_ArrayBase64:

	/* handler: j.ArrayBase64 type=[]uint8 kind=slice quoted=false*/

	{

		{
			if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for ", tok))
			}
		}

		if tok == fflib.FFTok_null {
			j.ArrayBase64 = nil
		} else {
			b := make([]byte, base64.StdEncoding.DecodedLen(fs.Output.Len()))
			n, err := base64.StdEncoding.Decode(b, fs.Output.Bytes())
			if err != nil {
				return fs.WrapErr(err)
			}

			v := reflect.ValueOf(&j.ArrayBase64).Elem()
			v.SetBytes(b[0:n])

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_ComplexArray:

	/* handler: j.ComplexArray type=[]*json.innerTest2 kind=slice quoted=false*/

	{

		{
			if tok != fflib.FFTok_left_brace && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for ", tok))
			}
		}

		if tok == fflib.FFTok_null {
			j.ComplexArray = nil
		} else {

			j.ComplexArray = []*innerTest2{}

			wantVal := true

			for {

				var tmpJComplexArray *innerTest2

				tok = fs.Scan()
				if tok == fflib.FFTok_error {
					goto tokerror
				}
				if tok == fflib.FFTok_right_brace {
					break
				}

				if tok == fflib.FFTok_comma {
					if wantVal == true {
						// TODO(pquerna): this isn't an ideal error message, this handles
						// things like [,,,] as an array value.
						return fs.WrapErr(fmt.Errorf("wanted value token, but got token: %v", tok))
					}
					continue
				} else {
					wantVal = true
				}

				/* handler: tmpJComplexArray type=*json.innerTest2 kind=ptr quoted=false*/

				{
					if tok == fflib.FFTok_null {

						tmpJComplexArray = nil

					} else {

						if tmpJComplexArray == nil {
							tmpJComplexArray = new(innerTest2)
						}

						err = tmpJComplexArray.UnmarshalJSONFFLexer(fs, fflib.FFParse_want_key)
						if err != nil {
							return err
						}
					}
					state = fflib.FFParse_after_value
				}

				j.ComplexArray = append(j.ComplexArray, tmpJComplexArray)

				wantVal = false
			}
		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_String:

	/* handler: j.String type=string kind=string quoted=false*/

	{

		{
			if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
			}
		}

		if tok == fflib.FFTok_null {

		} else {

			outBuf := fs.Output.Bytes()

			j.String = string(string(outBuf))

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_IntegerNumber:

	/* handler: j.IntegerNumber type=uint64 kind=uint64 quoted=false*/

	{
		if tok != fflib.FFTok_integer && tok != fflib.FFTok_null {
			return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for uint64", tok))
		}
	}

	{

		if tok == fflib.FFTok_null {

		} else {

			tval, err := fflib.ParseUint(fs.Output.Bytes(), 10, 64)

			if err != nil {
				return fs.WrapErr(err)
			}

			j.IntegerNumber = uint64(tval)

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_FloatNumber:

	/* handler: j.FloatNumber type=float64 kind=float64 quoted=false*/

	{
		if tok != fflib.FFTok_double && tok != fflib.FFTok_integer && tok != fflib.FFTok_null {
			return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for float64", tok))
		}
	}

	{

		if tok == fflib.FFTok_null {

		} else {

			tval, err := fflib.ParseFloat(fs.Output.Bytes(), 64)

			if err != nil {
				return fs.WrapErr(err)
			}

			j.FloatNumber = float64(tval)

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Boolean:

	/* handler: j.Boolean type=bool kind=bool quoted=false*/

	{
		if tok != fflib.FFTok_bool && tok != fflib.FFTok_null {
			return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for bool", tok))
		}
	}

	{
		if tok == fflib.FFTok_null {

		} else {
			tmpb := fs.Output.Bytes()

			if bytes.Compare([]byte{'t', 'r', 'u', 'e'}, tmpb) == 0 {

				j.Boolean = true

			} else if bytes.Compare([]byte{'f', 'a', 'l', 's', 'e'}, tmpb) == 0 {

				j.Boolean = false

			} else {
				err = errors.New("unexpected bytes for true/false value")
				return fs.WrapErr(err)
			}

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

wantedvalue:
	return fs.WrapErr(fmt.Errorf("wanted value token, but got token: %v", tok))
wrongtokenerror:
	return fs.WrapErr(fmt.Errorf("ffjson: wanted token: %v, but got token: %v output=%s", wantedTok, tok, fs.Output.String()))
tokerror:
	if fs.BigError != nil {
		return fs.WrapErr(fs.BigError)
	}
	err = fs.Error.ToError()
	if err != nil {
		return fs.WrapErr(err)
	}
	panic("ffjson-generated: unreachable, please report bug.")
done:

	return nil
}
