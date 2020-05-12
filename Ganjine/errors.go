/* For license and copyright information please see LEGAL file in repository */

package ganjine

import (
	"errors"
)

// Declare Errors Details
var (
	ErrHashIndexRecordNotExist = errors.New("Given recordID not exist in any storage devices")
	ErrHashIndexRecordManipulated = errors.New("HashIndex record has problem when engine try to read it from storage devices")

	ErrContentAlreadyExist = errors.New("This content was already exist")

	ErrCantPrepareStatement = errors.New("Can't prepare a new statement to database")

	ErrStoringDataNotComplete = errors.New("We have some problem in storing your data in our databases. Send your request again! If error exist contact SabzCity platform administrators")

	ErrDatabaseConnectionError = errors.New("Could not connect to database")

	ErrDatabasePingOut = errors.New("Error ocurred in Ping to database")
)
