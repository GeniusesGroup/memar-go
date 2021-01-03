/* For license and copyright information please see LEGAL file in repository */

package pehrest

import (
	persiaos "../PersiaOS-sdk"
	"../achaemenid"
	"../convert"
	etime "../earth-time"
	er "../error"
	"../ganjine"
	lang "../language"
	"../syllab"
)

const indexHashStructureID uint64 = 14141993521672016749

var indexHashStructure = ganjine.DataStructure{
	ID:                14141993521672016749,
	IssueDate:         1599286551,
	ExpiryDate:        0,
	ExpireInFavorOf:   "", // Other structure name
	ExpireInFavorOfID: 0,  // Other StructureID! Handy ID or Hash of ExpireInFavorOf!
	Status:            ganjine.DataStructureStatePreAlpha,
	Structure:         IndexHash{},

	Name: map[lang.Language]string{
		lang.LanguageEnglish: "IndexHash",
	},
	Description: map[lang.Language]string{
		lang.LanguageEnglish: "Store the hash index data by 32byte key and any 32byte value",
	},
	TAGS: []string{
		"",
	},
}

// IndexHash is standard structure to store any hash byte index!!
// It is simple secondary index e.g. hash(RecordStructureID, "user@email.com")
type IndexHash struct {
	/* Common header data */
	RecordID          [32]byte
	RecordStructureID uint64
	RecordSize        uint64
	WriteTime         etime.Time
	OwnerAppID        [32]byte

	/* Unique data */
	EarlierExpandTime   etime.Time
	LastExpandTime      etime.Time
	IndexValuesNumber   uint64     // IndexValues len
	IndexValuesCapacity uint64     // IndexValues cap.
	IndexValues         [][32]byte // Can store RecordsID or any data up to 32 byte length and store by time of added to index
}

// ReadHeader get needed data from storage and decode to given ih without IndexValues array data
func (ih *IndexHash) ReadHeader() (err *er.Error) {
	// Get first cluster of record to read header data
	var header []byte
	header, err = persiaos.StorageReadRecord(ih.RecordID, indexHashStructureID, 0, uint64(ih.syllabStackLen()))
	if err != nil {
		return
	}

	err = ih.syllabDecoder(header)
	if err != nil {
		return
	}

	return
}

// Pop return last RecordID in given index and delete it from index!
func (ih *IndexHash) Pop() (recordID [32]byte, err *er.Error) {
	// Get first cluster of record to read header data
	var header []byte
	header, err = persiaos.StorageReadRecord(ih.RecordID, indexHashStructureID, 0, uint64(ih.syllabStackLen()))
	if err != nil {
		return
	}

	err = ih.syllabDecoder(header)
	if err != nil {
		return
	}

	if ih.IndexValuesNumber > 1 {
		ih.IndexValuesNumber--
		var record []byte
		var recordOffset = uint64(ih.syllabStackLen()) + (ih.IndexValuesNumber * 32)
		record, err = persiaos.StorageReadRecord(ih.RecordID, indexHashStructureID, recordOffset, 32)
		if err != nil {
			return
		}
		copy(recordID[:], record[:])

		ih.WriteTime = etime.Now()
		var buf = make([]byte, ih.syllabStackLen())
		ih.syllabEncoderHeader(buf)
		err = persiaos.StorageWriteRecord(ih.RecordID, indexHashStructureID, 0, buf)
		err = persiaos.StorageWriteRecord(ih.RecordID, indexHashStructureID, recordOffset, make([]byte, 32))
	} else {
		err = persiaos.StorageDeleteRecord(ih.RecordID, indexHashStructureID)
		// err = ErrEndOfIndexRecord
	}

	return
}

// Peek return last recordID pushed to given index. unlike Pop() it won't delete it from index!
func (ih *IndexHash) Peek() (recordID [32]byte, err *er.Error) {
	// Get first cluster of record to read header data
	var header []byte
	header, err = persiaos.StorageReadRecord(ih.RecordID, indexHashStructureID, 0, uint64(ih.syllabStackLen()))
	if err != nil {
		return
	}

	err = ih.syllabDecoder(header)
	if err != nil {
		return
	}

	var record []byte
	var recordOffset = uint64(ih.syllabStackLen()) + ((ih.IndexValuesNumber - 1) * 32)
	record, err = persiaos.StorageReadRecord(ih.RecordID, indexHashStructureID, recordOffset, 32)
	if err != nil {
		return
	}
	copy(recordID[:], record[:])

	return
}

