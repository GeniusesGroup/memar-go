/* For license and copyright information please see the LEGAL file in the code repository */

package initialize

import (
	"os"

	http_p "memar/libgo/net/http/protocol"
	net_p "memar/net/protocol"
	cmd "memar/process/command"
	error_p "memar/process/error/protocol"
)

func (ser *serv) ServeHTTP(sk net_p.Socket, httpReq http_p.Request, httpRes http_p.Response) (err error_p.Error) {
	return
}

func (ser *serv) ServeCLA(arguments []string) (err error_p.Error) {
	var req Request
	_, err = req.FromCLA(arguments)
	if err != nil {
		return
	}

	var res Response
	res, err = ser.Process(nil, &req)
	if err != nil {
		return
	}

	var cwd = cmd.Cwd()

	for i, path := range res.path {
		var file, goErr = os.OpenFile(cwd+path, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0666)
		if goErr != nil {
			// err =
			break
		}
		_, goErr = file.Write(res.files[i])
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
