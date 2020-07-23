/* For license and copyright information please see LEGAL file in repository */

package achaemenid

import (
	"bytes"
	"context"
	"net"
)

type nodes struct {
	server       *Server
	localNode    Node
	Nodes        []Node // sort in nodeID
	nodesDetails []NodeDetails
}

// Init use to get other node data and make connection to each of them
func (n *nodes) Init(s *Server) (err error) {
	n.server = s

	n.nodesDetails, err = n.GetServerNodeDetails()
	if err != nil {
		return
	}

	return
}

// GetServerNodeDetails returns all platform nodes details.
func (n *nodes) GetServerNodeDetails() (nd []NodeDetails, err error) {
	if len(n.nodesDetails) == 0 {
		// TODO::: change to Giti network when it is ready to serve lookup domain to GP address.
		var ipsAddr []net.IPAddr
		ipsAddr, err = net.DefaultResolver.LookupIPAddr(context.Background(), n.server.Manifest.DomainName)
		if err != nil {
			// TODO::: block and try agin for 3 times??
		}

		// Check if this is first instance of platform app.
		if len(ipsAddr) == 1 {
			n.Nodes = []Node{
				Node{
					Conn:  nil, // Nil due to it is local node
					state: NodeStateStart,
				},
			}
		} else {
			var conn *Connection
			conn, err = n.server.Connections.MakeNewIPConnection(&ipsAddr[0])
			if err != nil {
				// TODO::: why fresh starting app can't make new connection???
			}
			var res *getServerNodeDetailsRes
			res, err = getServerNodeDetails(n.server, conn)
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

				if !bytes.Equal(res.nodes[i].IPAddr, n.server.Networks.localIP) {
					n.Nodes[i].Conn, err = n.server.Connections.MakeNewIPConnection(&net.IPAddr{IP: res.nodes[i].IPAddr})
				}
			}
		}

	}

	return n.nodesDetails, nil
}

type getServerNodeDetailsRes struct {
	nodes []NodeDetails
}

func getServerNodeDetails(s *Server, conn *Connection) (res *getServerNodeDetailsRes, err error) {
	// Make new request-response streams
	var reqStream, resStream *Stream
	reqStream, resStream, err = conn.MakeBidirectionalStream(0)
	if err != nil {
		return
	}

	// Set GetServerNodes ServiceID
	reqStream.ServiceID = 639492616

	err = SrpcOutcomeRequestHandler(s, reqStream)
	if err != nil {
		return
	}

	err = res.syllabDecoder(resStream.Payload)

	return
}

func (res *getServerNodeDetailsRes) syllabDecoder(buf []byte) error {
	// TODO:::
	return nil
}

// GetNodeByID returns exiting node.
func (n *nodes) GetNodeByID(nodeID uint64) *Node {
	return &n.Nodes[nodeID]
}
