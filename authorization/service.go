/* For license and copyright information please see LEGAL file in repository */

package authorization

// Service store needed data to authorize incoming service
type Service struct {
	CRUD     CRUD // CRUD == Create, Read, Update, Delete
	UserType UserType
}
