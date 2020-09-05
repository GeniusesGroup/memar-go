/* For license and copyright information please see LEGAL file in repository */

package ganjine

import (
	"unsafe"

	persiaos "../PersiaOS-sdk"
	etime "../earth-time"
)

const hashIndexStructureID uint64 = 2096939569836997748

// HashIndex is standard structure to store any hash byte index!!
// It is simple secondary index e.g. hash("user@email.com")
type HashIndex struct {
	/* Common header data */
	RecordID          [32]byte
	RecordStructureID uint64
	RecordSize        uint64
	WriteTime         int64
	OwnerAppID        [16]byte

	/* Unique data */
	ExpandTime      int64
	RecordsNumber   uint64     // RecordsID len
	RecordsCapacity uint64     // RecordsID cap.
	RecordsID       [][32]byte // Store by time of added to index
}

const (
	// hashIndexHeaderSize indicate start byte of first RecordID
	hashIndexHeaderSize = 96 // 72 + 8 + 8 + (4 + 4)

	firstBlockNumber  = 125 // to have first full 4K cluster >> 4096 = hashIndexHeaderSize + (125*32)
	commonBlockNumber = 128
)

// ReadHeader get needed data from storage and decode to given hi without RecordsID array data
func (hi *HashIndex) ReadHeader() (err error) {
	// Get first cluster of record to read header data
	var header []byte
	header, err = persiaos.ReadStorageRecord(hi.RecordID, 0, hashIndexHeaderSize)
	if err != nil {
		return
	}

	err = hi.syllabDecoder(header)
	if err != nil {
		return
	}

	return
}

// Pop return last RecordID in given index and delete it from index!
func (hi *HashIndex) Pop() (recordID [32]byte, err error) {
	// Get first cluster of record to read header data
	var header []byte
	header, err = persiaos.ReadStorageRecord(hi.RecordID, 0, 4096)
	if err != nil {
		return
	}

	err = hi.syllabDecoder(header)
	if err != nil {
		return
	}

	if hi.RecordsNumber > firstBlockNumber {
		hi.RecordsNumber--
		var record []byte
		var recordOffset = hashIndexHeaderSize + (hi.RecordsNumber * 32)
		record, err = persiaos.ReadStorageRecord(hi.RecordID, recordOffset, 32)
		if err != nil {
			return
		}
		copy(recordID[:], record[:])
		err = persiaos.WriteStorageRecord(hi.RecordID, 0, hi.syllabEncoderHeader())
		err = persiaos.WriteStorageRecord(hi.RecordID, recordOffset, make([]byte, 32))
	} else {
		if hi.RecordsNumber > 1 {
			hi.RecordsNumber--
			recordID = hi.RecordsID[hi.RecordsNumber]
			hi.RecordsID[hi.RecordsNumber] = [32]byte{}
			err = persiaos.SetStorageRecord(hi.syllabEncoderFull())
		} else {
			err = persiaos.DeleteStorageRecord(hi.RecordID)
			// err = ErrEndOfIndexRecord
		}
	}

	return
}

// Peek return last recordID pushed to given index. unlike Pop() don't delete it from index!
func (hi *HashIndex) Peek() (recordID [32]byte, err error) {
	// Get first cluster of record to read header data
	var header []byte
	header, err = persiaos.ReadStorageRecord(hi.RecordID, 0, 4096)
	if err != nil {
		return
	}

	err = hi.syllabDecoder(header)
	if err != nil {
		return
	}

	if hi.RecordsNumber > firstBlockNumber {
		var record []byte
		var recordOffset = hashIndexHeaderSize + (hi.RecordsNumber-1 * 32)
		record, err = persiaos.ReadStorageRecord(hi.RecordID, recordOffset, 32)
		if err != nil {
			return
		}
		copy(recordID[:], record[:])
	} else {
		recordID = hi.RecordsID[hi.RecordsNumber-1]
	}

	return
}

