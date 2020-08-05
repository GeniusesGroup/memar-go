/* For license and copyright information please see LEGAL file in repository */

package gsdk

import "../errors"

// Errors
var (
	ErrNoNodeAvailable = errors.New("NoNodeAvailable", "There isn't available node to handle requests")
)
