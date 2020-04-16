/* For license and copyright information please see LEGAL file in repository */

package engine

// Page :
type Page struct {
	Name     string // It must be unique e.g. product
	Icon       []byte
	Info       []Information
	LocaleInfo Information
}
