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
	lang "../language"
	"../pehrest"
	psdk "../pehrest-sdk"
	"../syllab"
)

const indexWordStructureID uint64 = 7165704319104393939

var indexWordStructure = ganjine.DataStructure{
	ID:                7165704319104393939,
	IssueDate:         1608786632,
	ExpiryDate:        0,
	ExpireInFavorOf:   "", // Other structure name
	ExpireInFavorOfID: 0,  // Other StructureID! Handy ID or Hash of ExpireInFavorOf!
	Status:            ganjine.DataStructureStatePreAlpha,
	Structure:         IndexWord{},

	Name: map[lang.Language]string{
		lang.LanguageEnglish: "Index Word",
	},
	Description: map[lang.Language]string{
		lang.LanguageEnglish: "store the word index data",
	},
	TAGS: []string{
		"",
	},
}

// IndexWord is standard structure to store the word index data!
type IndexWord struct {
	/* Common header data */
	RecordID          [32]byte
	RecordStructureID uint64
	RecordSize        uint64
	WriteTime         etime.Time
	OwnerAppID        [32]byte

	/* Unique data */
	Word               string `index-hash:"RecordID,RecordID[pair,RecordStructure],RecordID[pair,RecordSecondaryKey],RecordID[pair,RecordOwnerID]"` // Order of recordIDs index changed by algorithm in exact period of time!
	RecordStructure    uint64
	RecordPrimaryKey   [32]byte // Store any primary ID or any data up to 32 byte length e.g. ID
	RecordSecondaryKey [32]byte // Store any secondary ID or any data up to 32 byte length e.g. GroupID
	RecordOwnerID      [32]byte
	Tokens             []WordToken
}

// WordToken store detail about a word in the record to index
type WordToken struct {
	RecordID             [32]byte `json:",string"`
	RecordFieldID        uint8
	WordType             WordType
	WordOffsetInSentence uint64 //  Position of the word in the sentence
	WordOffsetInText     uint64 //  Position of the word in the text
	OffsetInSentence     uint64 //  First word charecter possition in the sentence
	OffsetInText         uint64 //  First word charecter possition in the text
}

// SaveNew method set some data and write entire IndexWord record with all indexes!
func (iw *IndexWord) SaveNew() (err *er.Error) {
	err = iw.Set()
	if err != nil {
		return
	}
	iw.IndexRecordIDForWord()
	iw.IndexRecordIDForWordRecordStructure()
	if iw.RecordSecondaryKey != [32]byte{} {
		iw.IndexRecordIDForWordRecordSecondaryKey()
	}
	iw.IndexRecordIDForWordRecordOwnerID()
	return
}

// SaveOrUpdate method set some data and write entire IndexWord record with all indexes or update exiting one!
func (iw *IndexWord) SaveOrUpdate() (err *er.Error) {
	var check = IndexWord{
		Word:             iw.Word,
		RecordPrimaryKey: iw.RecordPrimaryKey,
	}
	err = check.GetByWordRecordPrimaryKey()
	if err.Equal(ganjine.ErrRecordNotFound) {
		err = iw.SaveNew()
	} else {
		iw.Tokens = append(iw.Tokens, make([]WordToken, 0, len(iw.Tokens)+len(check.Tokens))...)
		for _, token := range check.Tokens {
			// TODO::: need to check first by RecordID??
			iw.Tokens = append(iw.Tokens, token)
		}
		err = iw.Set()
	}
	return
}

// Set method set some data and write entire IndexWord record!
func (iw *IndexWord) Set() (err *er.Error) {
	iw.RecordID = iw.hashWordRecordPrimaryKeyForRecordID()
	iw.RecordStructureID = indexWordStructureID
	iw.RecordSize = iw.syllabLen()
	iw.WriteTime = etime.Now()
	iw.OwnerAppID = achaemenid.Server.AppID

	var req = gs.SetRecordReq{
		Type:   gs.RequestTypeBroadcast,
		Record: iw.syllabEncoder(),
	}
	err = gsdk.SetRecord(&req)
	return
}

func (iw *IndexWord) hashWordRecordPrimaryKeyForRecordID() (hash [32]byte) {
	const field = "WordRecordPrimaryKey"
	var buf = make([]byte, 40+len(field)+len(iw.Word)) // 8+32
	syllab.SetUInt64(buf, 0, indexWordStructureID)
	copy(buf[8:], iw.RecordPrimaryKey[:])
	copy(buf[40:], field)
	copy(buf[40+len(field):], iw.Word)
	return sha512.Sum512_256(buf)
}

