# go.platform

go.platform intends to provide same functionality as python's platform package for gophers.

https://docs.python.org/3.5/library/platform.html

go.platform is a very young project, thus only a single function is implemented at the moment.
Stay tuned for updates.

# platform.LinuxDistribution() (dist, version, id, err)

Tries to determine the name, version and id of the Linux OS distribution
Supported dists = Arch, Debian, Ubuntu

    import "github.com/cemmanouilidis/go.plaform"
    dist, version, id, err := platform.LinuxDistribution()
    if dist == "arch" {
        // we are running on arch linux 
    }
    if dist == "ubuntu" {
        // we are running on ubuntu
        if version == "14.04" {
            // this is ubuntu 14.04
        }
    }
    if dist == "debian" {
        // we are running on debian
    }
