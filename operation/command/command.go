/* For license and copyright information please see the LEGAL file in the code repository */

package cmd

import (
	errs "memar/operation/command/errors"
	command_p "memar/operation/command/protocol"
	"memar/protocol"
)

// Command
//
//memar:impl memar/operation/command/protocol.Command
type Command struct {
	// parent is a parent command for this command.
	parent command_p.Command
	// Commands lists the available commands and help topics.
	// The order here is the order in which they are printed by 'go help'.
	// Note that subcommands are in general best avoided.
	subCommands []command_p.Command
}

//memar:impl memar/protocol.ObjectLifeCycle
func (c *Command) Init(parent command_p.Command, cmd ...command_p.Command) (err protocol.Error) {
	c.parent = parent
	// TODO::: check duplicate name usage
	c.subCommands = append(c.subCommands, cmd...)
	return
}

//memar:impl memar/operation/command/protocol.Command
func (c *Command) Runnable() bool                   { return false }
func (c *Command) Parent() command_p.Command        { return c.parent }
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
func (c *Command) ServeCLI() (err protocol.Error) {
	err = &errs.ErrServiceNotAcceptCLI
	return
}