// Get return related records ID to given index with offset and limit!
func (hi *HashIndex) Get(offset, limit uint64) (recordsID [][32]byte, err error) {
	// Get first cluster of record to read header data
	var header []byte
	header, err = persiaos.ReadStorageRecord(hi.RecordID, 0, 4096)
	if err != nil {
		return
	}

	err = hi.syllabDecoder(header)
	if err != nil {
		return
	}

	// If offset set to higher than exiting number of record always return last 32 records!
	if hi.RecordsNumber < offset {
		if hi.RecordsNumber > 32 {
			offset = hi.RecordsNumber - 32
			limit = 32
		} else {
			offset = 0
			limit = hi.RecordsNumber
		}
	}

	if limit == 0 {
		if offset == 0 {
			// it means returns all available RecordsID
			limit = hi.RecordsNumber
		} else {
			// it means returns all available RecordsID after offset
			limit = hi.RecordsNumber - offset
		}
	}

	if offset+limit > firstBlockNumber {
		offset = hashIndexHeaderSize + (offset * 32)
		limit = offset + (limit * 32)
		var record []byte
		record, err = persiaos.ReadStorageRecord(hi.RecordID, offset, limit)
		if err != nil {
			return
		}

		// We know (128*[32]byte) == [4096]byte is same size but Go don't let return buckets simply, so:
		recordsID = *(*[][32]byte)(unsafe.Pointer(&record))
		recordsID = recordsID[:limit]
	} else {
		recordsID = hi.RecordsID[offset:limit]
	}

	return
}

// Push add given RecordID to then end of given hash index!
func (hi *HashIndex) Push(recordID [32]byte) (err error) {
	var timeNow = etime.Now()

	// Get first cluster of record to read header data
	var record []byte
	record, err = persiaos.ReadStorageRecord(hi.RecordID, 0, hashIndexHeaderSize)
	if err != nil {
		// desire record not found. write new one!
		hi.RecordStructureID = hashIndexStructureID
		hi.RecordSize = 4096
		hi.WriteTime = timeNow
		// hi.OwnerAppID =  Must assign by caller!
		hi.RecordsNumber = 1
		hi.RecordsCapacity = firstBlockNumber
		hi.RecordsID = make([][32]byte, 1, firstBlockNumber*32)
		hi.RecordsID[0] = recordID
		err = persiaos.SetStorageRecord(hi.syllabEncoderFull())
		return
	}

	err = hi.syllabDecoder(record)
	if err != nil {
		return
	}

	// Check and expand record if needed
	if hi.RecordsNumber == hi.RecordsCapacity {
		hi.ExpandTime = timeNow
		var expandNumber = hi.calculateExpandNumber(timeNow)
		hi.RecordsCapacity += expandNumber
		err = persiaos.ExpandStorageRecord(hi.RecordID, expandNumber*32)
		if err != nil {
			return
		}
	}

	persiaos.WriteStorageRecord(hi.RecordID, hashIndexHeaderSize+(hi.RecordsNumber*32), recordID[:])

	hi.WriteTime = timeNow
	hi.RecordsNumber++
	err = persiaos.WriteStorageRecord(hi.RecordID, 0, hi.syllabEncoderHeader())
	return
}

// TODO::: Improve expand algorithm
func (hi *HashIndex) calculateExpandNumber(timeNow int64) (expandNumber uint64) {
	var ln = uint64(len(hi.RecordsID))
	var totalBlockNumber = (ln + 3) / commonBlockNumber
	if hi.ExpandTime+(60) > timeNow {
		if totalBlockNumber < 1000 {
			expandNumber = totalBlockNumber * commonBlockNumber
		} else if totalBlockNumber < 10000 {
			expandNumber = (totalBlockNumber / 2) * commonBlockNumber
		} else {
			expandNumber = (totalBlockNumber / 4) * commonBlockNumber
		}
	} else if hi.ExpandTime+(60*60) > timeNow {
		expandNumber = 100 * commonBlockNumber
	} else if hi.ExpandTime+(24*60*60) > timeNow {
		expandNumber = 50 * commonBlockNumber
	} else {
		expandNumber = 10 * commonBlockNumber
	}
	return
}

// Append add given RecordID with any logic need like expand!
func (hi *HashIndex) Append(recordID ...[32]byte) (err error) {
	return
}

