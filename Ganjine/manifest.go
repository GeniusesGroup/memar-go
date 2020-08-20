/* For license and copyright information please see LEGAL file in repository */

package ganjine

// Manifest store Ganjine manifest data
type Manifest struct {
	DataCentersClass  uint8  // 0:FirstClass 256:Low-Quality default:0
	ReplicationNumber uint8  // deafult:3

	TransactionTimeOut uint16 // in ms, default:500ms, Max 65.535s timeout
	NodeFailureTimeOut uint16 // in minute, default:60m, other corresponding node same range will replace failed node! not use in network failure, it is handy proccess!

	CachePercent uint8 // GC cached records when reach this size limit!
}
