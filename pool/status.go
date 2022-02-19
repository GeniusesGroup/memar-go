/* For license and copyright information please see LEGAL file in repository */

package pool

type PoolStatus int32

const (
	PoolStatus_Unset PoolStatus = iota
	PoolStatus_Running
	PoolStatus_Stopping
	PoolStatus_Stopped
)
