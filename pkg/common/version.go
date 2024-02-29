// File pretty much completely stolen from a generated
// Autostrada project I used a while ago. They don't
// require attribution but I still want to give them
// credit for their amazing project.
// https://autostrada.dev/

package common

import (
	"fmt"
	"runtime/debug"
)

// GetVersion returns the current git commit hash.
// If the version is modified the hash ends in "-dirty"
func GetVersion() string {
	var revision string
	var modified bool

	bi, ok := debug.ReadBuildInfo()
	if ok {
		for _, s := range bi.Settings {
			switch s.Key {
			case "vcs.revision":
				revision = s.Value
			case "vcs.modified":
				if s.Value == "true" {
					modified = true
				}
			}
		}
	}

	if revision == "" {
		return "unavailable"
	}

	if modified {
		return fmt.Sprintf("%s-dirty", revision)
	}

	return revision
}

// GetVersion returns the current git commit hash.
// This function does not add the "-dirty" string at the end of the hash.
func GetVersionPure() string {
	var revision string
	var modified bool

	bi, ok := debug.ReadBuildInfo()
	if ok {
		for _, s := range bi.Settings {
			switch s.Key {
			case "vcs.revision":
				revision = s.Value
			case "vcs.modified":
				if s.Value == "true" {
					modified = true
				}
			}
		}
	}

	if revision == "" {
		return "unavailable"
	}

	if modified {
		return revision
	}

	return revision
}
