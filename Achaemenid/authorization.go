/* For license and copyright information please see LEGAL file in repository */

package achaemenid

import "../errors"

// AccessControl : Use ABAC features for AccessControl fields.
// Must store arrays in sort for easy read and comparison
type AccessControl struct {
	AllowSocieties []uint32 // Part of GP address
	DenySocieties  []uint32 // Part of GP address
	AllowRouters   []uint32 // Part of GP address
	DenyRouters    []uint32 // Part of GP address
	AllowDays      []uint8
	DenyDays       []uint8
	AllowTime      []int64
	DenyTime       []int64
	AllowServices  []uint32 // ["ServiceID", "ServiceID"]
	DenyServices   []uint32 // ["ServiceID", "ServiceID"]
}

// Errors
var (
	ErrRequestServiceNotAllow = errors.New("RequestServiceNotAllow", "Request service is not in allow list of connection")
	ErrRequestServiceDenied   = errors.New("RequestServiceDenied", "Request service is in deny list of connection")

	ErrRequestByNotAllowSociety = errors.New("RequestByNotAllowSociety", "Request send by society that is not in allow list of connection")
	ErrRequestByDeniedSociety   = errors.New("RequestByDeniedSociety", "Request send by society that is in deny list of connection")

	ErrRequestByNotAllowRouter = errors.New("RequestByNotAllowRouter", "Request send by router that is not in allow list of connection")
	ErrRequestByDeniedRouter   = errors.New("RequestByDeniedRouter", "Request send by router that is in deny list of connection")
)

// AddDays store given days by check order.
func (ac *AccessControl) AddDays(allow, deny []uint8) {
}

// AddTime store given times by check order.
func (ac *AccessControl) AddTime(allow, deny []int64) {
	// Remove Useless Inner interval in When key in AccessControl.
	// e.g. 150000/160000 and 151010/153030 the second one is useless!
	// Iso8601 Time intervals <start>/<end> ["hhmmss/hhmmss", "hhmmss/hhmmss"]!!!
	// Just use GMT0!!!
}

func (ac *AccessControl) authorizeWhich(serviceID uint32) (err error) {
	var i int
	var notAuthorize bool

	var ln = len(ac.AllowServices)
	if ln > 0 {
		for i = 0; i < ln; i++ {
			if ac.AllowServices[i] == serviceID {
				goto DS
			} else {
				notAuthorize = true
			}
		}
		if notAuthorize {
			return ErrRequestServiceNotAllow
		}
	}

DS:
	ln = len(ac.DenyServices)
	if ln > 0 {
		for i = 0; i < ln; i++ {
			if ac.DenyServices[i] == serviceID {
				return ErrRequestServiceDenied
			}
		}
	}

	return
}

func (ac *AccessControl) authorizeWhen() (err error) {

	return
}

func (ac *AccessControl) authorizeWhere(societyID, RouterID uint32) (err error) {
	var i int
	var notAuthorize bool

	var ln = len(ac.AllowSocieties)
	if ln > 0 {
		for i = 0; i < ln; i++ {
			if ac.AllowSocieties[i] == societyID {
				goto DS
			} else {
				notAuthorize = true
			}
		}
		if notAuthorize {
			return ErrRequestByNotAllowSociety
		}
	}

DS:
	ln = len(ac.DenySocieties)
	if ln > 0 {
		for i = 0; i < ln; i++ {
			if ac.DenySocieties[i] == societyID {
				return ErrRequestByDeniedSociety
			}
		}
	}

	ln = len(ac.AllowRouters)
	if ln > 0 {
		for i = 0; i < ln; i++ {
			if ac.AllowSocieties[i] == RouterID {
				goto DR
			} else {
				notAuthorize = true
			}
		}
		if notAuthorize {
			return ErrRequestByNotAllowRouter
		}
	}

DR:
	ln = len(ac.DenyRouters)
	if ln > 0 {
		for i = 0; i < ln; i++ {
			if ac.DenyRouters[i] == RouterID {
				return ErrRequestByDeniedRouter
			}
		}
	}

	return
}

func (ac *AccessControl) authorizeWhat() (err error) {
	// TODO::: can be implement?

	// AllowRecords   [][16]byte // ["RecordUUID", "RecordUUID"]
	// DenyRecords    [][16]byte // ["RecordUUID", "RecordUUID"]
	return
}

func (ac *AccessControl) authorizeHow() (err error) {
	// TODO::: can be implement?

	// How []string
	return
}

func (ac *AccessControl) authorizeIf() (err error) {
	// TODO::: can be implement?

	// If  []string
	return
}
