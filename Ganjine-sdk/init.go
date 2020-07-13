/* For license and copyright information please see LEGAL file in repository */

package ganjine

import "../errors"

// Errors
var (
	ErrNoNodeAvailableToHandleRequests = errors.New("NoNodeAvailableToHandleRequests", "There isn't available node to handle requests")
)
