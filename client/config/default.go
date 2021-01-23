package config

func GetDefaultConfig() *Config {
	return &Config{
		Prompt: Prompt{
			HOffset: 1,
		},
	}
}
