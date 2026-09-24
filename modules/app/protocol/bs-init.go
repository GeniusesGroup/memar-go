/* For license and copyright information please see the LEGAL file in the code repository */

package ap

type (
	BS_Init_Request interface {
		Domain() Domain
		TemplateURI() string // org(organization), society(government), aggregator-org, game(or sub type of org??)
		Blank() bool         // raw (with little predefined codes)
	}

	BS_Init_Response interface {
		Path() []string  // path with file name start from project root
		Files() [][]byte // files data
	}
)
