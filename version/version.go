package version

import (
	"runtime/debug"
)

type BuildInfo struct {
	*debug.BuildInfo
}

func New() *BuildInfo {
	info, ok := debug.ReadBuildInfo()
	if !ok {
		return nil
	}

	return &BuildInfo{info}
}

func (info *BuildInfo) Revision() (string, bool) {
	rev := ""
	modified := false
	for _, setting := range info.Settings {
		switch setting.Key {
		case "vcs.revision":
			rev = setting.Value
		case "vcs.modified":
			if setting.Value == "true" {
				modified = true
			}
		}
	}

	return rev, modified
}

func (info *BuildInfo) VersionString(sumLen, revLen int) string {
	s := info.Main.Version
	if 0 < sumLen {
		sum := info.Main.Sum
		if sum != "" {
			s += "@" + headString(sum, sumLen)
		}
	}

	if 0 < revLen {
		rev := info.ReversionString(revLen, "-dirty")
		if rev != "" {
			s += "[" + rev + "]"
		}
	}

	return s
}

func (info *BuildInfo) ReversionString(n int, dirty string) string {
	rev, modified := info.Revision()
	if rev != "" {
		s := headString(rev, n)
		if modified {
			s += dirty
		}

		return s
	}

	return ""
}

func headString(s string, n int) string {
	if len(s) < n {
		return s
	}

	return s[:n]
}
