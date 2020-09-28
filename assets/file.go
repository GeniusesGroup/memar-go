/* For license and copyright information please see LEGAL file in repository */

package assets

import (
	"bytes"
	"compress/gzip"
	"hash/crc32"
	"regexp"
	"strconv"

	"../log"

	"github.com/tdewolff/minify"
	"github.com/tdewolff/minify/css"
	"github.com/tdewolff/minify/html"
	"github.com/tdewolff/minify/js"
	"github.com/tdewolff/minify/json"
	"github.com/tdewolff/minify/svg"
	"github.com/tdewolff/minify/xml"
)

// File store some data about a file!
type File struct {
	FullName     string
	Name         string
	Extension    string
	MimeType     string
	Dep          *Folder
	State        uint8
	Data         []byte
	CompressData []byte
	CompressType string
}

// File||Folder State
const (
	StateUnChanged uint8 = iota
	StateChanged
)

// var md5Hasher = md5.New()

// GetHashOfData return hash of f.Data
func (f *File) GetHashOfData() (hash string) {
	// Just want to differ two same file, So crc32 is more enough!
	// md5Hasher.Write(f.Data)
	// hash = hex.EncodeToString(md5Hasher.Sum(nil))
	// md5Hasher.Reset()
	return strconv.FormatUint(uint64(crc32.ChecksumIEEE(f.Data)), 10)
}

// AddHashToName add f.Data hash to its name and full name.
func (f *File) AddHashToName() {
	f.Name += "-" + f.GetHashOfData()
	f.FullName = f.Name + "." + f.Extension
}

// Copy returns a copy of the file.
func (f *File) Copy() *File {
	var file = File{
		FullName:  f.FullName,
		Name:      f.Name,
		Extension: f.Extension,
		MimeType:  f.MimeType,
		Dep:       f.Dep,
		State:     f.State,
		Data:      f.Data,
	}
	return &file
}

// DeepCopy returns a deep copy of the file.
func (f *File) DeepCopy() *File {
	var file = File{
		FullName:  f.FullName,
		Name:      f.Name,
		Extension: f.Extension,
		MimeType:  f.MimeType,
		Dep:       f.Dep,
		State:     f.State,
	}
	copy(file.Data, f.Data)
	return &file
}

var m = minify.New()
var htmlMinify = regexp.MustCompile(`\r?\n|(<!--.*?-->)|(<!--[\w\W\n\s]+?-->)`)

func init() {
	m.AddFunc("text/css", css.Minify)
	m.AddFunc("text/html", html.Minify)
	m.AddFunc("image/svg+xml", svg.Minify)
	m.AddFuncRegexp(regexp.MustCompile("^(application|text)/(x-)?(java|ecma)script$"), js.Minify)
	m.AddFuncRegexp(regexp.MustCompile("[/+]json$"), json.Minify)
	m.AddFuncRegexp(regexp.MustCompile("[/+]xml$"), xml.Minify)
}

// Minify replace file data with minify of them.
func (f *File) Minify() {
	// TODO::: have problem minify HTML >> https://github.com/tdewolff/minify/issues/318
	if f.Extension == "html" {
		f.Data = htmlMinify.ReplaceAll(f.Data, []byte{})
		return
	}
	var err error
	f.Data, err = m.Bytes(f.MimeType, f.Data)
	if err != nil {
		log.Warn("Minify", f.FullName, "occur this error:", err)
	}
}

// Supported compress types
const (
	CompressTypeGZIP = "gzip"
)

// Compress use to compress file data to mostly to use in serving file by servers.
func (f *File) Compress(compressType string) {
	// Check file type and compress just if it worth.
	switch f.Extension {
	case "png", "jpg", "gif", "jpeg", "mkv", "avi", "mp3", "mp4":
		f.CompressData = f.Data
		return
	}

	f.CompressType = compressType
	switch compressType {
	case CompressTypeGZIP:
		var b bytes.Buffer
		var gz = gzip.NewWriter(&b)
		gz.Write(f.Data)
		gz.Close()
		f.CompressData = b.Bytes()
	}
}

// ReplaceReq is request structure of Replace method.
type ReplaceReq struct {
	Data  string
	Start int
	End   int
}

// Replace replace given data in the file
func (f *File) Replace(data []ReplaceReq) {
	var totalAddedSize, addedSize int
	for _, d := range data {
		d.Start += totalAddedSize
		d.End += totalAddedSize

		var ln = len(d.Data)
		addedSize = ln - (d.End - d.Start)
		if addedSize > 0 {
			if (cap(f.Data) - len(f.Data)) < addedSize {
				f.Data = append(f.Data, make([]byte, addedSize)...)
			} else {
				// increase f.Data len
				f.Data = f.Data[:len(f.Data)+addedSize]
			}
		}
		totalAddedSize += addedSize

		copy(f.Data[d.End+addedSize:], f.Data[d.End:])
		copy(f.Data[d.Start:], d.Data)

		if addedSize < 0 {
			// decrease f.Data len
			f.Data = f.Data[:len(f.Data)+addedSize]
		}
	}
}
