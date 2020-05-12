/* For license and copyright information please see LEGAL file in repository */

package chaparkhane

// Connection can use by any type users
type Connection struct {
	PathID          []byte   // Can be this spec switch or MAC address in Ethernet spec
	ReversePathID   []byte   // Can be this spec switch in reverse or empty in Ethernet spec
	OSID            [4]byte  // Part of GP that lease by OS!
	Status          uint8    // 0:close 1:open 2:rate-limited 3:closing 4:opening 5:
	Weight          uint8    // 16 queue for priority weight of the connections exist.
	OwnerID         [16]byte // Can't change after creation. Guest=ConnectionPublicKey
	OwnerType       uint8    // 0:Guest, 1:Person, 2:Org, 3:App, ...
	OwnerPlatform   uint32   //
	MaxBandwidth    uint64   // use to tell the peer to slow down or packets will be drops in queues!
	RequestCount    uint64   // Use for PayAsGo strategy.
	BytesSent       uint64   // Counts the bytes of payload data sent.
	PacketsSent     uint64   // Counts packets sent.
	BytesReceived   uint64   // Counts the bytes of payload data Receive.
	PacketsReceived uint64   // Counts packets receive.
}

// NewConnection use to make new connection!
func NewConnection() *Connection {
	return &Connection{}
}
