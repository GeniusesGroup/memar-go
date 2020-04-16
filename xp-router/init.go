/* For license and copyright information please see LEGAL file in repository */

package xprouter

// New node get exiting shared database
// New XP node calculate path (latency and capacity) to all exiting node and send new information to all exiting node
// XPs can recalculate path and tell others about any change just if any change to their physical links!
// - multiple routes to the same place to be assigned the same cost and will cause traffic to be distributed evenly over those routes

var err error
err = UIP.CheckPacket()
if err != nil {
	// Send response or just ignore packet
	// TODO : DDOS!!??
	return
}