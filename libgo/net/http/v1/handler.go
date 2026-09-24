/* For license and copyright information please see the LEGAL file in the code repository */

package http_v1

import (
	"memar/libgo/net/http"
	"memar/libgo/net/http/headers"
	http_p "memar/libgo/net/http/protocol"
	net_p "memar/net/protocol"
	error_p "memar/process/error/protocol"
	"memar/process/log"
	service_p "memar/process/service/protocol"
)

// Handler is application layer HTTPv1 protocol handler.
type Handler struct {
	mux http_p.Handler
}

//memar:impl memar/computer/capsule/protocol.LifeCycle
func (self *Handler) Init(mux http_p.Handler) (err error_p.Error) {
	self.mux = mux
	return
}
func (self *Handler) Reinit() (err error_p.Error) {
	// err = "To make process simple and prevent data racing, You can't reuse `Handler`"
	return
}
func (self *Handler) Deinit() (err error_p.Error) {
	// self.mux.Deinit()
	return
}

// HandleIncomeRequest handle incoming HTTP request streams.
// It can use for architectures like restful, ...
// ServeHTTP handle incoming HTTP request.
// Developers can develop new one with its desire logic to more respect protocols like restful, ...
func (self *Handler) HandleIncomeRequest(sk net_p.Socket) (err error_p.Error) {
	var httpReq Request
	var httpRes Response
	err = httpReq.Init()
	err = httpRes.Init()

	err = httpReq.Decode(sk.ReceiveBuffer())
	if err != nil {
		httpRes.SetStatusByError(err)
	} else {
		err = self.mux.ServeHTTP(sk, &httpReq.Request, &httpRes.Response)
	}

	self.HandleOutcomeResponse(sk, &httpReq, &httpRes)
	return
}

// HandleOutcomeResponse use to handle outcome HTTP response stream.
func (self *Handler) HandleOutcomeResponse(sk net_p.Socket, httpReq *Request, httpRes *Response) (err error_p.Error) {
	// Do some global assignment to response
	httpRes.SetVersion(httpReq.Version())
	httpRes.Header_Set(headers.Key_Server, http.DefaultServer)
	// TODO::: add server time?? is it worth??

	err = httpRes.Encode(sk.SendBuffer())

	if log.CNF_DevelopingMode {
		// TODO::: httpReq||httpRes not serialized yet to log it! any idea to have better performance below??
		var reqAsStr, _ = httpReq.ToString()
		var resAsStr, _ = httpRes.ToString()
		log.Conf(httpReq, reqAsStr)
		log.Conf(httpRes, resAsStr)
	}
	return
}

// SendBidirectionalRequest use to handle outgoing HTTP request.
// It block caller until get response or error.
// Caller must sk.Close() the given socket or pool it for other request.
func SendBidirectionalRequest(sk net_p.Socket, service service_p.Service, httpReq *http.Request) (httpRes *http.Response, err error_p.Error) {
	err = httpReq.Encode(sk.SendBuffer())
	if err != nil {
		return
	}

	var skStatus = sk.State()
	for {
		var status = <-skStatus
		switch status {
		case net_p.Status_Timeout_Read, net_p.Status_Timeout_Write:
			// err =
			return
		case net_p.Status_ReceivedCompletely:
			var res http.Response
			res.Init()
			err = res.Decode(sk.ReceiveBuffer())
			if err == nil {
				err = res.GetError()
			}
			httpRes = &res
			return
		default:
			continue
		}
	}
}

// SendUnidirectionalRequest use to send outcome HTTP request and don't expect any response.
// It block caller until request send successfully or return error
// Caller must sk.Close() the given socket or pool it for other request.
func SendUnidirectionalRequest(sk net_p.Socket, service service_p.Service, httpReq *http.Request) (err error_p.Error) {
	err = httpReq.Encode(sk.SendBuffer())
	if err != nil {
		return
	}

	var skStatus = sk.State()
	for {
		var status = <-skStatus
		switch status {
		case net_p.Status_Timeout_Read, net_p.Status_Timeout_Write:
			// err =
			return
		case net_p.Status_SentCompletely:
			// Nothing to do. Just let execution go to sk.Close() and break the loop
			return
		default:
			continue
		}
	}
}
