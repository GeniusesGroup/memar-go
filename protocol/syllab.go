/* For license and copyright information please see the LEGAL file in the code repository */

package protocol

// Syllab is the interface that must implement by any struct to be a Syllab object transmittable over networks.
// Standards in https://github.com/GeniusesGroup/memar/blob/main/Syllab.md
type Syllab interface {
	SyllabMarshaler
	SyllabUnmarshaler
}

type SyllabUnmarshaler interface {
	// CheckSyllab usually just check LenOfSyllabStack not greater than len of given payload. and call just before decode payload.
	CheckSyllab(payload Buffer) (err Error)

	// FromSyllab ready given payload for get accessors methods.
	// - Due to strongly suggest to use fields get accessors methods, below method just change under hood buffer if it isn't struct.
	// - It can return LenOfSyllabStack()-1 as end of stack in payload, but it will be runtime logic. compiler can inline and do -1 in compile time easily.
	FromSyllab(payload Buffer, stackIndex uint32)
}

type SyllabMarshaler interface {
	// ToSyllab encode the struct pointer to Syllab format
	// in non embed struct usually `stackIndex = 0` & `heapIndex = {{rn}}.LenOfSyllabStack()` as heap start index || end of stack size.
	ToSyllab(payload Buffer, stackIndex, heapIndex uint32) (freeHeapIndex uint32)

	// LenAsSyllab return whole calculated length of Syllab encoded of the struct
	// default is simple as `return uint64({{RecName}}.LenOfSyllabStack() + {{RecName}}.LenOfSyllabHeap())`
	LenAsSyllab() uint64

	// LenOfSyllabStack return calculated stack length of Syllab encoded of the struct
	LenOfSyllabStack() uint32
	// LenOfSyllabStack return calculated heap length of Syllab encoded of the struct
	LenOfSyllabHeap() uint32
}
