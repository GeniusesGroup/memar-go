/* For license and copyright information please see the LEGAL file in the code repository */

package srpc

import (
	buffer_p "memar/buffer/protocol"
	codec_p "memar/codec/protocol"
	error_p "memar/process/error/protocol"
	net_p "memar/net/protocol"
	srpc_p "memar/net/srpc/protocol"
	service_p "memar/process/service/protocol"
)

// Handler ...
// Read more about this protocol : https://github.com/GeniusesGroup/memar/blob/main/sRPC.md
type Handler struct{}

// HandleIncomeRequest handle incoming sRPC request streams.
func (h *Handler) HandleIncomeRequest(sk net_p.Socket) (err error_p.Error) {
	var service = sk.OSI_ApplicationLayer().Service()
	if service == nil {
		// TODO:::
		return
	}

	// TODO::: can't easily call service and must schedule it by its weight.
	var ser, ok = service.(srpc_p.Handler)
	if ok {
		err = ser.ServeSRPC(sk)
		if err != nil {
			sk.OSI_ApplicationLayer().SetError(err)
		}
	} else {
		// TODO:::
	}
	return
}

// SendBidirectionalRequest use to send outcoming sRPC request.
// It block caller until get response or error.
// Caller must pool sk or close it.
func SendBidirectionalRequest(sk net_p.Socket, service service_p.Service, req codec_p.Codec) (res buffer_p.Buffer, err error_p.Error) {
	// TODO::: send service frame first

	// stream.SendRequest(syllab.NewCodec(req))
	err = req.Encode(sk.SendBuffer())
	if err != nil {
		return
	}

	for status := range sk.State() {
		switch status {
		case net_p.Status_Timeout_Read, net_p.Status_Timeout_Write:
			// err =
		case net_p.Status_ReceivedCompletely:
			res = sk.ReceiveBuffer()
			err = sk.OSI_ApplicationLayer().Error()
		default:
			continue
		}
		break
	}
	return
}

// SendUnidirectionalRequest use to send outcoming HTTP request and don't expect any response.
// It block caller until request send successfully or return error
// Caller must pool sk or close it.
func SendUnidirectionalRequest(sk net_p.Socket, service service_p.Service, req codec_p.Codec) (err error_p.Error) {
	// TODO::: send service frame first

	// stream.SendRequest(syllab.NewCodec(req))
	err = req.Encode(sk.SendBuffer())
	if err != nil {
		return
	}

	for status := range sk.State() {
		switch status {
		case net_p.Status_Timeout_Read, net_p.Status_Timeout_Write:
			// err =
		case net_p.Status_SentCompletely:
			// Nothing to do. Just let execution go to stream.Close() and break the loop
		default:
			continue
		}
		break
	}
	return
}
