/* For license and copyright information please see LEGAL file in repository */

package generator

import (
	"bytes"
	"hash/crc64"
	"strings"
	"text/template"

	"../assets"
)

var hash64 = crc64.New(crc64.MakeTable(crc64.ISO))

// MakeNewDatastoreFile use to make new service template file!
// Pass desire StructureName in ```kebab-case``` like ```person-authentication``` in file.Name
func MakeNewDatastoreFile(file *assets.File) (err error) {
	file.FullName = file.Name + ".go"
	file.Extension = "go"

	var rn string
	var ss = strings.Split(file.Name, "-")
	for i := 0; i < len(ss); i++ {
		rn += ss[i][0:1]
	}

	var sn = strings.Title(file.Name)
	sn = strings.ReplaceAll(sn, "-", "")

	// make hash64 ready to generate crc64 of sn for its ID
	hash64.Reset()
	hash64.Write([]byte(sn))

	var tempName = struct {
		StructureID        uint64
		StructureUpperName string
		StructureLowerName string
		ReceiverName       string
	}{
		StructureID:        hash64.Sum64(), // hash sn for its ID
		StructureUpperName: sn,
		StructureLowerName: strings.ToLower(sn[0:1]) + sn[1:],
		ReceiverName:       rn,
	}

	var sf = new(bytes.Buffer)
	err = datastoreFileTemplate.Execute(sf, tempName)
	if err != nil {
		return
	}

	file.Data = sf.Bytes()
	// Indicate file had been changed
	file.State = assets.StateChanged

	return
}

