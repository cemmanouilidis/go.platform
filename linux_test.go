package platform

import (
	"testing"

	platform "."
)

func TestLinuxDistributionArch(t *testing.T) {
	dist, version, id, err := platform.LinuxDistribution("test_data/linux_arch")
	if err != nil {
		t.Error("Failed to detect LinuxDistribution")
	}

	if dist != "arch" {
		t.Error("Unexpected dist found: ", dist)
	}

	if version != "" {
		t.Error("Unexpected version found: ", version)
	}

	if id != "" {
		t.Error("Unexpected id found: ", id)
	}
}

func TestLinuxDistributionUbuntu(t *testing.T) {
	dist, version, id, err := platform.LinuxDistribution("test_data/linux_ubuntu-14.04")
	if err != nil {
		t.Error("Failed to detect LinuxDistribution")
	}

	if dist != "ubuntu" {
		t.Error("Unexpected dist found: ", dist)
	}

	if version != "14.04" {
		t.Error("Unexpected version found: ", version)
	}

	if id != "trusty" {
		t.Error("Unexpected id found: ", id)
	}
}

func TestLinuxDistributionDummy(t *testing.T) {
	dist, version, id, err := platform.LinuxDistribution("test_data/linux_dummy/")
	if err != nil {
		t.Error("Unexpected error occored: ", err)
	}

	if dist != "unknown" {
		t.Error("Unexpected dist found: ", dist)
	}

	if version != "" {
		t.Error("Unexpected version found: ", version)
	}

	if id != "" {
		t.Error("Unexpected id found: ", id)
	}
}
