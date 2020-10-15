/* For license and copyright information please see LEGAL file in repository */

package gs

import (
	"../achaemenid"
	"../ganjine"
)

var (
	// Cluster store cluster data to use by services!
	cluster *ganjine.Cluster
	// Server store address location to server use by other part of app!
	server *achaemenid.Server
)

// Init use to register all available services to given server.
// Init must call in main file before use any methods!
func Init(s *achaemenid.Server, c *ganjine.Cluster) {
	server = s
	cluster = c

	s.Services.RegisterService(&DeleteRecordService)
	s.Services.RegisterService(&GetRecordService)
	s.Services.RegisterService(&ReadRecordService)
	s.Services.RegisterService(&SetRecordService)
	s.Services.RegisterService(&WriteRecordService)

	s.Services.RegisterService(&HashIndexDeleteKeyHistoryService)
	s.Services.RegisterService(&HashIndexDeleteKeyService)
	s.Services.RegisterService(&HashIndexDeleteValueService)
	s.Services.RegisterService(&HashIndexGetValuesNumberService)
	s.Services.RegisterService(&HashIndexGetValuesService)
	s.Services.RegisterService(&HashIndexListenToKeyService)
	s.Services.RegisterService(&HashIndexSetValueService)

	s.Services.RegisterService(&HashTransactionFinishService)
	s.Services.RegisterService(&HashTransactionGetValuesService)
	s.Services.RegisterService(&HashTransactionRegisterService)

	s.Services.RegisterService(&GetNodeDetailsService)
	// s.Services.RegisterService()
}

type requestType uint8

// Services request types
const (
	RequestTypeStandalone requestType = iota
	RequestTypeBroadcast
)
