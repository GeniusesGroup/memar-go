/* For license and copyright information please see LEGAL file in repository */

package generator

import (
	"bytes"
	"hash/crc64"
	"strings"
	"text/template"

	"../assets"
	etime "../earth-time"
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
		IssueDate          int64
	}{
		StructureID:        hash64.Sum64(), // hash sn for its ID
		StructureUpperName: sn,
		StructureLowerName: strings.ToLower(sn[0:1]) + sn[1:],
		ReceiverName:       rn,
		IssueDate:          etime.Now(),
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
	lang "../libgo/language"
	"../libgo/log"
	"../libgo/syllab"
)

const (
	{{.StructureLowerName}}StructureID uint64 = {{.StructureID}}
)

var {{.StructureLowerName}}Structure = ganjine.DataStructure{
	ID:                {{.StructureID}},
	IssueDate:         {{.IssueDate}},
	ExpiryDate:        0,
	ExpireInFavorOf:   "", // Other structure name
	ExpireInFavorOfID: 0,  // Other StructureID! Handy ID or Hash of ExpireInFavorOf!
	Status:            ganjine.DataStructureStatePreAlpha,
	Structure:         {{.StructureUpperName}}{},

	Name: map[lang.Language]string{
		lang.EnglishLanguage: "{{.StructureUpperName}}",
	},
	Description: map[lang.Language]string{
		lang.EnglishLanguage: "",
	},
	TAGS: []string{
		"",
	},
}

// {{.StructureUpperName}} ---Read locale description in {{.StructureLowerName}}Structure---
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
	ID               [16]byte ` + "`" + `ganjine:"Immutable,Unique"` + "`" + ` // {{.StructureUpperName}}ID
	OwnerID          [16]byte ` + "`" + `ganjine:"Immutable"` + "`" + ` // Owner User ID
	Remove me and add more data structure here! Also can delete two above items!
}

// Set method set some data and write entire {{.StructureUpperName}} record!
func ({{.ReceiverName}} *{{.StructureUpperName}}) Set() (err error) {
	{{.ReceiverName}}.RecordStructureID = {{.StructureLowerName}}StructureID
	{{.ReceiverName}}.RecordSize = {{.ReceiverName}}.syllabLen()
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
		if log.DebugMode {
			log.Debug("Ganjine - Set Record Error:", err)
		}
		// TODO::: Handle error situation
	}

	return
}

// GetByRecordID read existing record data by given RecordID!
func ({{.ReceiverName}} *{{.StructureUpperName}}) GetByRecordID() (err error) {
	var req = gs.GetRecordReq{
		RecordID: {{.ReceiverName}}.RecordID,
	}
	var res *gs.GetRecordRes
	res, err = gsdk.GetRecord(cluster, &req)
	if err != nil {
		return
	}

	err = {{.ReceiverName}}.syllabDecoder(res.Record)
	if err != nil {
		return
	}

	if {{.ReceiverName}}.RecordStructureID != {{.StructureLowerName}}StructureID {
		err = ganjine.ErrGanjineMisMatchedStructureID
	}
	return
}

// GetLastByID find and read last version of record by given ID
func ({{.ReceiverName}} *{{.StructureUpperName}}) GetLastByID() (err error) {
	var indexReq = &gs.HashIndexGetValuesReq{
		IndexKey: {{.ReceiverName}}.HashID(),
		Offset:    18446744073709551615,
		Limit:     1,
	}
	var indexRes *gs.HashIndexGetValuesRes
	indexRes, err = gsdk.HashIndexGetValues(cluster, indexReq)
	if err != nil {
		return
	}

	pn.RecordID = indexRes.IndexValues[0]
	err = pn.GetByRecordID()
	if err == ganjine.ErrGanjineMisMatchedStructureID {
		log.Warn("Platform collapsed!! HASH Collision Occurred on", {{.StructureLowerName}}StructureID)
	}
	return
}

// GetLastBy?????? find and read last version of record by given ??????
func ({{.ReceiverName}} *{{.StructureUpperName}}) GetLastBy??????() (err error) {
	var indexReq = &gs.HashIndexGetValuesReq{
		IndexKey: {{.ReceiverName}}.Hash??????(),
		Offset:    18446744073709551615,
		Limit:     1,
	}
	var indexRes *gs.HashIndexGetValuesRes
	indexRes, err = gsdk.HashIndexGetValues(cluster, indexReq)
	if err != nil {
		return err
	}

	copy(pn.ID[:], indexRes.IndexValues[0][:])
	err = pn.GetLastByID()
	if err == ganjine.ErrGanjineMisMatchedStructureID {
		log.Warn("Platform collapsed!! HASH Collision Occurred on", {{.StructureLowerName}}StructureID)
	}
	return
}

// Delete existing record just by given RecordID!
func ({{.ReceiverName}} *{{.StructureUpperName}}) Delete() (err error) {
	return
}

/*
	-- PRIMARY INDEXES --
*/

// IndexID index Unique-Field(ID???) chain to retrieve last record version fast later.
// Call in each update to the exiting record!
func ({{.ReceiverName}} *{{.StructureUpperName}}) IndexID() {
	var indexRequest = gs.HashIndexSetValueReq{
		Type:      gs.RequestTypeBroadcast,
		IndexKey: {{.ReceiverName}}.HashID(),
		IndexValue:  {{.ReceiverName}}.RecordID,
	}
	var err = gsdk.HashIndexSetValue(cluster, &indexRequest)
	if err != nil {
		if log.DebugMode {
			log.Debug("Ganjine - Set Index Error:", err)
		}
		// TODO::: we must retry more due to record wrote successfully!
	}
}


