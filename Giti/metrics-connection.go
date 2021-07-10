/* For license and copyright information please see LEGAL file in repository */

package giti

// MetricsConnection
type MetricsConnection interface {
	ServiceCalled()
	ServiceCallFail()
}