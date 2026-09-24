/* For license and copyright information please see the LEGAL file in the code repository */

package timer_errs

var TimerAlreadyStarted errTimerAlreadyStarted

type errTimerAlreadyStarted struct{}

//memar:impl memar/identifier/mediatype/protocol.Field_MediaType
func (dt *errTimerAlreadyStarted) MediaType() string {
	return domainBaseMediatype + "timer_already_started"
}
