/* For license and copyright information please see LEGAL file in repository */

package mediatype

import (
	"encoding/base64"
	"reflect"
	"strings"
	"unsafe"

	"golang.org/x/crypto/sha3"

	"../protocol"
)

// MediaType implement protocol.MediaType interface
// type "/" [tree "."] subtype ["+" suffix]* [";" parameter]
// https://datatracker.ietf.org/doc/html/rfc2046
type MediaType struct {
	uuid       [32]byte
	id         uint64
	idAsString string

	mediaType  string
	mainType   string
	tree       string
	subType    string
	suffix     string
	parameters []string

	fileExtension string

	referenceURI    string
	status          protocol.SoftwareStatus
	issueDate       protocol.TimeUnixSec
	expiryDate      protocol.TimeUnixSec
	expireInFavorOf protocol.MediaType // Other MediaType

	detail  map[protocol.LanguageID]*MediaTypeDetail
	details []protocol.MediaTypeDetail
}

func (mt *MediaType) Init(raw string) {
	mt.mediaType = raw
	// TODO::: complete extraction
	var i = strings.IndexByte(raw, '/')
	mt.mainType = raw[:i]
	mt.subType = raw[i+1:]
}

func (mt *MediaType) UUID() [32]byte                      { return mt.uuid }
func (mt *MediaType) ID() uint64                          { return mt.id }
func (mt *MediaType) IDasString() string                  { return mt.idAsString }
func (mt *MediaType) MediaType() string                   { return mt.mediaType }
func (mt *MediaType) Type() string                        { return mt.mainType }
func (mt *MediaType) Tree() string                        { return mt.tree }
func (mt *MediaType) SubType() string                     { return mt.subType }
func (mt *MediaType) Suffix() string                      { return mt.suffix }
func (mt *MediaType) Parameters() []string                { return mt.parameters }
func (mt *MediaType) FileExtension() string               { return mt.fileExtension }
func (mt *MediaType) Status() protocol.SoftwareStatus     { return mt.status }
func (mt *MediaType) IssueDate() protocol.TimeUnixSec     { return mt.issueDate }
func (mt *MediaType) ExpiryDate() protocol.TimeUnixSec    { return mt.expiryDate }
func (mt *MediaType) ExpireInFavorOf() protocol.MediaType { return mt.expireInFavorOf }
func (mt *MediaType) Details() []protocol.MediaTypeDetail { return mt.details }
func (mt *MediaType) Detail(lang protocol.LanguageID) protocol.MediaTypeDetail {
	return mt.detail[lang]
}

func (mt *MediaType) SetFileExtension(fileExtension string) *MediaType {
	mt.fileExtension = fileExtension
	// TODO::: if application use other package to register mediatypes, below logic will break the app functionality.
	poolByFileExtension[fileExtension] = mt
	// protocol.OS.RegisterMediaType(mt)
	return mt
}

func (mt *MediaType) SetInfo(status protocol.SoftwareStatus, issueDate protocol.TimeUnixSec, referenceURI string) *MediaType {
	mt.status = status
	mt.issueDate = issueDate
	mt.referenceURI = referenceURI
	return mt
}

// SetDetail add error text details to existing error and return it.
func (mt *MediaType) SetDetail(lang protocol.LanguageID, domain, summary, overview, userNote, devNote string, tags []string) *MediaType {
	var _, ok = mt.detail[lang]
	if ok {
		panic("Can't change MediaType detail after first set! Ask the holder to change details.")
	}

	var detail = MediaTypeDetail{
		languageID: lang,
		domain:     domain,
		summary:    summary,
		overview:   overview,
		userNote:   userNote,
		devNote:    devNote,
		tags:       tags,
	}
	mt.detail[lang] = &detail
	mt.details = append(mt.details, &detail)
	return mt
}

func (mt *MediaType) Expired(expiryDate protocol.TimeUnixSec, inFavorOf protocol.MediaType) *MediaType {
	mt.expiryDate = expiryDate
	mt.expireInFavorOf = inFavorOf
	return mt
}

func New(mediatype string) (mt *MediaType) {
	mt = &MediaType{
		detail: map[protocol.LanguageID]*MediaTypeDetail{},
	}
	mt.Init(mediatype)
	mt.uuid, mt.id = IDGenerator(mediatype)
	mt.idAsString = base64.RawURLEncoding.EncodeToString(mt.uuid[:8])

	// Check due to os can be nil almost in tests and benchmarks build
	if protocol.OS != nil {
		protocol.OS.RegisterMediaType(mt)
	}
	return
}

func IDGenerator(uri string) (uuid [32]byte, id uint64) {
	uuid = sha3.Sum256(unsafeStringToByteSlice(uri))
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
