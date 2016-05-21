#!/bin/sh

# Setup PKG_PATH
export PKG_PATH=http://mirrors.gigenet.com/pub/OpenBSD/5.9/packages/amd64

# Install essentials
pkg_add rsync-3.1.2 git curl pftop

# Fetch the OpenBSD Sources
cd /usr/src
curl -O -s http://mirrors.gigenet.com/pub/OpenBSD/5.9/src.tar.gz
curl -O -s http://mirrors.gigenet.com/OpenBSD/5.9/sys.tar.gz

# Extract Sources
tar xzf src.tar.gz
tar xzf sys.tar.gz

# Cleanup
rm src.tar.gz
rm sys.tar.gz

cd ~

mkdir setup
cd setup

curl -O https://github.com/abedra/routinator/releases/download/0.0.1/routinator
curl -O https://github.com/abedra/routinator/releases/download/0.0.1/firewall.example.json
curl -O https://github.com/abedra/routinator/releases/download/0.0.1/templates.tar.gz

tar xvzf templates.tar.gz

cd ~
rm install.sh

echo "Initial setup complete. Run setup/routinator to generate and install configs"

