/* For license and copyright information please see the LEGAL file in the code repository */

package net_p

import (
	identifier_p "memar/identifier/protocol"
)

type Field_SessionID interface {
	// It will use to get other data such as UserID, PeerID, DelegateUserID, ...
	SessionID() SessionID
}

type SessionID = identifier_p.UUID_Time
