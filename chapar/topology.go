/* For license and copyright information please see LEGAL file in repository */

package chapar

// MakeTopologyReq is request structure of MakeTopology()
type MakeTopologyReq struct {
	EndUser           uint
	GateWayPort       uint  // Last hop free port
	UserHopEfficiency uint8 // How many physical link to upper hop
	Efficiency        uint8 // How many physical link to upper hop
}

// MakeTopologyRes is response structure of MakeTopology()
type MakeTopologyRes struct {
	EndUserSwitch uint
	HopNumber     uint8
}

// MakeTopology calculate some data to implement Chapar network easily.
func MakeTopology(req *MakeTopologyReq) (res *MakeTopologyRes) {
	res = &MakeTopologyRes{}

	// TODO::: Calculate each hop free port to upper hop by req.Efficiency

	res.EndUserSwitch = req.EndUser / uint(255 - req.UserHopEfficiency + 1) 
	return
}
