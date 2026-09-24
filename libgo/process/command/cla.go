/* For license and copyright information please see the LEGAL file in the code repository */

package cmd

import (
	capsule_p "memar/computer/capsule/protocol"
	datatype_p "memar/computer/datatype/protocol"
	error_p "memar/process/error/protocol"
)

// These functions are helper to implement `command/protocol.Decoder` and `command/protocol.Encoder` easier.
//
//memar:impl memar/process/command/protocol.Decoder
func FromCLA(object capsule_p.Capsule, arguments []string) (remaining []string, err error_p.Error) {
	var flagSet FlagSet
	flagSet.Init(object, arguments)
	err = flagSet.Parse()
	remaining = flagSet.Args()
	return

}

//
//memar:impl memar/process/command/protocol.Encoder
func ToCLA(object capsule_p.Capsule) (arguments []string, err error_p.Error) {
	var fields []datatype_p.DataType = object.Fields()
	var ln = len(fields)
	if ln < 1 {
		// err =
		return
	}

	arguments = make([]string, 0, ln)
	for i := 0; i < ln; i++ {
		var field = fields[i]
		var fieldValue, _ = field.ToString()
		if fieldValue == "" {
			continue
		}
		var fieldName = field.Name()
		arguments = append(arguments, fieldName)
		arguments = append(arguments, fieldValue)
	}
	return
}
