/* For license and copyright information please see the LEGAL file in the code repository */

package net

import (
	"sync/atomic"

	"memar/protocol"
	"memar/time/unix"
)

// ConnectionsMetric store the connection metric data and implement protocol.ConnectionMetrics
type ConnectionsMetric struct {
	lastUsage           unix.Atomic
	openCount           atomic.Int64
	guestCount          atomic.Int64
	inUseCount          atomic.Int64
	idleCount           atomic.Int64
	waitCount           atomic.Int64
	closedCount         atomic.Int64
	guestClosedCount    atomic.Int64
	idleClosedCount     atomic.Int64
	waitClosedCount     atomic.Int64
	lifetimeClosedCount atomic.Int64
}

//libgo:impl memar/protocol.ConnectionsMetrics
func (cm *ConnectionsMetric) LastUsage() protocol.Time   { return &cm.lastUsage }
func (cm *ConnectionsMetric) OpenCount() int64           { return cm.openCount.Load() }
func (cm *ConnectionsMetric) GuestCount() int64          { return cm.guestCount.Load() }
func (cm *ConnectionsMetric) InUseCount() int64          { return cm.inUseCount.Load() }
func (cm *ConnectionsMetric) IdleCount() int64           { return cm.idleCount.Load() }
func (cm *ConnectionsMetric) WaitCount() int64           { return cm.waitCount.Load() }
func (cm *ConnectionsMetric) ClosedCount() int64         { return cm.closedCount.Load() }
func (cm *ConnectionsMetric) GuestClosedCount() int64    { return cm.guestClosedCount.Load() }
func (cm *ConnectionsMetric) IdleClosedCount() int64     { return cm.idleClosedCount.Load() }
func (cm *ConnectionsMetric) WaitClosedCount() int64     { return cm.waitClosedCount.Load() }
func (cm *ConnectionsMetric) LifetimeClosedCount() int64 { return cm.lifetimeClosedCount.Load() }

func (cm *ConnectionsMetric) ConnOpened() {
	cm.openCount.Add(1)
	cm.inUseCount.Add(1)
}
func (cm *ConnectionsMetric) GuestConnOpened() {
	cm.ConnOpened()
	cm.guestCount.Add(1)
}
func (cm *ConnectionsMetric) ConnIdled() {
	cm.inUseCount.Add(-1)
	cm.idleCount.Add(1)
}
func (cm *ConnectionsMetric) ConnWaited() {
	cm.inUseCount.Add(-1)
	cm.waitCount.Add(1)
}
func (cm *ConnectionsMetric) ConnClosed() {
	cm.inUseCount.Add(-1)
	cm.closedCount.Add(1)
}
func (cm *ConnectionsMetric) GuestConnClosed() {
	cm.ConnClosed()
	cm.guestClosedCount.Add(1)
}
func (cm *ConnectionsMetric) IdleConnClosed() {
	cm.closedCount.Add(1)
	cm.idleCount.Add(-1)
	cm.idleClosedCount.Add(1)
}
func (cm *ConnectionsMetric) WaitConnClosed() {
	cm.closedCount.Add(1)
	cm.waitCount.Add(-1)
	cm.waitClosedCount.Add(1)
}
func (cm *ConnectionsMetric) LifetimeConnClosed() {
	cm.ConnClosed()
	cm.lifetimeClosedCount.Add(1)
}
