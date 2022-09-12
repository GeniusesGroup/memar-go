/* For license and copyright information please see the LEGAL file in the code repository */

package uri

// URN implement protocol.URN interface
type URN struct {
	uri string
	nid string
	nss string
}

// SetDetail add error text details to existing error and return it.
func (u *URN) Init(urn string) {
	u.uri = urn
	// TODO:::
}

func (u *URN) URI() string    { return u.uri }
func (u *URN) Scheme() string { return "urn" }
func (u *URN) NID() string    { return u.nid }
func (u *URN) NSS() string    { return u.nss }
