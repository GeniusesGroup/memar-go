/* For license and copyright information please see the LEGAL file in the code repository */

package detail

import (
	"libgo/protocol"
)

// DS is the same as the Details.
// Use this type when embed in other struct to solve field & method same name problem(Details struct and Details() method) to satisfy interfaces.
type DS = Details

type Details struct {
	detail  map[protocol.LanguageID]protocol.Detail
	details []protocol.Detail
}

func (d *Details) Details() []protocol.Detail { return d.details }
func (d *Details) Detail(lang protocol.LanguageID) protocol.Detail {
	return d.detail[lang]
}

// SetDetail add error text details to existing error and return it.
func (d *Details) SetDetail(lang protocol.LanguageID, domain, summary, overview, userNote, devNote string, tags []string) {
	var _, ok = d.detail[lang]
	if ok {
		panic("/libgo/detail - Can't change detail after first set! Ask the holder to change details.")
	}

	var detail = Detail{
		languageID: lang,
		domain:     domain,
		summary:    summary,
		overview:   overview,
		userNote:   userNote,
		devNote:    devNote,
		tags:       tags,
	}
	if d.detail == nil {
		d.detail = make(map[protocol.LanguageID]protocol.Detail)
	}
	d.detail[lang] = &detail
	d.details = append(d.details, &detail)
}