// GetByRecordID method read all existing record data by given RecordID!
func (iw *IndexWord) GetByRecordID() (err *er.Error) {
	var req = gs.GetRecordReq{
		RecordID:          iw.RecordID,
		RecordStructureID: indexWordStructureID,
	}
	var res *gs.GetRecordRes
	res, err = gsdk.GetRecord(&req)
	if err != nil {
		return
	}

	err = iw.syllabDecoder(res.Record)
	if err != nil {
		return
	}

	if iw.RecordStructureID != indexWordStructureID {
		err = ganjine.ErrMisMatchedStructureID
	}
	return
}

// GetByWordRecordPrimaryKey find RecordsID by given Word+RecordPrimaryKey
func (iw *IndexWord) GetByWordRecordPrimaryKey() (err *er.Error) {
	iw.RecordID = iw.hashWordRecordPrimaryKeyForRecordID()
	err = iw.GetByRecordID()
	return
}

/*
	-- Search Methods --
*/

// FindRecordsIDByWord find RecordsID by given ID
func (iw *IndexWord) FindRecordsIDByWord(offset, limit uint64) (RecordsID [][32]byte, err *er.Error) {
	var indexReq = &pehrest.HashGetValuesReq{
		IndexKey: iw.hashWordforRecordID(),
		Offset:   offset,
		Limit:    limit,
	}
	var indexRes *pehrest.HashGetValuesRes
	indexRes, err = psdk.HashGetValues(indexReq)
	RecordsID = indexRes.IndexValues
	return
}

// FindRecordsIDByWordRecordStructure find RecordsID by given Word+RecordStructure
func (iw *IndexWord) FindRecordsIDByWordRecordStructure(offset, limit uint64) (RecordsID [][32]byte, err *er.Error) {
	var indexReq = &pehrest.HashGetValuesReq{
		IndexKey: iw.hashWordRecordStructureForRecordID(),
		Offset:   offset,
		Limit:    limit,
	}
	var indexRes *pehrest.HashGetValuesRes
	indexRes, err = psdk.HashGetValues(indexReq)
	RecordsID = indexRes.IndexValues
	return
}

// FindRecordsIDByWordSecondaryKey find RecordsID by given Word+SecondaryKey
func (iw *IndexWord) FindRecordsIDByWordSecondaryKey(offset, limit uint64) (RecordsID [][32]byte, err *er.Error) {
	var indexReq = &pehrest.HashGetValuesReq{
		IndexKey: iw.hashWordRecordSecondaryKeyForRecordID(),
		Offset:   offset,
		Limit:    limit,
	}
	var indexRes *pehrest.HashGetValuesRes
	indexRes, err = psdk.HashGetValues(indexReq)
	RecordsID = indexRes.IndexValues
	return
}

// FindRecordsIDByWordRecordOwnerID find RecordsID by given Word+RecordOwnerID
func (iw *IndexWord) FindRecordsIDByWordRecordOwnerID(offset, limit uint64) (RecordsID [][32]byte, err *er.Error) {
	var indexReq = &pehrest.HashGetValuesReq{
		IndexKey: iw.hashWordRecordOwnerIDForRecordID(),
		Offset:   offset,
		Limit:    limit,
	}
	var indexRes *pehrest.HashGetValuesRes
	indexRes, err = psdk.HashGetValues(indexReq)
	RecordsID = indexRes.IndexValues
	return
}

// FindByWordRecordStructure find  by given Word+RecordStructure
func (iw *IndexWord) FindByWordRecordStructure(offset, limit uint64) (phraseTokens []PhraseToken, err *er.Error) {
	var indexReq = &pehrest.HashGetValuesReq{
		IndexKey: iw.hashWordRecordStructureForRecordID(),
		Offset:   offset,
		Limit:    limit,
	}
	var indexRes *pehrest.HashGetValuesRes
	indexRes, err = psdk.HashGetValues(indexReq)
	var RecordsID = indexRes.IndexValues

	phraseTokens = make([]PhraseToken, len(RecordsID))
	for i := 0; i < len(RecordsID); i++ {
		iw.RecordID = RecordsID[i]
		iw.GetByRecordID()

		phraseTokens[i] = PhraseToken{
			RecordID:          iw.Tokens[len(iw.Tokens)-1].RecordID,
			RecordStructureID: iw.RecordStructure,
			RecordFieldID:     iw.Tokens[len(iw.Tokens)-1].RecordFieldID,
			RecordPrimaryKey:  iw.RecordPrimaryKey,
		}
	}
	return
}

/*
	-- PRIMARY INDEXES --
*/

