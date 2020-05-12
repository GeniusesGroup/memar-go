/* For license and copyright information please see LEGAL file in repository */

package services

import "../achaemenid"

// Init use to register all available services to given server.
func Init(s *achaemenid.Server) {
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
	s.Services.RegisterService(&setIndexService)
	s.Services.RegisterService(&setRecordService)
	s.Services.RegisterService(&writeRecordService)
	// s.Services.RegisterService()
}
