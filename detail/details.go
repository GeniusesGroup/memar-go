/* For license and copyright information please see the LEGAL file in the code repository */

package detail

import (
	"libgo/protocol"
)

type Details struct {
	// remove detail map and find by iterate details.
	details map[protocol.LanguageID]Detail
}

func (d *Details) Detail(lang protocol.LanguageID) protocol.Detail {
	// TODO::: if not exist??
	var dt = d.details[lang]
	// TODO::: below return style cause dt escape to heap! replace map with other hash table to prevent this heap alloc.
	return &dt
}

// SetDetail add error text details to existing error and return it.
func (d *Details) SetDetail(dt Detail) (err protocol.Error) {
	var lang = dt.Language()
	var _, ok = d.details[lang]
	if ok {
		panic("/libgo/detail - Can't change detail after first set! Ask the holder to change details.")
		// TODO::: error package need this package! we can't import it(import cycle problem) but it isn't our best practice to use panic any where.
		// err = &
		// return
	}

	if d.details == nil {
		d.details = make(map[protocol.LanguageID]Detail)
	}
	d.details[lang] = dt
	return
}
