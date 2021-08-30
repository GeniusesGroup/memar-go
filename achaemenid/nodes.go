/* For license and copyright information please see LEGAL file in repository */

package achaemenid

import (
	"bytes"
	"context"
	"net"

	"../log"
	"../uuid"
)

type nodes struct {
	LocalNode    *Node
	Nodes        []Node // sort in nodeID
	nodesDetails []NodeDetails
}

// Init use to get other node data and make connection to each of them
func (n *nodes) Init() (err protocol.Error) {
	var goErr error

	var localNode = Node{
		InstanceID: uuid.Random32Byte(),
		State:      NodeStateLocalNode,
	}

	if protocol.AppDevMode {
		n.LocalNode = &localNode
		n.Nodes = append(n.Nodes, localNode)
	} else {
		// TODO::: change to Giti network when it is ready to serve lookup domain to GP address.
		var ipsAddr []net.IPAddr
		ipsAddr, goErr = net.DefaultResolver.LookupIPAddr(context.Background(), Server.Manifest.DomainName)
		if goErr != nil {
			// TODO::: block and try agin for 3 times??
		}

		if protocol.AppDebugMode {
			log.Debug("Achaemenid - Available nodes addr:", ipsAddr)
		}

		// Check if this is first instance of platform app.
		if len(ipsAddr) == 1 {
			n.Nodes = append(n.Nodes, localNode)
			n.LocalNode = &n.Nodes[0]
		} else {
			var conn *Connection
			var ipAddr [16]byte
			copy(ipAddr[:], ipsAddr[0].IP)
			conn, err = Server.Connections.MakeNewIPConnection(ipAddr)
			if err != nil {
				// TODO::: why fresh starting app can't make new connection???
			}
			var res *getServerNodeDetailsRes
			res, err = getServerNodeDetails(conn)
			if err != nil {
				// TODO::: try other node to get platform nodes??
			}

			var ln = len(res.nodes)
			n.Nodes = make([]Node, ln)
			n.nodesDetails = make([]NodeDetails, ln)
			for i := 0; i < ln; i++ {
				n.Nodes[i].ID = res.nodes[i].ID
				n.nodesDetails[i].ID = res.nodes[i].ID

				n.nodesDetails[i].GPAddr = res.nodes[i].GPAddr
				n.nodesDetails[i].IPAddr = res.nodes[i].IPAddr

				if !bytes.Equal(res.nodes[i].IPAddr[:], Server.Networks.localIP[:]) {
					n.Nodes[i].Conn, err = Server.Connections.MakeNewIPConnection(res.nodes[i].IPAddr)
				}
			}
		}

		// Register local node to cluster
	}
	return
}

// GetServerNodeDetails returns all platform nodes details.
func (n *nodes) GetServerNodeDetails() (nd []NodeDetails) {
	return n.nodesDetails
}

type getServerNodeDetailsRes struct {
	nodes []NodeDetails
}

func getServerNodeDetails(conn *Connection) (res *getServerNodeDetailsRes, err protocol.Error) {
	// Make new request-response streams
	var st *Stream
	st, err = conn.MakeOutcomeStream(0)
	if err != nil {
		return
	}

	// Set GetServerNodes ServiceID
	st.Service = &Service{ID: 639492616}

	err = SrpcOutcomeRequestHandler(st)
	if err != nil {
		return
	}

	err = res.FromSyllab(st.OutcomePayload)
	return
}

func (res *getServerNodeDetailsRes) FromSyllab(buf []byte) protocol.Error {
	// TODO:::
	return nil
}

// GetNodeByID returns exiting node.
func (n *nodes) GetNodeByID(nodeID uint64) *Node {
	return &n.Nodes[nodeID]
}
