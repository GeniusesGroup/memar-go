/* For license and copyright information please see LEGAL file in repository */

package service

import "../protocol"

// Service store needed data to authorize incoming service
type Authorization struct {
	crud     protocol.CRUD // CRUD == Create, Read, Update, Delete
	userType protocol.UserType
}
