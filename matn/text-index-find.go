/* For license and copyright information please see LEGAL file in repository */

package matn

import (
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
func IndexTextFind(req *IndexTextFindReq) (res *IndexTextFindRes, err protocol.Error) {
	// TODO::: now just one word work!!!
	var w = &IndexWord{
		Word:            req.Term,
		RecordStructure: req.RecordStructure,
	}
	var phraseTokens []PhraseToken
	phraseTokens, err = w.FindByWordRecordStructure(req.PageNumber*10, 10)

	res = &IndexTextFindRes{
		// Tokens: *(*[10]PhraseToken)(unsafe.Pointer(&phraseTokens)),
	}
	// TODO::: fix to use unsafe
	for i := 0; i < len(phraseTokens); i++ {
		res.Tokens[i] = phraseTokens[i]
	}
	return
}

/*
	Request Encoders & Decoders
*/

// FromSyllab decode syllab to given IndexTextFindReq
func (req *IndexTextFindReq) FromSyllab(payload []byte, stackIndex uint32) {
	if uint32(len(buf)) < req.LenOfSyllabStack() {
		err = syllab.ErrShortArrayDecode
		return
	}

	req.Term = syllab.UnsafeGetString(buf, 0)
	req.RecordStructure = syllab.GetUInt64(buf, 8)
	req.PageNumber = syllab.GetUInt64(buf, 16)
	return
}

// ToSyllab encode given IndexTextFindReq to syllab format
func (req *IndexTextFindReq) ToSyllab(payload []byte, stackIndex, heapIndex uint32) (freeHeapIndex uint32)
	var hi uint32 = req.LenOfSyllabStack() // Heap index || Stack size!

	hi = syllab.SetString(buf, req.Term, 0, hi)
	syllab.SetUInt64(buf, 8, req.RecordStructure)
	syllab.SetUInt64(buf, 16, req.PageNumber)
	return
}

// LenOfSyllabStack return stack length of IndexTextFindReq
func (req *IndexTextFindReq) LenOfSyllabStack() uint32 {
	return 24
}

// LenOfSyllabHeap return heap length of IndexTextFindReq
func (req *IndexTextFindReq) LenOfSyllabHeap() (ln uint32) {
	ln += uint32(len(req.Term))
	return
}

// LenAsSyllab return whole length of IndexTextFindReq
func (req *IndexTextFindReq) LenAsSyllab() uint64 {
	return uint64(req.LenOfSyllabStack() + req.LenOfSyllabHeap())
}

// JSONDecoder decode json to given IndexTextFindReq
func (req *IndexTextFindReq) FromJSON(payload []byte) (err protocol.Error) {
	var decoder = json.DecoderUnsafeMinified{Buf: payload}
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

		if decoder.End() {
			return
		}
	}
	return
}

func (req *IndexTextFindReq) ToJSON(payload []byte) []byte {
	var encoder = json.Encoder{Buf: payload}
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

// FromSyllab decode syllab to given IndexTextFindRes
func (res *IndexTextFindRes) FromSyllab(payload []byte, stackIndex uint32) {
	if uint32(len(buf)) < res.LenOfSyllabStack() {
		err = syllab.ErrShortArrayDecode
		return
	}

	for i := 0; i < 10; i++ {
		res.Tokens[i].FromSyllab(buf, uint32(i)*res.Tokens[i].LenOfSyllabStack())
	}
	return
}

// ToSyllab encode given IndexTextFindRes to syllab format
func (res *IndexTextFindRes) ToSyllab(payload []byte, stackIndex, heapIndex uint32) (freeHeapIndex uint32)
	for i := 0; i < 10; i++ {
		res.Tokens[i].ToSyllab(buf, uint32(i)*res.Tokens[i].LenOfSyllabStack())
	}
	return
}

// LenOfSyllabStack return stack length of IndexTextFindRes
func (res *IndexTextFindRes) LenOfSyllabStack() uint32 {
	ln = 10 * res.Tokens[0].LenOfSyllabStack()
	return
}

// LenOfSyllabHeap return heap length of IndexTextFindRes
func (res *IndexTextFindRes) LenOfSyllabHeap() (ln uint32) {
	// ln += 10 * res.Tokens[0].LenOfSyllabHeap()
	return
}

// LenAsSyllab return whole length of IndexTextFindRes
func (res *IndexTextFindRes) LenAsSyllab() uint64 {
	return uint64(res.LenOfSyllabStack() + res.LenOfSyllabHeap())
}

// JSONDecoder decode json to given IndexTextFindRes
func (res *IndexTextFindRes) FromJSON(payload []byte) (err protocol.Error) {
	var decoder = json.DecoderUnsafeMinified{Buf: payload}
	for err == nil {
		var keyName = decoder.DecodeKey()
		switch keyName {
		case "Tokens":
			for i := 0; i < 10; i++ {
				err = res.Tokens[i].jsonDecoder(&decoder)
			}
		default:
			err = decoder.NotFoundKeyStrict()
		}

		if decoder.End() {
			return
		}
	}
	return
}

func (res *IndexTextFindRes) ToJSON(payload []byte) []byte {
	var encoder = json.Encoder{Buf: payload}

	var encoder = json.Encoder{
		Buf: make([]byte, 0, res.JSONLen()),
	}

	encoder.EncodeString(`{"Tokens":[`)
	for i := 0; i < 10; i++ {
		res.Tokens[i].ToJSON(&encoder)
		encoder.EncodeByte(',')
	}
	encoder.RemoveTrailingComma()

	encoder.EncodeString(`]}`)
}

// JSONLen return json needed len to encode!
func (res *IndexTextFindRes) JSONLen() (ln int) {
	ln = 10 * res.Tokens[0].LenAsJSON()
	ln += 13
	return
}
