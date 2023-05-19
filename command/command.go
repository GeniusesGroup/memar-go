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

func (c *Command) Init(parent protocol.Command, cmd ...protocol.Command) {
	c.parent = parent
	// TODO::: check duplicate name usage
	c.subCommands = append(c.subCommands, cmd...)
}

//libgo:impl libgo/protocol.Quiddity
func (c *Command) Name() string         { panic("Dev must implement Name() method to overwrite this method") }
func (c *Command) Abbreviation() string { return "" }
func (c *Command) Aliases() []string    { return nil }

//libgo:impl libgo/protocol.Command
func (c *Command) Runnable() bool                  { return false }
func (c *Command) Parent() protocol.Command        { return c.parent }
func (c *Command) SubCommands() []protocol.Command { return c.subCommands }
func (c *Command) SubCommand(name string) protocol.Command {
	// TODO::: intelligent suggestion or correction
	for _, cmd := range c.subCommands {
		if cmd.Name() == name {
			return cmd
		}
		if cmd.Abbreviation() == name {
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
	err = &ErrServiceNotAcceptCLI
	return
}
