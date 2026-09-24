/* For license and copyright information please see the LEGAL file in the code repository */

package cmd

import (
	error_p "memar/process/error/protocol"
	// command_errs "memar/process/command/errors"
	command_p "memar/process/command/protocol"
)

// Command
//
//memar:impl memar/process/command/protocol.Command
type Command struct {
	// parent is a parent command for this command.
	parent command_p.Command
	// Commands lists the available commands and help topics.
	// The order here is the order in which they are printed by 'go help'.
	// Note that subcommands are in general best avoided.
	subCommands []command_p.Command
}

//memar:impl memar/computer/capsule/protocol.LifeCycle
func (c *Command) Init(parent command_p.Command, cmd ...command_p.Command) (err error_p.Error) {
	c.parent = parent
	// TODO::: check duplicate name usage
	c.subCommands = append(c.subCommands, cmd...)
	return
}

//memar:impl memar/process/command/protocol.Command
func (c *Command) RunnableCommand() bool            { return false }
func (c *Command) ParentCommand() command_p.Command { return c.parent }
func (c *Command) SubCommands() []command_p.Command { return c.subCommands }
func (c *Command) SubCommand(name string) command_p.Command {
	// TODO::: intelligent suggestion or correction
	for _, cmd := range c.subCommands {
		var cmdName = cmd.Name()
		var cmdAbbr = cmd.Abbreviation()

		if cmdName == name {
			return cmd
		}
		if cmdAbbr == name {
			return cmd
		}
		for _, alias := range cmd.Aliases() {
			if alias == name {
				return cmd
			}
		}
	}
	return nil
}

// ServeCLI read and write to os.Stdin, os.Stdout, and os.Stderr files
func (c *Command) ServeCLI() (err error_p.Error) {
	// err = &command_errs.ErrServiceNotAcceptCLI
	return
}
