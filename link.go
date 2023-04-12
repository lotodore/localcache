//go:build !windows

package localcache

import (
	"os"
)

// Simply use native symlinks everywhere but on windows.

func Readlink(name string) (string, error) {
	return os.Readlink(name)
}

func Symlink(oldname, newname string) error {
	return os.Symlink(oldname, newname)
}

func Openlink(name string) (*os.File, error) {
	return os.Open(name)
}

func ReadlinkFile(name string) ([]byte, error) {
	return os.ReadFile(name)
}
