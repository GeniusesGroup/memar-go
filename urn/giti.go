/* For license and copyright information please see LEGAL file in repository */

package urn

import (
	"encoding/base64"
	"reflect"
	"unsafe"

	"golang.org/x/crypto/sha3"

	"../protocol"
)

// Giti implement protocol.GitiURN interface
// "urn:giti:{{domain-name}}:data-structure:{{data-structure-name}}"
// "urn:giti:{{domain-name}}:service:{{service-name}}"
// "urn:giti:{{domain-name}}:error:{{error-name}}"
type Giti struct {
	URN
	uuid       [32]byte
	id         uint64
	idAsString string
	domain     string
	scope      string
	name       string // It must be unique in the domain scope e.g. "product" in "page" scope of the "domain"
}

func NewGiti(urn string) (u *Giti) {
	u = &Giti{}
	u.Init(urn)
	return
}

// SetDetail add error text details to existing error and return it.
func (u *Giti) Init(urn string) {
	u.URN.uri = urn
	u.uuid, u.id = IDGenerator(urn)
	u.idAsString = base64.RawURLEncoding.EncodeToString(u.uuid[:8])
	// TODO:::
}

func (u *Giti) NID() string        { return "giti" }
func (u *Giti) UUID() [32]byte     { return u.uuid }
func (u *Giti) ID() uint64         { return u.id }
func (u *Giti) IDasString() string { return u.idAsString }
func (u *Giti) Domain() string     { return u.domain }
func (u *Giti) Scope() string      { return u.scope }
func (u *Giti) Name() string       { return u.name }

func CheckGitiID(urn *Giti, id uint64) (err protocol.Error) {
	if urn.id != id {
		// err = protocol.App.Log(protocol.LogType_Warning, "Wrong ID:", mt.URN().ID, "-- Calculate ID:", mt.URN.ID(), "--", ErrNotStandardStructureID)
	}
	return
}

func IDGenerator(urn string) (uuid [32]byte, id uint64) {
	uuid = sha3.Sum256(unsafeStringToByteSlice(urn))
	id = uint64(uuid[0]) | uint64(uuid[1])<<8 | uint64(uuid[2])<<16 | uint64(uuid[3])<<24 | uint64(uuid[4])<<32 | uint64(uuid[5])<<40 | uint64(uuid[6])<<48 | uint64(uuid[7])<<56
	return
}

func IDfromString(IDasString string) (id uint64, err protocol.Error) {
	var IDasSlice = unsafeStringToByteSlice(IDasString)
	var ID [8]byte
	var _, goErr = base64.RawURLEncoding.Decode(ID[:], IDasSlice)
	if goErr != nil {
		// err =
		return
	}
	id = uint64(ID[0]) | uint64(ID[1])<<8 | uint64(ID[2])<<16 | uint64(ID[3])<<24 | uint64(ID[4])<<32 | uint64(ID[5])<<40 | uint64(ID[6])<<48 | uint64(ID[7])<<56
	return
}

func unsafeStringToByteSlice(req string) (res []byte) {
	var reqStruct = (*reflect.StringHeader)(unsafe.Pointer(&req))
	var resStruct = (*reflect.SliceHeader)(unsafe.Pointer(&res))
	resStruct.Data = reqStruct.Data
	resStruct.Len = reqStruct.Len
	resStruct.Cap = reqStruct.Len
	return
}
