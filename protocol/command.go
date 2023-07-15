/* For license and copyright information please see the LEGAL file in the code repository */

package protocol

// Command is the interface that must implement by any struct to be a command service
type Command interface {
	// Init(parent Command, subCommands ...Command)

	// Runnable reports whether the command can be run; otherwise it is a documentation pseudo-command
	Runnable() bool

	// parent is the parent command for this command.
	// It can be nill for the root command.
	Parent() Command
	// SubCommand return a sub command by its name or alias that must use intelligent suggestion
	SubCommand(name string) Command
	// Commands lists the available commands and help topics.
	// The order here is the order in which they are printed by 'go help'.
	// Note that subcommands are in general best avoided.
	SubCommands() []Command

	CommandHandler
	ServiceDetails
}

// CommandHandler introduce CLI (command-line interface) service handler.
type CommandHandler interface {
	// ServeCLA or serve by command-line arguments might block the caller
	// Arguments list not include the command name.
	ServeCLA(arguments []string) (err Error)

	// read and write to e.g. os.Stdin, os.Stdout, and os.Stderr files
	// ServeCLI() (err Error)
}

type CommandLineArguments interface {
	FromCLA(arguments []string) (remaining []string, err Error)
	ToCLA() (arguments []string, err Error)
}
