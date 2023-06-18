/* For license and copyright information please see the LEGAL file in the code repository */

package authorization

import (
	"libgo/json"
	"libgo/protocol"
	"libgo/syllab"
	"libgo/time/utc"
)

// AccessControl store needed data to authorize a request to a platform!
// Use add methods due to arrays must store in sort for easy read and comparison
// Some protocols: ABAC, ACL, RBAC, ... features for AccessControl fields,
type AccessControl struct {
	// Authorize Where:
	AllowSocieties []uint32 // Part of GP address
	DenySocieties  []uint32 // Part of GP address
	AllowRouters   []uint32 // Part of GP address
	DenyRouters    []uint32 // Part of GP address

	// Authorize When:
	AllowWeekdays utc.Weekdays // Any day of the week
	DenyWeekdays  utc.Weekdays // Any day of the week
	AllowDayHours utc.DayHours // Any hour of the day
	DenyDayHours  utc.DayHours // Any hour of the day

	// Authorize Which:
	AllowServices []uint64      // ["ServiceID", "ServiceID"]
	DenyServices  []uint64      // ["ServiceID", "ServiceID"]
	AllowCRUD     protocol.CRUD // CRUD == Create, Read, Update, Delete
	DenyCRUD      protocol.CRUD // CRUD == Create, Read, Update, Delete

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

	ac.AllowWeekdays = utc.Weekdays_All
	ac.DenyWeekdays = utc.Weekdays_None
	ac.AllowDayHours = utc.DayHours_All
	ac.DenyDayHours = utc.DayHours_None

	ac.AllowServices = nil
	ac.DenyServices = nil
	ac.AllowCRUD = protocol.CRUD_All
	ac.DenyCRUD = protocol.CRUD_None
}

// AuthorizeWhich authorize by ServiceID and CRUD that allow or denied by access control table!
func (ac *AccessControl) AuthorizeWhich(serviceID protocol.MediaTypeID, crud protocol.CRUD) (err protocol.Error) {
	if !CheckCrud(ac.AllowCRUD, crud) {
		return &ErrCrudNotAllow
	}

	if CheckCrud(ac.DenyCRUD, crud) {
		return &ErrCRUDDenied
	}

	var i int
	var notAuthorize bool

	var ln = len(ac.AllowServices)
	if ln > 0 {
		for i = 0; i < ln; i++ {
			if ac.AllowServices[i] == uint64(serviceID) {
				goto DS
			} else {
				notAuthorize = true
			}
		}
		if notAuthorize {
			return &ErrServiceNotAllow
		}
	}

DS:
	ln = len(ac.DenyServices)
	if ln > 0 {
		for i = 0; i < ln; i++ {
			if ac.DenyServices[i] == uint64(serviceID) {
				return &ErrServiceDenied
			}
		}
	}

	return
}

// AuthorizeWhen --
func (ac *AccessControl) AuthorizeWhen(day utc.Weekdays, hours utc.DayHours) (err protocol.Error) {
	if !ac.AllowWeekdays.Check(day) {
		return &ErrDayNotAllow
	}

	if ac.DenyWeekdays.Check(day) {
		return &ErrDayDenied
	}

	if !ac.AllowDayHours.Check(hours) {
		return &ErrHourNotAllow
	}

	if ac.DenyDayHours.Check(hours) {
		return &ErrHourDenied
	}
	return
}

// AuthorizeWhere --
func (ac *AccessControl) AuthorizeWhere(societyID, RouterID uint32) (err protocol.Error) {
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
			return &ErrNotAllowSociety
		}
	}

DS:
	ln = len(ac.DenySocieties)
	if ln > 0 {
		for i = 0; i < ln; i++ {
			if ac.DenySocieties[i] == societyID {
				return &ErrDeniedSociety
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
			return &ErrNotAllowRouter
		}
	}

DR:
	ln = len(ac.DenyRouters)
	if ln > 0 {
		for i = 0; i < ln; i++ {
			if ac.DenyRouters[i] == RouterID {
				return &ErrDeniedRouter
			}
		}
	}

	return
}

// AuthorizeWhat --
func (ac *AccessControl) AuthorizeWhat() (err protocol.Error) {
	// TODO::: can be implement?

	// AllowRecords   [][16]byte // ["RecordUUID", "RecordUUID"]
	// DenyRecords    [][16]byte // ["RecordUUID", "RecordUUID"]
	return
}

