/* For license and copyright information please see the LEGAL file in the code repository */

package tcp

import (
	"io"
)

/*
********** protocol.Buffer interface **********
 */

func (s *Stream) ReadFrom(reader io.Reader) (n int64, err error) { return }
func (s *Stream) WriteTo(w io.Writer) (totalWrite int64, err error) {
	var writeLen int
	var data, _ = s.recv.buf.Marshal()
	writeLen, err = w.Write(data)
	totalWrite = int64(writeLen)
	return
}
