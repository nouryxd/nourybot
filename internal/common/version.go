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

	fmt.Println(revision)

	return revision
}