// AuthorizeHow --
func (ac *AccessControl) AuthorizeHow() (err protocol.Error) {
	// TODO::: can be implement?

	// How []string
	return
}

// AuthorizeIf --
func (ac *AccessControl) AuthorizeIf() (err protocol.Error) {
	// TODO::: can be implement?

	// If  []string
	return
}

//libgo:impl libgo/protocol.Syllab
//libgo:syllab
func (ac *AccessControl) FromSyllab(buf []byte, stackIndex uint32) {
	ac.AllowSocieties = syllab.GetUInt32Array(buf, stackIndex)
	ac.DenySocieties = syllab.GetUInt32Array(buf, stackIndex+8)
	ac.AllowRouters = syllab.GetUInt32Array(buf, stackIndex+16)
	ac.DenyRouters = syllab.GetUInt32Array(buf, stackIndex+24)

	ac.AllowWeekdays = utc.Weekdays(syllab.GetUInt8(buf, stackIndex+32))
	ac.DenyWeekdays = utc.Weekdays(syllab.GetUInt8(buf, stackIndex+33))
	ac.AllowDayHours = utc.DayHours(syllab.GetUInt32(buf, stackIndex+34))
	ac.DenyDayHours = utc.DayHours(syllab.GetUInt32(buf, stackIndex+38))

	ac.AllowServices = syllab.GetUInt64Array(buf, stackIndex+42)
	ac.DenyServices = syllab.GetUInt64Array(buf, stackIndex+50)
	ac.AllowCRUD = protocol.CRUD(syllab.GetUInt8(buf, stackIndex+58))
	ac.DenyCRUD = protocol.CRUD(syllab.GetUInt8(buf, stackIndex+59))
}
func (ac *AccessControl) ToSyllab(buf []byte, stackIndex, heapIndex uint32) (nextHeapAddr uint32) {
	heapIndex = syllab.SetUInt32Array(buf, ac.AllowSocieties, stackIndex, heapIndex)
	heapIndex = syllab.SetUInt32Array(buf, ac.DenySocieties, stackIndex+8, heapIndex)
	heapIndex = syllab.SetUInt32Array(buf, ac.AllowRouters, stackIndex+16, heapIndex)
	heapIndex = syllab.SetUInt32Array(buf, ac.DenyRouters, stackIndex+24, heapIndex)

	syllab.SetUInt8(buf, stackIndex+32, uint8(ac.AllowWeekdays))
	syllab.SetUInt8(buf, stackIndex+33, uint8(ac.DenyWeekdays))
	syllab.SetUInt32(buf, stackIndex+34, uint32(ac.AllowDayHours))
	syllab.SetUInt32(buf, stackIndex+38, uint32(ac.DenyDayHours))

	heapIndex = syllab.SetUInt64Array(buf, ac.AllowServices, stackIndex+42, heapIndex)
	heapIndex = syllab.SetUInt64Array(buf, ac.DenyServices, stackIndex+50, heapIndex)
	syllab.SetUInt8(buf, stackIndex+58, uint8(ac.AllowCRUD))
	syllab.SetUInt8(buf, stackIndex+59, uint8(ac.DenyCRUD))
	return heapIndex
}
func (ac *AccessControl) LenOfSyllabStack() uint32 {
	return 60 // 8+8+8+8+ 1+1+4+4+ 8+8+1+1
}
func (ac *AccessControl) LenOfSyllabHeap() (ln uint32) {
	ln += uint32(len(ac.AllowSocieties) * 4)
	ln += uint32(len(ac.DenySocieties) * 4)
	ln += uint32(len(ac.AllowRouters) * 4)
	ln += uint32(len(ac.DenyRouters) * 4)
	ln += uint32(len(ac.AllowServices) * 8)
	ln += uint32(len(ac.DenyServices) * 8)
	return
}
func (ac *AccessControl) LenAsSyllab() uint64 {
	return uint64(ac.LenOfSyllabStack() + ac.LenOfSyllabHeap())
}

//libgo:impl libgo/protocol.JSON
func (ac *AccessControl) FromJSON(payload []byte) (remaining []byte, err  protocol.Error) {
	var dec json.DecoderUnsafeMinified
	err = dec.Init(payload)
	if err != nil {
		return
	}
	err = ac.JSONDecoder(&dec)
	remaining = dec.Buf()
	return
}
func (ac *AccessControl) ToJSON(payload []byte) (added []byte, err protocol.Error) {
	var encoder json.Encoder
	err = encoder.Init(payload)
	if err != nil {
		return
	}
	err = ac.JSONEncoder(&encoder)
	added = encoder.Buf()
	return
}
func (ac *AccessControl) LenAsJSON() (ln int) {
	ln = len(ac.AllowSocieties) * 11
	ln += len(ac.DenySocieties) * 11
	ln += len(ac.AllowRouters) * 11
	ln += len(ac.DenyRouters) * 11
	ln += len(ac.AllowServices) * 21
	ln += len(ac.DenyServices) * 21
	ln += 247
	return
}