var datastoreFileTemplate = template.Must(template.New("datastoreFileTemplate").Parse(`/* For license and copyright information please see LEGAL file in repository */

package datastore

import (
	"crypto/sha512"

	etime "../libgo/earth-time"
	"../libgo/ganjine"
	gsdk "../libgo/ganjine-sdk"
	gs "../libgo/ganjine-services"
	"../libgo/syllab"
)

const (
	{{.StructureLowerName}}StructureID uint64 = {{.StructureID}}
	{{.StructureLowerName}}FixedSize   uint64 = ??? // 72 + ??? + (??? * 8) >> Common header + Unique data + vars add&&len
	{{.StructureLowerName}}State       uint8  = ganjine.DataStructureStatePreAlpha
)

// {{.StructureUpperName}} store 
type {{.StructureUpperName}} struct {
	/* Common header data */
	RecordID          [32]byte
	RecordStructureID uint64
	RecordSize        uint64
	WriteTime         int64
	OwnerAppID        [16]byte

	/* Unique data */
	AppInstanceID    [16]byte // Store to remember which app instance set||chanaged this record!
	UserConnectionID [16]byte // Store to remember which user connection set||chanaged this record!
	Remove me and add more data structure here! Also can delete two above items!
}

// Set method set some data and write entire {{.StructureUpperName}} record!
func ({{.ReceiverName}} *{{.StructureUpperName}}) Set() (err error) {
	{{.ReceiverName}}.RecordStructureID = {{.StructureLowerName}}StructureID
	{{.ReceiverName}}.RecordSize = {{.ReceiverName}}.len()
	{{.ReceiverName}}.WriteTime = etime.Now()
	{{.ReceiverName}}.OwnerAppID = server.Manifest.AppID

	var req = gs.SetRecordReq{
		Type:   gs.RequestTypeBroadcast,
		Record: {{.ReceiverName}}.syllabEncoder(),
	}
	{{.ReceiverName}}.RecordID = sha512.Sum512_256(req.Record[32:])
	copy(req.Record[0:], {{.ReceiverName}}.RecordID[:])

	err = gsdk.SetRecord(cluster, &req)
	if err != nil {
		// TODO::: Handle error situation
	}

	return
}

// GetByRecordID method read all existing record data by given RecordID!
func ({{.ReceiverName}} *{{.StructureUpperName}}) GetByRecordID() (err error) {
	var req = gs.GetRecordReq{
		RecordID: {{.ReceiverName}}.RecordID,
	}
	var res *gs.GetRecordRes
	res, err = gsdk.GetRecord(cluster, &req)
	if err != nil {
		return err
	}

	err = {{.ReceiverName}}.syllabDecoder(res.Record)
	if err != nil {
		return err
	}

	if {{.ReceiverName}}.RecordStructureID != {{.StructureLowerName}}StructureID {
		err = ganjine.ErrRecordMisMatchedStructureID
	}
	return
}

// GetBy?????? method find and read last version of record by given ??????
func ({{.ReceiverName}} *{{.StructureUpperName}}) GetBy??????() (err error) {
	var indexReq = &gs.FindRecordsReq{
		IndexHash: {{.ReceiverName}}.hash??????(),
		Offset:    18446744073709551615,
		Limit:     0,
	}
	var indexRes *gs.FindRecordsRes
	indexRes, err = gsdk.FindRecords(cluster, indexReq)
	if err != nil {
		return err
	}

	var ln = len(indexRes.RecordIDs)
	// TODO::: Need to handle this here?? if collision ocurred and last record ID is not our purpose??
	for {
		ln--
		{{.ReceiverName}}.RecordID = indexRes.RecordIDs[ln]
		err = {{.ReceiverName}}.GetByRecordID()
		if err != ganjine.ErrRecordMisMatchedStructureID {
			return
		}
	}
}

// Delete method use to delete existing record just by given RecordID!
func ({{.ReceiverName}} *{{.StructureUpperName}}) Delete() (err error) {
	return
}

// Index?????? index {{.ReceiverName}}.?????? to retrieve record fast later.
func ({{.ReceiverName}} *{{.StructureUpperName}}) Index??????() {
	var ??????Index = gs.SetIndexHashReq{
		Type:      gs.RequestTypeBroadcast,
		IndexHash: {{.ReceiverName}}.hash??????(),
		RecordID:  {{.ReceiverName}}.RecordID,
	}
	var err = gsdk.SetIndexHash(cluster, &??????Index)
	if err != nil {
		// TODO::: we must retry more due to record wrote successfully!
	}
}

func ({{.ReceiverName}} *{{.StructureUpperName}}) hash??????() (hash [32]byte) {
	var buf = make([]byte, ??) // 8+??

	buf[0] = byte({{.ReceiverName}}.RecordStructureID)
	buf[1] = byte({{.ReceiverName}}.RecordStructureID >> 8)
	buf[2] = byte({{.ReceiverName}}.RecordStructureID >> 16)
	buf[3] = byte({{.ReceiverName}}.RecordStructureID >> 24)
	buf[4] = byte({{.ReceiverName}}.RecordStructureID >> 32)
	buf[5] = byte({{.ReceiverName}}.RecordStructureID >> 40)
	buf[6] = byte({{.ReceiverName}}.RecordStructureID >> 48)
	buf[7] = byte({{.ReceiverName}}.RecordStructureID >> 56)

	remove me and add desire data to buff to index it||them

	return sha512.Sum512_256(buf)
}

func ({{.ReceiverName}} *{{.StructureUpperName}}) syllabDecoder(buf []byte) (err error) {
	if uint64(len(buf)) < {{.StructureLowerName}}FixedSize {
		err = syllab.ErrSyllabDecodingFailedSmallSlice
		return
	}

	copy({{.ReceiverName}}.RecordID[:], buf[:])
	{{.ReceiverName}}.RecordStructureID = uint64(buf[32]) | uint64(buf[33])<<8 | uint64(buf[34])<<16 | uint64(buf[35])<<24 | uint64(buf[36])<<32 | uint64(buf[37])<<40 | uint64(buf[38])<<48 | uint64(buf[39])<<56
	{{.ReceiverName}}.RecordSize = uint64(buf[40]) | uint64(buf[41])<<8 | uint64(buf[42])<<16 | uint64(buf[43])<<24 | uint64(buf[44])<<32 | uint64(buf[45])<<40 | uint64(buf[46])<<48 | uint64(buf[47])<<56
	{{.ReceiverName}}.WriteTime = int64(buf[48]) | int64(buf[49])<<8 | int64(buf[50])<<16 | int64(buf[51])<<24 | int64(buf[52])<<32 | int64(buf[53])<<40 | int64(buf[54])<<48 | int64(buf[55])<<56
	copy({{.ReceiverName}}.OwnerAppID[:], buf[56:])

	copy({{.ReceiverName}}.AppInstanceID[:], buf[72:])
	copy({{.ReceiverName}}.UserConnectionID[:], buf[88:])

	return
}

func ({{.ReceiverName}} *{{.StructureUpperName}}) syllabEncoder() (buf []byte) {
	buf = make([]byte, {{.ReceiverName}}.len())

	// copy(buf[0:], {{.ReceiverName}}.RecordID[:])
	buf[32] = byte({{.ReceiverName}}.RecordSize)
	buf[33] = byte({{.ReceiverName}}.RecordSize >> 8)
	buf[34] = byte({{.ReceiverName}}.RecordSize >> 16)
	buf[35] = byte({{.ReceiverName}}.RecordSize >> 24)
	buf[36] = byte({{.ReceiverName}}.RecordSize >> 32)
	buf[37] = byte({{.ReceiverName}}.RecordSize >> 40)
	buf[38] = byte({{.ReceiverName}}.RecordSize >> 48)
	buf[39] = byte({{.ReceiverName}}.RecordSize >> 56)
	buf[40] = byte({{.ReceiverName}}.RecordStructureID)
	buf[41] = byte({{.ReceiverName}}.RecordStructureID >> 8)
	buf[42] = byte({{.ReceiverName}}.RecordStructureID >> 16)
	buf[43] = byte({{.ReceiverName}}.RecordStructureID >> 24)
	buf[44] = byte({{.ReceiverName}}.RecordStructureID >> 32)
	buf[45] = byte({{.ReceiverName}}.RecordStructureID >> 40)
	buf[46] = byte({{.ReceiverName}}.RecordStructureID >> 48)
	buf[47] = byte({{.ReceiverName}}.RecordStructureID >> 56)
	buf[48] = byte({{.ReceiverName}}.WriteTime)
	buf[49] = byte({{.ReceiverName}}.WriteTime >> 8)
	buf[50] = byte({{.ReceiverName}}.WriteTime >> 16)
	buf[51] = byte({{.ReceiverName}}.WriteTime >> 24)
	buf[52] = byte({{.ReceiverName}}.WriteTime >> 32)
	buf[53] = byte({{.ReceiverName}}.WriteTime >> 40)
	buf[54] = byte({{.ReceiverName}}.WriteTime >> 48)
	buf[55] = byte({{.ReceiverName}}.WriteTime >> 56)
	copy(buf[56:], {{.ReceiverName}}.OwnerAppID[:])

	copy(buf[72:], {{.ReceiverName}}.AppInstanceID[:])
	copy(buf[88:], {{.ReceiverName}}.UserConnectionID[:])

	return
}

func ({{.ReceiverName}} *{{.StructureUpperName}}) len() uint64 {
	return {{.StructureLowerName}}FixedSize + uint64(len({{.ReceiverName}}.??????))
}

`))
