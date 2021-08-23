/* For license and copyright information please see LEGAL file in repository */

package file

import (
	"bytes"
	"compress/gzip"
	"hash/crc32"
	"io"
	"mime"
	"os"
	"regexp"
	"strconv"

	"../codec"
	"../giti"
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
	Folder       *Folder
	Path         string // File location in FileSystems include file name
	FullName     string
	Name         string
	Extension    string
	mediaType    string
	State        uint8
	data         []byte
	compressType string
}

// File||Folder State
const (
	StateUnSet uint8 = iota
	StateUnChanged
	StateChanged
)

// Data return full file data
func (f *File) Data() []byte {
	if f.data == nil {
		f.Update()
	}
	return f.data
}

// SetData set full file data with given one
func (f *File) SetData(data []byte) {
	f.data = data
	f.State = StateChanged
}

// AppendData append given data to end of exiting data
func (f *File) AppendData(data []byte) {
	f.data = append(f.data, data...)
	f.State = StateChanged
}

// Update use to read||update file from disk.
func (f *File) Update() (err error) {
	f.data, err = os.ReadFile(f.Path)
	return
}

// Save use to write file to the file system!
func (f *File) Save() (err error) {
	// Just write changed file
	if f.State == StateChanged {
		// Indicate state to not change to don't overwrite it again!
		f.State = StateUnChanged
		err = os.WriteFile(f.Path, f.data, 0700)
	}
	return
}

// Rename rename file name and full name
func (f *File) Rename(newName string) {
	f.Name = newName
	f.FullName = f.Name + "." + f.Extension
}

// Rename rename file name and full name
func (f *File) RenameFullName(fullName string) {
	f.FullName = fullName
	for i := len(f.FullName) - 1; i >= 0; i-- {
		if f.FullName[i] == '.' {
			f.Name = f.FullName[:i]
			f.Extension = f.FullName[i+1:]
			f.mediaType = mime.TypeByExtension(f.FullName[i:])
		}
	}
}

// CheckAndFix ...
func (f *File) CheckAndFix() {
	f.mediaType = mime.TypeByExtension("." + f.Extension)
	f.FullName = f.Name + "." + f.Extension
}

// var md5Hasher = md5.New()

// GetHashOfData return hash of f.data
func (f *File) GetHashOfData() (hash string) {
	// Just want to differ two same file, So crc32 is more enough!
	// md5Hasher.Write(f.data)
	// hash = hex.EncodeToString(md5Hasher.Sum(nil))
	// md5Hasher.Reset()
	return strconv.FormatUint(uint64(crc32.ChecksumIEEE(f.data)), 10)
}

// AddHashToName add f.data hash to its name and full name.
func (f *File) AddHashToName() {
	f.Name += "-" + f.GetHashOfData()
	f.FullName = f.Name + "." + f.Extension
}

// Copy returns a copy of the file.
func (f *File) Copy() *File {
	var file = File{
		Folder:    f.Folder,
		Path:      f.Path,
		FullName:  f.FullName,
		Name:      f.Name,
		Extension: f.Extension,
		mediaType: f.mediaType,
		State:     f.State,
		data:      f.data,
	}
	return &file
}

// DeepCopy returns a deep copy of the file.
func (f *File) DeepCopy() *File {
	var file = File{
		Folder:    f.Folder,
		Path:      f.Path,
		FullName:  f.FullName,
		Name:      f.Name,
		Extension: f.Extension,
		mediaType: f.mediaType,
		State:     f.State,
		data:      make([]byte, len(f.data)),
	}
	copy(file.data, f.data)
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
		f.data = htmlMinify.ReplaceAll(f.data, []byte{})
		return
	}
	var err error
	f.data, err = m.Bytes(f.mediaType, f.data)
	if err != nil && giti.AppDebugMode {
		log.Warn("Minify -", f.FullName, "occur this error:", err)
	}
}

// Compress use to compress file data to mostly to use in serving file by servers.
func (f *File) Compress(compressType string) {
	// Check file type and compress just if it worth.
	switch f.Extension {
	case "png", "jpg", "gif", "jpeg", "mkv", "avi", "mp3", "mp4":
		return
	}

	f.compressType = compressType
	switch compressType {
	case codec.CompressTypeGZIP:
		var b bytes.Buffer
		var gz = gzip.NewWriter(&b)
		gz.Write(f.data)
		gz.Close()
		f.data = b.Bytes()
	}
}

// Replace replace given data in the file
func (f *File) ReplaceAll(old, new []byte) {
	f.data = bytes.ReplaceAll(f.data, old, new)
	f.State = StateChanged
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
			if (cap(f.data) - len(f.data)) < addedSize {
				f.data = append(f.data, make([]byte, addedSize)...)
			} else {
				// increase f.data len
				f.data = f.data[:len(f.data)+addedSize]
			}
		}
		totalAddedSize += addedSize

		copy(f.data[d.End+addedSize:], f.data[d.End:])
		copy(f.data[d.Start:], d.Data)

		if addedSize < 0 {
			// decrease f.data len
			f.data = f.data[:len(f.data)+addedSize]
		}
	}
}

/*
********** giti.Codec interface **********
 */

func (f *File) MediaType() string    { return f.mediaType }
func (f *File) CompressType() string { return f.compressType }

func (f *File) Decode(buf giti.Buffer) (err giti.Error) {
	var data = buf.GetUnread()
	err = f.UnMarshal(data)
	return
}

// Encode write file to given buf.
// Pass buf with enough cap. Make buf by r.Len() or grow it by buf.Grow(r.Len())
func (f *File) Encode(buf giti.Buffer) {
	buf.Write(f.data)
}

// Marshal enecodes and return whole file data!
func (f *File) Marshal() (data []byte) {
	return f.data
}

// MarshalTo enecodes whole file data to given data and return it with new len!
func (f *File) MarshalTo(data []byte) []byte {
	data = append(data, f.data...)
	return data
}

// UnMarshal parses and decodes data of given data to f *File.
func (f *File) UnMarshal(data []byte) (err giti.Error) {
	// TODO::: MediaType??
	f.data = data
	return
}

// ReadFrom decodes f *File data by read from given io.Reader!
func (f *File) ReadFrom(reader io.Reader) (n int64, err error) {
	var buf bytes.Buffer
	n, err = io.Copy(&buf, reader)
	f.data = buf.Bytes()
	return
}

// WriteTo enecodes f *File data and write it to given io.Writer!
// Declare to respect io.WriterTo interface!
func (f *File) WriteTo(w io.Writer) (totalWrite int64, err error) {
	var writeLength int
	writeLength, err = w.Write(f.data)
	totalWrite = int64(writeLength)
	return
}

// Len return length of file
func (f *File) Len() (ln int) {
	return len(f.data)
}