//libgo:json
func (ac *AccessControl) JSONDecoder(decoder *json.DecoderUnsafeMinified) (err protocol.Error) {
	for err == nil {
		var keyName = decoder.DecodeKey()
		switch keyName {
		case "AllowSocieties":
			ac.AllowSocieties, err = decoder.DecodeUInt32SliceAsNumber()
		case "DenySocieties":
			ac.DenySocieties, err = decoder.DecodeUInt32SliceAsNumber()
		case "AllowRouters":
			ac.AllowRouters, err = decoder.DecodeUInt32SliceAsNumber()
		case "DenyRouters":
			ac.DenyRouters, err = decoder.DecodeUInt32SliceAsNumber()

		case "AllowWeekdays":
			var num uint8
			num, err = decoder.DecodeUInt8()
			ac.AllowWeekdays = utc.Weekdays(num)
		case "DenyWeekdays":
			var num uint8
			num, err = decoder.DecodeUInt8()
			ac.DenyWeekdays = utc.Weekdays(num)
		case "AllowDayHours":
			var num uint8
			num, err = decoder.DecodeUInt8()
			ac.AllowDayHours = utc.DayHours(num)
		case "DenyDayHours":
			var num uint8
			num, err = decoder.DecodeUInt8()
			ac.DenyDayHours = utc.DayHours(num)

		case "AllowServices":
			ac.AllowServices, err = decoder.DecodeUInt64SliceAsNumber()
		case "DenyServices":
			ac.DenyServices, err = decoder.DecodeUInt64SliceAsNumber()
		case "AllowCRUD":
			var num uint8
			num, err = decoder.DecodeUInt8()
			ac.AllowCRUD = protocol.CRUD(num)
		case "DenyCRUD":
			var num uint8
			num, err = decoder.DecodeUInt8()
			ac.DenyCRUD = protocol.CRUD(num)

		default:
			err = decoder.NotFoundKey()
		}

		decoder.FindEndToken()
		if decoder.CheckToken('}') {
			decoder.Offset(1)
			return
		}
	}
	return
}
func (ac *AccessControl) JSONEncoder(encoder *json.Encoder) (err protocol.Error) {
	encoder.EncodeString(`{"AllowSocieties":[`)
	encoder.EncodeUInt32SliceAsNumber(ac.AllowSocieties)

	encoder.EncodeString(`],"DenySocieties":[`)
	encoder.EncodeUInt32SliceAsNumber(ac.DenySocieties)

	encoder.EncodeString(`],"AllowRouters":[`)
	encoder.EncodeUInt32SliceAsNumber(ac.AllowRouters)

	encoder.EncodeString(`],"DenyRouters":[`)
	encoder.EncodeUInt32SliceAsNumber(ac.DenyRouters)

	encoder.EncodeString(`],"AllowWeekdays":`)
	encoder.EncodeUInt8(uint8(ac.AllowWeekdays))

	encoder.EncodeString(`,"DenyWeekdays":`)
	encoder.EncodeUInt8(uint8(ac.DenyWeekdays))

	encoder.EncodeString(`,"AllowDayHours":`)
	encoder.EncodeUInt64(uint64(ac.AllowDayHours))

	encoder.EncodeString(`,"DenyDayHours":`)
	encoder.EncodeUInt64(uint64(ac.DenyDayHours))

	encoder.EncodeString(`,"AllowServices":[`)
	encoder.EncodeUInt64SliceAsNumber(ac.AllowServices)

	encoder.EncodeString(`],"DenyServices":[`)
	encoder.EncodeUInt64SliceAsNumber(ac.DenyServices)

	encoder.EncodeString(`],"AllowCRUD":`)
	encoder.EncodeUInt8(uint8(ac.AllowCRUD))

	encoder.EncodeString(`,"DenyCRUD":`)
	encoder.EncodeUInt8(uint8(ac.DenyCRUD))

	encoder.EncodeByte('}')
	return
}

//libgo:impl libgo/protocol.CommandLineArguments
