/* For license and copyright information please see the LEGAL file in the code repository */

package new

import (
	"os"

	http_p "memar/libgo/net/http/protocol"
	net_p "memar/net/protocol"
	cmd "memar/process/command"
	error_p "memar/process/error/protocol"

	moderr "memar/modules/module/errors"
)

func (s *serv) ServeHTTP(sk net_p.Socket, httpReq http_p.Request, httpRes http_p.Response) (err error_p.Error) {
	// err =
	return
}

func (s *serv) ServeCLA(arguments []string) (err error_p.Error) {
	var req Request
	_, err = req.FromCLA(arguments)
	if err != nil {
		return
	}

	if req.Path() == "" {
		var cwd = cmd.Cwd()
		req.SetPath(cwd)
	}

	var res Response
	res, err = s.Process(nil, &req)
	if err != nil {
		return
	}

	var goErr error
	for _, path := range res.Dirs() {
		goErr = os.MkdirAll(path, os.ModeDir)
		if goErr != nil {
			err = &moderr.ErrCanNotCreateDir
			return
		}
	}

	for i, path := range res.FilesPath() {
		var file, goErr = os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0666)
		if goErr != nil {
			// err =
			break
		}
		_, goErr = file.Write(res.FilesData()[i])
		var goErr1 = file.Close()
		if goErr1 != nil {
			if goErr != nil {
				// err =
			} else {
				// err =
			}
			break
		}
	}

	return
}
