package version

import (
	"fmt"
	"runtime"
)

var (
	CodeVersion  string = "undefined"
	BuildVersion string = "undefined"
	CodeCommit   string = "undefined"
	BuildTime    string = "undefined"
)

type Version struct {
	Version      string `json:"version"`
	Commit       string `json:"commit"`
	BuildTime    string `json:"build_time"`
	BuildVersion string `json:"build_version"`
	GoVersion    string `json:"go_version"`
}

func Details() Version {
	return Version{
		Version:      CodeVersion,
		Commit:       CodeCommit,
		BuildTime:    BuildTime,
		BuildVersion: BuildVersion,
		GoVersion:    runtime.Version(),
	}
}

// String returns version details as pretty printed string.
func (v Version) String() string {
	return fmt.Sprintf(
		"profefe version %s, commit %s (%s), build %s, go version %s",
		v.Version,
		v.Commit,
		v.BuildTime,
		v.BuildVersion,
		v.GoVersion,
	)
}
