/* For license and copyright information please see the LEGAL file in the code repository */

package protocol

type Validation interface {
	Validate() (err Error)
}
