/* For license and copyright information please see LEGAL file in repository */

package srpc

import "../syllab"

/*
type pingFrame struct {
	ID int64 // Usually unix time
}
*/
type pingFrame []byte

func (f pingFrame) ID() int64         { return syllab.GetInt64(f, 0) }
func (f pingFrame) NextFrame() []byte { return f[8:] }
