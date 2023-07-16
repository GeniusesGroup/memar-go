/* For license and copyright information please see the LEGAL file in the code repository */

package tcp

// CCA or CongestionControlAlgorithm
// https://en.wikipedia.org/wiki/TCP_congestion_control
type CCA uint8

const (
	CongestionControlAlgorithm_Unset CCA = iota
	CongestionControlAlgorithm_Reno
	CongestionControlAlgorithm_Vegas
	CongestionControlAlgorithm_CUBIC
	CongestionControlAlgorithm_BBR

	// BIC-TCP is a sender-side-only change that ensures a linear RTT fairness
	// under large windows while offering both scalability and
	// bounded TCP-friendliness. The protocol combines two
	// schemes called additive increase and binary search
	// increase. When the congestion window is large, additive
	// increase with a large increment ensures linear RTT
	// fairness as well as good scalability. Under small
	// congestion windows, binary search increase provides TCP friendliness.
	CongestionControlAlgorithm_BIC
)
