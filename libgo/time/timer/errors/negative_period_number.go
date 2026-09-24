/* For license and copyright information please see the LEGAL file in the code repository */

package timer_errs

var ErrNegativePeriodNumber errNegativePeriodNumber

type errNegativePeriodNumber struct{}

//memar:impl memar/identifier/mediatype/protocol.Field_MediaType
func (dt *errNegativePeriodNumber) MediaType() string {
	return domainBaseMediatype + "negative_period_number"
}
