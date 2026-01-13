/* For license and copyright information please see the LEGAL file in the code repository */

package mediatype

// MediaType implement mediatype_p.MediaType interface
// MediaType embed to provide the interface methods when not need to implement by others.
// MT uses when embed in other struct to solve field & method same name problem(MediaType struct and MediaType() method) to satisfy interfaces.
type MT struct{}

//memar:impl memar/protocol.MediaType
func (mt *MT) MediaType() string { return "ERROR::: MediaType not indicated by developers" }
