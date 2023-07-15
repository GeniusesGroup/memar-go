/* For license and copyright information please see the LEGAL file in the code repository */

package protocol

type Len interface {
	Len() int // len in byte, -1 means it is dynamic and len is runtime base.
}

type Cap interface {
	Cap() int
}
