# Routinator 

This is the common set of configurations and scripts that I use for running my OpenBSD systems.  The most common use I have for OpenBSD is running it as a router/firewall.  See the [[Router/Firewall]] section below for details.

### What you get

After you run the installer, you will have a nice basic setup that is ready for use.  It will install the source tree and put it into the proper place.  There are scripts for updating the source tree, and also for rebuilding the kernel and the userland binaries.  There is a very basic pf configuration in the etc directory that will help get you going.

### Getting Started

Once you have OpenBSD installed, you just need to run this simple [[bootstrap script]] to get you going.  You can use lynx to pick it up:

    lynx https://github.com/abedra/routinator/raw/master/install.sh
    
Once you have the file downloaded just give it a whirl:

    sh install.sh
    
This will take some time depending on your internet connection.  After it is done you will have the following.

* git, curl, pftop, alpine installed.
* The routinator repository downloaded and linked into your home folder.
* The OpenBSD source tree in place and ready to be updated.

### Usage

##### Updating the sources

    script/update
    
##### Recompiling the kernel

    script/kernel
    reboot
    
##### Updating the system binaries

    script/userland
