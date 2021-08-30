/* For license and copyright information please see LEGAL file in repository */

package protocol

// CLIHandler is any object to be CLI (command-line interface) service handler.
type CLIHandler interface {
	ServeCLI(st Stream) (err Error)
}
