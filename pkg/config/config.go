package config

// GitHub defines the github configuration.
type GitHub struct {
	Token string `mapstructure:"token"`
}

// Logs defines the level and color for log configuration.
type Logs struct {
	Level  string `mapstructure:"level"`
	Pretty bool   `mapstructure:"pretty"`
	Color  bool   `mapstructure:"color"`
}

// Config defines the general configuration.
type Config struct {
	GitHub GitHub `mapstructure:"github"`
	Logs   Logs   `mapstructure:"log"`
}

// Load initializes a default configuration struct.
func Load() *Config {
	return &Config{}
}
