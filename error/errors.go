/* For license and copyright information please see LEGAL file in repository */

package error

import (
	"../protocol"
)

// Errors store
type Errors struct {
	poolByID  map[uint64]protocol.Error
	poolByURN map[string]protocol.Error
}

func (e *Errors) Init() {
	e.poolByID = make(map[uint64]protocol.Error, 512)
	e.poolByURN = make(map[string]protocol.Error, 512)
}

func (e *Errors) RegisterError(err protocol.Error) {
	var errID = err.URN().ID()
	if protocol.AppDevMode && e.poolByID[errID] != nil {
		// This condition will just be true in the dev phase.
		panic("Error id exist and used for other Error. Check it now for bad urn set or collision occurred" +
			"\nExiting error >> " + e.poolByID[errID].URN().URI() +
			"\nNew error >> " + err.URN().URI())
	}

	e.poolByID[errID] = err
	e.poolByURN[err.URN().URI()] = err
}

func (e *Errors) UnRegisterError(err protocol.Error) {
	delete(e.poolByID, err.URN().ID())
	delete(e.poolByURN, err.URN().URI())
}

// GetErrorByID returns desire error if exist or ErrNotFound!
func (e *Errors) GetErrorByID(id uint64) (err protocol.Error) {
	if id == 0 {
		return
	}
	var ok bool
	err, ok = e.poolByID[id]
	if !ok {
		err = ErrNotFound
	}
	return
}

// GetErrorByID returns desire error if exist or ErrNotFound!
func (e *Errors) GetErrorByURN(urn string) (err protocol.Error) {
	var ok bool
	err, ok = e.poolByURN[urn]
	if !ok {
		err = ErrNotFound
	}
	return
}
