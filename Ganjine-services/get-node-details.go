/* For license and copyright information please see LEGAL file in repository */

package gs

import (
	"../achaemenid"
	"../ganjine"
)

var getNodeDetailsService = achaemenid.Service{
	ID:                3707636027,
	URI:               "", // API services can set like "/apis?3707636027" but it is not efficient, find services by ID.
	Name:              "GetNodeDetails",
	IssueDate:         1596950889,
	ExpiryDate:        0,
	ExpireInFavorOf:   "",
	ExpireInFavorOfID: 0,
	Status:            achaemenid.ServiceStatePreAlpha,
	Description: []string{
		"Get node details that request receive",
	},
	TAGS:        []string{""},
	SRPCHandler: GetNodeDetailsSRPC,
}

// GetNodeDetailsSRPC is sRPC handler of GetNodeDetails service.
func GetNodeDetailsSRPC(s *achaemenid.Server, st *achaemenid.Stream) {
	if server.Manifest.DomainID != st.Connection.DomainID {
		// TODO::: Attack??
		st.ReqRes.Err = ErrNotAuthorizeGanjineRequest
		return
	}

	var res *ganjine.Node
	res, st.ReqRes.Err = GetNodeDetails(st)
	// Check if any error occur in bussiness logic
	if st.ReqRes.Err != nil {
		return
	}

	st.ReqRes.Payload = res.SyllabEncoder()
}

// GetNodeDetails returns local node details or related error!
func GetNodeDetails(st *achaemenid.Stream) (res *ganjine.Node, err error) {
	res, err = cluster.GetLocalNodeDetail()
	return
}
