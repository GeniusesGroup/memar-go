/* For license and copyright information please see LEGAL file in repository */

package matn

import (
	"../ganjine"
)

func init() {
	ganjine.Cluster.DataStructures.RegisterDataStructure(&indexPhraseStructure)
	ganjine.Cluster.DataStructures.RegisterDataStructure(&indexWordStructure)
	// ganjine.Cluster.DataStructures.RegisterDataStructure(&)

	// achaemenid.Server.Services.RegisterService()
}
