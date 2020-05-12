/* For license and copyright information please see LEGAL file in repository */

package chapar

import "testing"

func TestMiddleWaySwitch(t *testing.T) {
	var testPorts Ports
	// We add all ports for testing purpose! Otherwise below line must un-commenting
	// testPorts.Init()

	// add all port as virtual NIC to testPorts for testing purpose!
	var i byte
	for i = 0; i <= 255; i++ {
		var mws = middleWayBlockingSwitch{testPorts, i, nil}
		testPorts.RegisterPort(mws, i)
	}

	// Send 1 million unicast request to one port that unicast to other ports

	// Send 1 million broadcast request to one port

	// Send 1 million broadcast request to all port
}

type middleWayBlockingSwitch struct {
	Ports      Ports
	PortNumber byte
	Cache      [][]byte
}

// Send :
func (mws middleWayBlockingSwitch) Send(f Frame) {
}

// Receive detect and send frame to other port or ports!
// This switch logic usually implement in hardware level not software!!
func (mws middleWayBlockingSwitch) Receive(f Frame) {
	if f.IsBroadcastFrame() {
		// broadcast the frame to all ports!
		mws.Ports.BroadcastFrame(f, mws.PortNumber)
	} else {
		// Send frame to desire port queue
		mws.Ports.UnicastFrame(f, mws.PortNumber)
	}
}

type osBlockingSwitch struct {
	Ports      Ports
	PortNumber byte
}

// Send :
func (obs osBlockingSwitch) Send(f Frame) {
	return
}

// Receive :
func (obs osBlockingSwitch) Receive(f Frame) {
	switch f.GetNextHeader() {
	// GP
	case 0:

	// IPv4
	case 80:

	// IPv6
	case 86:
	}
}