// HashID hash {{.StructureLowerName}}StructureID + {{.ReceiverName}}.ID
func ({{.ReceiverName}} *{{.StructureUpperName}}) HashID() (hash [32]byte) {
	var buf = make([]byte, 24) // 8+16
	syllab.SetUInt64(buf, 0, {{.StructureLowerName}}StructureID)
	copy(buf[8:], {{.ReceiverName}}.ID[:])
	return sha512.Sum512_256(buf)
}

/*
	-- SECONDARY INDEXES --
*/

// IndexOwner index to retrieve all Unique-Field(ID???) owned by given OwnerID later.
// Don't call in update to an exiting record!
func ({{.ReceiverName}} *{{.StructureUpperName}}) IndexOwner() {
	var indexRequest = gs.HashIndexSetValueReq{
		Type:      gs.RequestTypeBroadcast,
		IndexKey: {{.ReceiverName}}.HashOwnerID(),
	}
	copy(indexRequest.IndexValue[:], {{.ReceiverName}}.ID[:])
	var err = gsdk.HashIndexSetValue(cluster, &indexRequest)
	if err != nil {
		if log.DebugMode {
			log.Debug("Ganjine - Set Index Error:", err)
		}
		// TODO::: we must retry more due to record wrote successfully!
	}
}

// HashOwnerID hash {{.StructureLowerName}}StructureID + OwnerID
func ({{.ReceiverName}} *{{.StructureUpperName}}) HashOwnerID() (hash [32]byte) {
	var buf = make([]byte, 24) // 8+16
	syllab.SetUInt64(buf, 0, {{.StructureLowerName}}StructureID)
	copy(buf[8:], {{.ReceiverName}}.OwnerID[:])
	return sha512.Sum512_256(buf)
}

// Index?????? index to retrieve all Unique-Field(ID???) owned by given ?????? later.
// Don't call in update to an exiting record!
func ({{.ReceiverName}} *{{.StructureUpperName}}) Index??????() {
	var indexRequest = gs.HashIndexSetValueReq{
		Type:      gs.RequestTypeBroadcast,
		IndexKey: {{.ReceiverName}}.Hash??????(),
	}
	copy(indexRequest.IndexValue[:], {{.ReceiverName}}.ID[:])
	var err = gsdk.HashIndexSetValue(cluster, &indexRequest)
	if err != nil {
		if log.DebugMode {
			log.Debug("Ganjine - Set Index Error:", err)
		}
		// TODO::: we must retry more due to record wrote successfully!
	}
}

// Hash?????? hash {{.StructureLowerName}}StructureID + ??????
func ({{.ReceiverName}} *{{.StructureUpperName}}) Hash??????() (hash [32]byte) {
	var buf = make([]byte, ??) // 8+??
	syllab.SetUInt64(buf, 0, {{.StructureLowerName}}StructureID)
	remove me and add desire data to buff to index it||them
	return sha512.Sum512_256(buf)
}

/*
	-- LIST FIELDS --
*/

// List??????  store all ?????? own by specific ??????.
// Don't call in update to an exiting record!
func ({{.ReceiverName}} *{{.StructureUpperName}}) List??????() {
	var indexRequest = gs.HashIndexSetValueReq{
		Type:      gs.RequestTypeBroadcast,
		IndexKey: {{.ReceiverName}}.Hash??????(),
	}
	copy(indexRequest.IndexValue[:], {{.ReceiverName}}.??????ID[:])
	var err = gsdk.HashIndexSetValue(cluster, &indexRequest)
	if err != nil {
		if log.DebugMode {
			log.Debug("Ganjine - Set Index Error:", err)
		}
		// TODO::: we must retry more due to record wrote successfully!
	}
}

// HashOwner??????Field hash {{.StructureLowerName}}StructureID + ?????? + "??????" field
func ({{.ReceiverName}} *{{.StructureUpperName}}) Hash??????Field() (hash [32]byte) {
	const field = "??????"
	var buf = make([]byte, 24+len(field)) // 8+16
	syllab.SetUInt64(buf, 0, {{.StructureLowerName}}StructureID)
	copy(buf[8:], {{.ReceiverName}}.??????[:])
	copy(buf[24:], field)
	return sha512.Sum512_256(buf)
}

/*
	-- Temporary INDEXES & LIST --
*/

// ??

/*
	-- Syllab Encoder & Decoder --
*/

func ({{.ReceiverName}} *{{.StructureUpperName}}) syllabDecoder(buf []byte) (err error) {
	return
}

func ({{.ReceiverName}} *{{.StructureUpperName}}) syllabEncoder() (buf []byte) {
	return
}

func ({{.ReceiverName}} *{{.StructureUpperName}}) syllabStackLen() (ln uint32) {
	return 0 // 72 + ??? + (?? * 8) >> Common header + Unique data + vars add&&len
}

func ({{.ReceiverName}} *{{.StructureUpperName}}) syllabHeapLen() (ln uint32) {
	return
}

func ({{.ReceiverName}} *{{.StructureUpperName}}) syllabLen() (ln uint64) {
	return uint64({{.ReceiverName}}.syllabStackLen() + {{.ReceiverName}}.syllabHeapLen())
}
`))
