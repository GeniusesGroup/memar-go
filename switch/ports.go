/* For license and copyright information please see LEGAL file in repository */

package frameswitch

/*
-------------------------------NOTICE:-------------------------------
This protocol must implement in hardware level not software!
We just want to show simpilicity of protocol and its functions here!
*/

const (
	// MaxPortNumber use to indicate max ports can be in Ports slice!
	MaxPortNumber byte = 255
)

// Port indicate a port object methods
type Port interface {
	Send(f Frame)
	Receive(f Frame)
}

// Ports store all available link port to other physical or logicall devices!
type Ports [MaxPortNumber]Port

// RegisterPort use to register new port on switch logic!
func (ports Ports) RegisterPort(p Port, portNumber byte) {
	ports[portNumber] = p
}

// UnRegisterPort use to un-register the port on switch logic!
func (ports Ports) UnRegisterPort(portNumber byte) {
	ports[portNumber] = nep
}

// UnicastFrame use to send a frame to the specific port!
func (ports Ports) UnicastFrame(f Frame, receivedPortNumber byte) {
	// The reasons of use SetSwitchPortNum method here:
	// - UnicastFrame : To be sure receive port is same with declaration one in frame, we replace it always!
	// - Rule&Security : To be sure physical network port is same on sender and receiver switch, we must set it again here!
	f.SetSwitchPortNum(f.GetNextHop(), receivedPortNumber)

	f.IncrementNextHop()

	ports[f.GetNextHop()].Send(f)
}

// BroadcastFrame use to send a frame to all port in Ports array!
func (ports Ports) BroadcastFrame(f Frame, receivedPortNumber byte) {
	// The reasons of use SetSwitchPortNum method here:
	// - BroadcastFrame : To improve performance, previous switch just send frame without declare port, we must declare it now!
	// - Rule&Security : To be sure physical network port is same on sender and receiver switch, we must set it again here!
	f.SetSwitchPortNum(f.GetNextHop(), receivedPortNumber)

	f.IncrementNextHop()

	var i byte
	for i = 0; i <= MaxPortNumber; i++ {
		// Send frame to desire port interface! Usually frame will put in queue!
		ports[i].Send(f)
	}
}

// Init use to initialize new Ports object otherwise panic will occur on un-register port call!
func (ports Ports) Init() {
	var i byte
	for i = 0; i <= MaxPortNumber; i++ {
		ports[i] = nep
	}
}

var nep nonExistPort

// nonExistPort use for default and empty switch port due to non of ports can be nil!
type nonExistPort struct{}

// Send use for default and empty switch port due to non of ports can be nil!
func (nep nonExistPort) Send(f Frame) {}

// Receive use for default and empty switch port due to non of ports can be nil!
func (nep nonExistPort) Receive(f Frame) {}
