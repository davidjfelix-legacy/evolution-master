#!/usr/bin/python
# In Yosemite and Mavericks, this is 2.7.10
from __future__ import print_function
from getpass import getpass
from subprocess import Popen
from distutils.util import strtobool


# 2to3
try:
    input = raw_input
except NameError:
    pass


def prompt_yes_no(prompt, tries=1, retry=False, retry_forever=False):
    should_retry = retry
    remaining_tries = tries
    while should_retry:
        response = input(prompt, " [Y/n]: ")
        try:
            ret = strtobool(response)
            return ret
        except ValueError:
            print("Response not valid for Y/n: ", response)
            remaining_tries -= 1
            if remaining_tries <= 0 and not retry_forever:
                raise ValueError("Response not valid for Y/n: ", response)


def needs_proxy():
    """Check to see if the system has an http proxy"""
    # TODO: try to connect
    return prompt_yes_no("Do you have an http proxy?", tries=3, retry=True)


def get_http_proxy():
    """Get proxy information from the user"""
    port = input("Proxy Port: ")
    host = input("Proxy Host: http://")
    has_auth = prompt_yes_no("Does your proxy require authentication?", tries=3, retry=True)

    if has_auth:
        usernm, passwd = get_auth()
        return "http://{}:{}@{}:{}".format(usernm, passwd, host, port)
    else:
        return "http://{}:{}".format(host, port)


def get_auth():
    """Get a username/password pair"""
    username = input("Username: ")
    password = getpass()
    return (username, password)


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


def main():
    if needs_proxy():
        get_http_proxy()


if __name__ == '__main__':
    main()
