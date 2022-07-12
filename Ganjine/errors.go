/* For license and copyright information please see LEGAL file in repository */

package ganjine

import (
	er "../error"
	"../protocol"
)

const domainEnglish = "Ganjine"
const domainPersian = "گنجینه"

// Errors
var (
	ErrBadNode   er.Error
	ErrWrongNode er.Error
	// ErrContentAlreadyExist     er.Error
	// ErrCantPrepareStatement    er.Error
	// ErrStoringDataNotComplete  er.Error
	// ErrDatabaseConnectionError er.Error
	// ErrDatabasePingOut         er.Error
)

func init() {
	ErrBadNode.Init("urn:giti:ganjine.protocol:error:bad-node")
	ErrBadNode.SetDetail(protocol.LanguageEnglish, domainEnglish, "Bad Node",
		"Given request can't proccess due to send to an non Ganjine node",
		"",
		"",
		nil)

	ErrWrongNode.Init("urn:giti:ganjine.protocol:error:wrong-node")
	ErrWrongNode.SetDetail(protocol.LanguageEnglish, domainEnglish, "Wrong Node",
		"Given request can't proccess due to send to a Ganjine node that not own the request range!",
		"",
		"",
		nil)

	// ErrContentAlreadyExist     = er.New("This content was already exist")
	// ErrCantPrepareStatement    = er.New("Can't prepare a new statement to database")
	// ErrStoringDataNotComplete  = er.New("We have some problem in storing your data in our databases. Send your request again! If error exist contact SabzCity platform administrators")
	// ErrDatabaseConnectionError = er.New("Could not connect to database")
	// ErrDatabasePingOut         = er.New("Error ocurred in Ping to database")
}
