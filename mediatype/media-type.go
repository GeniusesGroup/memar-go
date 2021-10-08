/* For license and copyright information please see LEGAL file in repository */

package mediatype

import (
	"strings"

	etime "../earth-time"
	"../protocol"
	"../urn"
)

// MediaType implement protocol.MediaType interface
type MediaType struct {
	urn          urn.Giti // Use ID instead of other data to improve efficiency. "urn:giti:{{domain-name}}:data-structure:{{data-structure-name}}"
	mediaType    string
	mainType     string
	subType      string
	extension    string // Use as file extension usually in windows os
	referenceURI string

	status          protocol.MediaTypeStatus
	issueDate       etime.Time
	expiryDate      etime.Time
	expireInFavorOf protocol.MediaType // Other MediaType
	structure       interface{}

	detail map[protocol.LanguageID]*MediaTypeDetail
}

// MediaTypeDetail store detail about an MediaType
type MediaTypeDetail struct {
	name        string
	description string
	tags        []string
}

func (mt *MediaType) URN() protocol.GitiURN               { return &mt.urn }
func (mt *MediaType) MediaType() string                   { return mt.mediaType }
func (mt *MediaType) MainType() string                    { return mt.mainType }
func (mt *MediaType) SubType() string                     { return mt.subType }
func (mt *MediaType) FileExtension() string               { return mt.extension }
func (mt *MediaType) Status() protocol.MediaTypeStatus    { return mt.status }
func (mt *MediaType) IssueDate() protocol.Time            { return mt.issueDate }
func (mt *MediaType) ExpiryDate() protocol.Time           { return mt.expiryDate }
func (mt *MediaType) ExpireInFavorOf() protocol.MediaType { return mt.expireInFavorOf }

func (mt *MediaType) SetInfo(status protocol.MediaTypeStatus, issueDate int64, structure interface{}) *MediaType {
	mt.status = status
	mt.issueDate = etime.Time(issueDate)
	mt.structure = structure
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

func (mt *MediaType) Expired(expiryDate int64, inFavorOf protocol.MediaType) MediaType {
	mt.expiryDate = etime.Time(expiryDate)
	mt.expireInFavorOf = inFavorOf
	return *mt
}

func New(urn, media, fileExtension, referenceURI string) (mt *MediaType) {
	mt = &MediaType{
		mediaType:    media,
		extension:    fileExtension,
		referenceURI: referenceURI,
		detail:       map[protocol.LanguageID]*MediaTypeDetail{},
	}
	mt.mainType, mt.subType = splitMediaType(media)
	mt.urn.Init(urn)

	protocol.OS.RegisterMediaType(mt)
	return
}

func splitMediaType(mt string) (mainType, subType string) {
	var i = strings.IndexByte(mt, '/')
	mainType = mt[:i]
	subType = mt[i+1:]
	return
}
