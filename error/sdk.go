/* For license and copyright information please see LEGAL file in repository */

package error

import (
	lang "../language"
	"../log"
)

type errors struct {
	poolByID map[uint64]*Error
	// poolByOrder store to access all errors in order
	poolByOrder []*Error

	jsSDK map[lang.Language][]byte
}

// GetErrorByCode returns desire error if exist or nil!
func (e *errors) GetErrorByCode(id uint64) (err *Error) {
	err = e.poolByID[id]
	if err == nil {
		return ErrNotFound
	}
	return
}

func (e *errors) AddError(err *Error) {
	var exitingError = Errors.poolByID[err.id]
	if exitingError != nil {
		log.Warn("Duplicate Error id exist, Check it now!!!!!!!!!!!!!!!!!")
		log.Warn("Exiting error >> ", *exitingError, " New error >> ", *err)
		return
	}

	e.updateJsSDK(err)

	Errors.poolByID[err.id] = err
	Errors.poolByOrder = append(Errors.poolByOrder, err)
}

func (e *errors) updateJsSDK(err *Error) {
	for lang, detail := range err.detail {
		e.jsSDK[lang] = append(e.jsSDK[lang], "GitiError.New(\""+err.idAsString+"\",\""+err.urn+"\",\""+detail.Domain+"\",\""+detail.Short+"\",\""+detail.Long+"\")\n"...)
	}
}

// it can return nil slice if not call UpdateJSSDK before call this method!
func (e *errors) GetErrorsInJsFormat(lang lang.Language) []byte {
	return e.jsSDK[lang]
}

// Errors store
var Errors = errors{
	poolByID:    map[uint64]*Error{},
	poolByOrder: make([]*Error, 0, 128),
	jsSDK:       map[lang.Language][]byte{},
}
