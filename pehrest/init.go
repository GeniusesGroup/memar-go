/* For license and copyright information please see LEGAL file in repository */

package pehrest

import (
	"../protocol"
)

const (
	DomainName = "index.protocol"
)

func init() {

	protocol.OS.RegisterMediaType(&indexHashStructure)
	// protocol.OS.RegisterMediaType(&)

	protocol.App.RegisterService(&HashDeleteKeyHistoryService)
	protocol.App.RegisterService(&HashDeleteKeyService)
	protocol.App.RegisterService(&HashDeleteValueService)
	protocol.App.RegisterService(&HashGetValuesNumberService)
	protocol.App.RegisterService(&HashGetValuesService)
	protocol.App.RegisterService(&HashListenToKeyService)
	protocol.App.RegisterService(&HashInsertValueService)
	protocol.App.RegisterService(&HashSetValueService)

	protocol.App.RegisterService(&HashTransactionFinishService)
	protocol.App.RegisterService(&HashTransactionGetValuesService)
	protocol.App.RegisterService(&HashTransactionRegisterService)

	// protocol.App.RegisterService()
}
