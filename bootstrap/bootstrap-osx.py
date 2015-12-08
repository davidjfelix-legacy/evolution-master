#!/usr/bin/python
# In Yosemite and Mavericks, this is 2.7.10
from getpass import getpass
from subprocess import Popen
from distutils.util import strtobool

# 2to3
try:
    input = raw_input
except NameError:
    pass

def get_http_proxy():
    port = input("Proxy Port: ")
    host = input("Proxy Host: http://")
    while True:
        resp = input("Does your proxy require authentication? [Y/n]: ")
        try:
            has_auth = strtobool(resp)
            break
        except ValueError:
            pass
    if has_auth:
        usernm, passwd = get_auth()
        return "http://{}:{}@{}:{}".format(usernm, passwd, host, port)
    else:
        return "http://{}:{}".format(host, port)

def get_auth():
    usernm = input("Username: ")
    passwd = getpass()
    return (usernm, passwd)

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
    
