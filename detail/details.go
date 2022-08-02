/* For license and copyright information please see LEGAL file in repository */

package detail

import (
	"github.com/GeniusesGroup/libgo/protocol"
)

type DetailsContainer struct {
	detail  map[protocol.LanguageID]protocol.Detail
	details []protocol.Detail
}

func (d *DetailsContainer) Details() []protocol.Detail { return d.details }
func (d *DetailsContainer) Detail(lang protocol.LanguageID) protocol.Detail {
	return d.detail[lang]
}

// SetDetail add error text details to existing error and return it.
func (d *DetailsContainer) SetDetail(lang protocol.LanguageID, domain, summary, overview, userNote, devNote string, tags []string) {
	var _, ok = d.detail[lang]
	if ok {
		panic("detail - Can't change detail after first set! Ask the holder to change details.")
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
