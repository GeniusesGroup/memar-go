/* For license and copyright information please see LEGAL file in repository */

package json

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"reflect"
	"testing"

	jsoniter "github.com/json-iterator/go"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
	fflib "github.com/pquerna/ffjson/fflib/v1"
)

/*
Benchmark1JsonDecode-8       	   22827	     45965 ns/op	    3000 B/op	      21 allocs/op
Benchmark1LibgoDecode-8      	  210447	      5665 ns/op	    2688 B/op	       1 allocs/op
Benchmark1EasyDecode-8       	  184730	      6001 ns/op	    2720 B/op	       2 allocs/op
Benchmark1FFDecode-8         	   87891	     13679 ns/op	    7140 B/op	       8 allocs/op
Benchmark1JsoniterDecode-8   	   78867	     15014 ns/op	    9154 B/op	       6 allocs/op

Benchmark1JsonEncode-8       	  167131	      7272 ns/op	    4609 B/op	       2 allocs/op
Benchmark1LibgoEncode-8      	  230906	      5157 ns/op	    3456 B/op	       1 allocs/op
Benchmark1EasyEncode-8       	  117720	     10600 ns/op	   18112 B/op	      11 allocs/op
Benchmark1FFEncode-8         	  107139	     11382 ns/op	   10070 B/op	      23 allocs/op
Benchmark1JsoniterEncode-8   	  163782	      7301 ns/op	    6666 B/op	       3 allocs/op

note1: This benchmark is not apple to apple due to EasyJson&&Libgo encode||decode array as base64 string!
note2: Libgo decoder performance better by: 8X from standard GO, 0.1X from EasyJson, 2.4X from FFJson and 2.7X from Jsoniter
note3: Libgo encoder performance better by: 1.4X from standard GO, 2.1X from EasyJson, 2.2X from FFJson and 1.4X from Jsoniter
*/

/*
	Test data
*/

type test1 struct {
	CaptchaID [16]byte // `json:",string"`
	Image     []byte   // `json:",string"`
}

var unMarshaledTest1 = test1{
	CaptchaID: [16]byte{167, 7, 56, 140, 146, 65, 70, 25, 183, 113, 230, 83, 166, 148, 108, 210},
}
var marshaledTest1 []byte
var marshaledTest1Libgo []byte
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
	marshaledTest1Libgo = unMarshaledTest1.libgoEncoder()
	marshaledTest1Easy, _ = unMarshaledTest1.easyEncoder()
	// marshaledTest1 = unMarshaledTest1.libgoEncoder()
	// fmt.Print("Syllab test initialized!!", "\n")
}

/*
	Decode Benchmark
*/

func Benchmark1JsonDecode(b *testing.B) {
	for n := 0; n < b.N; n++ {
		var t test1
		json.Unmarshal(marshaledTest1, &t)
	}
}

func Benchmark1LibgoDecode(b *testing.B) {
	for n := 0; n < b.N; n++ {
		var t test1
		t.libgoDecoder(marshaledTest1Libgo)
	}
}

func Benchmark1EasyDecode(b *testing.B) {
	for n := 0; n < b.N; n++ {
		var t test1
		t.easyDecoder(marshaledTest1Easy)
	}
}

func Benchmark1FFDecode(b *testing.B) {
	for n := 0; n < b.N; n++ {
		var t test1
		t.ffDecoder(marshaledTest1)
	}
}

func Benchmark1JsoniterDecode(b *testing.B) {
	for n := 0; n < b.N; n++ {
		var t test1
		jsoniter.Unmarshal(marshaledTest1, &t)
	}
}

/*
	Decode Tests
*/

