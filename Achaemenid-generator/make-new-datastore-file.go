/* For license and copyright information please see LEGAL file in repository */

package generator

import (
	"bytes"
	"go/format"
	"strings"
	"text/template"
	"time"
)

// MakeNewDatastoreFileReq is request structure of MakeNewDatastoreFile()
type MakeNewDatastoreFileReq struct {
	StructureName string
}

// MakeNewDatastoreFileRes is response structure of MakeNewDatastoreFile()
type MakeNewDatastoreFileRes struct {
	StructureFileName string
	StructureFile     []byte
}

// MakeNewDatastoreFile use to make new service template file!
func MakeNewDatastoreFile(req *MakeNewDatastoreFileReq) (res *MakeNewDatastoreFileRes, err error) {
	res = &MakeNewDatastoreFileRes{
		StructureFileName: req.StructureName + ".go",
	}

	req.StructureName = strings.Title(req.StructureName)
	req.StructureName = strings.ReplaceAll(req.StructureName, "-", "")

	var tempName = struct {
		StructureID        uint32
		StructureUpperName string
		StructureLowerName string
		IssueDate          int64
	}{
		StructureID:        0, // hash req.StructureName for its ID
		StructureUpperName: req.StructureName,
		StructureLowerName: strings.ToLower(req.StructureName[0:1]) + req.StructureName[1:],
		IssueDate:          time.Now().Unix(),
	}

	var sf = new(bytes.Buffer)
	err = datastoreFileTemplate.Execute(sf, tempName)
	if err != nil {
		return nil, err
	}

	res.StructureFile, err = format.Source(sf.Bytes())
	if err != nil {
		return nil, err
	}

	return res, nil
}

var datastoreFileTemplate = template.Must(template.New("datastoreFileTemplate").Parse(`
/* For license and copyright information please see LEGAL file in repository */

package datastore

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
	delete this line and add more data structure here!
}

// Set method use to write entire {{.StructureUpperName}} record!
// syl can't be nil otherwise panic occur!
func (syl *{{.StructureUpperName}}) Set() (err error) {
	return nil
}

// Get method use to read all existing record just by given RecordID!
// syl can't be nil otherwise panic occur!
func (syl *{{.StructureUpperName}}) Get() (err error) {
	return nil
}

// Delete method use to delete existing record just by given RecordID!
// syl can't be nil otherwise panic occur!
func (syl *{{.StructureUpperName}}) Delete() (err error) {
	return nil
}

func (syl *{{.StructureUpperName}}) syllabDecoder(buf []byte) (err error) {
	return nil
}

func (syl *{{.StructureUpperName}}) syllabEncoder(buf []byte) (err error) {
	return nil
}

`))
