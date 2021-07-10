/* For license and copyright information please see LEGAL file in repository */

package giti

// ReaderFrom is the interface that wraps the ReadFrom method.
//
// ReadFrom reads data from r until EOF or error.
// The return value n is the number of bytes read.
// Any error except EOF encountered during the read is also returned.
type ReaderFrom interface {
	ReadFrom(p []byte) (n int64, err Error)
}

// WriterTo is the interface that wraps the WriteTo & Len methods.
//
// WriteTo writes data to p until there's no more data to write as Len() or
// when an error occurs. Any error encountered during the write is also returned.
// Len return value n is the number of bytes that will written.
type WriterTo interface {
	WriteTo(p []byte) (err Error)
	Len() (n int)
}
