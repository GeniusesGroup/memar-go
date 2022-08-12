/* For license and copyright information please see the LEGAL file in the code repository */

package tcp

import (
	"io"
)

/*
********** protocol.Buffer interface **********
 */

func (s *Socket) ReadFrom(reader io.Reader) (n int64, err error) { return }
func (s *Socket) WriteTo(w io.Writer) (totalWrite int64, err error) {
	var writeLen int
	writeLen, err = w.Write(s.recv.buf)
	totalWrite = int64(writeLen)
	return
}