// Get return related records ID to given index with offset and limit!
func (ih *IndexHash) Get(offset, limit uint64) (indexValues [][32]byte, err *er.Error) {
	// Get first cluster of record to read header data
	var header []byte
	header, err = persiaos.StorageReadRecord(ih.RecordID, indexHashStructureID, 0, uint64(ih.syllabStackLen()))
	if err != nil {
		err = ganjine.ErrRecordNotFound
		return
	}

	err = ih.syllabDecoder(header)
	if err != nil {
		return
	}

	if limit > ih.IndexValuesNumber {
		offset = 0
		limit = ih.IndexValuesNumber
	} else {
		if offset > ih.IndexValuesNumber {
			// If offset set to ihgher than exiting number of record always return last records by given limit!
			offset = ih.IndexValuesNumber - limit
		}

		if limit == 0 {
			if offset == 0 {
				// it means returns all available IndexValues
				limit = ih.IndexValuesNumber
			} else {
				// it means returns all available IndexValues after offset
				limit = ih.IndexValuesNumber - offset
			}
		}
	}

	offset = uint64(ih.syllabStackLen()) + (offset * 32)
	limit = limit * 32
	var record []byte
	record, err = persiaos.StorageReadRecord(ih.RecordID, indexHashStructureID, offset, limit)
	if err != nil {
		return
	}

	// We know e.g. 128*[32]byte == [4096]byte is same size but Go is safe default and don't let return simply, so:
	indexValues = convert.UnsafeByteSliceTo32ByteArraySlice(record)
	return
}

// Push add given RecordID to then end of given hash index!
func (ih *IndexHash) Push(recordID [32]byte) (err *er.Error) {
	var timeNow = etime.Now()

	// Get first cluster of record to read header data
	var record []byte
	record, err = persiaos.StorageReadRecord(ih.RecordID, indexHashStructureID, 0, uint64(ih.syllabStackLen()))
	if err != nil {
		if err.Equal(persiaos.ErrPersiaOSStorageNotExist) {
			// desire record not found. write new one!
			ih.RecordStructureID = indexHashStructureID
			ih.RecordSize = uint64(ih.syllabStackLen()) + (1 * 32)
			ih.WriteTime = timeNow
			ih.OwnerAppID = achaemenid.Server.AppID
			// ih.EarlierExpandTime = timeNow
			ih.LastExpandTime = timeNow
			ih.IndexValuesNumber = 1
			ih.IndexValuesCapacity = 1
			ih.IndexValues = make([][32]byte, 1)
			ih.IndexValues[0] = recordID
			err = persiaos.StorageSetRecord(ih.syllabEncoderFull())
		} else {
			// err =
		}
		return
	}

	err = ih.syllabDecoder(record)
	if err != nil {
		return
	}

	// Check and expand record if needed
	if ih.IndexValuesNumber >= ih.IndexValuesCapacity {
		var expandNumber = ih.calculateExpandNumber(timeNow)
		var expandSize = expandNumber * 32
		ih.RecordSize += expandSize
		ih.EarlierExpandTime = ih.LastExpandTime
		ih.LastExpandTime = timeNow
		ih.IndexValuesCapacity += expandNumber
		err = persiaos.StorageExpandRecord(ih.RecordID, indexHashStructureID, expandSize)
		if err != nil {
			return
		}
	}

	persiaos.StorageWriteRecord(ih.RecordID, indexHashStructureID, uint64(ih.syllabStackLen())+(ih.IndexValuesNumber*32), recordID[:])

	ih.WriteTime = timeNow
	ih.IndexValuesNumber++
	ih.syllabEncoderHeader(record)
	err = persiaos.StorageWriteRecord(ih.RecordID, indexHashStructureID, 0, record)
	return
}

// TODO::: Improve expand algorithm
func (ih *IndexHash) calculateExpandNumber(timeNow etime.Time) (expandNumber uint64) {
	var ln = ih.IndexValuesCapacity
	if ih.LastExpandTime-ih.EarlierExpandTime < (60 * 60) { // Expanded twice in less than 60 minutes
		if timeNow < ih.LastExpandTime+(60*60) { // Last expand earlier than 60 minutes
			expandNumber = ln
		} else if timeNow < ih.LastExpandTime+(24*60*60) { // Last expand earlier than 1 day
			expandNumber = ln / 2
		} else { // Last expand more than 1 day
			expandNumber = 8
		}
	} else if ih.LastExpandTime-ih.EarlierExpandTime < (24 * 60 * 60) { // Expanded twice in less than 1 day
		if timeNow < ih.LastExpandTime+(24*60*60) { // Last expand earlier than 1 day
			expandNumber = ln / 4
		} else { // Last expand more than 1 day
			expandNumber = 4
		}
	} else if ih.LastExpandTime-ih.EarlierExpandTime < (7 * 24 * 60 * 60) { // Expanded twice in less than 1 week
		if timeNow < ih.LastExpandTime+(24*60*60) { // Last expand earlier than 1 day
			expandNumber = ln / 8
		} else { // Last expand more than 1 day
			expandNumber = 1
		}
	}
	if expandNumber == 0 { // Usually means expanded twice in more than 1 week
		expandNumber = 1
	}
	return
}

// Append add given RecordID with any logic need like expand!
func (ih *IndexHash) Append(recordID ...[32]byte) (err *er.Error) {
	return
}

