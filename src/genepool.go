package main

import "github.com/hashicorp/hcl/hcl/ast"

type GenepoolConfig struct {
	GitRepositoryURL string   `hcl:"git"`
	Genes            []string `hcl:"genes"`
}

func ParseGenepoolConfigs(list *ast.ObjectList) ([]*GenepoolConfig, error) {
	return nil, nil
}
