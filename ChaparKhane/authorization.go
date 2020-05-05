/* For license and copyright information please see LEGAL file in repository */

package chaparkhane

// AccessControl : Use ABAC features for AccessControl fields.
// Must store arrays in sort for easy read and comparison
type AccessControl struct {
	// Remove Useless Inner interval in When key in AccessControl.
	// e.g. 150000/160000 and 151010/153030 the second one is useless!
	// Iso8601 Time intervals <start>/<end> ["hhmmss/hhmmss", "hhmmss/hhmmss"]!!!
	// Just use GMT0!!!
	When  []uint64
	Where [][16]byte // ["UIP", "UIP"]!
	Which []uint32   // ["ServiceID", "ServiceID"] Just in specific AppID
	How   []string   //
	What  [][16]byte // ["RecordUUID", "RecordUUID"]
	If    []string   //
}

// AuthorizeWhich will authorize by Connection.AccessControl.Which
func (st *Stream) AuthorizeWhich() {
	// Check requested user have enough access
	var notAuthorize bool
	for _, service := range st.Connection.AccessControl.Which {
		if service == st.ServiceID {
			notAuthorize = false
			break
		} else {
			notAuthorize = true
		}
	}
	if notAuthorize == true {
		// err =
		return
	}
}

// AuthorizeWhen will authorize by Connection.AccessControl.When
func (st *Stream) AuthorizeWhen() {}

// AuthorizeWhere will authorize by Connection.AccessControl.Where
func (st *Stream) AuthorizeWhere() {
	var notAuthorize bool
	for _, ip := range st.Connection.AccessControl.Where {
		// TODO : ip may contain zero padding!! org may restricted user to isp not subnet nor even device!!
		if ip == st.Connection.UIPAddress {
			notAuthorize = false
			break
		} else {
			notAuthorize = true
		}
	}
	if notAuthorize == true {
		// err =
		return
	}
}
