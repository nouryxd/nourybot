package common

import "runtime/debug"

// GetVersion returns the current git commit hash.
func GetGoVersion() string {
	var version string

	bi, ok := debug.ReadBuildInfo()
	if ok {
		version = bi.GoVersion
	}

	return version
}

// GetVersion returns the current git commit hash.
func GetCommit() string {
	var revision string
	var short string

	bi, ok := debug.ReadBuildInfo()
	if ok {
		for _, s := range bi.Settings {
			switch s.Key {
			case "vcs.revision":
				revision = s.Value
				short = revision[:7]
			}
		}
	}

	if revision == "" {
		return "unavailable"
	}

	return short
}
