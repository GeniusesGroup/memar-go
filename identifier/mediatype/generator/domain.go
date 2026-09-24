/* For license and copyright information please see the LEGAL file in the code repository */

package mediatype_generator

// e.g. "domain/memar.scm.geniuses.group; package=errors; type=error; name="duplicate_identifier""
const domainTemplate = "domain/{{.Domain}}; package={{.Package}}; type={{.Type}}; name={{.Name}}"

type Domain struct {
	f_domain string
	f_package string
	f_type string
	f_name string
}


