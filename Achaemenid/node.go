/* For license and copyright information please see LEGAL file in repository */

package achaemenid

// Node store details about a node that part of the platfrom.
type Node struct {
	ID              uint64
	dataCenterID    [16]byte
	storageCapacity uint64 // In bytes, Max 16EB(Exabyte) that more enough for one node capacity. 0 means service only node.
	Conn            *Connection
	server          *Server
	state           uint8
}

// Node State
const (
	NodeStateStop uint8 = iota
	NodeStateStoping
	NodeStateStart
	NodeStateStarting
	NodeStateStable
	NodeStateNotResponse
)

// NodeDetails ...
type NodeDetails struct {
	ID     uint64
	GPAddr [16]byte
	IPAddr []byte
}
