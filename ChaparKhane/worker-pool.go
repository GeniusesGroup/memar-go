/* For license and copyright information please see LEGAL file in repository */

package chaparkhane

// WorkerPool use to queue and do works!
type WorkerPool struct {
	ReceivePool        [65536][]byte
	LastReceivePoolUse uint16
	LastReceivePoolDo  uint16
	SendWeight0        [1024]*Stream
}

// RegisterStreamToSend use to register a stream to send automatically to other side.
func (wp WorkerPool) RegisterStreamToSend(st *Stream) (err error) {
	// First Check st.Connection.Status to ability send stream over it
	
	return nil
}
