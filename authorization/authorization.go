/* For license and copyright information please see LEGAL file in repository */

package authorization

import (
	etime "../earth-time"
	er "../error"
)

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
	AllowWeekdays etime.Weekdays // Any day of the week
	DenyWeekdays  etime.Weekdays // Any day of the week
	AllowDayhours etime.Dayhours // Any hour of the day
	DenyDayhours  etime.Dayhours // Any hour of the day

	// Authorize Which:
	AllowServices []uint32 // ["ServiceID", "ServiceID"]
	DenyServices  []uint32 // ["ServiceID", "ServiceID"]
	AllowCRUD     CRUD     // CRUD == Create, Read, Update, Delete
	DenyCRUD      CRUD     // CRUD == Create, Read, Update, Delete

	// Authorize What:
	// Authorize How:
	// Authorize If:
}

// GiveFullAccess set some data to given AccessControl to be full access
func (ac *AccessControl) GiveFullAccess() {
	ac.AllowSocieties = nil
	ac.DenySocieties = nil
	ac.AllowRouters = nil
	ac.DenyRouters = nil

	ac.AllowWeekdays = etime.WeekdaysAll
	ac.DenyWeekdays = etime.WeekdaysNone
	ac.AllowDayhours = etime.DayhoursAll
	ac.DenyDayhours = etime.DayhoursNone

	ac.AllowServices = nil
	ac.DenyServices = nil
	ac.AllowCRUD = CRUDAll
	ac.DenyCRUD = CRUDNone
}

// AuthorizeWhich authorize by ServiceID and CRUD that allow or denied by access control table!
func (ac *AccessControl) AuthorizeWhich(serviceID uint32, crud CRUD) (err *er.Error) {
	if !ac.AllowCRUD.Check(crud) {
		return ErrCrudNotAllow
	}

	if ac.DenyCRUD.Check(crud) {
		return ErrCRUDDenied
	}

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
			return ErrServiceNotAllow
		}
	}

DS:
	ln = len(ac.DenyServices)
	if ln > 0 {
		for i = 0; i < ln; i++ {
			if ac.DenyServices[i] == serviceID {
				return ErrServiceDenied
			}
		}
	}

	return
}

// AuthorizeWhen --
func (ac *AccessControl) AuthorizeWhen(day etime.Weekdays, hours etime.Dayhours) (err *er.Error) {
	if !ac.AllowWeekdays.Check(day) {
		return ErrDayNotAllow
	}

	if ac.DenyWeekdays.Check(day) {
		return ErrDayDenied
	}

	if !ac.AllowDayhours.Check(hours) {
		return ErrHourNotAllow
	}

	if ac.DenyDayhours.Check(hours) {
		return ErrHourDenied
	}
	return
}

// AuthorizeWhere --
func (ac *AccessControl) AuthorizeWhere(societyID, RouterID uint32) (err *er.Error) {
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
			return ErrNotAllowSociety
		}
	}

DS:
	ln = len(ac.DenySocieties)
	if ln > 0 {
		for i = 0; i < ln; i++ {
			if ac.DenySocieties[i] == societyID {
				return ErrDeniedSociety
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
			return ErrNotAllowRouter
		}
	}

DR:
	ln = len(ac.DenyRouters)
	if ln > 0 {
		for i = 0; i < ln; i++ {
			if ac.DenyRouters[i] == RouterID {
				return ErrDeniedRouter
			}
		}
	}

	return
}

// AuthorizeWhat --
func (ac *AccessControl) AuthorizeWhat() (err *er.Error) {
	// TODO::: can be implement?

	// AllowRecords   [][16]byte // ["RecordUUID", "RecordUUID"]
	// DenyRecords    [][16]byte // ["RecordUUID", "RecordUUID"]
	return
}

// AuthorizeHow --
func (ac *AccessControl) AuthorizeHow() (err *er.Error) {
	// TODO::: can be implement?

	// How []string
	return
}

// AuthorizeIf --
func (ac *AccessControl) AuthorizeIf() (err *er.Error) {
	// TODO::: can be implement?

	// If  []string
	return
}
