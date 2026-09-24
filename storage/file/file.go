/* For license and copyright information please see the LEGAL file in the code repository */

package file

import (
	"hash/crc32"
	"strconv"
	"strings"

	codec_p "memar/codec/protocol"
	"memar/codec/compress"
	cts "memar/codec/compress-types"
	compress_p "memar/codec/compress/protocol"
	error_p "memar/process/error/protocol"
	"memar/libgo/codec/minify"
	file_p "memar/storage/file/protocol"
)

func FindByRelativeFrom(file file_p.File, relativePath string) (desireFile file_p.File) {
	var locPart = strings.Split(relativePath, "/")
	if len(locPart) < 2 {
		return
	}
	var desireDir = file.Metadata().ParentDirectory()
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

func AddHashToFileName(file file_p.File) {
	var nameWithoutExtension = file.Metadata().FileNameWithoutExtension()
	var fileExt = file.Metadata().FileExtension()

	// Just want to differ two same file, So crc32 is more enough!
	// var md5Hasher = md5.New()
	// md5Hasher.Write(file.Data().Marshal())
	// hash = hex.EncodeToString(md5Hasher.Sum(nil))
	// md5Hasher.Reset()
	var fileData, _ = file.Data().Marshal()
	var hashOfFileData = strconv.FormatUint(uint64(crc32.ChecksumIEEE(fileData)), 10)

	file.Rename(nameWithoutExtension + "-" + hashOfFileData + "." + fileExt)
}

// Minify replace file data with minify of them if possible.
func Minify(file file_p.File) (err error_p.Error) {
	err = minify.Minify(file.Data())
	return
}

// Compress creates new files in the directory with desire compress algorithm
func Compress(file file_p.File, contentEncodings []string, options compress.Options) (err error_p.Error) {
	// Check file type and compress just if it worth.
	var fe = file.Metadata().FileExtension()
	switch fe {
	case "png", "jpg", "gif", "jpeg", "mkv", "avi", "mp3", "mp4":
		return
	}

	var parentDir = file.ParentDirectory()
	var ct compress_p.CompressType
	for _, ce := range contentEncodings {
		ct, err = cts.GetByContentEncoding(ce)
		if err != nil {
			continue
		}
		var compressor codec_p.Codec
		compressor, err = ct.Compress(file.Data(), options)
		if err != nil {
			return
		}
		var compressedFileName = file.Metadata().FileName() + "." + ct.FileExtension()
		var compressedFile file_p.File
		compressedFile, err = parentDir.File(compressedFileName)
		if err != nil {
			return
		}
		compressedFile.Data().Decode(compressor)
	}
	return
}

// ReplaceReq is request structure of Replace method.
type ReplaceReq struct {
	Data  string
	Start int
	End   int
}

// Replace replace given data in the file
func ReplaceLocation(file file_p.File, data []ReplaceReq) (err error_p.Error) {
	var fileData []byte
	fileData, err = file.Data().Marshal()
	if err != nil {
		return
	}

	var totalAddedSize, addedSize int
	for _, d := range data {
		d.Start += totalAddedSize
		d.End += totalAddedSize

		var ln = len(d.Data)
		addedSize = ln - (d.End - d.Start)
		if addedSize > 0 {
			if (cap(fileData) - len(fileData)) < addedSize {
				fileData = append(fileData, make([]byte, addedSize)...)
			} else {
				// increase fileData len
				fileData = fileData[:len(fileData)+addedSize]
			}
		}
		totalAddedSize += addedSize

		copy(fileData[d.End+addedSize:], fileData[d.End:])
		copy(fileData[d.Start:], d.Data)

		if addedSize < 0 {
			// decrease fileData len
			fileData = fileData[:len(fileData)+addedSize]
		}
	}

	_, err = file.Data().Unmarshal(fileData)
	return
}
