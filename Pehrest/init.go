/* For license and copyright information please see LEGAL file in repository */

package pehrest

import (
	"../achaemenid"
	"../ganjine"
)

func init() {

	ganjine.Cluster.DataStructures.RegisterDataStructure(&indexHashStructure)
	// ganjine.Cluster.DataStructures.RegisterDataStructure(&)

	achaemenid.Server.Services.RegisterService(&HashDeleteKeyHistoryService)
	achaemenid.Server.Services.RegisterService(&HashDeleteKeyService)
	achaemenid.Server.Services.RegisterService(&HashDeleteValueService)
	achaemenid.Server.Services.RegisterService(&HashGetValuesNumberService)
	achaemenid.Server.Services.RegisterService(&HashGetValuesService)
	achaemenid.Server.Services.RegisterService(&HashListenToKeyService)
	achaemenid.Server.Services.RegisterService(&HashSetValueService)

	achaemenid.Server.Services.RegisterService(&HashTransactionFinishService)
	achaemenid.Server.Services.RegisterService(&HashTransactionGetValuesService)
	achaemenid.Server.Services.RegisterService(&HashTransactionRegisterService)

	// achaemenid.Server.Services.RegisterService()
}
