/* For license and copyright information please see LEGAL file in repository */

package generator

import (
	"bytes"
	"go/format"
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

	file.Data, err = format.Source(sf.Bytes())
	// Indicate file had been changed
	file.State = assets.StateChanged

	return
}

var datastoreFileTemplate = template.Must(template.New("datastoreFileTemplate").Parse(`
/* For license and copyright information please see LEGAL file in repository */

package datastore

import (
	"time"

	ganjine "../../libgo/ganjine-sdk"
	"../../libgo/uuid"
)

const {{.StructureLowerName}}StructureID uint64 = {{.StructureID}}

// {{.StructureUpperName}} store 
type {{.StructureUpperName}} struct {
	/* Common header data */
	Checksum          [32]byte
	RecordID          [16]byte
	RecordSize        uint64
	RecordStructureID uint64
	WriteTime         int64
	OwnerAppID        [16]byte

	/* Unique data */
	AppConnectionID  [16]byte // Store to remember which app instance connection set||chanaged this record!
	UserConnectionID [16]byte // Store to remember which user connection set||chanaged this record!
	delete this line and add more data structure here! Also can delete two above items!
}

// Set method use to write entire {{.StructureUpperName}} record!
// {{.ReceiverName}} can't be nil otherwise panic occur!
func ({{.ReceiverName}} *{{.StructureUpperName}}) Set() (err error) {
	{{.ReceiverName}}.RecordID = uuid.NewV4()
	{{.ReceiverName}}.RecordStructureID = {{.StructureLowerName}}StructureID
	{{.ReceiverName}}.WriteTime = time.Now().Unix()
	{{.ReceiverName}}.OwnerAppID = server.Manifest.AppID

	var req = ganjine.SetRecordReq{
		RecordID: pa.RecordID,
	}
	req.Record = pa.syllabEncoder(0)

	err = ganjine.SetRecord(server, cluster, &req)
	if err != nil {
		// TODO::: Handle error situation
	}
	return
}

// Get method use to read all existing record just by given RecordID!
// {{.ReceiverName}} can't be nil otherwise panic occur!
func ({{.ReceiverName}} *{{.StructureUpperName}}) Get() (err error) {
	// TODO::: First read from local OS (related lib) as cache

	// If not exist in cache read from DataStore
	var req = ganjine.GetRecordReq{
		RecordID: pa.RecordID,
	}
	var res *ganjine.GetRecordRes
	res, err = ganjine.GetRecord(server, cluster, &req)
	if err != nil {
		return err
	}

	// TODO::: Write to local OS as cache if not enough storage exist do GC(Garbage Collector)

	err = pa.syllabDecoder(res.Record)

	return
}

// Delete method use to delete existing record just by given RecordID!
// {{.ReceiverName}} can't be nil otherwise panic occur!
func ({{.ReceiverName}} *{{.StructureUpperName}}) Delete() (err error) {
	return
}

func ({{.ReceiverName}} *{{.StructureUpperName}}) syllabDecoder(buf []byte) (err error) {
	return
}

func ({{.ReceiverName}} *{{.StructureUpperName}}) syllabEncoder(offset int) (buf []byte) {
	return
}

`))
