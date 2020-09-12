/* For license and copyright information please see LEGAL file in repository */

package json

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"reflect"
	"testing"

	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
	fflib "github.com/pquerna/ffjson/fflib/v1"
)

/*
Benchmark1JsonDecode-8    	   29338	     40414 ns/op	    2616 B/op	      21 allocs/op
Benchmark1LibgoDecode-8   	  218322	      5580 ns/op	    2304 B/op	       1 allocs/op
Benchmark1EasyDecode-8    	 2592272	       459 ns/op	      96 B/op	       4 allocs/op
Benchmark1FFDecode-8      	   95296	     12783 ns/op	    6755 B/op	       8 allocs/op

Benchmark1JsonEncode-8    	  136446	      8768 ns/op	    6018 B/op	       2 allocs/op
Benchmark1LibgoEncode-8   	  176578	      6941 ns/op	    4864 B/op	       1 allocs/op
Benchmark1EasyEncode-8    	  102252	     11819 ns/op	   19520 B/op	      11 allocs/op
Benchmark1FFEncode-8      	   98433	     12219 ns/op	   10069 B/op	      23 allocs/op

note1: This benchmark is not apple to apple e.g. easyJson failed on decoding due to it can't decode array as array, it decode array as base64 string!
*/

/*
	Test data
*/

type test1 struct {
	CaptchaID [16]byte
	Image     []byte
}

var marshaledTest1 = []byte(`{"CaptchaID":[167,7,56,140,146,65,70,25,183,113,230,83,166,148,108,210],"Image":"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"}`)
var unMarshaledTest1 = test1{
	CaptchaID: [16]byte{167, 7, 56, 140, 146, 65, 70, 25, 183, 113, 230, 83, 166, 148, 108, 210},
	Image:     []byte("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"),
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
		t.libgoDecoder(marshaledTest1)
	}
}

func Benchmark1EasyDecode(b *testing.B) {
	for n := 0; n < b.N; n++ {
		var t test1
		t.easyDecoder(marshaledTest1)
	}
}

func Benchmark1FFDecode(b *testing.B) {
	for n := 0; n < b.N; n++ {
		var t test1
		t.ffDecoder(marshaledTest1)
	}
}

/*
	Decode Tests
*/

func Test1JsonDecode(b *testing.T) {
	var t test1
	var err = json.Unmarshal(marshaledTest1, &t)
	if t.CaptchaID != unMarshaledTest1.CaptchaID || err != nil {
		b.Fail()
	}
}

func Test1LibgoDecode(b *testing.T) {
	var t test1
	var err = t.libgoDecoder(marshaledTest1)
	if t.CaptchaID != unMarshaledTest1.CaptchaID || err != nil {
		fmt.Print(err)
		b.Fail()
	}
}

func Test1EasyDecode(b *testing.T) {
	var t test1
	var err = t.easyDecoder(marshaledTest1)
	if t.CaptchaID != unMarshaledTest1.CaptchaID || err != nil {
		b.Fail()
	}
}

func Test1FFDecode(b *testing.T) {
	var t test1
	var err = t.ffDecoder(marshaledTest1)
	if t.CaptchaID != unMarshaledTest1.CaptchaID || err != nil {
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

/*
	Encode Tests
*/

// TODO:::

/*
	libgo Encoder and decoder (this package)
*/

func (t *test1) libgoDecoder(buf []byte) (err error) {
	var decoder = Decoder{
		Buf: buf,
	}
	for {
		decoder.IterationCheckStartMinifed()
		switch decoder.Buf[0] { // Just check first letter first!
		case 'C':
			decoder.Buf = decoder.Buf[12:]
			t.CaptchaID, err = decoder.Decode16ByteArrayAsNumber()
			if err != nil {
				return
			}
		case 'I':
			decoder.Buf = decoder.Buf[8:]
			t.Image, err = decoder.DecodeSliceAsBase64()
			if err != nil {
				return
			}
		}

		if len(decoder.Buf) < 2 {
			// Reach last item!
			return
		}
		err = decoder.IterationCheckEnd()
		if err != nil {
			return
		}
	}
}

func (t *test1) libgoEncoder() []byte {
	var base64Len = base64.StdEncoding.EncodedLen(len(t.Image))
	var encoder = Encoder{
		Buf: make([]byte, 0, 90+base64Len), // Fixed-Size-Data >> 90 = 14+((16*3)+15)+11+...+2
	}

	encoder.EncodeString(`{"CaptchaID":[`)
	encoder.EncodeSliceAsNumber(t.CaptchaID[:])

	encoder.EncodeString(`],"Image":"`)
	encoder.EncodeSliceAsBase64(t.Image, base64Len)
	encoder.EncodeString(`"}`)

	return encoder.Buf
}

/*
	EasyJson Encoder and decoder
	https://github.com/mailru/easyjson
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
