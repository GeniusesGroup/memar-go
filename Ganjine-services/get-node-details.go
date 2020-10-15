/* For license and copyright information please see LEGAL file in repository */

package gs

import (
	"../achaemenid"
	"../ganjine"
	lang "../language"
)

// GetNodeDetailsService store details about GetNodeDetails service
var GetNodeDetailsService = achaemenid.Service{
	ID:                3707636027,
	IssueDate:         1596950889,
	ExpiryDate:        0,
	ExpireInFavorOf:   "",
	ExpireInFavorOfID: 0,
	Status:            achaemenid.ServiceStatePreAlpha,

	Name: map[lang.Language]string{
		lang.EnglishLanguage: "GetNodeDetails",
	},
	Description: map[lang.Language]string{
		lang.EnglishLanguage: "Get node details that request receive",
	},
	TAGS: []string{""},

	SRPCHandler: GetNodeDetailsSRPC,
}

// GetNodeDetailsSRPC is sRPC handler of GetNodeDetails service.
func GetNodeDetailsSRPC(st *achaemenid.Stream) {
	if server.Manifest.DomainID != st.Connection.DomainID {
		// TODO::: Attack??
		st.Err = ganjine.ErrGanjineNotAuthorizeRequest
		return
	}

	var res *ganjine.Node
	res, st.Err = GetNodeDetails(st)
	// Check if any error occur in bussiness logic
	if st.Err != nil {
		return
	}

	st.OutcomePayload = res.SyllabEncoder()
}

// GetNodeDetails returns local node details or related error!
func GetNodeDetails(st *achaemenid.Stream) (res *ganjine.Node, err error) {
	res, err = cluster.GetLocalNodeDetail()
	return
}
