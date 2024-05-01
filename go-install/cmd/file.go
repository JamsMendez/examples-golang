package cmd

import (
	"fmt"
	"os"
	"path"
)

const (
	tmplName  = "go-%s.tar.gz"
	goDirName = "go"

	DstInstall = "/usr/local"
)

func getName(version string) string {
	return fmt.Sprintf(tmplName, version)
}

func getFile(version string) (*os.File, error) {
	name := getName(version)
	file, err := os.OpenFile(name, os.O_CREATE|os.O_RDWR, 0o666)
	if err != nil {
		return nil, err
	}

	return file, nil
}

func removeFile(version string) error {
	name := getName(version)
	info, err := os.Stat(name)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}

		return err
	}

	if info.IsDir() {
		return fmt.Errorf("%s is not file", name)
	}

	return os.Remove(name)
}

func removeGoDir(dst string) error {
	name := path.Join(dst, goDirName)
	info, err := os.Stat(name)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}

		return err
	}

	if !info.IsDir() {
		return fmt.Errorf("%s is not directory", name)
	}

	return os.RemoveAll(name)
}
