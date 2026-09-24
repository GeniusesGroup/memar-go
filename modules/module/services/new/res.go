/* For license and copyright information please see the LEGAL file in the code repository */

package new

// Response
type Response struct {
	dirsPath  []string
	filesPath  []string
	filesData [][]byte
}

func (res *Response) Dirs() []string      { return res.dirsPath }
func (res *Response) FilesPath() []string { return res.filesPath }
func (res *Response) FilesData() [][]byte { return res.filesData }
