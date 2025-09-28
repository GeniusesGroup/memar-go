/* For license and copyright information please see LEGAL file in repository */

package pool

import (
	"sync"
	"sync/atomic"
	"time"

	"../protocol"
)

// LIFO pool, i.e. the most recently Put() item will return in call Get()
// Such a scheme keeps CPU caches hot (in theory).
//
// Due to LIFO behaviors and MaxItems logic can't implement with sync/atomic like sync.Pool
type LIFO struct {
	MaxItems    int64
	MinItems    int64
	MaxIdles    int64
	IdleTimeout protocol.Duration
	// New optionally specifies a function to generate
	// a value when Get would otherwise return nil.
	// It may not be changed concurrently with calls to Get.
	New func() interface{}
	// CloseFunc optionally specifies a function to call when an item want to drop after timeout
	CloseFunc func(item interface{})

	itemsCount  int64      // idles + actives
	mutex       sync.Mutex // lock below fields until new line
	hotItem     interface{}
	idleItems   []interface{}
	victimItems []interface{}

	state int32
}

func (p *LIFO) Init() {
	var preStatus = p.SetState(PoolStatus_Running)
	if preStatus != PoolStatus_Unset {
		panic("[pool] LIFO instance can't reuse")
	}
	if p.CloseFunc == nil {
		p.CloseFunc = func(item interface{}) {}
	}
	p.idleItems = make([]interface{}, p.MaxItems/4)
	go p.Clean()
}

func (p *LIFO) State() PoolStatus { return PoolStatus(atomic.LoadInt32(&p.state)) }
func (p *LIFO) SetState(state PoolStatus) (pre PoolStatus) {
	return PoolStatus(atomic.SwapInt32(&p.state, int32(state)))
}
func (p *LIFO) Len() (ln int) {
	p.mutex.Lock()
	ln = len(p.idleItems) + len(p.victimItems)
	p.mutex.Unlock()
	return
}

func (p *LIFO) Get() (item interface{}) {
	if p.isStop() {
		return
	}

	item = p.popHead()
	if item == nil && p.New != nil {
		item = p.makeNew()
	}
	return
}

func (p *LIFO) Put(item interface{}) {
	if p.isStop() {
		p.CloseFunc(item)
		return
	}

	p.pushHead(item)
}

func (p *LIFO) Close() {
	var preStatus = p.SetState(PoolStatus_Stopping)
	if preStatus != PoolStatus_Running {
		panic("[pool] LIFO instance wasn't started to call Close()")
	}

	p.mutex.Lock()
	for i := 0; i < len(p.idleItems); i++ {
		var item = p.idleItems[i]
		if item != nil {
			p.CloseFunc(item)
		}
	}
	p.mutex.Unlock()

	p.SetState(PoolStatus_Stopped)
}

// Clean items but not by exact p.IdleTimeout. To improve performance we choose two window clean up.
// Means some idle items can live up to double of p.IdleTimeout
func (p *LIFO) Clean() {
	for {
		time.Sleep(time.Duration(p.IdleTimeout))
		if p.isStop() {
			break
		}

		p.mutex.Lock()
		var vi = p.victimItems
		p.victimItems = p.idleItems
		p.idleItems = make([]interface{}, p.MaxItems/4)
		p.mutex.Unlock()

		atomic.AddInt64(&p.itemsCount, -int64(len(vi)))

		// Close items outside of p.mutex
		for i := 0; i < len(vi); i++ {
			p.CloseFunc(vi[i])
		}
	}
}

func (p *LIFO) isStop() bool {
	var pStatus = p.State()
	if pStatus == PoolStatus_Stopping || pStatus == PoolStatus_Stopped {
		return true
	}
	return false
}

func (p *LIFO) popHead() (item interface{}) {
	p.mutex.Lock()
	if p.hotItem != nil {
		item = p.hotItem
		p.hotItem = nil
	} else {
		var ln = len(p.idleItems) - 1
		if ln > -1 {
			item = p.idleItems[ln]
			p.idleItems = p.idleItems[:ln]
		} else {
			ln = len(p.victimItems) - 1
			if ln > -1 {
				item = p.victimItems[ln]
				p.victimItems = p.victimItems[:ln]
			}
		}
	}
	p.mutex.Unlock()
	return
}

func (p *LIFO) pushHead(item interface{}) {
	p.mutex.Lock()
	if p.hotItem == nil {
		p.hotItem = item
	} else {
		p.idleItems = append(p.idleItems, item)
	}
	p.mutex.Unlock()
}

func (p *LIFO) makeNew() (item interface{}) {
	var ic = atomic.AddInt64(&p.itemsCount, 1)
	if ic <= p.MaxItems {
		item = p.New()
	} else {
		atomic.AddInt64(&p.itemsCount, -1)
	}
	return
}
