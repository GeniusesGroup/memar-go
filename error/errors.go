/* For license and copyright information please see LEGAL file in repository */

package error

import (
	"../giti"
	"../log"
)

type errors struct {
	poolByID map[uint64]*Error
	jsSDK    map[giti.LanguageID][]byte
}

// GetErrorByID returns desire error if exist or nil!
func (e *errors) GetErrorByID(id uint64) (err giti.Error) {
	var ok bool
	err, ok = e.poolByID[id]
	if !ok {
		err = ErrNotFound
	}
	return
}

func (e *errors) addError(err *Error) {
	var exitingError = e.poolByID[err.id]
	if exitingError != nil {
		log.Warn("Duplicate Error id exist, Check it now!!!!!!!!!!!!!!!!!")
		log.Warn("Exiting error >> ", exitingError.URN(), " New error >> ", *err)
		return
	}

	e.poolByID[err.id] = err
	e.updateJsSDK(err)
}

func (e *errors) updateJsSDK(err *Error) {
	for lang, detail := range err.detail {
		e.jsSDK[lang] = append(e.jsSDK[lang], "GitiError.New(\""+err.idAsString+"\",\""+err.urn+"\").SetDetail(\""+detail.domain+"\",\""+detail.short+"\",\""+detail.long+"\",\""+detail.userAction+"\",\""+detail.devAction+"\")\n"...)
	}
}

// it can return nil slice if not call updateJsSDK before call this method!
func (e *errors) GetErrorsInJsFormat(lang giti.LanguageID) []byte {
	return e.jsSDK[lang]
}

// Errors store
var Errors = errors{
	poolByID: make(map[uint64]*Error, 512),
	jsSDK:    map[giti.LanguageID][]byte{},
}