// DeleteRecord use to delete given record ID form given indexHash!
func (ih *IndexHash) DeleteRecord() (err *er.Error) {
	// Do for i=0 as local node
	err = persiaos.StorageDeleteRecord(ih.RecordID, indexHashStructureID)
	return
}

// Delete use to delete given record ID form given indexHash!
func (ih *IndexHash) Delete(recordID [32]byte) (err *er.Error) {
	var record []byte
	record, err = persiaos.StorageGetRecord(ih.RecordID, indexHashStructureID)
	if err != nil {
		return
	}

	err = ih.syllabDecoder(record)
	if err != nil {
		return
	}

	var ln = len(ih.IndexValues)
	for i := 0; i < ln; i++ {
		if ih.IndexValues[i] == recordID {
			copy(ih.IndexValues[i:], ih.IndexValues[i+1:])
			ih.IndexValuesNumber--
		}
	}
	for i := uint64(len(ih.IndexValues) - 1); i >= ih.IndexValuesNumber; i-- {
		ih.IndexValues[i] = [32]byte{}
	}

	err = persiaos.StorageSetRecord(ih.syllabEncoderFull())
	return
}

// Deletes use to delete given records ID form given indexHash!
func (ih *IndexHash) Deletes(indexValues [][32]byte) (err *er.Error) {
	// Get first cluster of record to read header data
	var record []byte
	record, err = persiaos.StorageGetRecord(ih.RecordID, indexHashStructureID)
	if err != nil {
		return
	}

	err = ih.syllabDecoder(record)
	if err != nil {
		return
	}

	var ln = len(ih.IndexValues)
	var inputLn = len(indexValues)
	for i := 0; i < ln; i++ {
		for j := 0; j < inputLn; j++ {
			if ih.IndexValues[i] == indexValues[j] {
				copy(ih.IndexValues[i:], ih.IndexValues[i+1:])
				ih.IndexValuesNumber--
			}
		}
	}
	for i := uint64(len(ih.IndexValues) - 1); i >= ih.IndexValuesNumber; i-- {
		ih.IndexValues[i] = [32]byte{}
	}

	err = persiaos.StorageSetRecord(ih.syllabEncoderFull())
	return
}

func (ih *IndexHash) syllabDecoder(buf []byte) (err *er.Error) {
	if uint32(len(buf)) < ih.syllabStackLen() {
		err = syllab.ErrSyllabDecodeSmallSlice
		return
	}

	copy(ih.RecordID[:], buf[0:])
	ih.RecordStructureID = syllab.GetUInt64(buf, 32)
	ih.RecordSize = syllab.GetUInt64(buf, 40)
	ih.WriteTime = etime.Time(syllab.GetInt64(buf, 48))
	copy(ih.OwnerAppID[:], buf[56:])

	if ih.RecordStructureID != indexHashStructureID {
		err = ErrRecordNotValid
		return
	}

	ih.EarlierExpandTime = etime.Time(syllab.GetInt64(buf, 88))
	ih.LastExpandTime = etime.Time(syllab.GetInt64(buf, 96))
	ih.IndexValuesNumber = syllab.GetUInt64(buf, 104)
	ih.IndexValuesCapacity = syllab.GetUInt64(buf, 112)
	// Break syllab rules and don't get IndexValues Add&&len
	buf = buf[ih.syllabStackLen():]
	ih.IndexValues = convert.UnsafeByteSliceTo32ByteArraySlice(buf)
	return
}

func (ih *IndexHash) syllabEncoderFull() (buf []byte) {
	buf = make([]byte, uint64(ih.syllabLen()))
	copy(buf[ih.syllabStackLen():], convert.Unsafe32ByteArraySliceToByteSlice(ih.IndexValues))
	ih.syllabEncoderHeader(buf)
	return
}

func (ih *IndexHash) syllabEncoderHeader(buf []byte) {
	copy(buf[0:], ih.RecordID[:])
	syllab.SetUInt64(buf, 32, indexHashStructureID)
	syllab.SetUInt64(buf, 40, ih.RecordSize)
	syllab.SetInt64(buf, 48, int64(ih.WriteTime))
	copy(buf[56:], ih.OwnerAppID[:])

	syllab.SetInt64(buf, 88, int64(ih.EarlierExpandTime))
	syllab.SetInt64(buf, 96, int64(ih.LastExpandTime))
	syllab.SetUInt64(buf, 104, ih.IndexValuesNumber)
	syllab.SetUInt64(buf, 112, ih.IndexValuesCapacity)
	// Break syllab rules and don't set IndexValues Add&&len
}

func (ih *IndexHash) syllabStackLen() (ln uint32) {
	return 120 // 88 + 8 + 8 + 8 + 8 + ?  !!don't need IndexValues Add&&len!!
}

func (ih *IndexHash) syllabHeapLen() (ln uint32) {
	ln += uint32(ih.IndexValuesCapacity * 32)
	return
}

func (ih *IndexHash) syllabLen() (ln uint64) {
	return uint64(ih.syllabStackLen() + ih.syllabHeapLen())
}
