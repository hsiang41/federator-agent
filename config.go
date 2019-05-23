package operator

import (
	"github.com/containers-ai/alameda/pkg/utils/log"
)

type JobExecutor struct {
	Name                string  `mapstructure:"name"`
	ScheduledTaskSpec   string  `mapstructure:"schedule-spec"`
	LibPath             string  `mapstructure:"lib-path"`
	LibConfiguration    string  `mapstructure:"lib-configuration"`
}

// Config defines configurations
type Config struct {
	Log              *log.Config      `mapstructure:"log"`
	InputJobs        []JobExecutor    `mapstructure:"input_jobs"`
	OutputJobs       []JobExecutor    `mapstructure:"output_jobs"`
}

// NewConfig returns Config objecdt
func NewConfig() Config {
	c := Config{
	}
	c.init()
	return c
}

func (c *Config) init() {
	defaultLogConfig := log.NewDefaultConfig()
	c.Log = &defaultLogConfig
}

func (c Config) Validate() error {
	return nil
}
