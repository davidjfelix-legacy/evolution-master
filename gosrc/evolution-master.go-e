package main

import (
	"fmt"
	cli "github.com/codegangsta/cli"
	"os"
	"os/exec"
)


func runEvolutionMaster(ctx *cli.Context) {
	gitCloneGenepool("https://github.com/hatchery/Brood2", "/opt/Brood2")
}


func gitCloneGenepool(url string, path string) error {
	err := os.RemoveAll(path)
	if  err != nil {
		return err
	}
	out, _ := exec.Command("git", "clone", url, path).CombinedOutput()
	fmt.Printf("%s\n", out)
	return nil
}


func main() {
	evo := cli.NewApp()
	evo.Name = "Evolution Master"
	evo.Usage = "Provision your development machine."
	evo.Version = "0.1.0"
	evo.Authors = []cli.Author{
		{
			Name:  "David J Felix",
			Email: "felix.davidj@gmail.com",
		},
	}
	evo.Copyright = "MIT"
	evo.Flags = []cli.Flag{
		cli.StringFlag{
			Name: "root-dir, d",
			Value: "/opt/evolution-master",
			Usage: "directory that evolution-master is installed to",
			EnvVar: "EVOLUTION_MASTER_PATH",
		},
		cli.StringFlag{
			Name: "broodfile, f",
			Value: "https://raw.githubusercontent.com/hatchery/Brood2/master/Brood.hcl",
			Usage: "file that evolution-master uses to build its brood. Can be URL or local file",
			EnvVar: "EVOLUTION_MASTER_BROODFILE",
		},
		cli.StringFlag{
			Name: "http_proxy, p",
			Value: "",
			Usage: "proxy string for access to remote files",
			EnvVar: "HTTPS_PROXY,HTTP_PROXY,https_proxy,http_proxy",
		},
	}
	evo.Action = runEvolutionMaster
	evo.Run(os.Args)
}
