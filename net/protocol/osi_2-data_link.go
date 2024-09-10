/* For license and copyright information please see the LEGAL file in the code repository */

package net_p

import (
	object_p "memar/computer/language/object/protocol"
)

/*
**********************************************************************************
Link - (OSI Layer 2: Data Link) - Device to Device Connection

https://en.wikipedia.org/wiki/Data_link_layer
**********************************************************************************
*/

// OSI_DataLink use to network hardware devices in a computers or connect two or more computers.
type OSI_DataLink interface {
	object_p.LifeCycle
	Framer
	NetworkAddress // string form of address (for example, "MAC://aa:bb:cc:dd:ee:ff", "Chapar://[1:242:20]")
	FrameWriter
}
