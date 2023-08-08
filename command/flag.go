/* For license and copyright information please see the LEGAL file in the code repository */

package cmd

import (
	"strings"

	errs "memar/command/errors"
	"memar/protocol"
)

// A FlagSet represents a set of defined fields
type FlagSet struct {
	object protocol.Object
	fields []protocol.DataType
	parsed []protocol.DataType
	args   []string // arguments after flags
}

// argument list should not include the commands name.
//
// field names must be unique within a FlagSet. An attempt to define a flag whose
// name is already in use will cause panic.
//
//memar:impl memar/protocol.ObjectLifeCycle
func (f *FlagSet) Init(ob protocol.Object, arguments []string) (err protocol.Error) {
	f.object = ob
	f.fields = ob.Fields()
	err = f.checkFields()
	if err != nil {
		return
	}
	if f.parsed == nil {
		f.parsed = make([]protocol.DataType, 0, len(f.fields))
	}
	f.args = arguments
	return
}
func (f *FlagSet) Reinit() (err protocol.Error) {
	f.fields = nil
	f.parsed = f.parsed[:0]
	f.args = nil
	return
}
func (f *FlagSet) Deinit() (err protocol.Error) {
	return
}

// Parse parses flag definitions from f.args.
// It Must be called after Init called and before flags are accessed by the program.
func (f *FlagSet) Parse() (err protocol.Error) {
	for len(f.args) > 0 {
		err = f.parseOne()
		if err != nil {
			return
		}
	}
	return
}

// NFlag returns the number of flags that have been set.
func (f *FlagSet) NFlag() int { return len(f.fields) }

// NArg is the number of arguments remaining after flags have been processed.
func (f *FlagSet) NArg() int { return len(f.args) }

// Args returns the non-flag arguments.
func (f *FlagSet) Args() []string { return f.args }

// Arg returns the i'th argument. Arg(0) is the first remaining argument
// after flags have been processed. Arg returns an empty string if the
// requested element does not exist.
func (f *FlagSet) Arg(i int) string {
	if i < 0 || i >= len(f.args) {
		return ""
	}
	return f.args[i]
}

// VisitAll visits the flags in given order, calling fn for each.
// It visits all flags, even those not set.
func (f *FlagSet) VisitAll(fn func(protocol.DataType) (breaking bool)) {
	for _, flag := range f.fields {
		var br = fn(flag)
		if br {
			return
		}
	}
}

// Visit visits the flags in given order, calling fn for each.
// It visits only those flags that have been set.
func (f *FlagSet) Visit(fn func(protocol.DataType) (breaking bool)) {
	for _, field := range f.parsed {
		var br = fn(field)
		if br {
			return
		}
	}
}

// Lookup returns the Field of the named flag, returning nil if none exists.
func (f *FlagSet) Lookup(name string) protocol.DataType {
	for _, field := range f.fields {
		if field.Name() == name || field.Abbreviation() == name {
			return field
		}
	}
	return nil
}

// Set sets the value of the named flag.
func (f *FlagSet) Set(name, value string) (err protocol.Error) {
	var flag = f.Lookup(name)
	if flag == nil {
		return &errs.ErrFlagNotFound
	}

	err = flag.FromString(value)
	if err != nil {
		return
	}
	f.parsed = append(f.parsed, flag)
	return
}

func (f *FlagSet) checkFields() (err protocol.Error) {
	var fieldsName = make([]string, 0, len(f.fields))

	for _, field := range f.fields {
		var fieldName = field.Name()

		// Flag must not begin "-" or contain "=".
		if strings.HasPrefix(fieldName, "-") {
			panic(fieldName + " flag begins with -")
		} else if strings.Contains(fieldName, "=") {
			panic(fieldName + " flag %q contains =")
		}

		fieldsName = append(fieldsName, fieldName)
	}

	for i := 0; i < len(fieldsName); i++ {
		for j := i + 1; j < len(fieldsName); j++ {
			var fieldName = fieldsName[i]
			if fieldName == fieldsName[j] {
				// Happens only if flags are declared with identical names
				if fieldName == "" {
					panic(fieldName + " flag redefined.")
				} else {
					panic(fieldName + "flag redefined as" + fieldsName[j])
				}
			}
		}
	}
	return
}

// parseOne parses one flag
func (f *FlagSet) parseOne() (err protocol.Error) {
	if len(f.args) == 0 {
		return nil
	}

	var s = f.args[0]
	if len(s) < 2 || s[0] != '-' {
		return &errs.ErrFlagBadSyntax
	}
	var numMinuses = 1
	if s[1] == '-' {
		numMinuses++
		if len(s) == 2 { // "--" terminates the flags
			f.args = f.args[1:]
			return &errs.ErrFlagBadSyntax
		}
	}
	var name = s[numMinuses:]
	if len(name) == 0 || name[0] == '-' || name[0] == '=' {
		return &errs.ErrFlagBadSyntax
	}

	// it's a flag. does it have an argument?
	f.args = f.args[1:]
	var value string
	for i := 1; i < len(name); i++ { // equals cannot be first
		if name[i] == '=' {
			value = name[i+1:]
			name = name[0:i]
			break
		}
	}

	err = f.checkAndSet(name, value)
	return
}

// checkAndSet check given value is correct or get from f, and sets the value of the named flag.
func (f *FlagSet) checkAndSet(name, value string) (err protocol.Error) {
	var hasValue = len(value) > 0
	if !hasValue {
		var flag = f.Lookup(name)
		if flag == nil {
			return &errs.ErrFlagNotFound
		}

		var pf, ok = flag.(protocol.DataType_Primitive)
		if ok && pf.Primitive() == protocol.DataType_PrimitiveKind_Boolean { // special case: doesn't need an arg
			value = "true"
		} else if len(f.args) > 0 { // It must have a value, which might be the next argument.
			// value is the next arg
			hasValue = true
			value, f.args = f.args[0], f.args[1:]
		} else {
			return &errs.ErrFlagNeedsAnArgument
		}
	}

	err = f.Set(name, value)
	return
}
