/* For license and copyright information please see the LEGAL file in the code repository */

package services

import (
	file_p "memar/storage/file/protocol"
)

// UpdateServiceFile use to update service file and complete or edit some auto generate part.
func UpdateServiceFile(file file_p.File) (err error) {
	// Update handler function
	// Update encoders
	// Update validator

	const tagName = "valid"

	// Validate each data with tagName with related function from validators folder.
	// valid data can be required, optional, ...

	return
}
