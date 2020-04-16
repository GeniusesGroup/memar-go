/* For license and copyright information please see LEGAL file in repository */

package persiadb

import (
	"unsafe"

	persiaos "../PersiaOS-sdk"
)

/*
We can simple use this structure but It isn't efficient neither computing(RAM,CPU) nor storage devices!
type HashIndex struct {
	HashIndex map[[32]byte][][32]byte // Store by time of added to index
}
*/

// DefaultHashIndex use to store hash index data of DB instance!
var DefaultHashIndex HashIndex

// HashIndex is standard structure to store any hash byte index!!
// It is simple secondary index e.g. hash("user@email.com")
type HashIndex struct {
	/* Common header data */
	Checksum          [32]byte
	RecordID          [16]byte
	RecordSize        uint64
	RecordStructureID uint64
	WriteTime         int64
	OwnerAppID        [16]byte
	/* Unique data */
	IndexBucketsNumber      uint32
	StartOfIndexBuckets     uint64
	StartOfRecordsIDBuckets uint64        // Location of first cluster of
	ReAllocateTime          uint64        // Time to know for some algorithm like how many bit add to hash bit number!
	Status                  uint8         // Stable, Splitting, Re-Allocate
	OldHashIndex            *HashIndex    // Just exist in Re-Allocate proccess!
	ZeroPadding             []byte        // Just add to Buckets start at first of each storage cluster to improve write efficency!
	IndexBuckets            []IndexBucket // Cache them all to RAM, If enough RAM available to app!
	RecordsIDBuckets        []RecordsIDBucket
}

// IndexBucket store base hash index data. It get 384byte storage space!
type IndexBucket struct {
	IndexHash             [8][32]byte // Any hash making algorithm e.g. SHA-256, ... . Key of hash table!
	RecordsIDBucketOffset [8]uint64   // Start cluster location of bucket in this record!
	RecordNumber          [8]uint64   // use to know number of RecordsIDBucket and RecordID inside it!
}

// RecordsIDBucket store 4096Byte (4K cluster storage engines) of RecordID equal 256 number of it!
// Maybe some slot not use and be free for long time but it is necessary for performance!
type RecordsIDBucket struct {
	RecordsID [256][16]byte // Store by time of added to index
}

// New use to make new secondary index record!
func (hi *HashIndex) New() {
	// In large cluster or large size of this record it is better to write it in dedicate storage device if exist!

	// calculate IndexBucketBitNumber by some factor like available storage cpacity, RAM, ...
	// In most cases if enough RAM availabe start with 20 bit, because reallocate is not good and easy!
	// Below 16 bit is good just for one node cluster not real multi node one!
	// Higher 28 bit is not efficient for one node because more than 128GB RAM needed for index and other server proccesses!
	// - bit>> buckets num.	>> IndexHash		>> IndexBucket size(MB) >> Min Record size(MB) ~
	// - 8	>> 256			>> 2,048			>> 0.09375				>> 8.09375
	// - 16	>> 65536		>> 524,288			>> 24					>> 2,072
	// - 20	>> 1,048,576	>> 8,388,608		>> 384					>> 33,152
	// - 24	>> 16,777,216	>> 134,217,728		>> 6,144				>> 530,432
	// - 28	>> 268,435,456	>> 2,147,483,648	>> 98,304				>> 8,486,912
	// - 32	>> 4,294,967,296>> 34,359,738,368	>> 1,572,864			>> 135,790,592
}

// Open use to get exiting secondary index record data without RecordsIDBuckets part!
// We don't cache RecordsIDBuckets in RAM, Because they use rarely and it isn't hardware efficency!
func (hi *HashIndex) Open(recordID [16]byte) error {
	if hi == nil {
		hi = &DefaultHashIndex
	}

	// Get first cluster of record to read header data
	var buckets [][4096]byte
	var err error
	buckets, err = persiaos.ReadStorageRecord(recordID, 0, 1)
	if err != nil {
		return ErrHashIndexRecordNotExist
	}

	// Copy header data to structure!
	var header [4096]byte = buckets[0]
	copy(hi.Checksum[:], header[0:])
	copy(hi.RecordID[:], header[32:])
	hi.RecordSize = uint64(header[48]) | uint64(header[49])<<8 | uint64(header[50])<<16 | uint64(header[51])<<24 |
		uint64(header[52])<<32 | uint64(header[53])<<40 | uint64(header[54])<<48 | uint64(header[55])<<56
	hi.RecordStructureID = uint64(header[56]) | uint64(header[57])<<8 | uint64(header[58])<<16 | uint64(header[59])<<24 |
		uint64(header[60])<<32 | uint64(header[61])<<40 | uint64(header[62])<<48 | uint64(header[63])<<56
	hi.WriteTime = int64(header[56]) | int64(header[57])<<8 | int64(header[58])<<16 | int64(header[59])<<24 |
		int64(header[60])<<32 | int64(header[61])<<40 | int64(header[62])<<48 | int64(header[63])<<56
	copy(hi.OwnerAppID[:], header[64:])

	// Copy unique data of hash index record
	hi.IndexBucketsNumber = uint32(header[80]) | uint32(header[81])<<8 | uint32(header[82])<<16 | uint32(header[83])<<24
	hi.StartOfIndexBuckets = uint64(header[84]) | uint64(header[85])<<8 | uint64(header[86])<<16 | uint64(header[87])<<24 |
		uint64(header[88])<<32 | uint64(header[89])<<40 | uint64(header[90])<<48 | uint64(header[91])<<56
	hi.StartOfRecordsIDBuckets = uint64(header[92]) | uint64(header[93])<<8 | uint64(header[94])<<16 | uint64(header[95])<<24 |
		uint64(header[96])<<32 | uint64(header[97])<<40 | uint64(header[98])<<48 | uint64(header[99])<<56

	// Read IndexBuckets clusters of HashIndex record
	buckets, err = persiaos.ReadStorageRecord(recordID, hi.StartOfIndexBuckets, hi.StartOfRecordsIDBuckets)
	if err != nil {
		return ErrHashIndexRecordManipulated
	}

	// TODO : Below code is not write efficient, bucket overlap on clusters on writes! 384Byte*10bucket = 3840Byte < 4096Byte
	hi.IndexBuckets = *(*[]IndexBucket)(unsafe.Pointer(&buckets))

	return nil
}

