/* For license and copyright information please see LEGAL file in repository */

package giti

// SocietyID indicate society ID with some usefull methods
type SocietyID uint32

const (
	SocietyUnSet SocietyID = iota
	SocietyPersia
	SocietyIran
)

// SocietyUUID = sha512.Sum512_256([]byte(SocietyName))