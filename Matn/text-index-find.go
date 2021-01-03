/* For license and copyright information please see LEGAL file in repository */

package matn

import (
	"unsafe"

	er "../error"
	"../json"
	"../syllab"
)

// IndexTextFindReq is request structure of IndexTextFind()
type IndexTextFindReq struct {
	Term            string
	RecordStructure uint64
	PageNumber      uint64
}

// IndexTextFindRes is response structure of IndexTextFind()
type IndexTextFindRes struct {
	Tokens [10]PhraseToken
}

// IndexTextFind return index data of given terms if any exist
func IndexTextFind(req *IndexTextFindReq) (res *IndexTextFindRes, err *er.Error) {
	// TODO::: now just one word work!!!
	var w = &IndexWord{
		Word:            req.Term,
		RecordStructure: req.RecordStructure,
	}
	var phraseTokens []PhraseToken
	phraseTokens, err = w.FindByWordRecordStructure(req.PageNumber*10, 10)

	res = &IndexTextFindRes{
		Tokens: *(*[10]PhraseToken)(unsafe.Pointer(&phraseTokens)),
	}
	return
}

/*
	Request Encoders & Decoders
*/

// SyllabDecoder decode syllab to given IndexTextFindReq
func (req *IndexTextFindReq) SyllabDecoder(buf []byte) (err *er.Error) {
	if uint32(len(buf)) < req.SyllabStackLen() {
		err = syllab.ErrSyllabDecodeSmallSlice
		return
	}

	req.Term = syllab.UnsafeGetString(buf, 0)
	req.RecordStructure = syllab.GetUInt64(buf, 8)
	req.PageNumber = syllab.GetUInt64(buf, 16)
	return
}

// SyllabEncoder encode given IndexTextFindReq to syllab format
func (req *IndexTextFindReq) SyllabEncoder(buf []byte) {
	var hi uint32 = req.SyllabStackLen() // Heap index || Stack size!

	hi = syllab.SetString(buf, req.Term, 0, hi)
	syllab.SetUInt64(buf, 8, req.RecordStructure)
	syllab.SetUInt64(buf, 16, req.PageNumber)
	return
}

// SyllabStackLen return stack length of IndexTextFindReq
func (req *IndexTextFindReq) SyllabStackLen() (ln uint32) {
	return 24
}

// SyllabHeapLen return heap length of IndexTextFindReq
func (req *IndexTextFindReq) SyllabHeapLen() (ln uint32) {
	ln += uint32(len(req.Term))
	return
}

// SyllabLen return whole length of IndexTextFindReq
func (req *IndexTextFindReq) SyllabLen() (ln uint64) {
	return uint64(req.SyllabStackLen() + req.SyllabHeapLen())
}

// JSONDecoder decode json to given IndexTextFindReq
func (req *IndexTextFindReq) JSONDecoder(buf []byte) (err *er.Error) {
	var decoder = json.DecoderUnsafeMinifed{
		Buf: buf,
	}
	for err == nil {
		var keyName = decoder.DecodeKey()
		switch keyName {
		case "Term":
			req.Term, err = decoder.DecodeString()
		case "RecordStructure":
			req.RecordStructure, err = decoder.DecodeUInt64()
		case "PageNumber":
			req.PageNumber, err = decoder.DecodeUInt64()
		default:
			err = decoder.NotFoundKeyStrict()
		}

		if len(decoder.Buf) < 3 {
			return
		}
	}
	return
}

// JSONEncoder encode given IndexTextFindReq to json format.
func (req *IndexTextFindReq) JSONEncoder() (buf []byte) {
	var encoder = json.Encoder{
		Buf: make([]byte, 0, req.JSONLen()),
	}

	encoder.EncodeString(`{"Term":"`)
	encoder.EncodeString(req.Term)

	encoder.EncodeString(`","RecordStructure":`)
	encoder.EncodeUInt64(req.RecordStructure)

	encoder.EncodeString(`,"PageNumber":`)
	encoder.EncodeUInt64(req.PageNumber)

	encoder.EncodeByte('}')
	return encoder.Buf
}

// JSONLen return json needed len to encode!
func (req *IndexTextFindReq) JSONLen() (ln int) {
	ln = len(req.Term)
	ln += 84
	return
}

/*
	Response Encoders & Decoders
*/

// SyllabDecoder decode syllab to given IndexTextFindRes
func (res *IndexTextFindRes) SyllabDecoder(buf []byte) (err *er.Error) {
	if uint32(len(buf)) < res.SyllabStackLen() {
		err = syllab.ErrSyllabDecodeSmallSlice
		return
	}

	for i := 0; i < 10; i++ {
		res.Tokens[i].syllabDecoder(buf, uint32(i)*res.Tokens[i].syllabStackLen())
	}
	return
}

// SyllabEncoder encode given IndexTextFindRes to syllab format
func (res *IndexTextFindRes) SyllabEncoder(buf []byte) {
	var hi uint32 = res.SyllabStackLen() // Heap index || Stack size!

	for i := 0; i < 10; i++ {
		res.Tokens[i].syllabEncoder(buf, uint32(i)*res.Tokens[i].syllabStackLen(), hi)
	}
	return
}

// SyllabStackLen return stack length of IndexTextFindRes
func (res *IndexTextFindRes) SyllabStackLen() (ln uint32) {
	ln = 10 * res.Tokens[0].syllabStackLen()
	return
}

// SyllabHeapLen return heap length of IndexTextFindRes
func (res *IndexTextFindRes) SyllabHeapLen() (ln uint32) {
	// ln += 10 * res.Tokens[0].syllabHeapLen()
	return
}

// SyllabLen return whole length of IndexTextFindRes
func (res *IndexTextFindRes) SyllabLen() (ln uint64) {
	return uint64(res.SyllabStackLen() + res.SyllabHeapLen())
}

// JSONDecoder decode json to given IndexTextFindRes
func (res *IndexTextFindRes) JSONDecoder(buf []byte) (err *er.Error) {
	var decoder = json.DecoderUnsafeMinifed{
		Buf: buf,
	}
	for err == nil {
		var keyName = decoder.DecodeKey()
		switch keyName {
		case "Tokens":
			for i := 0; i < 10; i++ {
				err = res.Tokens[i].jsonDecoder(decoder)
			}
		default:
			err = decoder.NotFoundKeyStrict()
		}

		if len(decoder.Buf) < 3 {
			return
		}
	}
	return
}

// JSONEncoder encode given IndexTextFindRes to json format.
func (res *IndexTextFindRes) JSONEncoder() (buf []byte) {
	var encoder = json.Encoder{
		Buf: make([]byte, 0, res.JSONLen()),
	}

	encoder.EncodeString(`{"Tokens":[`)
	for i := 0; i < 10; i++ {
		res.Tokens[i].jsonEncoder(encoder)
	}
	encoder.RemoveTrailingComma()

	encoder.EncodeString(`]}`)
	return encoder.Buf
}

// JSONLen return json needed len to encode!
func (res *IndexTextFindRes) JSONLen() (ln int) {
	ln = 10 * res.Tokens[0].jsonLen()
	ln += 13
	return
}
