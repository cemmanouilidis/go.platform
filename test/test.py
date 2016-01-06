#!/usr/bin/env python3

import unittest
import subprocess


class TestLinuxDistribution(unittest.TestCase):
    def testUbuntu1404(self):
        _, output = subprocess.getstatusoutput(
            "docker run --rm -it -v `pwd`:/pkg ubuntu:14.04 /pkg/platform"
        )
        self.assertEqual("ubuntu 14.04 trusty", output)

    def testUbuntu1204(self):
        _, output = subprocess.getstatusoutput(
            "docker run --rm -it -v `pwd`:/pkg ubuntu:12.04 /pkg/platform"
        )
        self.assertEqual("ubuntu 12.04 precise", output)

    def testDebianJessie(self):
        _, output = subprocess.getstatusoutput(
            "docker run --rm -it -v `pwd`:/pkg debian:jessie /pkg/platform"
        )
        self.assertEqual("debian 8.2 ", output)


if __name__ == "__main__":
    unittest.main(verbosity=2)
