package env

import "os"

// Get returns the value of a specified environment variable.
// If the environment variable does not exist, the default value is returned.
func Get(name string, defaultValue string) string {
	if value := os.Getenv(name); value != "" {
		return value
	}
	return defaultValue
}
