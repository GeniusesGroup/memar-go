/* For license and copyright information please see the LEGAL file in the code repository */

package errors

import (
	packageErrors "memar/errors/errors"
	"memar/protocol"
)

// Errors store
type Errors struct {
	poolByID        map[protocol.MediaTypeID]protocol.Error
	poolByMediaType map[string]protocol.Error
}

//memar:impl memar/protocol.ObjectLifeCycle
func (e *Errors) Init() (err protocol.Error) {
	e.poolByID = make(map[protocol.MediaTypeID]protocol.Error, 512)
	e.poolByMediaType = make(map[string]protocol.Error, 512)
	return
}
func (e *Errors) Reinit() (err protocol.Error) {
	return
}
func (e *Errors) Deinit() (err protocol.Error) {
	return
}

func (e *Errors) RegisterError(errorToRegister protocol.Error) (err protocol.Error) {
	var errID = errorToRegister.ID()

	if protocol.AppMode_Dev {
		if errID == 0 {
			err = &packageErrors.ErrNotProvideIdentifier
			return
		}
		if e.poolByID[errID] != nil {
			err = &packageErrors.ErrDuplicateIdentifier
			return
		}
	}

	e.poolByID[errID] = errorToRegister
	e.poolByMediaType[errorToRegister.ToString()] = errorToRegister
	return
}

func (e *Errors) UnRegisterError(err protocol.Error) {
	delete(e.poolByID, err.ID())
	delete(e.poolByMediaType, err.ToString())
}

// GetErrorByID returns desire error if exist or ErrNotFound!
func (e *Errors) GetErrorByID(id protocol.MediaTypeID) (err protocol.Error) {
	if id == 0 {
		return
	}
	var ok bool
	err, ok = e.poolByID[id]
	if !ok {
		err = &packageErrors.ErrNotFound
	}
	return
}

// GetErrorByMediaType returns desire error if exist or ErrNotFound!
func (e *Errors) GetErrorByMediaType(mt string) (err protocol.Error) {
	var ok bool
	err, ok = e.poolByMediaType[mt]
	if !ok {
		err = &packageErrors.ErrNotFound
	}
	return
}