// Close use to calculate & write checksum of secondary index record and nil si!
func (hi *HashIndex) Close() {}

// reAllocate use to expand||shrink index buckets capacity for exiting secondary index record!
func (hi *HashIndex) reAllocate() {
	// By ReAllocateTime we decide how many bit add to IndexBucketBitNumber
	// and re-allocate capacity to size that don't need to re-allocate tomorrow of this proccess!
	// It is non blocking opperation and call by set index
}

// GetIndexRecords use to get related records ID to given index with offset and limit!
func (hi *HashIndex) GetIndexRecords(indexHash [32]byte, offset, limit uint64) (recordsID [][256][16]byte) {
	var off, num = hi.GetRecordsIDBucketInfo(indexHash)
	// If offset set to higher than exiting number of record always means dev want last 256 records!
	if num < offset {
		offset = num - 256
	}

	var startBucket uint64 = off + (offset / 256)
	var endBucket uint64 = off + ((offset + limit) / 256)
	var buckets [][4096]byte
	buckets, _ = persiaos.ReadStorageRecord(hi.RecordID, startBucket, endBucket)

	// We know [256][16]byte == [4096]byte is same size but Go don't let return buckets simply, so:
	return *(*[][256][16]byte)(unsafe.Pointer(&buckets))
}

func (hi *HashIndex) getIndexBucketID(indexHash [32]byte) uint32 {
	var bucketID uint32 = uint32(indexHash[28]) | uint32(indexHash[29])<<8 | uint32(indexHash[30])<<16 | uint32(indexHash[31])<<24
	// use % of just 32bit right of indexHash to detect bucketID!
	return bucketID % hi.IndexBucketsNumber
}

// GetRecordsIDBucketInfo use to get number of recordsID register for specific IndexHash
func (hi *HashIndex) GetRecordsIDBucketInfo(indexHash [32]byte) (offset, number uint64) {
	var indexBucket IndexBucket = hi.IndexBuckets[hi.getIndexBucketID(indexHash)]
	for i := 0; i < 8; i++ {
		if indexBucket.IndexHash[i] == indexHash {
			return indexBucket.RecordsIDBucketOffset[i], indexBucket.RecordNumber[i]
		}
	}

	// Check if HashIndex is in Re-Allocate proccess!
	if hi.OldHashIndex != nil {
		return hi.OldHashIndex.GetRecordsIDBucketInfo(indexHash)
	}

	// if indexHash not found it return zero
	return 0, 0
}

// SetIndexRecord set given record ID in given index hash!
func (hi *HashIndex) SetIndexRecord(indexHash [32]byte, recordID [16]byte) {
	var bucketID uint32 = hi.getIndexBucketID(indexHash)
	var indexBucket IndexBucket = hi.IndexBuckets[bucketID]
	for i := 0; i < 8; i++ {
		if indexBucket.IndexHash[i] == indexHash {
			indexBucket.RecordNumber[i]++
			hi.writeIndexBucket(indexBucket, bucketID)
			hi.writeIndexRecord(recordID, indexBucket.RecordsIDBucketOffset[i], indexBucket.RecordNumber[i])
		} else if indexBucket.IndexHash[i] == [32]byte{0} {
			indexBucket.IndexHash[i] = indexHash
			indexBucket.RecordNumber[i]++
			hi.writeIndexBucket(indexBucket, bucketID)
			// TODO: Get first free bucket offset
			hi.writeIndexRecord(recordID, indexBucket.RecordsIDBucketOffset[i], indexBucket.RecordNumber[i])
			// If indexHash not exist and after set this indexHash just one slot exist in IndexBucket we must call ReAllocate method
			if i >= 6 {
				hi.reAllocate()
			}
		}
	}
}

func (hi *HashIndex) writeIndexBucket(indexBucket IndexBucket, bucketID uint32) {
	var bucketLoc uint64 = hi.StartOfIndexBuckets + uint64(bucketID)
	var data = [][4096]byte{*(*[4096]byte)(unsafe.Pointer(&indexBucket))}
	persiaos.WriteStorageRecord(hi.RecordID, bucketLoc, data)
}

// writeIndexRecord set given record ID in given index hash!
func (hi *HashIndex) writeIndexRecord(recordID [16]byte, offset, number uint64) {
	// Write changes to related storage device
}

// DeleteIndexHash use to delete given IndexHash and all related RecordsID!
func (hi *HashIndex) DeleteIndexHash(indexHash [32]byte) {
	return
}

// DeleteRecordID use to delete given record ID form given indexHash!
func (hi *HashIndex) DeleteRecordID(indexHash [32]byte, recordID [][16]byte) {
	return
}
