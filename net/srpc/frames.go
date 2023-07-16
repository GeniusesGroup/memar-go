/* For license and copyright information please see LEGAL file in repository */

package srpc

import (
	"../protocol"
)

func HandleFrames(conn protocol.Connection, frames []byte) (err protocol.Error) {
	for len(frames) > 0 {
		var frame = frame(frames)
		switch frame.Type() {
		case frameTypePadding:
			var paddingFrame = paddingFrame(frame.Payload())
			// Nothing to do, just ignore padding data
			frames = paddingFrame.NextFrame()
		case frameTypePing:
			var pingFrame = pingFrame(frame.Payload())
			// TODO:::
			frames = pingFrame.NextFrame()
		case frameTypeCallService:
			var serviceFrame = serviceFrame(frame.Payload())
			err = callService(conn, serviceFrame)
			if err != nil {
				return
			}
			frames = serviceFrame.NextFrame()
		case frameTypeOpenStream:
			var openStreamFrame = openStreamFrame(frame.Payload())
			err = openStream(conn, openStreamFrame)
			if err != nil {
				return
			}
			frames = openStreamFrame.NextFrame()
		case frameTypeData:
			var dataFrame = dataFrame(frame.Payload())
			err = appendData(conn, dataFrame)
			if err != nil {
				return
			}
			frames = dataFrame.NextFrame()
		case frameTypeSignature:
			var signatureFrame = signatureFrame(frame.Payload())
			err = registerStreamSignature(conn, signatureFrame)
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
