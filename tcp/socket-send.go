/* For license and copyright information please see LEGAL file in repository */

package tcp

// Send Sequence Space
//
//                    1         2          3          4
//               ----------|----------|----------|----------
//                      SND.UNA    SND.NXT    SND.UNA
//                                           +SND.WND
//
//         1 - old sequence numbers which have been acknowledged
//         2 - sequence numbers of unacknowledged data
//         3 - sequence numbers allowed for new data transmission
//         4 - future sequence numbers which are not yet allowed
type sendSequenceSpace struct {
	una  uint32 // send unacknowledged
	next uint32
	wnd  uint16 // send window
	up   bool   // send urgent pointer
	wl1  uint32 // segment sequence number used for last window update
	wl2  uint32 // segment acknowledgment number used for last window update
	iss  uint32 // initial send sequence number
	// buf    []byte Don't need it, because we don't need to copy buffer between kernel and userpspace
}
