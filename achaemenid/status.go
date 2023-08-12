/* For license and copyright information please see the LEGAL file in the code repository */

package achaemenid

import (
	"memar/protocol"
	"sync"
)

type status struct {
	status             protocol.ApplicationStatus
	statusListeners    []chan protocol.ApplicationStatus
	statusChangeLocker sync.Mutex
}

func (s *status) Status() protocol.ApplicationStatus { return s.status }
func (s *status) NotifyStatus(notifyBy chan protocol.ApplicationStatus) {
	s.statusListeners = append(s.statusListeners, notifyBy)
}

//memar:impl memar/protocol.ObjectLifeCycle
func (s *status) Init() (err protocol.Error)   { return }
func (s *status) Reinit() (err protocol.Error) { return }
func (s *status) Deinit() (err protocol.Error) { return }

func (s *status) changeState(status protocol.ApplicationStatus) {
	s.statusChangeLocker.Lock()
	defer s.statusChangeLocker.Unlock()
	s.status = status
	for _, listener := range s.statusListeners {
		// Can't be blocking if listener hasn't enough capacity buffer??
		select {
		case listener <- status:
		default:
		}
	}
}
