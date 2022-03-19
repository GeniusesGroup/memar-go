/* For license and copyright information please see LEGAL file in repository */

package raw

import (
	"io"

	compress ".."
	"../../protocol"
)

type comDecom struct {
	data    []byte
	reader  protocol.Reader
	readLen int
}

/*
********** protocol.Codec interface **********
 */

func (r *comDecom) MediaType() protocol.MediaType       { return nil }
func (r *comDecom) CompressType() protocol.CompressType { return &RAW }
func (r *comDecom) Len() (ln int)                       { return len(r.data) }

func (r *comDecom) Decode(reader protocol.Reader) (err protocol.Error) {
	if r.data == nil && r.readLen > 0 {
		r.data = make([]byte, r.readLen)
		io.ReadFull(reader, r.data)
	}
	return
}
func (r *comDecom) Encode(writer protocol.Writer) (err protocol.Error) {
	var _, goErr = r.WriteTo(writer)
	if goErr != nil {
		// err =
	}
	return
}
func (r *comDecom) Marshal() (data []byte)       { return r.data }
func (r *comDecom) MarshalTo(data []byte) []byte { return append(data, r.data...) }
func (r *comDecom) Unmarshal(data []byte) (err protocol.Error) {
	err = compress.ErrSourceNotChangeable
	return
}
func (r *comDecom) UnmarshalFrom(data []byte) (remaining []byte, err protocol.Error) {
	err = compress.ErrSourceNotChangeable
	return
}

/*
********** io package interfaces **********
 */

func (r *comDecom) ReadFrom(reader io.Reader) (n int64, err error) {
	err = compress.ErrSourceNotChangeable
	return
}
func (r *comDecom) WriteTo(w io.Writer) (totalWrite int64, err error) {
	var writeLen int
	writeLen, err = w.Write(r.data)
	totalWrite = int64(writeLen)
	return
}
