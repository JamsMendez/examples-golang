package cmd

import (
	"archive/tar"
	"compress/gzip"
	"errors"
	"io"
	"log"
	"os"
	"path/filepath"
)

func decompress(dst string, r io.Reader) error {
	gzr, err := gzip.NewReader(r)
	if err != nil {
		return err
	}

	defer func() {
		err = gzr.Close()
		if err != nil {
			log.Println("gzip close error: ", err)
		}
	}()

	tr := tar.NewReader(gzr)

	for {
		var header *tar.Header
		header, err = tr.Next()
		if err != nil {
			if errors.Is(err, io.EOF) {
				return nil
			}

			return err
		}

		if header == nil {
			continue
		}

		target := filepath.Join(dst, header.Name)

		switch header.Typeflag {

		case tar.TypeDir:
			if _, err = os.Stat(target); err != nil {
				if err = os.MkdirAll(target, 0o755); err != nil {
					return err
				}
			}

		case tar.TypeReg:
			var file *os.File
			file, err = os.OpenFile(target, os.O_CREATE|os.O_RDWR, os.FileMode(header.Mode))
			if err != nil {
				return err
			}

			if _, err = io.Copy(file, tr); err != nil {
				if errClose := file.Close(); errClose != nil {
					log.Println("decompress file copy error: ", errClose)
				}

				return err
			}

			if err = file.Close(); err != nil {
				log.Println("decompress file copy error: ", err)
			}

		}
	}
}
