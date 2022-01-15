package config

type Config struct {
	Port int
}

func NewConfig() *Config {
	// in actual practice this config is read from the file or other sources
	// for demo purpose hardcoding the config
	return &Config{
		Port: 8080,
	}
}