// IndexRecordIDForWord save RecordID chain for ID+Language
// Call in each update to the exiting record!
func (iw *IndexWord) IndexRecordIDForWord() {
	var indexRequest = pehrest.HashSetValueReq{
		Type:       gs.RequestTypeBroadcast,
		IndexKey:   iw.hashWordforRecordID(),
		IndexValue: iw.RecordID,
	}
	var err = psdk.HashSetValue(&indexRequest)
	if err != nil {
		// TODO::: we must retry more due to record wrote successfully!
	}
}

func (iw *IndexWord) hashWordforRecordID() (hash [32]byte) {
	const field = "Word"
	var buf = make([]byte, 8+len(field)+len(iw.Word))
	syllab.SetUInt64(buf, 0, indexWordStructureID)
	copy(buf[8:], field)
	copy(buf[8+len(field):], iw.Word)
	return sha512.Sum512_256(buf[:])
}

/*
	-- SECONDARY INDEXES --
*/

// IndexRecordIDForWordRecordStructure save RecordID chain for Word+RecordStructure
// Don't call in update to an exiting record!
func (iw *IndexWord) IndexRecordIDForWordRecordStructure() {
	var indexRequest = pehrest.HashSetValueReq{
		Type:       gs.RequestTypeBroadcast,
		IndexKey:   iw.hashWordRecordStructureForRecordID(),
		IndexValue: iw.RecordID,
	}
	var err = psdk.HashSetValue(&indexRequest)
	if err != nil {
		// TODO::: we must retry more due to record wrote successfully!
	}
}

func (iw *IndexWord) hashWordRecordStructureForRecordID() (hash [32]byte) {
	const field = "WordRecordStructure"
	var buf = make([]byte, 8+len(field)+len(iw.Word)) // 8+8
	syllab.SetUInt64(buf, 0, indexWordStructureID)
	syllab.SetUInt64(buf, 8, iw.RecordStructure)
	copy(buf[16:], field)
	copy(buf[16+len(field):], iw.Word)
	return sha512.Sum512_256(buf)
}

// IndexRecordIDForWordRecordSecondaryKey save RecordID chain for Word+RecordSecondaryKey
// Don't call in update to an exiting record!
func (iw *IndexWord) IndexRecordIDForWordRecordSecondaryKey() {
	var indexRequest = pehrest.HashSetValueReq{
		Type:       gs.RequestTypeBroadcast,
		IndexKey:   iw.hashWordRecordSecondaryKeyForRecordID(),
		IndexValue: iw.RecordID,
	}
	var err = psdk.HashSetValue(&indexRequest)
	if err != nil {
		// TODO::: we must retry more due to record wrote successfully!
	}
}

func (iw *IndexWord) hashWordRecordSecondaryKeyForRecordID() (hash [32]byte) {
	const field = "WordRecordSecondaryKey"
	var buf = make([]byte, 40+len(field)+len(iw.Word)) // 8+32
	syllab.SetUInt64(buf, 0, indexWordStructureID)
	copy(buf[8:], iw.RecordSecondaryKey[:])
	copy(buf[40:], field)
	copy(buf[40+len(field):], iw.Word)
	return sha512.Sum512_256(buf)
}

// IndexRecordIDForWordRecordOwnerID save RecordID chain for Word+RecordOwnerID
// Don't call in update to an exiting record!
func (iw *IndexWord) IndexRecordIDForWordRecordOwnerID() {
	var indexRequest = pehrest.HashSetValueReq{
		Type:       gs.RequestTypeBroadcast,
		IndexKey:   iw.hashWordRecordOwnerIDForRecordID(),
		IndexValue: iw.RecordID,
	}
	var err = psdk.HashSetValue(&indexRequest)
	if err != nil {
		// TODO::: we must retry more due to record wrote successfully!
	}
}

func (iw *IndexWord) hashWordRecordOwnerIDForRecordID() (hash [32]byte) {
	const field = "WordRecordOwnerID"
	var buf = make([]byte, 40+len(field)+len(iw.Word)) // 8+32
	syllab.SetUInt64(buf, 0, indexWordStructureID)
	copy(buf[8:], iw.RecordOwnerID[:])
	copy(buf[40:], field)
	copy(buf[40+len(field):], iw.Word)
	return sha512.Sum512_256(buf)
}

/*
	-- Syllab Encoder & Decoder --
*/

