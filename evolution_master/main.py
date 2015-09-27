from __future__ import print_function
import sys
from pygit2 import clone_repository
import yaml
from hashlib import sha1
from os.path import expanduser
from shutil import rmtree

repo_types = ['git', 'hg']

def xin(key, keys, dic):
    """Tests if exclusively one key of keys is contained in dic"""
    return (key in keys) and (not any(xkey in dic for xkey in set(keys) - set(key)))

def main(config_file):
    rmtree(expanduser("~") + "/.evolution-master/")
    # Parse YAML
    config = yaml.load(config_file)
    #TODO: make this parallel
    for genepool in config.get('genepools', []):
        if xin('git', repo_types, genepool):
            repo = genepool.get('git')
            repo_folder = expanduser("~") + "/.evolution-master/" + sha1(repo).hashdigest()
            clone_repository(repo, repo_folder)
    ## Fetch all Git repos
    ## Check that specified modules exist within git repos
    ## Import modules in order specified within broodfile
    for module in modules:
        try:
            __import__(module)
            
    rmtree(expanduser("~") + "/.evolution-master/")
    
    
if __name__ == "__main__":
    main()
