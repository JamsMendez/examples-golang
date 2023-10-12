package gofs

import (
	"io/fs"
	"os"
	"path/filepath"
)

func Find(path string) []string {
	fsys := os.DirFS(path)

	files := []string{}

	_ = fs.WalkDir(fsys, ".", func(p string, _ fs.DirEntry, err error) error {
		if err != nil {
			return fs.SkipDir
		}

		if filepath.Ext(p) == ".txt" {
			files = append(files, p)
		}

		return nil
	})

	return files
}

func Files(fsys fs.FS) (paths []string) {
	_ = fs.WalkDir(fsys, ".", func(path string, _ fs.DirEntry, _ error) error {
		if filepath.Ext(path) == ".txt" {
			paths = append(paths, path)
		}

		return nil
	})

	return paths
}
