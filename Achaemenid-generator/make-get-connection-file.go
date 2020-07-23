/* For license and copyright information please see LEGAL file in repository */

package generator

import (
	"bytes"
	"go/format"
	"text/template"

	"../assets"
)

// MakeGetConnectionFile use to make get-connection.go file.
func MakeGetConnectionFile(file *assets.File) (err error) {
	file.FullName = "get-connection.go"
	file.Name = "get-connection"
	file.Extension = "go"
	file.State = assets.StateChanged

	var mt = new(bytes.Buffer)
	if err = getConnectionFileTemplate.Execute(mt, ""); err != nil {
		return
	}

	file.Data, err = format.Source(mt.Bytes())
	if err != nil {
		return
	}

	return
}

var getConnectionFileTemplate = template.Must(template.New("main").Parse(`
/* For license and copyright information please see LEGAL file in repository */

package main

import "./libgo/achaemenid"

// getConnectionsByID returns available connection by given data if any exist.
func getConnectionsByID(connID [16]byte) (conn *achaemenid.Connection) {
	// TODO::: complete logic otherwise every new connection make as guest
	return
}

// getConnectionsByUserID returns available connection by given data if any exist.
func getConnectionsByUserID(userID, appID, thingID [16]byte) (conn *achaemenid.Connection) {
	// TODO::: complete logic otherwise every new connection make as guest
	return
}
`))
