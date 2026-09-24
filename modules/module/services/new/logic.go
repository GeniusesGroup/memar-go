/* For license and copyright information please see the LEGAL file in the code repository */

package new

import (
	"memar/modules/module/internal/generator"
	mp "memar/modules/module/protocol"
	net_p "memar/net/protocol"
	error_p "memar/process/error/protocol"
)

func (s *serv) Process(sk net_p.Socket, req mp.BS_New_Request) (res Response, err error_p.Error) {
	var gen generator.New
	err = gen.Init(req.Path())
	if err != nil {
		return
	}

	res.dirsPath = gen.DirsPath()
	res.filesPath = gen.FilesPath()
	res.filesData = gen.FilesData()
	return
}
