/* For license and copyright information please see LEGAL file in repository */

package srpc

import "../syllab"

/*
	********************PAY ATTENTION:*******************
	We don't suggest use these 2 func instead use chaparkhane to autogenerate needed code before compile time
	and reduce runtime proccess to improve performance of the app and gain max performance from this protocol!
*/

// MarshalPacket use to encode automatically the value of s to the payload buffer.
func MarshalPacket(p []byte, id uint32, s interface{}) (err error) {
	// Set ServiceID to payload
	SetID(p, id)

	// encode s to p by syllab encoder
	err = syllab.MarshalSyllab(p[4:], s)

	return err
}

// UnMarshalPacket use to decode automatically payload and stores the result
// in the value pointed to by s.
func UnMarshalPacket(p []byte, expectedMinLen int, s interface{}) (id uint32, err error) {
	err = CheckPacket(p, expectedMinLen)
	if err != nil {
		return 0, err
	}

	// Get ErrorID from payload
	id = GetID(p)

	// decode payload to s by syllab encoder
	err = syllab.UnMarshalSyllab(p[4:], s)

	return id, err
}
