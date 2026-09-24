/* For license and copyright information please see the LEGAL file in the code repository */

package timer_errs

var TimerNotInit errTimerNotInit

type errTimerNotInit struct{}

//memar:impl memar/identifier/mediatype/protocol.Field_MediaType
func (dt *errTimerNotInit) MediaType() string {
	return domainBaseMediatype + "timer_not_initialize"
}
