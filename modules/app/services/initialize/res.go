/* For license and copyright information please see the LEGAL file in the code repository */

package initialize

// Response :
type Response struct {
	path  []string
	files [][]byte
}

func (res *Response) Path() []string  { return res.path }
func (res *Response) Files() [][]byte { return res.files }
