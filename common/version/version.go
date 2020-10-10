package version

import "fmt"

// Version stores the current version of the application
// Override during a build with: -ldflags "-X version.Version=x.y.z"
var Version string

// Println outputs the current version to stdout
func Println() {
	fmt.Printf("Version: %v\n", Version)
}
