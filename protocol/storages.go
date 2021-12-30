/* For license and copyright information please see LEGAL file in repository */

package protocol

// Storages is the interface that can implement by any Application to provide storage mechanism.
type Storages interface {
	Objects() StorageObjects      // Distributed objects storage
	LocalObjects() StorageObjects // Local object storage
	ObjectsCache() StorageObjects // Local cached object storage. Keep them on volatile memories and just save very common on non-volatile storages.

	Files() FileDirectory      // Distributed files storage
	LocalFiles() FileDirectory // Local file storage
	FilesCache() FileDirectory // Local cached files storage. Keep them on volatile memories and just save very common on non-volatile storages.

	Records() StorageRecords      // Distributed records storage
	LocalRecords() StorageRecords // Local records storage
	RecordsCache() StorageRecords // Local cached records storage. Keep them on volatile memories and just save very common on non-volatile storages.

	KeyValues() StorageKeyValue      // Distributed key/value storage
	LocalKeyValues() StorageKeyValue // Local key/value storage
	KeyValuesCache() StorageKeyValue // Local cached key/value storage. Keep them on volatile memories and just save very common on non-volatile storages.
}
