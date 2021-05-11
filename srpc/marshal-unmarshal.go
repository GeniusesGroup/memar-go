/* For license and copyright information please see LEGAL file in repository */

package srpc

import "../syllab"

/*
	********************PAY ATTENTION:*******************
	We don't suggest use these 2 func instead use achaemenid-generator to autogenerate needed code before compile time
	and reduce runtime proccess to improve performance of the app and gain max performance from this protocol!
*/

// MarshalPacket use to encode automatically the value of s to the payload buffer.
func MarshalPacket(id uint64, s interface{}) (p []byte, err error) {
	// encode s to p by syllab encoder
	p, err = syllab.Marshal(s, p[8:])

	// Set ServiceID to first of payload
	SetID(p, id)

	return
}

// UnMarshalPacket use to decode automatically payload and stores the result
// in the value pointed to by s.
func UnMarshalPacket(p []byte, expectedMinLen int, s interface{}) (id uint64, err error) {
	err = CheckPacket(p, expectedMinLen)
	if err != nil {
		return
	}

	// Get ErrorID from payload
	id = GetID(p)

	// decode payload to s by syllab encoder
	err = syllab.UnMarshal(p[8:], s)

	return
}
