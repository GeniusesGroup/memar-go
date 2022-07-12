/* For license and copyright information please see LEGAL file in repository */

package matn

import (
	"golang.org/x/crypto/sha3"

	"../ganjine"
	"../json"
	"../pehrest"
	"../protocol"
	"../syllab"
	"../time/utc"
)

const indexPhraseStructureID uint64 = 736712670881955651

var indexPhraseStructure = ds.DataStructure{
	URN:             "urn:giti:matn.protocol:data-structure:index-phrase",
	ID:              indexPhraseStructureID,
	IssueDate:       1608786641,
	ExpiryDate:      0,
	ExpireInFavorOf: "",
	Status:          protocol.Software_PreAlpha,
	Structure:       IndexPhrase{},

	Name: map[protocol.LanguageID]string{
		protocol.LanguageEnglish: "Index Phrase",
	},
	Description: map[protocol.LanguageID]string{
		protocol.LanguageEnglish: "Store the phrase index",
	},
	TAGS: []string{
		"",
	},
}

// IndexPhrase is standard structure to store any hash byte index!!
// It is simple secondary index e.g. hash("user@email.com")
type IndexPhrase struct {
	RecordID   [32]byte
	Terms      []string `index-hash:"RecordID[pair,PageNumber]"` // Array must be order to retrievable!
	PageNumber uint64
	Tokens     [10]PhraseToken // Order of PhraseTokens index changed by algorithm in exact period of time!
}

// PhraseToken store detail about a word in the record to index
type PhraseToken struct {
	RecordID          [32]byte `json:",string"`
	RecordStructureID uint64
	RecordFieldID     uint8
	RecordPrimaryKey  [32]byte `json:",string"` // Store any primary ID or any data up to 32 byte length
	// Don't need Snippet text due to it is not web search engine!
}

// SaveNew method set some data and write entire IndexPhrase record with all indexes!
func (ip *IndexPhrase) SaveNew() (err protocol.Error) {
	err = ip.Set()
	if err != nil {
		return
	}
	ip.IndexRecordIDForTermsPageNumber()
	return
}

// Set method set some data and write entire IndexPhrase record!
func (ip *IndexPhrase) Set() (err protocol.Error) {
	ip.RecordStructureID = indexPhraseStructureID
	ip.RecordSize = ip.LenAsSyllab()
	ip.WriteTime = utc.Now()
	ip.OwnerAppID = protocol.OS.AppManifest().AppUUID()

	var req = ganjine.SetRecordReq{
		Type:   ganjine.RequestTypeBroadcast,
		Record: ip.ToSyllab(),
	}
	ip.RecordID = sha3.Sum256(req.Record[32:])
	copy(req.Record[0:], ip.RecordID[:])

	err = gsdk.SetRecord(&req)
	return
}

// GetByRecordID method read all existing record data by given RecordID!
func (ip *IndexPhrase) GetByRecordID() (err protocol.Error) {
	var req = ganjine.GetRecordReq{
		RecordID:          ip.RecordID,
		RecordStructureID: indexPhraseStructureID,
	}
	var res *ganjine.GetRecordRes
	res, err = gsdk.GetRecord(&req)
	if err != nil {
		return
	}

	err = ip.FromSyllab(res.Record)
	if err != nil {
		return
	}

	if ip.RecordStructureID != indexPhraseStructureID {
		err = ganjine.ErrMisMatchedStructureID
	}
	return
}

/*
	-- Search Methods --
*/

// FindRecordsIDByTermsPageNumber find RecordsID by given Terms+PageNumber
func (ip *IndexPhrase) FindRecordsIDByTermsPageNumber(offset, limit uint64) (RecordsID [][32]byte, err protocol.Error) {
	var indexReq = &pehrest.HashGetValuesReq{
		IndexKey: ip.hashTermsPageNumberForRecordID(),
		Offset:   offset,
		Limit:    limit,
	}
	var indexRes *pehrest.HashGetValuesRes
	indexRes, err = psdk.HashGetValues(indexReq)
	RecordsID = indexRes.IndexValues
	return
}

/*
	-- PRIMARY INDEXES --
*/

// IndexRecordIDForTermsPageNumber save RecordID chain for Terms+PageNumber
// Call in each update to the exiting record!
func (ip *IndexPhrase) IndexRecordIDForTermsPageNumber() {
	var indexRequest = pehrest.HashSetValueReq{
		Type:       ganjine.RequestTypeBroadcast,
		IndexKey:   ip.hashTermsPageNumberForRecordID(),
		IndexValue: ip.RecordID,
	}
	var err = psdk.HashSetValue(&indexRequest)
	if err != nil {
		// TODO::: we must retry more due to record wrote successfully!
	}
}

func (ip *IndexPhrase) hashTermsPageNumberForRecordID() (hash [32]byte) {
	const field = "TermsPageNumber"
	var bufLen = 16 + len(field)
	for _, t := range ip.Terms {
		bufLen += len(t)
	}
	var buf = make([]byte, bufLen)
	syllab.SetUInt64(buf, 0, indexPhraseStructureID)
	syllab.SetUInt64(buf, 8, ip.PageNumber)
	copy(buf[16:], field)
	bufLen = 16 + len(field)
	for _, t := range ip.Terms {
		copy(buf[bufLen:], t)
		bufLen += len(t)
	}
	return sha3.Sum256(buf[:])
}

