/* For license and copyright information please see LEGAL file in repository */

package chapar

/*
-------------------------------NOTICE:-------------------------------
This protocol usually implement in hardware level not software!
We just want to show simpilicity of protocol and its functions here!
*/

// Port indicate a port object methods
type Port interface {
	PortNumber() (num byte)
	Send(frame []byte)
	Receive(frame []byte)
}

// Ports store all available link port to other physical or logicall devices!
// 256 is max ports that Chapar protocol support directly!!
type Ports [256]Port

// RegisterPort registers new port on given ports pool!
func (ports Ports) RegisterPort(p Port) {
	ports[p.PortNumber()] = p
}

// UnRegisterPort delete the port by given port number on given ports pool!
func (ports Ports) UnRegisterPort(portNumber byte) {
	ports[portNumber] = nep
}

// HandleFrame handles a frame by given port!
func (ports Ports) HandleFrame(p Port, frame []byte) {
	IncrementNextHop(frame, p.PortNumber())

	if IsBroadcastFrame(frame) {
		// send the frame to all ports as BroadcastFrame!
		var i byte
		for i = 0; i <= 255; i++ {
			// Send frame to desire port interface! Usually frame will put in queue!
			ports[i].Send(frame)
		}
	} else {
		// send the frame to the specific port as UnicastFrame!
		ports[GetNextHop(frame)].Send(frame)
	}
}

// Init initializes new Ports object otherwise panic will occur on un-register port call!
func (ports Ports) Init() {
	var i byte
	for i = 0; i <= 255; i++ {
		ports[i] = nep
	}
}

var nep nonExistPort

// nonExistPort use for default and empty switch port due to non of ports can be nil!
type nonExistPort struct{}

// Send use for default and empty switch port due to non of ports can be nil!
func (nep nonExistPort) PortNumber() (num byte) { return }

// Send use for default and empty switch port due to non of ports can be nil!
func (nep nonExistPort) Send(frame []byte) {}

// Receive use for default and empty switch port due to non of ports can be nil!
func (nep nonExistPort) Receive(frame []byte) {}
