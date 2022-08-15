/* For license and copyright information please see the LEGAL file in the code repository */

package cmd

import (
	"os"
	"path/filepath"
	"strings"
)

var (
	cwd string
	// Executable returns the path name for the executable that started the current process
	ed string
)

func init() {
	var ex, err = os.Executable()
	if err != nil {
		panic("cannot determine executable directory: " + err.Error())
	}
	ed = filepath.Dir(ex)

	cwd, err = os.Getwd()
	if err != nil {
		panic("cannot determine current working directory: " + err.Error())
	}

	// TODO::: 3 way to indicate current working directory
	// fmt.Println(os.Executable())
	// fmt.Println(os.Getwd())
	// fmt.Println(os.Args[0])
}

// Cwd returns the current working directory at the time of the first call.
func Cwd() string { return cwd }

// ED or Executable Directory returns the path name for the executable that started the current process
func ED() string { return ed }

// ShortPath returns an absolute or relative name for path, whatever is shorter.
func ShortPath(path string) string {
	if rel, err := filepath.Rel(Cwd(), path); err == nil && len(rel) < len(path) {
		return rel
	}
	return path
}

// RelPaths returns a copy of paths with absolute paths
// made relative to the current directory if they would be shorter.
func RelPaths(paths []string) []string {
	var out []string
	for _, p := range paths {
		rel, err := filepath.Rel(Cwd(), p)
		if err == nil && len(rel) < len(p) {
			p = rel
		}
		out = append(out, p)
	}
	return out
}

// IsTestFile reports whether the source file is a set of tests and should therefore
// be excluded from coverage analysis.
func IsTestFile(file string) bool {
	// We don't cover tests, only the code they test.
	return strings.HasSuffix(file, "_test.go")
}
