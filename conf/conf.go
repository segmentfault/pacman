package conf

// StaticConfig the config that do not update after application run
// Deprecated: suggested use Config instead
type StaticConfig interface {
	// LoadAndSet load config from anywhere and set to conf struct
	LoadAndSet(conf interface{}) error
}

// Config interface determines the common methods for parsing configuration from specified resources
type Config interface {
	Parse(any) error
	Get(string) any
	GetBool(string) bool
	GetString(string) string
	GetInt(string) int
}
