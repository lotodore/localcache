//go:build !windows

package localcache

import "os"

func Readlink(name string) (string, error) {
	return os.Readlink(name)
}

func Symlink(oldname, newname string) error {
	return os.Symlink(oldname, newname)
}
