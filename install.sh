#!/bin/sh

# Setup PKG_PATH
export PKG_PATH=ftp://ftp.openbsd.org/pub/OpenBSD/5.9/packages/amd64

# Install essentials
pkg_add rsync-3.1.2 git curl pftop

# Fetch this repository
mkdir src
cd src
git clone git://github.com/abedra/routinator
cd ~

# Install
rm ~/.profile
cp ~/src/routinator/home/.profile ~/.profile

# Fetch the OpenBSD Sources
cd /usr/src
curl -O ftp://ftp.openbsd.org/pub/OpenBSD/5.9/src.tar.gz
curl -O ftp://ftp.openbsd.org/pub/OpenBSD/5.9/sys.tar.gz

# Extract Sources
tar xzf src.tar.gz
tar xzf sys.tar.gz

# Cleanup
rm src.tar.gz
rm sys.tar.gz

cd ~
rm install.sh

# Finish
. ./.profile

cd src/routinator
make
./configurator
cp etc/* /etc

echo "Installation and setup complete. Reboot for all changes to take effect."

