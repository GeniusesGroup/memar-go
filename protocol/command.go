/* For license and copyright information please see the LEGAL file in the code repository */

package protocol

// Command is the interface that must implement by any struct to be a command service
type Command interface {
	// Init(parent Command, subCommands ...Command)

	Name() string // e.g. init
	// These are not suggested to the user in the shell completion,
	// but accepted if entered manually.
	Aliases() []string // e.g. []string{"initialize", "initialise", "create"}

	// UsageLine is the one-line usage message.
	// Recommended syntax is as follow:
	//   [ ] identifies an optional argument. Arguments that are not enclosed in brackets are required.
	//   ... indicates that you can specify multiple values for the previous argument.
	//   |   indicates mutually exclusive information. You can use the argument to the left of the separator or the
	//       argument to the right of the separator. You cannot use both arguments in a single use of the command.
	//   { } delimits a set of mutually exclusive arguments when one of the arguments is required. If the arguments are
	//       optional, they are enclosed in brackets ([ ]).
	// Example: '<app> add [-F file | -D dir]... [-f format] profile'
	UsageLine() string

	// Runnable reports whether the command can be run; otherwise it is a documentation pseudo-command
	Runnable() bool

	// parent is a parent command for this command.
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

// CommandHandler is any object to be CLI (command-line interface) service handler.
type CommandHandler interface {
	// ServeCLA or serve by command-line arguments might block the caller
	// Arguments list not include the command name.
	ServeCLA(arguments []string) (err Error)

	// read and write to os.Stdin, os.Stdout, and os.Stderr files
	// ServeCLI() (err Error)
}

type CommandLineArguments interface {
	FromCLA(arguments []string) (remaining []string, err Error)
	ToCLA() (arguments []string, err Error)
}
