#!/bin/sh

# Setup a path to fetch packages
export PKG_PATH=ftp://ftp.openbsd.org/pub/OpenBSD/5.5/packages/i386

# Install essentials
pkg_add git curl pftop mutt

# Fetch this repository
git clone git://github.com/abedra/routinator

# Install
rm ~/.profile
cp routinator/home/.profile ~
ln -sf routinator/home/script ~/script

# Fetch the OpenBSD Sources
cd /usr/src
curl -O ftp://ftp.openbsd.org/pub/OpenBSD/5.5/src.tar.gz
curl -O ftp://ftp.openbsd.org/pub/OpenBSD/5.5/sys.tar.gz

# Extract Sources
tar xzf src.tar.gz
tar xzf sys.tar.gz
cd ~

echo "Setup complete. Make sure to run script/update to update your source tree."

