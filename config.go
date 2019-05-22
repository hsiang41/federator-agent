package operator

import (
	"github.com/containers-ai/alameda/operator/datahub"
	"github.com/containers-ai/alameda/pkg/utils/log"
)

// Config defines configurations
type Config struct {
	Log              *log.Config      `mapstructure:"log"`
	Datahub          *datahub.Config  `mapstructure:"datahub"`
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
	c.Datahub = datahub.NewConfig()
}

func (c Config) Validate() error {
	return nil
}
