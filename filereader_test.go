package platform

import (
	"testing"
)

func TestFileReader(t *testing.T) {
	entries := []string{}

	path := "test_data/linux_arch/etc/os-release"
	for line := range fileReader(path) {
		entries = append(entries, line)
	}

	if len(entries) != 8 {
		t.Error("Unexpected number of lines retrieved: #", len(entries))
	}
}
