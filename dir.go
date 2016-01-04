package platform

import (
	"os"
	"path/filepath"
)

func dirRead(path string) (FileList, error) {
	files := make([]string, 0)
	scan := func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// ignore dirs
		if info.IsDir() {
			return nil
		}

		files = append(files, path)

		return nil
	}

	err := filepath.Walk(path, scan)

	if err != nil {
		return nil, err
	}

	return FileList(files), nil
}
