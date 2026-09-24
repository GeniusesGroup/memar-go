/* For license and copyright information please see the LEGAL file in the code repository */

package services

// accessor methods or Getters & Setters
// `get:"public" set:"private"`
// `methods:"get; private set;"`

// public		The code is accessible for all classes
// private		The code is only accessible within the same class
// protected	The code is accessible within the same class, or in a class that is inherited from that class. You will learn more about inheritance in a later chapter
// internal		The code is only accessible within its own assembly, but not from another assembly. You will learn more about this in a later chapter

// https://github.com/masaushi/accessory

type Test struct {
	fieldA string `get:"public" set:"private" json:"a"`
}

func (t *Test) FieldA() string     { return t.fieldA }
func (t *Test) SetFieldA(f string) { t.fieldA = f }
