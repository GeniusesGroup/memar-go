/* For license and copyright information please see LEGAL file in repository */

package ganjine

import (
	"../errors"
)

// Declare Errors Details
var (
	ErrHashIndexRecordNil         = errors.New("HashIndexRecordNil", "Given hash record can't be nil")
	ErrHashIndexRecordNotValid    = errors.New("HashIndexRecordNotValid", "Given recordID exist in storage devices but has diffrent StructureID")
	ErrHashIndexRecordNotExist    = errors.New("HashIndexRecordNotExist", "Given recordID not exist in any storage devices")
	ErrHashIndexRecordManipulated = errors.New("HashIndexRecordManipulated", "HashIndex record has problem when engine try to read it from storage devices")

	ErrNodeNotGanjineNode = errors.New("NodeNotGanjineNode", "Given request can't proccess due to send to an non Ganjine node")

	// ErrContentAlreadyExist     = errors.New("This content was already exist")
	// ErrCantPrepareStatement    = errors.New("Can't prepare a new statement to database")
	// ErrStoringDataNotComplete  = errors.New("We have some problem in storing your data in our databases. Send your request again! If error exist contact SabzCity platform administrators")
	// ErrDatabaseConnectionError = errors.New("Could not connect to database")
	// ErrDatabasePingOut         = errors.New("Error ocurred in Ping to database")
)
