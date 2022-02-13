/* For license and copyright information please see LEGAL file in repository */

package log

import (
	"../protocol"
)

type Listeners struct {
	listeners map[listenerKey]protocol.LogListener
}

type listenerKey struct {
	level  protocol.LogType
	domain string
}

func (l *Listeners) init() {
	l.listeners = make(map[listenerKey]protocol.LogListener, 1024)
}

func (l *Listeners) sendToListeners(event protocol.LogEvent) {
	var level = event.Level()
	var domain = event.Domain()

	var key = listenerKey{
		level: level,
		domain: domain,
	}
	var ls = l.listeners[key]
	if ls != nil {
		ls.LogEventHandler(event)
	}

	key = listenerKey{
		level: level,
		// domain: "",
	}
	ls = l.listeners[key]
	if ls != nil {
		ls.LogEventHandler(event)
	}

	key = listenerKey{
		level: protocol.Log_Unset,
		domain: domain,
	}
	ls = l.listeners[key]
	if ls != nil {
		ls.LogEventHandler(event)
	}
}

func (l *Listeners) RegisterListener(level protocol.LogType, domain string, ll protocol.LogListener) (err protocol.Error) {
	var key = listenerKey{
		level: level,
		domain: domain,
	}
	l.listeners[key] = ll
	return
}
