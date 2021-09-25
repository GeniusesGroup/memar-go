/* For license and copyright information please see LEGAL file in repository */

package file

import (
	"hash/crc32"
	"strconv"
	"strings"

	"../protocol"
)

func FindByRelativeFrom(file protocol.File, relativePath string) (desireFile protocol.File) {
	var locPart = strings.Split(relativePath, "/")
	if len(locPart) < 2 {
		return
	}
	var desireDir = file.ParentDirectory()
	for i := 0; i < len(locPart)-1; i++ { // -1 due have file name at end of locPart
		switch locPart[i] {
		case ".":
			// noting to do!
		case "..":
			desireDir = desireDir.ParentDirectory()
		default:
			desireDir, _ = desireDir.Directory(locPart[i])
			if desireDir == nil {
				// err =
				return
			}
		}
	}

	var fileName = locPart[len(locPart)-1]
	desireFile, _ = desireDir.File(fileName)
	return
}

func AddHashToFileName(file protocol.File) {
	var nameWithoutExtension = file.MetaData().URI().NameWithoutExtension()
	var fileExt = file.MetaData().URI().Extension()

	// Just want to differ two same file, So crc32 is more enough!
	// var md5Hasher = md5.New()
	// md5Hasher.Write(f.data)
	// hash = hex.EncodeToString(md5Hasher.Sum(nil))
	// md5Hasher.Reset()
	var hashOfFileData = strconv.FormatUint(uint64(crc32.ChecksumIEEE(file.Data().Marshal())), 10)

	file.Rename(nameWithoutExtension + "-" + hashOfFileData + "." + fileExt)
}
