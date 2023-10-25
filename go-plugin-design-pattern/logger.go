package main

// Plugin implementation
type LoggerPlugin struct {
	isActive bool
}

func (p *LoggerPlugin) Activate() error {
	p.isActive = true
	return nil
}

func (p *LoggerPlugin) Deactivate() error {
	p.isActive = false
	return nil
}

func (p *LoggerPlugin) Execute() string {
	if p.isActive {
		return "Logging data ..."
	}

	return "Plugin is not active"
}
