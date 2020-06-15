/* For license and copyright information please see LEGAL file in repository */

package generator

import (
	"strconv"
	"strings"

	"../assets"
)

// UpdateDatastoreFile use to update datastore file and complete or edit some auto generate part.
func UpdateDatastoreFile(file *assets.File) (err error) {
	var sn = strings.Title(file.Name)
	sn = strings.ReplaceAll(sn, "-", "")

	// make hash64 ready to generate crc64 of sn for its ID
	hash64.Reset()
	hash64.Write([]byte(sn))

	file.DataString = strings.Replace(file.DataString, "StructureID uint64 = 0", "StructureID uint64 = "+strconv.FormatUint(hash64.Sum64(), 10), 1)

	// Indicate file had been changed
	file.State = assets.StateChanged
	file.Data = []byte(file.DataString)

	return
}
