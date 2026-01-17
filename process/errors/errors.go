/* For license and copyright information please see the LEGAL file in the code repository */

package errors

import (
	datatype_p "memar/computer/datatype/protocol"
	error_p "memar/process/error/protocol"
	errors_errs "memar/process/errors/errors"
	"memar/process/log"
)

func Register(er error_p.Error) (err error_p.Error)   { return errors.Register(er) }
func UnRegister(er error_p.Error) (err error_p.Error) { return errors.UnRegister(er) }
func GetByID(id datatype_p.ID) (err error_p.Error)    { return errors.GetByID(id) }
func GetByMediaType(mt string) (err error_p.Error)    { return errors.GetByMediaType(mt) }

var errors = errorsPool{
	poolByID:        make(map[datatype_p.ID]error_p.Error, 256),
	poolByMediaType: make(map[string]error_p.Error, 256),
}

type errorsPool struct {
	poolByID        map[datatype_p.ID]error_p.Error
	poolByMediaType map[string]error_p.Error
}

func (self *errorsPool) Register(errorToRegister error_p.Error) (err error_p.Error) {
	var errID = errorToRegister.DataTypeID()

	if log.CNF_DevelopingMode {
		if errID == 0 {
			err = &errors_errs.NotProvideIdentifier
			return
		}
		if self.poolByID[errID] != nil {
			err = &errors_errs.DuplicateIdentifier
			return
		}
	}

	self.poolByID[errID] = errorToRegister
	self.poolByMediaType[errorToRegister.MediaType()] = errorToRegister
	return
}

func (self *errorsPool) UnRegister(er error_p.Error) (err error_p.Error) {
	delete(self.poolByID, er.DataTypeID())
	delete(self.poolByMediaType, er.MediaType())
	return
}

// GetErrorByID returns desire error if exist or ErrNotFound!
func (self *errorsPool) GetByID(id datatype_p.ID) (err error_p.Error) {
	if id == 0 {
		return
	}
	var ok bool
	err, ok = self.poolByID[id]
	if !ok {
		err = &errors_errs.ErrNotFound
	}
	return
}

// GetErrorByMediaType returns desire error if exist or ErrNotFound!
func (self *errorsPool) GetByMediaType(mt string) (err error_p.Error) {
	var ok bool
	err, ok = self.poolByMediaType[mt]
	if !ok {
		err = &errors_errs.ErrNotFound
	}
	return
}