/*
	-- Syllab Encoder & Decoder --
*/

func (ip *IndexPhrase) FromSyllab(payload []byte, stackIndex uint32) {
	if uint32(len(payload)) < ip.LenOfSyllabStack() {
		err = syllab.ErrShortArrayDecode
		return
	}

	ip.Terms = syllab.UnsafeGetStringArray(payload, 88)
	ip.PageNumber = syllab.GetUInt64(payload, 96)

	var si uint32 = 104
	for i := 0; i < 10; i++ {
		ip.Tokens[i].FromSyllab(payload, si)
		si += uint32(ip.Tokens[i].LenAsSyllab())
	}
	return
}

func (ip *IndexPhrase) ToSyllab(payload []byte, stackIndex, heapIndex uint32) (freeHeapIndex uint32) {
	heapIndex = syllab.SetStringArray(payload, ip.Terms, 88, heapIndex)
	syllab.SetUInt64(payload, 96, ip.PageNumber)

	var si uint32 = 104
	for i := 0; i < 10; i++ {
		ip.Tokens[i].ToSyllab(payload, si)
		si += uint32(ip.Tokens[i].LenAsSyllab())
	}
	return heapIndex
}

func (ip *IndexPhrase) LenOfSyllabStack() (ln uint32) {
	ln = 104
	ln += uint32(len(ip.Tokens)) * ip.Tokens[0].LenOfSyllabStack()
	return
}

func (ip *IndexPhrase) LenOfSyllabHeap() (ln uint32) {
	for i := 0; i < len(ip.Terms); i++ {
		ln += uint32(len(ip.Terms[i]))
	}
	// ln += uint32(len(ip.Tokens)) * ip.Tokens[0].LenOfSyllabHeap
	return
}

func (ip *IndexPhrase) LenAsSyllab() uint64 {
	return uint64(ip.LenOfSyllabStack() + ip.LenOfSyllabHeap())
}

/*
	-- PhraseToken Encoder & Decoder --
*/

func (pt *PhraseToken) FromSyllab(buf []byte, stackIndex uint32) {
	copy(pt.RecordID[:], buf[stackIndex:])
	pt.RecordStructureID = syllab.GetUInt64(buf, stackIndex+32)
	pt.RecordFieldID = syllab.GetUInt8(buf, stackIndex+40)
	copy(pt.RecordPrimaryKey[:], buf[stackIndex+41:])
}

func (pt *PhraseToken) ToSyllab(buf []byte, stackIndex uint32) {
	copy(buf[0:], pt.RecordID[stackIndex:])
	syllab.SetUInt64(buf, stackIndex+32, pt.RecordStructureID)
	syllab.SetUInt8(buf, stackIndex+40, pt.RecordFieldID)
	copy(buf[stackIndex+41:], pt.RecordPrimaryKey[:])
}

func (pt *PhraseToken) LenOfSyllabStack() uint32 {
	return 73
}

func (pt *PhraseToken) LenOfSyllabHeap() (ln uint32) {
	return
}

func (pt *PhraseToken) LenAsSyllab() uint64 {
	return uint64(pt.LenOfSyllabStack() + pt.LenOfSyllabHea)
}

func (pt *PhraseToken) FromJSON(decoder *json.DecoderUnsafeMinified) (err protocol.Error) {
	for err == nil {
		var keyName = decoder.DecodeKey()
		switch keyName {
		case "RecordID":
			err = decoder.DecodeByteArrayAsBase64(pt.RecordID[:])
		case "RecordStructureID":
			pt.RecordStructureID, err = decoder.DecodeUInt64()
		case "RecordFieldID":
			pt.RecordFieldID, err = decoder.DecodeUInt8()
		case "RecordPrimaryKey":
			err = decoder.DecodeByteArrayAsBase64(pt.RecordPrimaryKey[:])
		default:
			err = decoder.NotFoundKeyStrict()
		}

		if decoder.End() {
			return
		}
	}
	return
}

func (pt *PhraseToken) ToJSON(payload []byte) []byte {
	var encoder = json.Encoder{Buf: payload}

	encoder.EncodeString(`{"RecordID":"`)
	encoder.EncodeByteSliceAsBase64(pt.RecordID[:])

	encoder.EncodeString(`","RecordStructureID":`)
	encoder.EncodeUInt64(pt.RecordStructureID)

	encoder.EncodeString(`,"RecordFieldID":`)
	encoder.EncodeUInt8(pt.RecordFieldID)

	encoder.EncodeString(`,"RecordPrimaryKey":"`)
	encoder.EncodeByteSliceAsBase64(pt.RecordPrimaryKey[:])

	encoder.EncodeString(`"}`)
	return encoder.Buf
}

func (pt *PhraseToken) LenAsJSON() (ln int) {
	ln = 184
	return
}
