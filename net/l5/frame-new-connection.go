/* For license and copyright information please see the LEGAL file in the code repository */

package l5

// Three way to make new connection
// - Guest: provide some data like cipher-suite to open the connection
// - New: Ask peer to get authentication data from given society
// - Load: Ask to load exiting connection by provide some detail like UserID+DelegateUserID

// This frame must not to encrypt??
// Must be the first frame in the first Packet and not mean PacketID==0 to improve the connection to know the state??
