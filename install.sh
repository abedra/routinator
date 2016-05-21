#!/bin/sh

# Setup PKG_PATH
export PKG_PATH=http://mirrors.gigenet.com/pub/OpenBSD/5.9/packages/amd64

# Install essentials
echo "Installing necessary packages"
pkg_add rsync-3.1.2 git curl pftop > /dev/null 2>&1

# Fetch the OpenBSD Sources
echo "Downloading src and sys"
cd /usr/src
curl -O -s http://mirrors.gigenet.com/pub/OpenBSD/5.9/src.tar.gz
curl -O -s http://mirrors.gigenet.com/OpenBSD/5.9/sys.tar.gz

# Extract Sources

tar xzf src.tar.gz > /dev/null 2>&1
tar xzf sys.tar.gz > /dev/null 2>&1

# Cleanup
rm src.tar.gz
rm sys.tar.gz

cd ~

echo "Downloading routinator program"
mkdir setup
cd setup

curl -O -L -s https://github.com/abedra/routinator/releases/download/0.0.1/routinator_openbsd_amd64
curl -O -L -s https://github.com/abedra/routinator/releases/download/0.0.1/firewall.example.json
curl -O -L -s https://github.com/abedra/routinator/releases/download/0.0.1/templates.tar.gz

chmod +x routinator_openbsd_amd64

tar xzf templates.tar.gz > /dev/null 2>&1

cd ~

echo "Cleaning up"
rm install.sh

echo "Initial setup complete. Run setup/routinator to generate and install configs"

