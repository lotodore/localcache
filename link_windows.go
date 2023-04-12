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

func Openlink(name string) (*os.File, error) {
	entry, err := Readlink(name)
	if err != nil {
		return nil, err
	}
	return os.Open(entry)
}

func ReadlinkFile(name string) ([]byte, error) {
	entry, err := Readlink(name)
	if err != nil {
		return nil, err
	}
	return os.ReadFile(entry)
}
