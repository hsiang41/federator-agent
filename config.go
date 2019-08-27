package operator

type JobExecutor struct {
	Name                string  `mapstructure:"name"`
	ScheduledTaskSpec   string  `mapstructure:"schedule-spec"`
	LibPath             string  `mapstructure:"lib-path"`
	LibConfiguration    string  `mapstructure:"lib-configuration"`
}

type LogConf struct {
	SetLogCaller bool   `mapstructure:"set-logcallers"`
	OutputLevel string  `mapstructure:"output-level"` // debug, info, warn, error, fatal, none
}

// Config defines configurations
type Config struct {
	Log              *LogConf                  `mapstructure:"log"`
	InputJobs        map[string]JobExecutor    `mapstructure:"input_jobs"`
	OutputJobs       map[string]JobExecutor    `mapstructure:"output_jobs"`
}

// NewConfig returns Config objecdt
func NewConfig() Config {
	c := Config{
	}
	c.init()
	return c
}

func (c *Config) init() {
	defaultLogConfig := &LogConf{SetLogCaller:false, OutputLevel:"info"}
	c.Log = defaultLogConfig
}

func (c Config) Validate() error {
	return nil
}
