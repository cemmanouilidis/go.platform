package platform

import (
	"testing"
)

func TestDirRead(f *testing.T) {
	fileList, err := dirRead("test_data/dir/")

	if err != nil {
		f.Error("Unexpected error occured: ", err)
	}

	if len(fileList) != 4 {
		f.Error("Unexpected number of files return by DirRead: ", len(fileList))
	}
}
