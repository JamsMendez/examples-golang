package main

import "fmt"

// Plugin interface
type Plugin interface {
	Activate() error
	Deactivate() error
	Execute() string
}

type Core struct {
	plugin map[string]Plugin
}

func (c *Core) Register(name string, plugin Plugin) {
	if c.plugin == nil {
		c.plugin = make(map[string]Plugin)
	}

	c.plugin[name] = plugin
}

func (c *Core) Active(name string) error {
	if plugin, ok := c.plugin[name]; ok {
		return plugin.Activate()
	}

	return fmt.Errorf("Plugin %s not found", name)
}

func (c *Core) Execute(name string) (string, error) {
	if plugin, ok := c.plugin[name]; ok {
		return plugin.Execute(), nil
	}

	return "", fmt.Errorf("Plugin %s not found", name)
}
