/* For license and copyright information please see LEGAL file in repository */

package generator

// UpdateDatastoreFileReq is request structure of UpdateDatastoreFile()
type UpdateDatastoreFileReq struct {
	DatastoreFile []byte
}

// UpdateDatastoreFileRes is response structure of UpdateDatastoreFile()
type UpdateDatastoreFileRes struct {
	DatastoreFile []byte
}

// UpdateDatastoreFile use to update datastore file and complete or edit some auto generate part.
func UpdateDatastoreFile(req *UpdateDatastoreFileReq) (res *UpdateDatastoreFileRes, err error) {
	return
}
