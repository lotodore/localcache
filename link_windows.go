package localcache

import (
	"os"
)

// On windows, use custom link files, because symlinks require special privileges.

func Readlink(name string) (string, error) {
	link, err := os.ReadFile(name)
	if err != nil {
		return "", err
	}
	return string(link), nil
}

func Symlink(oldname, newname string) error {
	return os.WriteFile(newname, []byte(oldname), 0644)
}
