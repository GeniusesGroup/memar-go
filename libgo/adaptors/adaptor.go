/* For license and copyright information please see the LEGAL file in the code repository */

package codec

import (
	"io"

	error_adaptor "memar/libgo/adaptors/error"
	error_p "memar/process/error/protocol"
)

type ReadAdaptor func(p []byte) (n int, err error)

func (ra ReadAdaptor) Read(p []byte) (n int, err error_p.Error) {
	var goErr error
	n, goErr = ra(p)
	err = error_adaptor.ToError(goErr)
	return
}

type ReaderAdaptor struct {
	io.Reader
}

func (ra ReaderAdaptor) Read(p []byte) (n int, err error_p.Error) {
	var goErr error
	n, goErr = ra.Reader.Read(p)
	err = error_adaptor.ToError(goErr)
	return
}
