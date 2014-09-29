# Routinator

This is the common set of configurations and scripts that I use for
running my OpenBSD systems.  The most common use I have for OpenBSD is
running it as a router/firewall. See the [[Router/Firewall]] section
below for details.

### What you get

After you run the installer, you will have a nice basic setup that is
ready for use.  It will install the source tree and put it into the
proper place.  There are scripts for updating the source tree, and
also for rebuilding the kernel and the userland binaries.  There is a
very basic pf configuration in the etc directory that will help get
you going.

### Getting Started

Once you have OpenBSD installed, run this
[bootstrap script](https://github.com/abedra/routinator/raw/master/install.sh)
to get you going.

```sh
$ ftp https://raw.githubusercontent.com/abedra/routinator/master/install.sh
$ sh install.sh
```

This will take some time depending on your internet connection.  After
it is done you will have the following.

* The routinator repository downloaded and linked in to root's home folder.
* The OpenBSD source tree in place and ready to be updated.

### Usage

##### Updating the sources

```sh
$ script/update
```

##### Recompiling the kernel

```sh
$ script/kernel
$ reboot
```

##### Updating the system binaries

```sh
$ script/userland
```
