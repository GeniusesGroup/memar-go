/* For license and copyright information please see LEGAL file in repository */

package matn

import (
	"crypto/sha512"

	"../achaemenid"
	etime "../earth-time"
	er "../error"
	"../ganjine"
	gsdk "../ganjine-sdk"
	gs "../ganjine-services"
	"../json"
	lang "../language"
	"../pehrest"
	psdk "../pehrest-sdk"
	"../syllab"
)

const indexPhraseStructureID uint64 = 504952921428380083

var indexPhraseStructure = ganjine.DataStructure{
	ID:                504952921428380083,
	IssueDate:         1608786641,
	ExpiryDate:        0,
	ExpireInFavorOf:   "", // Other structure name
	ExpireInFavorOfID: 0,  // Other StructureID! Handy ID or Hash of ExpireInFavorOf!
	Status:            ganjine.DataStructureStatePreAlpha,
	Structure:         IndexPhrase{},

	Name: map[lang.Language]string{
		lang.LanguageEnglish: "Index Phrase",
	},
	Description: map[lang.Language]string{
		lang.LanguageEnglish: "Store the phrase index",
	},
	TAGS: []string{
		"",
	},
}

// IndexPhrase is standard structure to store any hash byte index!!
// It is simple secondary index e.g. hash("user@email.com")
type IndexPhrase struct {
	/* Common header data */
	RecordID          [32]byte
	RecordStructureID uint64
	RecordSize        uint64
	WriteTime         etime.Time
	OwnerAppID        [32]byte

	/* Unique data */
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
func (ip *IndexPhrase) SaveNew() (err *er.Error) {
	err = ip.Set()
	if err != nil {
		return
	}
	ip.IndexRecordIDForTermsPageNumber()
	return
}

// Set method set some data and write entire IndexPhrase record!
func (ip *IndexPhrase) Set() (err *er.Error) {
	ip.RecordStructureID = indexPhraseStructureID
	ip.RecordSize = ip.syllabLen()
	ip.WriteTime = etime.Now()
	ip.OwnerAppID = achaemenid.Server.AppID

	var req = gs.SetRecordReq{
		Type:   gs.RequestTypeBroadcast,
		Record: ip.syllabEncoder(),
	}
	ip.RecordID = sha512.Sum512_256(req.Record[32:])
	copy(req.Record[0:], ip.RecordID[:])

	err = gsdk.SetRecord(&req)
	return
}

// GetByRecordID method read all existing record data by given RecordID!
func (ip *IndexPhrase) GetByRecordID() (err *er.Error) {
	var req = gs.GetRecordReq{
		RecordID:          ip.RecordID,
		RecordStructureID: indexPhraseStructureID,
	}
	var res *gs.GetRecordRes
	res, err = gsdk.GetRecord(&req)
	if err != nil {
		return
	}

	err = ip.syllabDecoder(res.Record)
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
func (ip *IndexPhrase) FindRecordsIDByTermsPageNumber(offset, limit uint64) (RecordsID [][32]byte, err *er.Error) {
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
		Type:       gs.RequestTypeBroadcast,
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
	return sha512.Sum512_256(buf[:])
}

/*
	-- Syllab Encoder & Decoder --
*/

func (ip *IndexPhrase) syllabDecoder(buf []byte) (err *er.Error) {
	if uint32(len(buf)) < ip.syllabStackLen() {
		err = syllab.ErrSyllabDecodeSmallSlice
		return
	}

	copy(ip.RecordID[:], buf[0:])
	ip.RecordStructureID = syllab.GetUInt64(buf, 32)
	ip.RecordSize = syllab.GetUInt64(buf, 40)
	ip.WriteTime = etime.Time(syllab.GetInt64(buf, 48))
	copy(ip.OwnerAppID[:], buf[56:])

	ip.Terms = syllab.UnsafeGetStringArray(buf, 88)
	ip.PageNumber = syllab.GetUInt64(buf, 96)

	var si uint32 = 104
	for i := 0; i < 10; i++ {
		ip.Tokens[i].syllabDecoder(buf, si)
		si += uint32(ip.Tokens[i].syllabLen())
	}
	return
}

func (ip *IndexPhrase) syllabEncoder() (buf []byte) {
	buf = make([]byte, ip.syllabLen())
	var hi uint32 = ip.syllabStackLen() // Heap index || Stack size!

	// copy(buf[0:], ip.RecordID[:])
	syllab.SetUInt64(buf, 32, ip.RecordStructureID)
	syllab.SetUInt64(buf, 40, ip.RecordSize)
	syllab.SetInt64(buf, 48, int64(ip.WriteTime))
	copy(buf[56:], ip.OwnerAppID[:])

	hi = syllab.SetStringArray(buf, ip.Terms, 88, hi)
	syllab.SetUInt64(buf, 96, ip.PageNumber)

	var si uint32 = 104
	for i := 0; i < 10; i++ {
		ip.Tokens[i].syllabEncoder(buf, si, 0)
		si += uint32(ip.Tokens[i].syllabLen())
	}
	return
}

func (ip *IndexPhrase) syllabStackLen() (ln uint32) {
	ln = 104
	ln += uint32(len(ip.Tokens)) * ip.Tokens[0].syllabStackLen()
	return
}

func (ip *IndexPhrase) syllabHeapLen() (ln uint32) {
	for i := 0; i < len(ip.Terms); i++ {
		ln += uint32(len(ip.Terms[i]))
	}
	// ln += uint32(len(ip.Tokens)) * ip.Tokens[0].syllabHeapLen()
	return
}

func (ip *IndexPhrase) syllabLen() (ln uint64) {
	return uint64(ip.syllabStackLen() + ip.syllabHeapLen())
}

/*
	-- PhraseToken Encoder & Decoder --
*/

func (pt *PhraseToken) syllabDecoder(buf []byte, stackIndex uint32) {
	copy(pt.RecordID[:], buf[stackIndex:])
	pt.RecordStructureID = syllab.GetUInt64(buf, stackIndex+32)
	pt.RecordFieldID = syllab.GetUInt8(buf, stackIndex+40)
	copy(pt.RecordPrimaryKey[:], buf[stackIndex+41:])
}

func (pt *PhraseToken) syllabEncoder(buf []byte, stackIndex, heapIndex uint32) (nextHeapAddr uint32) {
	copy(buf[0:], pt.RecordID[stackIndex:])
	syllab.SetUInt64(buf, stackIndex+32, pt.RecordStructureID)
	syllab.SetUInt8(buf, stackIndex+40, pt.RecordFieldID)
	copy(buf[stackIndex+41:], pt.RecordPrimaryKey[:])
	return heapIndex
}

func (pt *PhraseToken) syllabStackLen() (ln uint32) {
	return 73
}

func (pt *PhraseToken) syllabHeapLen() (ln uint32) {
	return
}

func (pt *PhraseToken) syllabLen() (ln uint64) {
	return uint64(pt.syllabStackLen() + pt.syllabHeapLen())
}

func (pt *PhraseToken) jsonDecoder(decoder json.DecoderUnsafeMinifed) (err *er.Error) {
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

		if len(decoder.Buf) < 3 {
			return
		}
	}
	return
}

func (pt *PhraseToken) jsonEncoder(encoder json.Encoder) {
	encoder.EncodeString(`{"RecordID":"`)
	encoder.EncodeByteSliceAsBase64(pt.RecordID[:])

	encoder.EncodeString(`","RecordStructureID":`)
	encoder.EncodeUInt64(pt.RecordStructureID)

	encoder.EncodeString(`,"RecordFieldID":`)
	encoder.EncodeUInt8(pt.RecordFieldID)

	encoder.EncodeString(`,"RecordPrimaryKey":"`)
	encoder.EncodeByteSliceAsBase64(pt.RecordPrimaryKey[:])

	encoder.EncodeString(`"}`)
}

func (pt *PhraseToken) jsonLen() (ln int) {
	ln = 184
	return
}
