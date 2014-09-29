#!/bin/sh

# Setup PKG_PATH
export PKG_PATH=ftp://ftp.openbsd.org/pub/OpenBSD/5.5/packages/i386

# Install essentials
pkg_add git curl pftop

# Fetch this repository
mkdir src
cd src
git clone git://github.com/abedra/routinator
cd ~

# Install
rm ~/.profile
ln -sf ~/src/routinator/home/.profile ~/.profile

# Fetch the OpenBSD Sources
cd /usr/src
curl -O ftp://ftp.openbsd.org/pub/OpenBSD/5.5/src.tar.gz
curl -O ftp://ftp.openbsd.org/pub/OpenBSD/5.5/sys.tar.gz

# Extract Sources
tar xzf src.tar.gz
tar xzf sys.tar.gz

# Cleanup
rm src.tar.gz
rm sys.tar.gz

cd ~
rm install.sh

# Finish
source .profile
echo "Setup complete. Make sure to run script/update to update your source tree."

