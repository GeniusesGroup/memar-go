/* For license and copyright information please see LEGAL file in repository */

package protocol

// StoragesDistributed is the interface that can implement by any Application to provide distributed storage mechanism.
type StoragesDistributed interface {
	Objects() StorageObjects
	Files() FileDirectory
	Records() StorageRecords
	KeyValues() StorageKeyValue
}

// StoragesLocal is the interface that can implement by any Application to provide local storage mechanism.
type StoragesLocal interface {
	LocalObjects() StorageObjects
	LocalFiles() FileDirectory
	LocalRecords() StorageRecords
	LocalKeyValues() StorageKeyValue
}

// StoragesCache is the interface that can implement by any Application to provide cache storage mechanism.
// Keep them on volatile memories and just save very common on non-volatile storages.
// All GUI & edge nodes of any software must use cache in many cases to improve performance.
type StoragesCache interface {
	ObjectsCache() StorageObjects
	FilesCache() FileDirectory
	RecordsCache() StorageRecords
	KeyValuesCache() StorageKeyValue
}
