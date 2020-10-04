/* For license and copyright information please see LEGAL file in repository */

package authorization

// AccessControl store needed data to authorize a request to a platform!
// Use add methods due to arrays must store in sort for easy read and comparison
// Some standards: ABAC, ACL, RBAC, ... features for AccessControl fields,
type AccessControl struct {
	// Authorize Where:
	AllowSocieties []uint32 // Part of GP address
	DenySocieties  []uint32 // Part of GP address
	AllowRouters   []uint32 // Part of GP address
	DenyRouters    []uint32 // Part of GP address

	// Authorize When:
	AllowDays Day      // Any day of the week
	DenyDays  Day      // Any day of the week
	AllowTime []uint16 // Any Time of the Day in minute
	DenyTime  []uint16 // Any Time of the Day in minute

	// Authorize Which:
	AllowServices []uint32 // ["ServiceID", "ServiceID"]
	DenyServices  []uint32 // ["ServiceID", "ServiceID"]
	AllowCRUD     CRUD     // CRUD == Create, Read, Update, Delete
	DenyCRUD      CRUD     // CRUD == Create, Read, Update, Delete

	// Authorize What:
	// Authorize How:
	// Authorize If:
}

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

// AuthorizeWhich authorize by ServiceID and CRUD that allow or denied by access control table!
func (ac *AccessControl) AuthorizeWhich(serviceID uint32, crud CRUD) (err error) {
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
			return ErrAuthorizationServiceNotAllow
		}
	}

DS:
	ln = len(ac.DenyServices)
	if ln > 0 {
		for i = 0; i < ln; i++ {
			if ac.DenyServices[i] == serviceID {
				return ErrAuthorizationServiceDenied
			}
		}
	}

	if ac.AllowCRUD != CRUDAll {
		// TODO:::
	}

	if ac.DenyCRUD != CRUDNone {
		// TODO:::
	}

	return
}

// AuthorizeWhen --
func (ac *AccessControl) AuthorizeWhen(day Day, time int64) (err error) {

	return
}

// AuthorizeWhere --
func (ac *AccessControl) AuthorizeWhere(societyID, RouterID uint32) (err error) {
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
			return ErrAuthorizationNotAllowSociety
		}
	}

DS:
	ln = len(ac.DenySocieties)
	if ln > 0 {
		for i = 0; i < ln; i++ {
			if ac.DenySocieties[i] == societyID {
				return ErrAuthorizationDeniedSociety
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
			return ErrAuthorizationNotAllowRouter
		}
	}

DR:
	ln = len(ac.DenyRouters)
	if ln > 0 {
		for i = 0; i < ln; i++ {
			if ac.DenyRouters[i] == RouterID {
				return ErrAuthorizationDeniedRouter
			}
		}
	}

	return
}

// AuthorizeWhat --
func (ac *AccessControl) AuthorizeWhat() (err error) {
	// TODO::: can be implement?

	// AllowRecords   [][16]byte // ["RecordUUID", "RecordUUID"]
	// DenyRecords    [][16]byte // ["RecordUUID", "RecordUUID"]
	return
}

// AuthorizeHow --
func (ac *AccessControl) AuthorizeHow() (err error) {
	// TODO::: can be implement?

	// How []string
	return
}

// AuthorizeIf --
func (ac *AccessControl) AuthorizeIf() (err error) {
	// TODO::: can be implement?

	// If  []string
	return
}
