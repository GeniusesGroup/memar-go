/* For license and copyright information please see the LEGAL file in the code repository */

package srpc

import (
	"memar/protocol"
)

// Handler ...
// Read more about this protocol : https://github.com/GeniusesGroup/memar/blob/main/sRPC.md
type Handler struct{}

// HandleIncomeRequest handle incoming sRPC request streams.
func (h *Handler) HandleIncomeRequest(sk protocol.Socket) (err protocol.Error) {
	var service = sk.ApplicationLayer().Service()
	if service == nil {
		// TODO:::
		return
	}

	// TODO::: can't easily call service and must schedule it by its weight.
	var ser, ok = service.(protocol.SRPCHandler)
	if ok {
		var res protocol.Codec
		res, err = ser.ServeSRPC(sk, sk)
		if err != nil {
			sk.ApplicationLayer().SetError(err)
			return
		}
		sk.ApplicationLayer().SetResponse(res)
		sk.Decode(res)
	} else {
		// TODO:::
	}
	return
}

// SendBidirectionalRequest use to send outcoming sRPC request.
// It block caller until get response or error.
// Caller must pool sk or close it.
func SendBidirectionalRequest(sk protocol.Socket, service protocol.Service, req protocol.Codec) (res protocol.Codec, err protocol.Error) {
	// TODO::: send service frame first

	// stream.SendRequest(syllab.NewCodec(req))
	_, err = sk.Decode(req)
	if err != nil {
		return
	}

	for status := range sk.State() {
		switch status {
		case protocol.NetworkStatus_Timeout:
			// err =
		case protocol.NetworkStatus_ReceivedCompletely:
			res = sk
			err = sk.ApplicationLayer().Error()
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
func SendUnidirectionalRequest(sk protocol.Socket, service protocol.Service, req protocol.Codec) (err protocol.Error) {
	// TODO::: send service frame first

	// stream.SendRequest(syllab.NewCodec(req))
	_, err = sk.Decode(req)
	if err != nil {
		return
	}

	for status := range sk.State() {
		switch status {
		case protocol.NetworkStatus_Timeout:
			// err =
		case protocol.NetworkStatus_SentCompletely:
			// Nothing to do. Just let execution go to stream.Close() and break the loop
		default:
			continue
		}
		break
	}
	return
}
