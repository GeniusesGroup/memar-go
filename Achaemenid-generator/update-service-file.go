/* For license and copyright information please see LEGAL file in repository */

package generator

import (
	"hash/crc32"
	"strconv"
	"strings"

	"../assets"
)

// UpdateServiceFile use to update service file and complete or edit some auto generate part.
func UpdateServiceFile(file *assets.File) (err error) {
	var sn = strings.Title(file.Name)
	sn = strings.ReplaceAll(sn, "-", "")

	file.DataString = strings.Replace(file.DataString, "ID:                0", "ID:                "+strconv.FormatUint(uint64(crc32.ChecksumIEEE([]byte(sn))), 10), 1)

	// Check and update service detail essentially service ID!
	// Update handler function
	// Update encoders
	// Update validator

	const tagName = "valid"

	// Validate each data with tagName with related function from validators folder.
	// valid data can be required, optional, ...

	// Indicate file had been changed
	file.Status = assets.StateChanged
	file.Data = []byte(file.DataString)

	return
}
