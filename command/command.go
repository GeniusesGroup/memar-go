/* For license and copyright information please see the LEGAL file in the code repository */

package cmd

import (
	"libgo/protocol"
)

type Command struct {
	// parent is a parent command for this command.
	parent protocol.Command
	// Commands lists the available commands and help topics.
	// The order here is the order in which they are printed by 'go help'.
	// Note that subcommands are in general best avoided.
	subCommands []protocol.Command
}

//libgo:impl libgo/protocol.ObjectLifeCycle
func (c *Command) Init(parent protocol.Command, cmd ...protocol.Command) (err protocol.Error) {
	c.parent = parent
	// TODO::: check duplicate name usage
	c.subCommands = append(c.subCommands, cmd...)
	return
}

//libgo:impl libgo/protocol.Command
func (c *Command) Runnable() bool                  { return false }
func (c *Command) Parent() protocol.Command        { return c.parent }
func (c *Command) SubCommands() []protocol.Command { return c.subCommands }
func (c *Command) SubCommand(name string) protocol.Command {
	// TODO::: intelligent suggestion or correction
	for _, cmd := range c.subCommands {
		var cmdDetail = cmd.Detail(protocol.AppLanguage)
		var cmdName = cmdDetail.Name()
		var cmdAbbr = cmdDetail.Abbreviation()

		if cmdName == name {
			return cmd
		}
		if cmdAbbr == name {
			return cmd
		}
		for _, alias := range cmdDetail.Aliases() {
			if alias == name {
				return cmd
			}
		}
	}
	return nil
}

// ServeCLI read and write to os.Stdin, os.Stdout, and os.Stderr files
func (c *Command) ServeCLI() (err protocol.Error) {
	err = &ErrServiceNotAcceptCLI
	return
}
