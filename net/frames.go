/* For license and copyright information please see the LEGAL file in the code repository */

package net

import (
	error_p "memar/process/error/protocol"
	"memar/net/chapar"
	"memar/net/gp"
	"memar/net/l4"
	net_p "memar/net/protocol"
	"memar/net/sec"
	"memar/net/srpc"
)

// HandleFrames Don't hold frames, So sender can reuse frames for any purpose.
func HandleFrames(np net_p.Packet) (err error_p.Error) {
	var sk net_p.Socket
	var frames, _ = np.Peek()

	for len(frames) > 0 {
		var frame = frame(frames)
		switch frame.FrameType() {
		case net_p.FrameType_Chapar:
			var chaparFrame = chapar.Frame(frame.Payload())
			err = chaparFrame.Do(sk)
			if err != nil {
				return
			}
			frames = chaparFrame.NextFrame()
		case net_p.FrameType_GP:
			var gpFrame = gp.Frame(frame.Payload())
			err = gpFrame.Do(sk)
			if err != nil {
				return
			}
			frames = gpFrame.NextFrame()
		case net_p.FrameType_Padding:
			var paddingFrame = sec.PaddingFrame(frame.Payload())
			// Nothing to do, just ignore padding data
			frames = paddingFrame.NextFrame()
		case net_p.FrameType_CallService:
			var callServiceFrame = srpc.ServiceFrame(frame.Payload())
			err = callServiceFrame.Do(sk)
			if err != nil {
				return
			}
			frames = callServiceFrame.NextFrame()
		case net_p.FrameType_OpenStream:
			var openStreamFrame = srpc.OpenStreamFrame(frame.Payload())
			err = openStreamFrame.Do(sk)
			if err != nil {
				return
			}
			frames = openStreamFrame.NextFrame()
		case net_p.FrameType_Data:
			var dataFrame = l4.DataFrame(frame.Payload())
			err = dataFrame.Do(sk)
			if err != nil {
				return
			}
			frames = dataFrame.NextFrame()
		case net_p.FrameType_Error:
			var errorFrame = srpc.ErrorFrame(frame.Payload())
			err = errorFrame.Do(sk)
			if err != nil {
				return
			}
			frames = errorFrame.NextFrame()
		case net_p.FrameType_Signature:
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
