/* For license and copyright information please see LEGAL file in repository */

package mediatype

import (
	"strings"

	"../protocol"
	"../urn"
)

// MediaType implement protocol.MediaType interface
// type "/" [tree "."] subtype ["+" suffix]* [";" parameter]
type MediaType struct {
	id           uint64
	mediaType    string
	mainType     string
	tree         string
	subType      string
	suffix       string
	parameter    string
	extension    string // Use as file extension usually in windows os
	referenceURI string

	status          protocol.SoftwareStatus
	issueDate       protocol.TimeUnixSec
	expiryDate      protocol.TimeUnixSec
	expireInFavorOf protocol.MediaType // Other MediaType

	detail map[protocol.LanguageID]*MediaTypeDetail
}

// MediaTypeDetail store detail about an MediaType
type MediaTypeDetail struct {
	name        string
	description string
	tags        []string
}

func (mt *MediaType) ID() uint64                          { return mt.id }
func (mt *MediaType) MediaType() string                   { return mt.mediaType }
func (mt *MediaType) Type() string                        { return mt.mainType }
func (mt *MediaType) Tree() string                        { return mt.tree }
func (mt *MediaType) SubType() string                     { return mt.subType }
func (mt *MediaType) Suffix() string                      { return mt.suffix }
func (mt *MediaType) Parameter() string                   { return mt.parameter }
func (mt *MediaType) FileExtension() string               { return mt.extension }
func (mt *MediaType) Status() protocol.SoftwareStatus     { return mt.status }
func (mt *MediaType) IssueDate() protocol.TimeUnixSec     { return mt.issueDate }
func (mt *MediaType) ExpiryDate() protocol.TimeUnixSec    { return mt.expiryDate }
func (mt *MediaType) ExpireInFavorOf() protocol.MediaType { return mt.expireInFavorOf }

func (mt *MediaType) SetInfo(status protocol.SoftwareStatus, issueDate protocol.TimeUnixSec, referenceURI string) *MediaType {
	mt.status = status
	mt.issueDate = issueDate
	mt.referenceURI = referenceURI
	return mt
}

func (mt *MediaType) SetDetail(lang protocol.LanguageID, name, description string, tags []string) *MediaType {
	var ok bool
	_, ok = mt.detail[lang]
	if ok {
		// Can't change service detail after first set! Ask service holder to change details!!
		return mt
	}
	mt.detail[lang] = &MediaTypeDetail{
		name:        name,
		description: description,
		tags:        tags,
	}
	return mt
}

func (mt *MediaType) Expired(expiryDate protocol.TimeUnixSec, inFavorOf protocol.MediaType) MediaType {
	mt.expiryDate = expiryDate
	mt.expireInFavorOf = inFavorOf
	return *mt
}

func New(mediatype, fileExtension string) (mt *MediaType) {
	mt = &MediaType{
		mediaType: mediatype,
		extension: fileExtension,
		detail:    map[protocol.LanguageID]*MediaTypeDetail{},
	}
	mt.mainType, mt.subType = extractMediaType(mediatype)
	_, mt.id = urn.IDGenerator(mediatype)

	// Check due to os can be nil almost in tests and benchmarks build
	if protocol.OS != nil {
		protocol.OS.RegisterMediaType(mt)
	}
	return
}

func extractMediaType(mt string) (mainType, subType string) {
	// TODO::: complete extraction
	var i = strings.IndexByte(mt, '/')
	mainType = mt[:i]
	subType = mt[i+1:]
	return
}
