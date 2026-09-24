/* For license and copyright information please see the LEGAL file in the code repository */

package operation_p

type Field_Weight interface {
	// Use to queue requests by its weights in the same priority
	OperationWeight() Weight
}

// Weight indicate to use by any queue systems like network sockets to implement priority/weight mechanism.
type Weight uint8

// Weights
const (
	Weight_Unset Weight = iota
	Weight_Lowest
	Weight_Lower
	Weight_Low
	Weight_Normal
	// TODO::: add more queue for priority weights
	Weight_TimeSensitive Weight = 255 // Call related service in each received packet. VoIP, IPTV, Sensors data, ...
)
