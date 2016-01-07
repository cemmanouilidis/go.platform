package platform

import (
	"fmt"
	"testing"

	platform "."
)

func TestLinuxDistribution(t *testing.T) {
	data := [][]string{
		//test_data/linux-XX folder, exp. dist, exp. version, exp. id
		{"arch", "arch", "", ""},
		{"centos-7", "centos", "7", "Core"},
		{"debian-jessie", "debian", "8.2", ""},
		{"fedora-23", "fedora", "23", "Twenty Three"},
		{"ubuntu-14.04", "ubuntu", "14.04", "trusty"},
		{"dummy", "fantasyOS", "42", ""},
		{"unix", "", "", ""},
	}

	for _, d := range data {
		t.Log("Test", d[1])

		path := fmt.Sprintf("test_data/linux_%s", d[0])
		dist, version, id, err := platform.LinuxDistribution(path)
		if err != nil {
			t.Error("Unexpected error occured ", err)
		}

		if dist != d[1] {
			t.Error("Unexpected dist found: ", dist)
		}

		if version != d[2] {
			t.Error("Unexpected version found: ", version)
		}

		if id != d[3] {
			t.Error("Unexpected id found: ", id)
		}
	}
}
