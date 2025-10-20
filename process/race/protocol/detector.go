/* For license and copyright information please see the LEGAL file in the code repository */

package race_p

import (
	event_p "memar/process/event/protocol"
	error_p "memar/process/error/protocol"
)

type Detector interface {
	event_p.Target[Report_Event]
}

// Static means when something is done not at Runtime but could be done anytime e.g. commit, PR, pre-compile, compile, ...
// 
// https://en.wikipedia.org/wiki/Static_program_analysis
type Detector_Static interface {
	Detector

	// DetectRaces use to `Analyze` programs for data races.
	// TODO::: Need options?? e.g. ignorePatterns, AccessType, ProgrammingLanguage, ...
	DetectRaces(ast string, options any) (err error_p.Error)

	AnalyzeVariable(ast string, varName string) (err error_p.Error)
}

// Dynamic means when something is done at Runtime
// 
// https://en.wikipedia.org/wiki/Dynamic_program_analysis
type Detector_Dynamic interface {
	Detector

	// Enable re-enables(Start) handling(Monitoring) of race events in the current thread.
	Enable() (err error_p.Error)
	// Disable disables handling of race synchronization events in the current thread.
	Disable() (err error_p.Error)
}
