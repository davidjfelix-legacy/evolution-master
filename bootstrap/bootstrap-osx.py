#!/usr/bin/python
# In Yosemite and Mavericks, this is 2.7.10
from subprocess import Popen


def install_xcode_clt_pkg(pkg_file):
    Popen(['sudo', 'install', '-pkg', pkg_file, '-target', '/']).wait()


def mount_xcode_dmg(dmg_file):
    # FIXME: actually return real mount point... not this ghetto thing
    Popen(['hdiutil', 'attach', dmg_file]).wait()
    return "/Volumes/" + dmg_file.split("/")[-1]


def install_xcode_cloud():
    Popen(['xcode-select', '--install']).wait()


def install_brew():
    # FIXME: this won't work by default... need to actually use a shell to get the content
    Popen(['ruby', '-e', '"$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install)"']).wait()
    
