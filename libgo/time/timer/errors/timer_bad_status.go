/* For license and copyright information please see the LEGAL file in the code repository */

package timer_errs

var TimerBadStatus errTimerBadStatus

type errTimerBadStatus struct{}

//memar:impl memar/identifier/mediatype/protocol.Field_MediaType
func (dt *errTimerBadStatus) MediaType() string {
	return domainBaseMediatype + "timer_bad_status"
}
