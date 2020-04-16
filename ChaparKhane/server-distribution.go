/* For license and copyright information please see LEGAL file in repository */

package chaparkhane

// DomainRequestLog is use to improve intelligence of dev-app starting in needed area like cloud, fog, edge,...!
type DomainRequestLog struct {
	UserID    [16]byte
	UserUIP   [16]byte
	Domain    string
	DomainUIP [16]byte
}
