/* For license and copyright information please see the LEGAL file in the code repository */

package authorization

import (
	er "libgo/error"
)

// Errors
var (
	// User
	ErrUserNotAllow     er.Error
	ErrUserNotOwnRecord er.Error

	// Society
	ErrNotAllowSociety er.Error
	ErrDeniedSociety   er.Error

	// Router
	ErrNotAllowRouter er.Error
	ErrDeniedRouter   er.Error

	// Day
	ErrDayNotAllow er.Error
	ErrDayDenied   er.Error

	// Hour
	ErrHourNotAllow er.Error
	ErrHourDenied   er.Error

	// Service
	ErrServiceNotAllow er.Error
	ErrServiceDenied   er.Error

	// CRUD
	ErrCrudNotAllow er.Error
	ErrCRUDDenied   er.Error

	// Delegate
	ErrNotAllowToDelegate    er.Error
	ErrNotAllowToNotDelegate er.Error
)

func init() {
	// User
	ErrUserNotAllow.Init("domain/authorization.protocol; type=error; name=user-not-allow")
	ErrUserNotOwnRecord.Init("domain/authorization.protocol; type=error; name=user-not-own-record")

	// Society
	ErrNotAllowSociety.Init("domain/authorization.protocol; type=error; name=not-allow-society")
	ErrDeniedSociety.Init("domain/authorization.protocol; type=error; name=denied-society")

	// Router
	ErrNotAllowRouter.Init("domain/authorization.protocol; type=error; name=not-allow-router")
	ErrDeniedRouter.Init("domain/authorization.protocol; type=error; name=denied-router")

	// Day
	ErrDayNotAllow.Init("domain/authorization.protocol; type=error; name=day-not-allow")
	ErrDayDenied.Init("domain/authorization.protocol; type=error; name=day-denied")

	// Hour
	ErrHourNotAllow.Init("domain/authorization.protocol; type=error; name=hour-not-allow")
	ErrHourDenied.Init("domain/authorization.protocol; type=error; name=hour-denied")

	// Service
	ErrServiceNotAllow.Init("domain/authorization.protocol; type=error; name=service-not-allow")
	ErrServiceDenied.Init("domain/authorization.protocol; type=error; name=service-denied")

	// CRUD
	ErrCrudNotAllow.Init("domain/authorization.protocol; type=error; name=crud-not-allow")
	ErrCRUDDenied.Init("domain/authorization.protocol; type=error; name=crud-denied")

	// Delegate
	ErrNotAllowToDelegate.Init("domain/authorization.protocol; type=error; name=not-allow-to-delegate")
	ErrNotAllowToNotDelegate.Init("domain/authorization.protocol; type=error; name=not-allow-to-not-delegate")
}
