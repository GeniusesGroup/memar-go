/* For license and copyright information please see LEGAL file in repository */

package ganjine

import (
	"time"
	"unsafe"

	persiaos "../PersiaOS-sdk"
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
	RecordNumber uint64
	Padding      uint64     // just add to have full 4K cluster >> 4096 = hashIndexHeaderSize + (125*32)
	RecordsID    [][32]byte // Store by time of added to index
}

const (
	// hashIndexHeaderSize indicate start byte of first RecordID
	hashIndexHeaderSize = 72 + 8 + 8 + (4 + 4)

	firstBlockNumber  = 125
	commonBlockNumber = 128
)

// GetHeader get data from storage and decode to given hi
func (hi *HashIndex) GetHeader() (err error) {
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

// GetIndexRecords return related records ID to given index with offset and limit!
func (hi *HashIndex) GetIndexRecords(offset, limit uint64) (recordsID [][32]byte, err error) {
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

	// If offset set to higher than exiting number of record always return last 128 records!
	if hi.RecordNumber < offset {
		if hi.RecordNumber > 128 {
			offset = hi.RecordNumber - 128
			limit = 128
		} else {
			offset = 0
			limit = hi.RecordNumber
		}
	}

	// if limit==0 it means returns all available RecordsID
	if limit == 0 {
		limit = hi.RecordNumber
	}

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

	return
}

// AppendRecordID add given RecordID with any logic need like shrink!
func (hi *HashIndex) AppendRecordID(recordID [32]byte) (err error) {
	// Get first cluster of record to read header data
	var record []byte
	record, err = persiaos.ReadStorageRecord(hi.RecordID, 0, hashIndexHeaderSize)
	if err != nil {
		// desire record not found. write new one!
		hi.RecordStructureID = hashIndexStructureID
		hi.RecordSize = 4096
		hi.WriteTime = time.Now().Unix()
		// hi.OwnerAppID =  Must assign by caller!
		hi.RecordNumber = 1
		hi.RecordsID = make([][32]byte, 1, firstBlockNumber*32)
		hi.RecordsID[0] = recordID
		err = persiaos.SetStorageRecord(hi.syllabEncoderFull())
		return
	}

	err = hi.syllabDecoder(record)
	if err != nil {
		return
	}

	if hi.RecordNumber == uint64(len(hi.RecordsID)) {
		record, err = persiaos.GetStorageRecord(hi.RecordID)
		if err != nil {
			return
		}

		err = hi.syllabDecoder(record)
		if err != nil {
			return
		}

		hi.shrink()

		hi.WriteTime = time.Now().Unix()
		hi.RecordNumber++
		hi.RecordsID[hi.RecordNumber] = recordID

		err = persiaos.SetStorageRecord(hi.syllabEncoderFull())
	} else {
		persiaos.WriteStorageRecord(hi.RecordID, hashIndexHeaderSize+(hi.RecordNumber*32), recordID[:])

		hi.WriteTime = time.Now().Unix()
		hi.RecordNumber++
		persiaos.WriteStorageRecord(hi.RecordID, 0, hi.syllabEncoderHeader())
	}

	return
}

func (hi *HashIndex) shrink() {
	var timeNow = time.Now().Unix()
	if hi.WriteTime+(60) > timeNow {
		hi.RecordsID = append(hi.RecordsID, make([][32]byte, 20*commonBlockNumber)...)
	} else if hi.WriteTime+(60*60) > timeNow {
		hi.RecordsID = append(hi.RecordsID, make([][32]byte, 8*commonBlockNumber)...)
	} else {
		hi.RecordsID = append(hi.RecordsID, make([][32]byte, commonBlockNumber)...)
	}
}

// DeleteRecordID use to delete given record ID form given indexHash!
func (hi *HashIndex) DeleteRecordID(recordID [32]byte) (err error) {
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
			hi.RecordNumber--
		}
	}
	for i := uint64(len(hi.RecordsID) - 1); i >= hi.RecordNumber; i-- {
		hi.RecordsID[i] = [32]byte{}
	}

	err = persiaos.SetStorageRecord(hi.syllabEncoderFull())
	return
}

// DeleteRecordsID use to delete given records ID form given indexHash!
func (hi *HashIndex) DeleteRecordsID(recordsID [][32]byte) (err error) {
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
				hi.RecordNumber--
			}
		}
	}
	for i := uint64(len(hi.RecordsID) - 1); i >= hi.RecordNumber; i-- {
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
	hi.RecordNumber = uint64(buf[72]) | uint64(buf[73])<<8 | uint64(buf[74])<<16 | uint64(buf[75])<<24 |
		uint64(buf[76])<<32 | uint64(buf[77])<<40 | uint64(buf[78])<<48 | uint64(buf[79])<<56
	// hi.Padding =  syllab.GetUInt64(buf[80:])
	// var recordsIDAdd = syllab.GetUInt32(buf[88:])
	var recordsIDLen uint32 = uint32(buf[92]) | uint32(buf[93])<<8 | uint32(buf[94])<<16 | uint32(buf[95])<<24
	buf = buf[hashIndexHeaderSize:]
	hi.RecordsID = *(*[][32]byte)(unsafe.Pointer(&buf))
	hi.RecordsID = hi.RecordsID[:recordsIDLen]

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
	buf[72] = byte(hi.RecordNumber)
	buf[73] = byte(hi.RecordNumber >> 8)
	buf[74] = byte(hi.RecordNumber >> 16)
	buf[75] = byte(hi.RecordNumber >> 24)
	buf[76] = byte(hi.RecordNumber >> 32)
	buf[77] = byte(hi.RecordNumber >> 40)
	buf[78] = byte(hi.RecordNumber >> 48)
	buf[79] = byte(hi.RecordNumber >> 56)
	// buf[80] = byte(hi.Padding)
	// buf[81] = byte(hi.Padding >> 8)
	// buf[82] = byte(hi.Padding >> 16)
	// buf[83] = byte(hi.Padding >> 24)
	// buf[84] = byte(hi.Padding >> 32)
	// buf[85] = byte(hi.Padding >> 40)
	// buf[86] = byte(hi.Padding >> 48)
	// buf[87] = byte(hi.Padding >> 56)
	// const recordsIDAdd = 96
	// buf[88] = byte(recordsIDAdd)
	// buf[89] = byte(recordsIDAdd >> 8)
	// buf[90] = byte(recordsIDAdd >> 16)
	// buf[91] = byte(recordsIDAdd >> 24)
	var ln = len(hi.RecordsID)
	buf[92] = byte(ln)
	buf[93] = byte(ln >> 8)
	buf[94] = byte(ln >> 16)
	buf[95] = byte(ln >> 24)
}
