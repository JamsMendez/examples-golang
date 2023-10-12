package gofs

import (
	"archive/zip"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"testing"
	"testing/fstest"

	"github.com/google/go-cmp/cmp"
)

func BenchmarkFileOnDisk(b *testing.B) {
	fsys := os.DirFS("testdata")
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = Files(fsys)
	}
}

func BenchmarkInMapFS(b *testing.B) {
	fsys := fstest.MapFS{
		"file.txt":                    {},
		"sub_folder_1/sub_file_1.txt": {},
		"sub_folder_2/sub_file_1.txt": {},
		"sub_folder_2/sub_file_3.txt": {},
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = Files(fsys)
	}
}

func TestFilesInZip(t *testing.T) {
	t.Parallel()

	fsys, err := zip.OpenReader("testdata.zip")
	if err != nil {
		t.Fatal(err)
	}

	want := []string{
		"first.txt",
		"testdata/files/other.txt",
	}

	got := Files(fsys)
	if !cmp.Equal(want, got) {
		cmp.Diff(want, got)
	}
}

func TestFiles(t *testing.T) {
	t.Parallel()

	fsys := os.DirFS("testdata")

	want := []string{
		"first.txt",
		"testdata/files/other.txt",
	}

	got := Files(fsys)

	if !cmp.Equal(want, got) {
		cmp.Diff(want, got)
	}
}

func TestFilesInMapFS(t *testing.T) {
	t.Parallel()

	fsys := fstest.MapFS{
		"file.txt":                    {},
		"sub_folder_1/sub_file_1.txt": {},
		"sub_folder_2/sub_file_1.txt": {},
		"sub_folder_2/sub_file_3.txt": {},
	}

	want := []string{
		"file.txt",
		"sub_folder_1/sub_file_1.txt",
		"sub_folder_2/sub_file_1.txt",
		"sub_folder_2/sub_file_3.txt",
	}

	got := Files(fsys)
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestGetFilesV1(t *testing.T) {
	t.Parallel()

	want := []string{
		"first.txt",
		"testdata/files/other.txt",
	}

	got := Find("testdata")

	if cmp.Equal(want, got) {
		t.Fatal(cmp.Diff(want, got))
	}
}

func TestGetMatchesWalkDir(t *testing.T) {
	t.Parallel()

	size := 0
	fsys := os.DirFS("testdata")

	err := fs.WalkDir(fsys, ".", func(path string, _ fs.DirEntry, err error) error {
		if err != nil {
			fmt.Println(err)

			return fs.SkipDir
		}

		if filepath.Ext(path) == ".txt" {
			size++
		}

		return nil
	})

	if err != nil {
		t.Fatalf("expected error nil, got %v", err)
	}

	if size != 2 {
		t.Fatalf("expected size == 2, got %d", size)
	}
}

func _TestGetMatchesGlob(t *testing.T) {
	t.Parallel()

	fsys := os.DirFS("testdata")

	// no search recursive
	matches, err := fs.Glob(fsys, "*.txt")
	if err != nil {
		t.Fatal(err)
	}

	size := len(matches)
	if size != 2 {
		t.Fatalf("expected size == 2, got %d", size)
	}
}
