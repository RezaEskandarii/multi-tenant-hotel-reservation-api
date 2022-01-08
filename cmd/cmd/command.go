package cmd

import (
	"flag"
)

type Command struct {
	fs   *flag.FlagSet
	name string
	Func func() error
}

func NewCommand(cmdName string, defaultVal string, usage string, fn func() error) *Command {
	c := &Command{
		fs:   flag.NewFlagSet(cmdName, flag.ContinueOnError),
		Func: fn,
		name: cmdName,
	}
	c.fs.StringVar(&c.name, cmdName, defaultVal, usage)
	commands := []string{cmdName}
	c.fs.Parse(commands)
	return c
}

// Name returns name of command
func (c *Command) Name() string {
	return c.fs.Name()
}

// Run run flag
func (c *Command) Run() error {
	if c.Func != nil {
		return c.Func()
	}
	return nil
}
