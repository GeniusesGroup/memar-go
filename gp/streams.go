/* For license and copyright information please see LEGAL file in repository */

package gp

import (
	"sync"

	"../protocol"
)

// StreamPool set & get streams in a pool by ID!
type StreamPool struct {
	mutex               sync.Mutex                 // TODO::: it is not efficient way and need more work
	p                   map[uint32]protocol.Stream // key is Stream.ID
	freeIncomeStreamID  uint32
	freeOutcomeStreamID uint32
	totalOpenedStreams  uint32 // Manifest.TechnicalInfo.MaxStreamConnectionDaily
}

// Init initialize the pool
func (sp *StreamPool) Init() {
	sp.p = make(map[uint32]protocol.Stream)
}

// OutcomeStream make the stream and returns it!
func (sp *StreamPool) OutcomeStream(service protocol.Service) (stream protocol.Stream, err protocol.Error) {
	// TODO::: Check stream isn't closed!!
	return
}

// Stream returns Stream from pool if exists by given ID!
func (sp *StreamPool) Stream(id uint32) protocol.Stream {
	// TODO::: Check stream isn't closed!!
	return sp.p[id]
}

// RegisterStream save given Stream to pool
func (sp *StreamPool) RegisterStream(st protocol.Stream) {
	sp.mutex.Lock()
	sp.p[st.GetID()] = st
	sp.mutex.Unlock()
}
