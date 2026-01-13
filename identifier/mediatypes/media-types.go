/* For license and copyright information please see the LEGAL file in the code repository */

package mediatypes

import (
	"memar/protocol"
)

func Register(mt protocol.MediaType) (err protocol.Error) { return mts.Register(mt) }
func ByMediaType(mediaType string) (mt protocol.MediaType, err protocol.Error) {
	return mts.GetByMediaType(mediaType)
}
func ByID(id protocol.MediaTypeID) (mt protocol.MediaType, err protocol.Error) {
	return mts.GetByID(id)
}
func ByFileExtension(ex string) (mt protocol.MediaType, err protocol.Error) {
	return mts.GetByFileExtension(ex)
}

var mts = mediaTypes{
	poolByMediaType:     map[string]protocol.MediaType{},
	poolByID:            map[protocol.MediaTypeID]protocol.MediaType{},
	poolByFileExtension: map[string]protocol.MediaType{},
}

type mediaTypes struct {
	poolByMediaType     map[string]protocol.MediaType
	poolByID            map[protocol.MediaTypeID]protocol.MediaType
	poolByFileExtension map[string]protocol.MediaType
}

// RegisterMediaType register given MediaType to the application
func (mts *mediaTypes) Register(mt protocol.MediaType) (err protocol.Error) {
	// TODO::: change panics to error
	if mt.MainType() == "" && mt.SubType() == "" {
		panic("Mediatype doesn't has main and sub type. Can't register empty mediatype.")
	}

	// TODO::: lock??
	mts.poolByMediaType[mt.ToString()] = mt
	mts.poolByID[mt.ID()] = mt
	var fe = mt.FileExtension()
	if fe != "" {
		mts.poolByFileExtension[fe] = mt
	}
	return
}

func (mts *mediaTypes) GetByMediaType(mediaType string) (mt protocol.MediaType, err protocol.Error) {
	mt = mts.poolByMediaType[mediaType]
	// if mt == nil {
	// err = &errs.ErrNotFound
	// }
	return
}

func (mts *mediaTypes) GetByID(id protocol.MediaTypeID) (mt protocol.MediaType, err protocol.Error) {
	mt = mts.poolByID[id]
	// if mt == nil {
	// err = &errs.ErrNotFound
	// }
	return
}
func (mts *mediaTypes) GetByFileExtension(ex string) (mt protocol.MediaType, err protocol.Error) {
	mt = mts.poolByFileExtension[ex]
	// if mt == nil {
	// err = &errs.ErrNotFound
	// }
	return
}
