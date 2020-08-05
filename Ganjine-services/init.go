/* For license and copyright information please see LEGAL file in repository */

package gs

import (
	"../achaemenid"
	"../errors"
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

	s.Services.RegisterService(&deleteIndexHashHistoryService)
	s.Services.RegisterService(&deleteIndexHashRecordService)
	s.Services.RegisterService(&deleteIndexHashService)
	s.Services.RegisterService(&deleteRecordService)
	s.Services.RegisterService(&findRecordsConsistentlyService)
	s.Services.RegisterService(&findRecordsService)
	s.Services.RegisterService(&finishTransactionService)
	s.Services.RegisterService(&getIndexHashNumberService)
	s.Services.RegisterService(&getRecordService)
	s.Services.RegisterService(&listenToIndexService)
	s.Services.RegisterService(&readRecordService)
	s.Services.RegisterService(&registerTransactionService)
	s.Services.RegisterService(&setIndexHashService)
	s.Services.RegisterService(&setRecordService)
	s.Services.RegisterService(&writeRecordService)
	// s.Services.RegisterService()
}

type requestType uint8

// Services request types
const (
	RequestTypeStandalone requestType = iota
	RequestTypeBroadcast
)

// Package errors
var (
	ErrNotAuthorizeGanjineRequest = errors.New("NotAuthorizeGanjineRequest", "Given request to ganjine services came from connection that not authorize to request this service")
)
