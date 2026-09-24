/* For license and copyright information please see the LEGAL file in the code repository */

package timer_errs

import (
	"memar/process/errors"
)

var TimerRacyAccess errTimerRacyAccess

type errTimerRacyAccess struct{ }

//memar:impl memar/identifier/mediatype/protocol.Field_MediaType
func (dt *errTimerRacyAccess) MediaType() string {
	return domainBaseMediatype + "timer_racy_access"
}
