package main

import (
	"github.com/hashicorp/hcl/hcl/ast"
	"path"
	"os"
	"fmt"
	"os/exec"
	"log"
)

type GenepoolConfig struct {
	Name             string
	GitRepositoryURL string   `hcl:"git"`
	Genes            []string `hcl:"genes"`
}

func ParseGenepoolConfigs(list *ast.ObjectList) ([]*GenepoolConfig, error) {
	return nil, nil
}

func (g *GenepoolConfig) FetchGenepool(rootDir string) error{
	cloneDir := path.Join(rootDir, g.Name)

	// Check if the directory already exists
	if _, err := os.Stat(cloneDir); err == nil {
		return fmt.Errorf("genepool directory already exists")
	}

	// Attempt to git clone the repo and echo its output
	// FIXME: perhaps split the stderr and stdout here
	out, err := exec.Command("git", "clone", g.GitRepositoryURL, cloneDir).CombinedOutput()
	log.Printf("%s\n", out)
	if err != nil {
		return fmt.Errorf("could not fetch genepool")
	}
	return nil
}
