/* For license and copyright information please see the LEGAL file in the code repository */

package protocol

/*
**********************************************************************************
Physical - (OSI Layer 1: Physical) - network interface

https://en.wikipedia.org/wiki/Physical_layer
**********************************************************************************
*/

// OSI_Physical or NetworkInterface or NetworkPhysical_Connection or Hardware2Hardware_Connection
// In Linux (or Unix) world, most network interfaces, such as eth0 and ppp0, are associated to a physical device that is in charge or transmitting
// and receiving data packets. However, there are exceptions to this rule,
// and some logical network interface doesn't feature any physical packet transmission;
// the most known examples are the shaper and eql interfaces.
// This article shows how such ``virtual'' interfaces attach to the kernel and to the packet transmission mechanism.
// From the kernel's point of view, a network interface is a software object that can process outgoing packets,
// and the actual transmission mechanism remains hidden inside the interface driver.
type OSI_Physical interface {
	// Send transmitting in blocking mode and block caller until physical layer driver copy packet to its hardware.
	// Return error not related to packet situation, just about any hardware error.
	// A situation might be occur that the port available when a packet queued but when the time to send is come,
	// the port broken and sender don't know about this.
	// Due to speed matters in link layer, and it is very rare situation, it is better to ignore suddenly port unavailability.
	Send(packet Network_Packet) (err Error)

	ObjectLifeCycle
	DataType
	Network_Framer
	NetworkMTU
	// NetworkAddress
	Network_FrameWriter
}

// NetworkMTU
type NetworkMTU interface {
	MTU() int
}
