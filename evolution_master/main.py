from __future__ import print_function
import sys
import git2
import yaml

repos = set(['git', 'hg'])

def main(config_file):
    # Parse YAML
    config = yaml.load(config_file)
    for genepool in config.get('genepools', []):
        if 'git' in genepool and not any(s in genepool for s in repos - set('git')):
            genepool.get('git')
    ## Fetch all Git repos
    ## Check that specified modules exist within git repos
    ## Import modules in order specified within broodfile
    for module in modules:
        try:
            __import__(module)
    
    
if __name__ == "__main__":
    main()
