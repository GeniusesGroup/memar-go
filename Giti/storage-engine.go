/* For license and copyright information please see LEGAL file in repository */

package giti

// StorageEngine is the interface that must implement by any Application!
type StorageEngine interface {
	Init() // Register app or load last running app data by appID & userID
	RegisterObjectStructure(StructureID uint64)
	NewObject(uuid [32]byte, structureID uint64) (object Object, err Error)
	GetObject(uuid [32]byte, structureID uint64) (object Object, err Error)
	SetObject(object Syllab) (err Error)
	DeleteObject(uuid [32]byte, structureID uint64) (err Error)
}
