/* For license and copyright information please see the LEGAL file in the code repository */

package protocol

// STG is default global protocol.Storages like window global variable in browsers.
// You must assign to it by any object implement protocol.Storages on your main.go file. Suggestion:
// protocol.STG = &ganjine.Storages
var STG Storages

type Storages interface {
	StoragesLocal
	StoragesCache
	// Server specific applications
	StoragesDistributed
}

// StoragesDistributed is the interface that can implement by any Application to provide distributed storage mechanism.
type StoragesDistributed interface {
	Objects() StorageObjects
	Files() FileDirectory
	Records() StorageRecords
	KeyValues() StorageKeyValue
}

// StoragesLocal is the interface that can implement by any Application to provide local storage mechanism.
type StoragesLocal interface {
	Local_Objects() StorageObjects
	Local_Files() FileDirectory
	Local_Records() StorageRecords
	Local_KeyValues() StorageKeyValue

	// It is like old kernels file systems that shared between all applications.
	// It serve by the GUI app that run on the OS as default user interface.
	Local_SharedFiles() FileDirectory
}

// StoragesCache is the interface that can implement by any Application to provide cache storage mechanism.
// Keep them on volatile memories and just save very common on non-volatile storages.
// All GUI & edge nodes of any software must use cache in many cases to improve performance.
type StoragesCache interface {
	Cache_Objects() StorageObjects
	Cache_Files() FileDirectory
	Cache_Records() StorageRecords
	Cache_KeyValues() StorageKeyValue
}
