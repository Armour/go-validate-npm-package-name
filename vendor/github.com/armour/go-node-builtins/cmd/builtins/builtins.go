// Package builtins provides function to return nodejs coreModules list.
package builtins

import (
	"fmt"

	"github.com/Masterminds/semver"
)

// GetVersion returns a list of coreModules for specific nodejs version.
func GetVersion(version string) ([]string, error) {
	v, err := semver.NewVersion(version)
	if err != nil {
		return nil, fmt.Errorf("Error parsing version: %s", err)
	}

	coreModules := []string{
		"assert",
		"buffer",
		"child_process",
		"cluster",
		"console",
		"constants",
		"crypto",
		"dgram",
		"dns",
		"domain",
		"events",
		"fs",
		"http",
		"https",
		"module",
		"net",
		"os",
		"path",
		"punycode",
		"querystring",
		"readline",
		"repl",
		"stream",
		"string_decoder",
		"sys",
		"timers",
		"tls",
		"tty",
		"url",
		"util",
		"vm",
		"zlib",
	}

	v2, err := semver.NewVersion("6.0.0")
	if err != nil {
		return nil, fmt.Errorf("Error parsing version: %s", err)
	}
	if v.LessThan(v2) {
		coreModules = append(coreModules, "freelist")
	}

	newModules := []struct {
		version     string
		coreModules []string
	}{
		{"1.0.0", []string{"v8"}},
		{"1.1.0", []string{"process"}},
		{"8.1.0", []string{"async_hooks"}},
		{"8.4.0", []string{"http2"}},
		{"8.5.0", []string{"perf_hooks"}},
	}

	for _, n := range newModules {
		v2, err := semver.NewVersion(n.version)
		if err != nil {
			return nil, fmt.Errorf("Error parsing version: %s", err)
		}
		if v.GreaterThan(v2) || v.Equal(v2) {
			for _, m := range n.coreModules {
				coreModules = append(coreModules, m)
			}
		}
	}

	return coreModules, nil
}