func Test1JsonDecode(b *testing.T) {
	var t test1
	var err = json.Unmarshal(marshaledTest1, &t)
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

func Test1LibgoDecode(b *testing.T) {
	var t test1
	var err = t.libgoDecoder(marshaledTest1Libgo)
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

func Test1EasyDecode(b *testing.T) {
	var t test1
	var err = t.easyDecoder(marshaledTest1Easy)
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

func Test1FFDecode(b *testing.T) {
	var t test1
	var err = t.ffDecoder(marshaledTest1)
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

func Test1JsoniterDecode(b *testing.T) {
	var t test1
	var err = jsoniter.Unmarshal(marshaledTest1, &t)
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

func Benchmark1JsonEncode(b *testing.B) {
	for n := 0; n < b.N; n++ {
		json.Marshal(&unMarshaledTest1)
	}
}

func Benchmark1LibgoEncode(b *testing.B) {
	for n := 0; n < b.N; n++ {
		unMarshaledTest1.libgoEncoder()
	}
}

func Benchmark1EasyEncode(b *testing.B) {
	for n := 0; n < b.N; n++ {
		unMarshaledTest1.easyEncoder()
	}
}

func Benchmark1FFEncode(b *testing.B) {
	for n := 0; n < b.N; n++ {
		unMarshaledTest1.ffEncoder()
	}
}

func Benchmark1JsoniterEncode(b *testing.B) {
	for n := 0; n < b.N; n++ {
		jsoniter.Marshal(&unMarshaledTest1)
	}
}

/*
	Encode Tests
*/

func Test1LibgoEncode(t *testing.T) {
	var buf []byte
	buf = unMarshaledTest1.libgoEncoder()
	if !bytes.Equal(buf, marshaledTest1Libgo) {
		t.Error("Encoded unMarshaledTest1 not same\n")
		t.Error("len--cap of test: ", len(buf), "--", cap(buf), "\n")
		t.Error("len--cap of base: ", len(marshaledTest1Libgo), "--", cap(marshaledTest1Libgo), "\n")
		t.Error(string(buf), "\n")
		t.Error(string(marshaledTest1Libgo), "\n")
		t.Fail()
	}
	// fmt.Print(string(buf))
}

// TODO:::

/*
	libgo Encoder and decoder (this package)
*/

func (t *test1) libgoDecoder(buf []byte) (err error) {
	var decoder = DecoderUnsafeMinifed{
		Buf: buf,
	}
	for len(decoder.Buf) > 2 {
		decoder.Offset(2)       // remove >>	'{"' 	&& 		',"'	due to don't need them
		switch decoder.Buf[0] { // Just check first letter first!
		case 'C':
			decoder.SetFounded()
			decoder.Offset(12)
			err = decoder.DecodeArrayAsBase64(t.CaptchaID[:])
			if err != nil {
				return
			}
		case 'I':
			decoder.SetFounded()
			decoder.Offset(8)
			t.Image, err = decoder.DecodeSliceAsBase64()
			if err != nil {
				return
			}
		}

		err = decoder.IterationCheck()
		if err != nil {
			return
		}
	}
	return
}

func (t *test1) libgoEncoder() []byte {
	var encoder = Encoder{
		Buf: make([]byte, 0, t.jsonLen()),
	}

	encoder.EncodeString(`{"CaptchaID":"`)
	encoder.EncodeByteSliceAsBase64(t.CaptchaID[:])

	encoder.EncodeString(`","Image":"`)
	encoder.EncodeByteSliceAsBase64(t.Image)

	encoder.EncodeString(`"}`)
	return encoder.Buf
}

func (t *test1) jsonLen() (ln int) {
	ln = 27                                 // len(`{"CaptchaID":"","Image":""}`)
	ln += base64.StdEncoding.EncodedLen(16) // [16]bye array
	ln += base64.StdEncoding.EncodedLen(len(t.Image))
	return
}

/*
	EasyJson Encoder and decoder
	https://github.com/mailru/easyjson
	>> easyjson -all ./benchmark1_test.go
*/

func (t test1) easyEncoder() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson42239ddeEncodeBench(&w, t)
	return w.Buffer.BuildBytes(), w.Error
}

// UnmarshalJSON supports json.Unmarshaler interface
func (t *test1) easyDecoder(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson42239ddeDecodeBench(&r, t)
	return r.Error()
}

func easyjson42239ddeDecodeBench(in *jlexer.Lexer, out *test1) {
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
		case "CaptchaID":
			if in.IsNull() {
				in.Skip()
			} else {
				copy(out.CaptchaID[:], in.Bytes())
			}
		case "Image":
			if in.IsNull() {
				in.Skip()
				out.Image = nil
			} else {
				out.Image = in.Bytes()
			}
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
func easyjson42239ddeEncodeBench(out *jwriter.Writer, in test1) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"CaptchaID\":"
		out.RawString(prefix[1:])
		out.Base64Bytes(in.CaptchaID[:])
	}
	{
		const prefix string = ",\"Image\":"
		out.RawString(prefix)
		out.Base64Bytes(in.Image)
	}
	out.RawByte('}')
}

/*
	ffJson Encoder and decoder
	https://github.com/pquerna/ffjson
	>> ffjson ./benchmark1_test.go
*/

func (t *test1) ffEncoder() ([]byte, error) {
	var buf fflib.Buffer
	if t == nil {
		buf.WriteString("null")
		return buf.Bytes(), nil
	}
	err := t.MarshalJSONBuf(&buf)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (t *test1) MarshalJSONBuf(buf fflib.EncodingBuffer) error {
	if t == nil {
		buf.WriteString("null")
		return nil
	}
	var err error
	var obj []byte
	_ = obj
	_ = err
	buf.WriteString(`{"CaptchaID":`)
	buf.WriteString(`[`)
	for i, v := range t.CaptchaID {
		if i != 0 {
			buf.WriteString(`,`)
		}
		fflib.FormatBits2(buf, uint64(v), 10, false)
	}
	buf.WriteString(`]`)
	buf.WriteString(`,"Image":`)
	if t.Image != nil {
		buf.WriteString(`"`)
		{
			enc := base64.NewEncoder(base64.StdEncoding, buf)
			enc.Write(reflect.Indirect(reflect.ValueOf(t.Image)).Bytes())
			enc.Close()
		}
		buf.WriteString(`"`)
	} else {
		buf.WriteString(`null`)
	}
	buf.WriteByte('}')
	return nil
}

const (
	ffjttest1base = iota
	ffjttest1nosuchkey

	ffjttest1CaptchaID

	ffjttest1Image
)

var ffjKeytest1CaptchaID = []byte("CaptchaID")

var ffjKeytest1Image = []byte("Image")

func (t *test1) ffDecoder(input []byte) error {
	fs := fflib.NewFFLexer(input)
	return t.UnmarshalJSONFFLexer(fs, fflib.FFParse_map_start)
}

// UnmarshalJSONFFLexer fast json unmarshall - template ffjson
func (t *test1) UnmarshalJSONFFLexer(fs *fflib.FFLexer, state fflib.FFParseState) error {
	var err error
	currentKey := ffjttest1base
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
				currentKey = ffjttest1nosuchkey
				state = fflib.FFParse_want_colon
				goto mainparse
			} else {
				switch kn[0] {

				case 'C':

					if bytes.Equal(ffjKeytest1CaptchaID, kn) {
						currentKey = ffjttest1CaptchaID
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'I':

					if bytes.Equal(ffjKeytest1Image, kn) {
						currentKey = ffjttest1Image
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				}

				if fflib.SimpleLetterEqualFold(ffjKeytest1Image, kn) {
					currentKey = ffjttest1Image
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffjKeytest1CaptchaID, kn) {
					currentKey = ffjttest1CaptchaID
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				currentKey = ffjttest1nosuchkey
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

				case ffjttest1CaptchaID:
					goto handle_CaptchaID

				case ffjttest1Image:
					goto handle_Image

				case ffjttest1nosuchkey:
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

handle_CaptchaID:

	/* handler: j.CaptchaID type=[16]uint8 kind=array quoted=false*/

	{

		{
			if tok != fflib.FFTok_left_brace && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for ", tok))
			}
		}

		t.CaptchaID = [16]uint8{}

		if tok != fflib.FFTok_null {
			wantVal := true

			idx := 0
			for {

				var tmpJCaptchaID uint8

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

				/* handler: tmpJCaptchaID type=uint8 kind=uint8 quoted=false*/

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

						tmpJCaptchaID = uint8(tval)

					}
				}

				// Standard json.Unmarshal ignores elements out of array bounds,
				// that what we do as well.
				if idx < 16 {
					t.CaptchaID[idx] = tmpJCaptchaID
					idx++
				}

				wantVal = false
			}
		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Image:

	/* handler: t.Image type=[]uint8 kind=slice quoted=false*/

	{

		{
			if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for ", tok))
			}
		}

		if tok == fflib.FFTok_null {
			t.Image = nil
		} else {
			b := make([]byte, base64.StdEncoding.DecodedLen(fs.Output.Len()))
			n, err := base64.StdEncoding.Decode(b, fs.Output.Bytes())
			if err != nil {
				return fs.WrapErr(err)
			}

			v := reflect.ValueOf(&t.Image).Elem()
			v.SetBytes(b[0:n])

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
