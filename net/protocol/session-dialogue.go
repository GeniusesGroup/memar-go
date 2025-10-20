/* For license and copyright information please see the LEGAL file in the code repository */

package net_p

import (
	datatype_p "memar/computer/datatype/protocol"
)

type Field_DialogueType interface {
	DialogueType() OSI_Session_DialogueType
}

// Dialogue is a discussion intended to produce an agreement
// https://en.wikipedia.org/wiki/Session_layer#Dialogue_control
type OSI_Session_DialogueType uint8

const (
	OSI_Session_DialogueType_Unset      OSI_Session_DialogueType = iota
	OSI_Session_DialogueType_FullDuplex                          // allowing communication in opposite directions simultaneously
	OSI_Session_DialogueType_HalfDuplex                          // information can be sent in only one direction at a time (two way alternate)
	OSI_Session_DialogueType_Simplex                             // one way (Monolog)
)