// Delete use to delete given record ID form given indexHash!
func (hi *HashIndex) Delete(recordID [32]byte) (err error) {
	var record []byte
	record, err = persiaos.GetStorageRecord(hi.RecordID)
	if err != nil {
		return
	}

	err = hi.syllabDecoder(record)
	if err != nil {
		return
	}

	var ln = len(hi.RecordsID)
	for i := 0; i < ln; i++ {
		if hi.RecordsID[i] == recordID {
			copy(hi.RecordsID[i:], hi.RecordsID[i+1:])
			hi.RecordsNumber--
		}
	}
	for i := uint64(len(hi.RecordsID) - 1); i >= hi.RecordsNumber; i-- {
		hi.RecordsID[i] = [32]byte{}
	}

	err = persiaos.SetStorageRecord(hi.syllabEncoderFull())
	return
}

// Deletes use to delete given records ID form given indexHash!
func (hi *HashIndex) Deletes(recordsID [][32]byte) (err error) {
	// Get first cluster of record to read header data
	var record []byte
	record, err = persiaos.GetStorageRecord(hi.RecordID)
	if err != nil {
		return
	}

	err = hi.syllabDecoder(record)
	if err != nil {
		return
	}

	var ln = len(hi.RecordsID)
	var inputLn = len(recordsID)
	for i := 0; i < ln; i++ {
		for j := 0; j < inputLn; j++ {
			if hi.RecordsID[i] == recordsID[j] {
				copy(hi.RecordsID[i:], hi.RecordsID[i+1:])
				hi.RecordsNumber--
			}
		}
	}
	for i := uint64(len(hi.RecordsID) - 1); i >= hi.RecordsNumber; i-- {
		hi.RecordsID[i] = [32]byte{}
	}

	err = persiaos.SetStorageRecord(hi.syllabEncoderFull())
	return
}

func (hi *HashIndex) syllabDecoder(buf []byte) (err error) {
	// Copy header data to structure!
	hi.RecordStructureID = uint64(buf[32]) | uint64(buf[33])<<8 | uint64(buf[34])<<16 | uint64(buf[35])<<24 |
		uint64(buf[36])<<32 | uint64(buf[37])<<40 | uint64(buf[38])<<48 | uint64(buf[39])<<56
	hi.RecordSize = uint64(buf[40]) | uint64(buf[41])<<8 | uint64(buf[42])<<16 | uint64(buf[43])<<24 |
		uint64(buf[44])<<32 | uint64(buf[45])<<40 | uint64(buf[46])<<48 | uint64(buf[47])<<56
	hi.WriteTime = int64(buf[48]) | int64(buf[49])<<8 | int64(buf[50])<<16 | int64(buf[51])<<24 |
		int64(buf[52])<<32 | int64(buf[53])<<40 | int64(buf[54])<<48 | int64(buf[55])<<56
	copy(hi.OwnerAppID[:], buf[56:])

	// Copy unique data of hash index record
	hi.ExpandTime = int64(buf[72]) | int64(buf[73])<<8 | int64(buf[74])<<16 | int64(buf[75])<<24 |
		int64(buf[76])<<32 | int64(buf[77])<<40 | int64(buf[78])<<48 | int64(buf[79])<<56
	hi.RecordsNumber = uint64(buf[80]) | uint64(buf[81])<<8 | uint64(buf[82])<<16 | uint64(buf[83])<<24 |
		uint64(buf[84])<<32 | uint64(buf[85])<<40 | uint64(buf[86])<<48 | uint64(buf[87])<<56
	// use recordsIDAdd space to expand recordsIDLen to Uint64
	hi.RecordsCapacity = uint64(buf[88]) | uint64(buf[89])<<8 | uint64(buf[90])<<16 | uint64(buf[91])<<24 |
		uint64(buf[92])<<32 | uint64(buf[93])<<40 | uint64(buf[94])<<48 | uint64(buf[95])<<56
	buf = buf[hashIndexHeaderSize:]
	hi.RecordsID = *(*[][32]byte)(unsafe.Pointer(&buf))
	hi.RecordsID = hi.RecordsID[:hi.RecordsNumber]

	if hi.RecordStructureID != hashIndexStructureID {
		err = ErrHashIndexRecordNotValid
		return
	}

	return
}

