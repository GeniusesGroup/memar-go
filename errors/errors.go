/* For license and copyright information please see the LEGAL file in the code repository */

package errors

import (
	"memar/protocol"
)

func Register(er protocol.Error) (err protocol.Error)      { return errors.Register(er) }
func UnRegister(er protocol.Error) (err protocol.Error)    { return errors.UnRegister(er) }
func GetByID(id protocol.MediaTypeID) (err protocol.Error) { return errors.GetByID(id) }
func GetByMediaType(mt string) (err protocol.Error)        { return errors.GetByMediaType(mt) }

var errors errors_

type errors_ struct {
	poolByID        map[protocol.MediaTypeID]protocol.Error
	poolByMediaType map[string]protocol.Error
}

//memar:impl memar/protocol.ObjectLifeCycle
func (e *errors_) Init() (err protocol.Error) {
	e.poolByID = make(map[protocol.MediaTypeID]protocol.Error, 256)
	e.poolByMediaType = make(map[string]protocol.Error, 256)
	return
}
func (e *errors_) Reinit() (err protocol.Error) {
	return
}
func (e *errors_) Deinit() (err protocol.Error) {
	return
}

func (e *errors_) Register(errorToRegister protocol.Error) (err protocol.Error) {
	var errID = errorToRegister.ID()

	if protocol.AppMode_Dev {
		if errID == 0 {
			err = &ErrNotProvideIdentifier
			return
		}
		if e.poolByID[errID] != nil {
			err = &ErrDuplicateIdentifier
			return
		}
	}

	e.poolByID[errID] = errorToRegister
	e.poolByMediaType[errorToRegister.ToString()] = errorToRegister
	return
}

func (e *errors_) UnRegister(er protocol.Error) (err protocol.Error) {
	delete(e.poolByID, er.ID())
	delete(e.poolByMediaType, er.ToString())
	return
}

// GetErrorByID returns desire error if exist or ErrNotFound!
func (e *errors_) GetByID(id protocol.MediaTypeID) (err protocol.Error) {
	if id == 0 {
		return
	}
	var ok bool
	err, ok = e.poolByID[id]
	if !ok {
		err = &ErrNotFound
	}
	return
}

// GetErrorByMediaType returns desire error if exist or ErrNotFound!
func (e *errors_) GetByMediaType(mt string) (err protocol.Error) {
	var ok bool
	err, ok = e.poolByMediaType[mt]
	if !ok {
		err = &ErrNotFound
	}
	return
}
