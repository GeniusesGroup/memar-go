/* For license and copyright information please see LEGAL file in repository */

package protocol

const (
	SyllabDomain = "syllab.protocol"

	SyllabServiceCompleteGoSyllabMethods = "urn:giti:syllab.protocol:service:complete-go-syllab-methods"
	SyllabServiceGenerateGoSyllabMethods = "urn:giti:syllab.protocol:service:generate-go-syllab-methods"
)

// Syllab is the interface that must implement by any struct to be a Syllab object tranmitable over networks!
// Standards in https://github.com/GeniusesGroup/RFCs/blob/master/Syllab.md
type Syllab interface {
	// CheckSyllab usually just check LenOfSyllabStack not greater than len of given payload. and call just before decode payload.
	CheckSyllab(payload []byte) (err Error)

	// ToSyllab encode the struct pointer to Syllab format
	// in non embed struct usually `stackIndex = 0` & `heapIndex = {{rn}}.LenOfSyllabStack()` as heap start index || end of stack size!
	ToSyllab(payload []byte, stackIndex, heapIndex uint32) (freeHeapIndex uint32)
	// FromSyllab decode Syllab to the struct pointer!
	FromSyllab(payload []byte, stackIndex uint32)

	// LenAsSyllab return whole calculated length of Syllab encoded of the struct
	// default is simple as `return uint64({{RecName}}.LenOfSyllabStack() + {{RecName}}.LenOfSyllabHeap())`
	LenAsSyllab() uint64

	// LenOfSyllabStack return calculated stack length of Syllab encoded of the struct
	LenOfSyllabStack() uint32
	// LenOfSyllabStack return calculated heap length of Syllab encoded of the struct
	LenOfSyllabHeap() uint32
}
