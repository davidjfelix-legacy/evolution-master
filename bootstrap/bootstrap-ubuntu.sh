#!/usr/bin/env bash
# Plan on migrating this to system python (2.7?)

# The version of Python to use
version="3.5.0"

# Get packages required for bootstrapping (this currently includes building python, unfortunately)
sudo apt-get update
sudo apt-get -y install git \
  build-essential \
  libncursesw5-dev \
  libreadline-gplv2-dev \
  libssl-dev \
  libgdbm-dev \
  libc6-dev \
  libsqlite3-dev \
  tk-dev \
  libbz2-dev

# Create the directory that all this will live in
sudo mkdir -p /opt/evolution-master

# Make a temporary directory in /tmp/ and download pythong source
dir=$(mktemp -d "evolution-master.python.XXXXXXXXXX" --tmpdir)
wget https://www.python.org/ftp/python/$version/Python-$version.tar.xz -O $dir/python.tar.xz

# Unxz & untar python Source
cd $dir
tar xvf $dir/Python-$version.tar.xz
cd $dir/Python-$version

# Configure and build python in /tmp then deploy it to /opt
$dir/Python-$version/configure --prefix=/opt/evolution-master/python
make
sudo make install
