package env

import (
	"fmt"
	"os"
	"strconv"
)

// Get returns the value of a specified environment variable.
// If the environment variable does not exist, the default value is returned.
func Get(name string, defaultValue string) string {
	if value := os.Getenv(name); value != "" {
		return value
	}
	return defaultValue
}

// GetInt returns the value of a specified environment variable as an int.
// If the environment variable does not exist, the default value is returned.
// If the environment variable value is not an integer, an error is returned.
func GetInt(name string, defaultValue int) (int, error) {
	if value := os.Getenv(name); value != "" {
		return strconv.Atoi(value)
	}
	return defaultValue, nil
}

// MustGetInt returns the value of a specified environment variable as an int.
// If the environment variable does not exist, the default value is returned.
// If the environment variable value is not an integer, the function panicss.
func MustGetInt(name string, defaultValue int) int {
	value, err := GetInt(name, defaultValue)
	if err != nil {
		panic(fmt.Sprintf("invalid value for env var %q: %v", name, err))
	}
	return value
}