func (hi *HashIndex) syllabEncoderFull() (buf []byte) {
	var temp = append(make([][32]byte, 3), hi.RecordsID...)
	buf = *(*[]byte)(unsafe.Pointer(&temp))
	hi.syllabEncoderCopy(buf)
	return
}

func (hi *HashIndex) syllabEncoderHeader() (buf []byte) {
	buf = make([]byte, hashIndexHeaderSize)
	hi.syllabEncoderCopy(buf)
	return
}

func (hi *HashIndex) syllabEncoderCopy(buf []byte) {
	// Copy header data to buf!
	copy(buf[:], hi.RecordID[:])
	buf[32] = byte(hi.RecordStructureID)
	buf[33] = byte(hi.RecordStructureID >> 8)
	buf[34] = byte(hi.RecordStructureID >> 16)
	buf[35] = byte(hi.RecordStructureID >> 24)
	buf[36] = byte(hi.RecordStructureID >> 32)
	buf[37] = byte(hi.RecordStructureID >> 40)
	buf[38] = byte(hi.RecordStructureID >> 48)
	buf[39] = byte(hi.RecordStructureID >> 56)
	buf[40] = byte(hi.RecordSize)
	buf[41] = byte(hi.RecordSize >> 8)
	buf[42] = byte(hi.RecordSize >> 16)
	buf[43] = byte(hi.RecordSize >> 24)
	buf[44] = byte(hi.RecordSize >> 32)
	buf[45] = byte(hi.RecordSize >> 40)
	buf[46] = byte(hi.RecordSize >> 48)
	buf[47] = byte(hi.RecordSize >> 56)
	buf[48] = byte(hi.WriteTime)
	buf[49] = byte(hi.WriteTime >> 8)
	buf[50] = byte(hi.WriteTime >> 16)
	buf[51] = byte(hi.WriteTime >> 24)
	buf[52] = byte(hi.WriteTime >> 32)
	buf[53] = byte(hi.WriteTime >> 40)
	buf[54] = byte(hi.WriteTime >> 48)
	buf[55] = byte(hi.WriteTime >> 56)
	copy(buf[56:], hi.OwnerAppID[:])

	// Copy unique data of hash index record to buf
	buf[72] = byte(hi.ExpandTime)
	buf[73] = byte(hi.ExpandTime >> 8)
	buf[74] = byte(hi.ExpandTime >> 16)
	buf[75] = byte(hi.ExpandTime >> 24)
	buf[76] = byte(hi.ExpandTime >> 32)
	buf[77] = byte(hi.ExpandTime >> 40)
	buf[78] = byte(hi.ExpandTime >> 48)
	buf[79] = byte(hi.ExpandTime >> 56)
	buf[80] = byte(hi.RecordsNumber)
	buf[81] = byte(hi.RecordsNumber >> 8)
	buf[82] = byte(hi.RecordsNumber >> 16)
	buf[83] = byte(hi.RecordsNumber >> 24)
	buf[84] = byte(hi.RecordsNumber >> 32)
	buf[85] = byte(hi.RecordsNumber >> 40)
	buf[86] = byte(hi.RecordsNumber >> 48)
	buf[87] = byte(hi.RecordsNumber >> 56)
	// use recordsIDAdd space to expand recordsIDLen to Uint64
	buf[88] = byte(hi.RecordsCapacity)
	buf[89] = byte(hi.RecordsCapacity >> 8)
	buf[90] = byte(hi.RecordsCapacity >> 16)
	buf[91] = byte(hi.RecordsCapacity >> 24)
	buf[92] = byte(hi.RecordsCapacity >> 32)
	buf[93] = byte(hi.RecordsCapacity >> 40)
	buf[94] = byte(hi.RecordsCapacity >> 48)
	buf[95] = byte(hi.RecordsCapacity >> 56)
}
