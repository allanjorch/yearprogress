package core

// Config holds user-facing display options. A config file loader can populate this later.
type Config struct {
	ShowEvenDayLabel bool
}

func DefaultConfig() Config {
	return Config{
		ShowEvenDayLabel: false,
	}
}