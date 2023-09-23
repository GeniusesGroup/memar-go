/* For license and copyright information please see the LEGAL file in the code repository */

package std

import (
	"io"
)

/*
********** protocol.Buffer interface **********
 */

func (sk *Socket) ReadFrom(reader io.Reader) (n int64, err error) { return }
func (sk *Socket) WriteTo(w io.Writer) (totalWrite int64, err error) {
	var writeLen int
	var data, _ = sk.Marshal()
	writeLen, err = w.Write(data)
	totalWrite = int64(writeLen)
	return
}
