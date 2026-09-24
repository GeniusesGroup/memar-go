/* For license and copyright information please see the LEGAL file in the code repository */

package generator

import (
	"bytes"
	"go/format"
	file_p "memar/storage/file/protocol"
	"text/template"
)

// MakeConnectionsFile use to make connections.go file in root of repo.
func MakeConnectionsFile(file file_p.File) (fileData []byte, err error) {
	// file.Rename("connections.go")

	var mt = new(bytes.Buffer)
	if err = getConnectionFileTemplate.Execute(mt, ""); err != nil {
		return
	}

	fileData, err = format.Source(mt.Bytes())
	return
}

var getConnectionFileTemplate = template.Must(template.New("main").Parse(`/* For license and copyright information please see the LEGAL file in the code repository */

package main

import (
	"memar/achaemenid"
	er "memar/process/error"
)

// GetConnectionsByID returns available connection by given data
func getConnectionsByID(connID [32]byte) (conn *achaemenid.Connection) {
	var err error
	var uac = datastore.UserAppConnection{
		ID: connID,
	}
	err = uac.GetLastByID()
	if err != nil {
		return
	}
	conn = &achaemenid.Connection{
		/* Connection data */
		Server: &server,
		ID:     uac.ID,
		State:  achaemenid.StateLoaded,
		Weight: uac.Weight,

		/* Peer data */
		// Peer Location
		GPAddr:  uac.GPAddr,
		IPAddr:  uac.IPAddr,
		ThingID: uac.ThingID,
		// Peer Identifiers
		UserID:           uac.UserID,
		UserType:         uac.UserType,
		DelegateUserID:   uac.DelegateUserID,
		DelegateUserType: uac.DelegateUserType,

		/* Security data */
		PeerPublicKey: uac.PeerPublicKey,
		// Cipher        crypto.Cipher
		AccessControl: uac.AccessControl,

		/* Metrics data */
		LastUsage:             uac.LastUsage,
		PacketPayloadSize:     uac.PacketPayloadSize,
		MaxBandwidth:          uac.MaxBandwidth,
		ServiceCallCount:      uac.ServiceCallCount,
		BytesSent:             uac.BytesSent,
		PacketsSent:           uac.PacketsSent,
		BytesReceived:         uac.BytesReceived,
		PacketsReceived:       uac.PacketsReceived,
		FailedPacketsReceived: uac.FailedPacketsReceived,
		FailedServiceCall:     uac.FailedServiceCall,
	}
	conn.StreamPool.Init()
	return
}

// getConnectionsByUserIDThingID returns available connection by given data
func getConnectionsByUserIDThingID(userID, thingID [32]byte) (conn *achaemenid.Connection) {
	var err error_p.Error
	var uac = datastore.UserAppConnection{
		UserID:  userID,
		ThingID: thingID,
	}
	err = uac.GetLastByUserIDThingID()
	if err != nil {
		return
	}
	conn = &achaemenid.Connection{
		/* Connection data */
		Server: &server,
		ID:     uac.ID,
		State:  achaemenid.StateLoaded,
		Weight: uac.Weight,

		/* Peer data */
		// Peer Location
		GPAddr:  uac.GPAddr,
		IPAddr:  uac.IPAddr,
		ThingID: uac.ThingID,
		// Peer Identifiers
		UserID:           uac.UserID,
		UserType:         uac.UserType,
		DelegateUserID:   uac.DelegateUserID,
		DelegateUserType: uac.DelegateUserType,

		/* Security data */
		PeerPublicKey: uac.PeerPublicKey,
		// Cipher        crypto.Cipher
		AccessControl: uac.AccessControl,

		/* Metrics data */
		LastUsage:             uac.LastUsage,
		PacketPayloadSize:     uac.PacketPayloadSize,
		MaxBandwidth:          uac.MaxBandwidth,
		ServiceCallCount:      uac.ServiceCallCount,
		BytesSent:             uac.BytesSent,
		PacketsSent:           uac.PacketsSent,
		BytesReceived:         uac.BytesReceived,
		PacketsReceived:       uac.PacketsReceived,
		FailedPacketsReceived: uac.FailedPacketsReceived,
		FailedServiceCall:     uac.FailedServiceCall,
	}
	conn.StreamPool.Init()
	return
}

// saveConnection get a connection and save it its data by platform rules
func saveConnection(conn *achaemenid.Connection) (err error) {
	var uac = datastore.UserAppConnection{
		/* Unique data */
		AppInstanceID:    achaemenid.App.Nodes.LocalNode.InstanceID,
		UserConnectionID: conn.ID,
		Status:           datastore.UserAppConnectionUpdate,
		// Description:      conn.Description,

		/* Connection data */
		ID:     conn.ID,
		Weight: conn.Weight,

		/* Peer data */
		// Peer Location
		GPAddr:  conn.GPAddr,
		IPAddr:  conn.IPAddr,
		ThingID: conn.ThingID,
		// Peer Identifiers
		UserID:           conn.UserID,
		UserType:         conn.UserType,
		DelegateUserID:   conn.DelegateUserID,
		DelegateUserType: conn.DelegateUserType,

		/* Security data */
		PeerPublicKey: conn.PeerPublicKey,
		AccessControl: conn.AccessControl,

		// Metrics data
		LastUsage:             conn.LastUsage,
		PacketPayloadSize:     conn.PacketPayloadSize,
		MaxBandwidth:          conn.MaxBandwidth,
		ServiceCallCount:      conn.ServiceCallCount,
		BytesSent:             conn.BytesSent,
		PacketsSent:           conn.PacketsSent,
		BytesReceived:         conn.BytesReceived,
		PacketsReceived:       conn.PacketsReceived,
		FailedPacketsReceived: conn.FailedPacketsReceived,
		FailedServiceCall:     conn.FailedServiceCall,
	}
	err = uac.Set()
	if err != nil {
		return
	}
	uac.IndexRecordIDForID()
	if conn.State == achaemenid.StateNew {
		conn.State = achaemenid.StateLoaded
		if conn.UserType == user_p.Type_Unset {
			uac.IndexIDForUserTypeDaily()
		} else {
			uac.IndexIDForUserID()
			if uac.ThingID != [32]byte{} {
				uac.IndexIDForThingID()
				uac.IndexIDForUserIDThingID()
				uac.ListThingIDForUserID()
				uac.ListUserIDForThingID()
			}
		}
	}
	return
}
`))
