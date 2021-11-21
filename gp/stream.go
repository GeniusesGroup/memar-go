/* For license and copyright information please see LEGAL file in repository */

package gp

import (
	etime "../earth-time"
	"../protocol"
)

// Stream use to send or receive data on specific connection.
// It can pass to logic layer to give data access to developer!
// Data flow can be up to down (parse raw income data) or down to up (encode app data with respect MTU)
// If OutcomePayload not present stream is UnidirectionalStream otherwise it is BidirectionalStream!
type Stream struct {
	id          uint32 // Even number for Peer(who start connection). Odd number for server(who accept connection).
	connection  *Connection
	protocolID  protocol.NetworkApplicationProtocolID
	service     protocol.Service
	incomeData  protocol.Codec
	outcomeData protocol.Codec

	/* State */
	status protocol.ConnectionState      // States locate in const of this file.
	state  chan protocol.ConnectionState // States locate in const of this file.
	Weight protocol.ConnectionWeight     // 16 queue for priority weight of the streams exist.

	/* Metrics */
	TotalPacket     uint32 // Expected packets count that must received!
	PacketReceived  uint32 // Count of packets received!
	LastPacketID    uint32 // Last send or received Packet use to know order of packets!
	PacketDropCount uint8  // Count drop packets to prevent some attacks type!
}

func (st *Stream) ID() uint32                                        { return st.id }
func (st *Stream) Connection() protocol.Connection                   { return st.connection }
func (st *Stream) Service() protocol.Service                         { return st.service }
func (st *Stream) ProtocolID() protocol.NetworkApplicationProtocolID { return st.protocolID }
func (st *Stream) Status() protocol.ConnectionState                  { return st.status }
func (st *Stream) State() chan protocol.ConnectionState              { return st.state }
func (st *Stream) IncomeData() protocol.Codec                        { return st.incomeData }
func (st *Stream) OutcomeData() protocol.Codec                       { return st.outcomeData }
func (st *Stream) SetIncomeData(codec protocol.Codec)                { st.incomeData = codec }
func (st *Stream) SetOutcomeData(codec protocol.Codec)               { st.outcomeData = codec }

// SetService check and if it is first time register given service for the stream.
func (st *Stream) SetService(service protocol.Service) {
	if st.service != nil {
		st.service = service
	}
}


// SendRequest use for default and empty switch port due to non of ports can be nil!
// SendAndWait register stream in send pool and block caller until response ready to read.
func (st *Stream) SendRequest() (err protocol.Error) {
	
	var outcomeData = stream.OutcomeData()
	if outcomeData != nil {
	}
	
	// st.Send()

	// Listen to response stream and decode error ID and return it to caller
	var responseStatus protocol.ConnectionState = <-st.State()
	if responseStatus == protocol.ConnectionStateReady {

	} else {

	}

	// Last Close response stream!
	st.Close()
	return
}

// SendResponse register stream in send pool. Usually use to send response stream.
func (st *Stream) SendResponse() (err protocol.Error) {
	// First Check st.Connection.Status to ability send stream over it

	// TODO::: remove this check when remove IP support
	// if st.Connection.IPAddr != [16]byte{} {
	// Send by IP
	// } else {
	// Send stream by GP
	// }

	// send stream by weight

	/* Metrics data */
	// c.BytesSent += uint64(len(st.OutcomePayload))

	// Last Close stream!
	st.Close()
	return
}

// CloseStream delete given Stream from pool
func (st *Stream) Close() {
	sp.mutex.Lock()
	delete(st.connection.StreamPool.p, st.id)
	sp.mutex.Unlock()
}

// SetState change state of stream and send notification on stream StateChannel.
func (st *Stream) SetState(state protocol.ConnectionState) {
	// atomic.StoreUInt64(&st.State, state)
	st.State = state
	// notify stream listener that stream state has been changed!
	st.StateChannel <- state
}

// Authorize authorize request by data in related stream and connection.
func (st *Stream) Authorize() (err protocol.Error) {
	// if st.Connection().UserID() != protocol.OS.AppManifest().AppUUID() {
	// 	// TODO::: Attack??
	// 	err = ErrNotAuthorizeRequest
	// 	return
	// }

	err = st.Service.Authorization.UserType.Check(st.Connection.UserType)
	if err != nil {
		return
	}

	err = st.Connection.AccessControl.AuthorizeWhen(etime.Now().Weekdays(), etime.Now().DayHours())
	if err != nil {
		return
	}
	err = st.Connection.AccessControl.AuthorizeWhich(st.Service.ID, st.Service.Authorization.CRUD)
	if err != nil {
		return
	}
	err = st.Connection.AccessControl.AuthorizeWhere(st.Connection.GPAddr.GetSocietyID(), st.Connection.GPAddr.GetRouterID())
	if err != nil {
		return
	}
	return
}

// MakeNewStream use to make new stream without any connection exist yet!
// TODO::: due to IPv4&&IPv6 support we need this func! Remove it when remove those support!
func MakeNewStream() (st *Stream, err protocol.Error) {
	// TODO::: check server allow make new connection and has enough resource

	st = &Stream{
		// ID:           0,
		State:        protocol.ConnectionStateOpening,
		StateChannel: make(chan protocol.ConnectionState),
	}
	return
}
