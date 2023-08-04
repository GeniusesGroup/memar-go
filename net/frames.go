/* For license and copyright information please see the LEGAL file in the code repository */

package net

import (
	"memar/net/chapar"
	"memar/net/gp"
	"memar/net/l4"
	"memar/net/sec"
	"memar/net/srpc"
	"memar/protocol"
)

// HandleFrames Don't hold frames, So sender can reuse frames for any purpose.
func HandleFrames(frames []byte) (err protocol.Error) {
	var sk protocol.Socket

	for len(frames) > 0 {
		var frame = frame(frames)
		switch frame.FrameType() {
		case protocol.Network_FrameType_Chapar:
			var chaparFrame = chapar.Frame(frame.Payload())
			err = chaparFrame.Do(sk)
			if err != nil {
				return
			}
			frames = chaparFrame.NextFrame()
		case protocol.Network_FrameType_GP:
			var gpFrame = gp.Frame(frame.Payload())
			err = gpFrame.Do(sk)
			if err != nil {
				return
			}
			frames = gpFrame.NextFrame()
		case protocol.Network_FrameType_Padding:
			var paddingFrame = sec.PaddingFrame(frame.Payload())
			// Nothing to do, just ignore padding data
			frames = paddingFrame.NextFrame()
		case protocol.Network_FrameType_CallService:
			var callServiceFrame = srpc.ServiceFrame(frame.Payload())
			err = callServiceFrame.Do(sk)
			if err != nil {
				return
			}
			frames = callServiceFrame.NextFrame()
		case protocol.Network_FrameType_OpenStream:
			var openStreamFrame = srpc.OpenStreamFrame(frame.Payload())
			err = openStreamFrame.Do(sk)
			if err != nil {
				return
			}
			frames = openStreamFrame.NextFrame()
		case protocol.Network_FrameType_Data:
			var dataFrame = l4.DataFrame(frame.Payload())
			err = dataFrame.Do(sk)
			if err != nil {
				return
			}
			frames = dataFrame.NextFrame()
		case protocol.Network_FrameType_Error:
			var errorFrame = srpc.ErrorFrame(frame.Payload())
			err = errorFrame.Do(sk)
			if err != nil {
				return
			}
			frames = errorFrame.NextFrame()
		case protocol.Network_FrameType_Signature:
			var signatureFrame = sec.SignatureFrame(frame.Payload())
			err = signatureFrame.Do(sk)
			if err != nil {
				return
			}
			frames = signatureFrame.NextFrame()
		default:
			// TODO:::
		}
	}
	return
}