func (iw *IndexWord) syllabDecoder(buf []byte) (err *er.Error) {
	if uint32(len(buf)) < iw.syllabStackLen() {
		err = syllab.ErrSyllabDecodeSmallSlice
		return
	}
	var i, add, ln uint32 // index, address and len of strings, slices, maps, ...

	copy(iw.RecordID[:], buf[0:])
	iw.RecordStructureID = syllab.GetUInt64(buf, 32)
	iw.RecordSize = syllab.GetUInt64(buf, 40)
	iw.WriteTime = etime.Time(syllab.GetInt64(buf, 48))
	copy(iw.OwnerAppID[:], buf[56:])

	iw.Word = syllab.UnsafeGetString(buf, 88)
	iw.RecordStructure = syllab.GetUInt64(buf, 96)
	copy(iw.RecordPrimaryKey[:], buf[104:])
	copy(iw.RecordSecondaryKey[:], buf[136:])
	copy(iw.RecordOwnerID[:], buf[168:])

	add = syllab.GetUInt32(buf, 200)
	ln = syllab.GetUInt32(buf, 204)
	iw.Tokens = make([]WordToken, ln)
	for i = 0; i < ln; i++ {
		iw.Tokens[i].syllabDecoder(buf, add)
		add += uint32(iw.Tokens[i].syllabLen())
	}
	return
}

func (iw *IndexWord) syllabEncoder() (buf []byte) {
	buf = make([]byte, iw.syllabLen())
	var hi uint32 = iw.syllabStackLen() // Heap index || Stack size!

	copy(buf[0:], iw.RecordID[:])
	syllab.SetUInt64(buf, 32, iw.RecordStructureID)
	syllab.SetUInt64(buf, 40, iw.RecordSize)
	syllab.SetInt64(buf, 48, int64(iw.WriteTime))
	copy(buf[56:], iw.OwnerAppID[:])

	hi = syllab.SetString(buf, iw.Word, 88, hi)
	syllab.SetUInt64(buf, 96, iw.RecordStructure)
	copy(buf[104:], iw.RecordPrimaryKey[:])
	copy(buf[136:], iw.RecordSecondaryKey[:])
	copy(buf[168:], iw.RecordOwnerID[:])

	syllab.SetUInt32(buf, 200, hi)
	syllab.SetUInt32(buf, 204, uint32(len(iw.Tokens)))
	for i := 0; i < len(iw.Tokens); i++ {
		iw.Tokens[i].syllabEncoder(buf, hi, 0)
		hi += uint32(iw.Tokens[i].syllabLen())
	}
	return
}

func (iw *IndexWord) syllabStackLen() (ln uint32) {
	ln = 208
	ln += uint32(len(iw.Tokens)) * iw.Tokens[0].syllabStackLen()
	return
}

func (iw *IndexWord) syllabHeapLen() (ln uint32) {
	ln += uint32(len(iw.Word))
	// ln += uint32(len(iw.Tokens)) * iw.Tokens[0].syllabHeapLen()
	return
}

func (iw *IndexWord) syllabLen() (ln uint64) {
	return uint64(iw.syllabStackLen() + iw.syllabHeapLen())
}

/*
	-- Syllab Encoder & Decoder --
*/

func (wt *WordToken) syllabDecoder(buf []byte, stackIndex uint32) {
	copy(wt.RecordID[:], buf[stackIndex:])
	wt.RecordFieldID = syllab.GetUInt8(buf, stackIndex+32)
	wt.WordType = WordType(syllab.GetUInt16(buf, stackIndex+33))
	wt.WordOffsetInSentence = syllab.GetUInt64(buf, stackIndex+35)
	wt.WordOffsetInText = syllab.GetUInt64(buf, stackIndex+43)
	wt.OffsetInSentence = syllab.GetUInt64(buf, stackIndex+51)
	wt.OffsetInText = syllab.GetUInt64(buf, stackIndex+59)
}

func (wt *WordToken) syllabEncoder(buf []byte, stackIndex, heapIndex uint32) (nextHeapAddr uint32) {
	copy(buf[0:], wt.RecordID[stackIndex:])
	syllab.SetUInt8(buf, stackIndex+32, wt.RecordFieldID)
	syllab.SetUInt16(buf, stackIndex+33, uint16(wt.WordType))
	syllab.SetUInt64(buf, stackIndex+35, wt.WordOffsetInSentence)
	syllab.SetUInt64(buf, stackIndex+43, wt.WordOffsetInText)
	syllab.SetUInt64(buf, stackIndex+51, wt.OffsetInSentence)
	syllab.SetUInt64(buf, stackIndex+59, wt.OffsetInText)
	return heapIndex
}

func (wt *WordToken) syllabStackLen() (ln uint32) {
	return 67
}

func (wt *WordToken) syllabHeapLen() (ln uint32) {
	return
}

func (wt *WordToken) syllabLen() (ln uint64) {
	return uint64(wt.syllabStackLen() + wt.syllabHeapLen())
}
