/* For license and copyright information please see LEGAL file in repository */

package dos

import (
	"bytes"
	"compress/gzip"
	"io"
	"mime"
	goos "os"

	"../../codec"
	"../../minify"
	"../../protocol"
)

// File store some data about a file!
type File struct {
	metadata        fileMetaData
	parentDirectory *FileDirectory
	mediaType       string
	compressType    string
	data            []byte
}

func (f *File) MetaData() protocol.FileMetaData         { return &f.metadata }
func (f *File) Data() protocol.FileData                 { return f }
func (f *File) ParentDirectory() protocol.FileDirectory { return f.parentDirectory }

func (f *File) init(path string) (err protocol.Error) {
	f.metadata.uri.Init(path)
	f.mediaType = mime.TypeByExtension("." + f.metadata.uri.Extension())
	return
}

// Save use to write file to the file system!
func (f *File) Save() (err protocol.Error) {
	// Just write changed file
	if f.data != nil {
		var goErr = goos.WriteFile(f.metadata.uri.Path(), f.data, 0700)
		if goErr != nil {
			// err =
		}
	}
	return
}

// Rename rename file name and full name
func (f *File) Rename(newName string) {
	delete(f.parentDirectory.files, f.metadata.uri.Name())
	f.parentDirectory.files[newName] = f
	f.metadata.uri.Rename(newName)
}

// Minify replace file data with minify of them if possible.
func (f *File) Minify() (err protocol.Error) {
	err = minify.Minify(f)
	return
}

// Compress use to compress file data to mostly to use in serving file by servers.
func (f *File) Compress(compressType string) {
	// Check file type and compress just if it worth.
	switch f.metadata.uri.Extension() {
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

// Replace replace all old data in the file with new one
func (f *File) Replace(old, new []byte, n int) {
	f.data = bytes.Replace(f.data, old, new, n)
}

// ReplaceReq is request structure of Replace method.
type ReplaceReq struct {
	Data  string
	Start int
	End   int
}

// Replace replace given data in the file
func (f *File) ReplaceLocation(data []ReplaceReq) {
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
********** protocol.FileData interface **********
 */

// Prepend add given data to start of exiting data that add earlier
func (f *File) Prepend(data []byte) {
	if f.data != nil {
		f.data = append(data, f.data...)
	} else {
		// TODO:::
	}
}

// Append append given data to end of exiting data that add earlier
func (f *File) Append(data []byte) {
	if f.data != nil {
		f.data = append(f.data, data...)
	} else {
		var goErr error
		var file *goos.File
		file, goErr = goos.OpenFile(f.metadata.URI().URI(), goos.O_APPEND|goos.O_WRONLY, 0700)
		if goErr != nil {
			// if goos.IsNotExist(goErr) {
			// 	return ErrStorageNotExist
			// }
			// if goos.IsPermission(goErr) {
			// 	return ErrStorageNotAuthorize
			// }
			// return ErrStorageDeviceProblem
			return
		}

		_, goErr = file.Write(data)
		var err1 = file.Close()
		if goErr == nil {
			goErr = err1
		}
	}
}

/*
********** protocol.Codec interface **********
 */

func (f *File) MediaType() string    { return f.mediaType }
func (f *File) CompressType() string { return f.compressType }

func (f *File) Decode(reader io.Reader) (err protocol.Error) {
	var buf bytes.Buffer
	var _, goErr = io.Copy(&buf, reader)
	if goErr != nil {
		// err =
	}
	err = f.Unmarshal(buf.Bytes())
	return
}

// Encode write file to given writer
func (f *File) Encode(writer io.Writer) (err error) { _, err = f.WriteTo(writer); return }

// Marshal enecodes and return whole file data!
func (f *File) Marshal() (data []byte) {
	if f.data != nil {
		return f.data
	}
	data, _ = goos.ReadFile(f.metadata.uri.Path())
	return
}

// MarshalTo enecodes whole file data to given data and return it with new len!
func (f *File) MarshalTo(data []byte) []byte {
	data = append(data, f.data...)
	return data
}

// Unmarshal save given data to the file. It will overwritten exiting data.
func (f *File) Unmarshal(data []byte) (err protocol.Error) {
	// TODO::: MediaType??
	f.data = data
	return
}

// Len return length of file
func (f *File) Len() (ln int) { return len(f.data) }

/*
********** io package interfaces **********
 */

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
