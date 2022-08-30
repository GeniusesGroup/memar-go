/* For license and copyright information please see the LEGAL file in the code repository */

package cmd

import (
	"strings"

	"github.com/GeniusesGroup/libgo/protocol"
)

// A FlagSet represents a set of defined fields
//
// field names must be unique within a FlagSet. An attempt to define a flag whose
// name is already in use will cause a panic.
type FlagSet struct {
	fields []protocol.Field
	parsed []protocol.Field
	args   []string // arguments after flags
}

// argument list should not include the commands name.
func (f *FlagSet) Init(fields []protocol.Field, arguments []string) {
	f.fields = fields
	f.checkFields()
	if f.parsed == nil {
		f.parsed = make([]protocol.Field, 0, len(f.fields))
	}
	f.args = arguments
}
func (f *FlagSet) Reinit() {
	f.fields = nil
	f.parsed = f.parsed[:0]
	f.args = nil
}
func (f *FlagSet) Deinit() { f.Reinit() }

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
func (f *FlagSet) VisitAll(fn func(protocol.Field)) {
	for _, flag := range f.fields {
		fn(flag)
	}
}

// Visit visits the flags in given order, calling fn for each.
// It visits only those flags that have been set.
func (f *FlagSet) Visit(fn func(protocol.Field)) {
	for _, field := range f.parsed {
		fn(field)
	}
}

// Lookup returns the Field of the named flag, returning nil if none exists.
func (f *FlagSet) Lookup(name string) protocol.Field {
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
		return &ErrFlagNotFound
	}

	var hasValue = len(value) > 0

	if flag.Type() == protocol.FieldType_Boolean && !hasValue { // special case: doesn't need an arg
		value = "true"
	} else {
		// It must have a value, which might be the next argument.
		if !hasValue && len(f.args) > 0 {
			// value is the next arg
			hasValue = true
			value, f.args = f.args[0], f.args[1:]
		}
		if !hasValue {
			return &ErrFlagNeedsAnArgument
		}
	}

	err = flag.FromString(value)
	if err != nil {
		return
	}
	f.parsed = append(f.parsed, flag)
	return nil
}

func (f *FlagSet) checkFields() {
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
}

// parseOne parses one flag
func (f *FlagSet) parseOne() (err protocol.Error) {
	if len(f.args) == 0 {
		return nil
	}

	var s = f.args[0]
	if len(s) < 2 || s[0] != '-' {
		return &ErrFlagBadSyntax
	}
	var numMinuses = 1
	if s[1] == '-' {
		numMinuses++
		if len(s) == 2 { // "--" terminates the flags
			f.args = f.args[1:]
			return &ErrFlagBadSyntax
		}
	}
	var name = s[numMinuses:]
	if len(name) == 0 || name[0] == '-' || name[0] == '=' {
		return &ErrFlagBadSyntax
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

	err = f.Set(name, value)
	return
}
