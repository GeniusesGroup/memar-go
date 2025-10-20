/* For license and copyright information please see the LEGAL file in the code repository */

package race_p

import (
  thread_p "memar/computer/runtime/thread/protocol"
	event_p "memar/process/event/protocol"
)

// Event is race report event
type Event interface {
	thread_p.Field_ThreadID
  runtime_p.Stack

  //   variable: string;
  //   accessType: 'read' | 'write';

  //   line: number;
  //   variable: string;
  //   accessType: 'read' | 'write';
  //   conflictingThreads: string[];
  //   suggestion: string;

  // { file: 'main.cpp', line: 42, description: 'Unprotected write' }


	// IsPositive() bool
  // IsNegative() bool
  
  event_p.Event
}
