/* For license and copyright information please see the LEGAL file in the code repository */

package cmd

import (
	"libgo/protocol"
)

// These functions are helper to implement libgo/protocol.CommandLineArguments easier.
func FromCLA(object protocol.Object, arguments []string) (remaining []string, err protocol.Error) {
	var flagSet FlagSet
	flagSet.Init(object, arguments)
	err = flagSet.Parse()
	remaining = flagSet.Args()
	return

}
func ToCLA(object protocol.Object) (arguments []string, err protocol.Error) {
	var fields []protocol.DataType = object.Fields()
	var ln = len(fields)
	if ln < 1 {
		// err =
		return
	}

	arguments = make([]string, 0, ln)
	for i := 0; i < ln; i++ {
		var field = fields[i]
		var fieldValue = field.ToString()
		if fieldValue == "" {
			continue
		}
		var fieldName = field.Detail(protocol.AppLanguage).Name()
		arguments = append(arguments, fieldName)
		arguments = append(arguments, fieldValue)
	}
	return
}
