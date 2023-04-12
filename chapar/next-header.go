/* For license and copyright information please see the LEGAL file in the code repository */

package chapar

import (
	"libgo/protocol"
)

type NextHeaderID byte

// https://github.com/GeniusesGroup/RFCs/blob/master/Chapar.md#next-header-standard-supported-protocols
const (
	NextHeaderID_Unset NextHeaderID = iota
	NextHeaderID_sRPC
	NextHeaderID_GP
	NextHeaderID_IPv6
	NextHeaderID_NTP

	NextHeaderID_Experimental1 NextHeaderID = 251 // Use for non supported protocols like IPv4, ...
	NextHeaderID_Experimental2 NextHeaderID = 252
	NextHeaderID_Experimental3 NextHeaderID = 253
	NextHeaderID_Experimental4 NextHeaderID = 254
	NextHeaderID_Experimental5 NextHeaderID = 255
)

func NetworkLink_NextHeaderIDToChaparNextHeaderID(nhID protocol.NetworkLink_NextHeaderID) NextHeaderID {
	if nhID > 255 {
		return NextHeaderID_Unset
	}
	return NextHeaderID(nhID)
}
