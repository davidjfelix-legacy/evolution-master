from __future__ import print_function
from pygit2 import clone_repository
import yaml
from hashlib import sha1
from os.path import expanduser
from shutil import rmtree

repo_types = ['git', 'hg']


def xin(key, keys, dic):
    """Tests if exclusively one key of keys is contained in dic"""
    return (key in keys) and (
    not any(xkey in dic for xkey in set(keys) - set(key)))


def parse_config_file(config_file):
    with open(config_file, 'r') as f:
        return f.readlines()


def get_genepool_listing(config):
    config = yaml.load(config)
    return config.get('genepools', [])


def fetch_genepool_repo(repo, repo_folder):
    clone_repository(repo, repo_folder)


def main(config_file):
    rmtree(expanduser("~") + "/.evolution-master/")
    config = parse_config_file(config_file)
    for genepool in get_genepool_listing(config):
        if xin('git', repo_types, genepool):
            repo = genepool.get('git')
            repo_folder = expanduser("~") + "/.evolution-master/" + sha1(
                repo).hashdigest()
            clone_repository(repo, repo_folder)
    ## Fetch all Git repos
    ## Check that specified modules exist within git repos
    ## Import modules in order specified within broodfile
    modules = []

    for module in modules:
        try:
            __import__(module)
        except:
            pass

    rmtree(expanduser("~") + "/.evolution-master/")


if __name__ == "__main__":
    main()
