package main

import (
	"fmt"
	cli "github.com/codegangsta/cli"
	"os"
	"os/exec"
	"strings"
)

func bootstrap(ctx *cli.Context) {
}

func splice(ctx *cli.Context) {
	if len(ctx.Args()) != 1 {
		// FIXME: log here
		return
	}

	broodPath := ctx.Args().First()
	var config string

	if strings.HasPrefix(broodPath, "https://") || strings.HasPrefix(broodPath, "http://") {
		// FIXME: recover error and log here
		config, _ = loadWebConfig(broodPath, ctx.String("http_proxy"))
	} else {
		// FIXME: recover error and log here
		config, _ = loadFileConfig(broodPath)
	}
	fmt.Printf("%s\n", config)
}

func loadFileConfig(path string) (string, error) {
	return "", nil
}

func loadWebConfig(url string, proxy string) (string, error) {
	return "", nil
}

func autoreclaim(ctx *cli.Context) {
}

func runEvolutionMaster(ctx *cli.Context) {
	gitCloneGenepool("https://github.com/hatchery/Brood2", "/opt/Brood2")
}

func gitCloneGenepool(url string, path string) error {
	err := os.RemoveAll(path)
	if err != nil {
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
			Name:   "root-dir, d",
			Value:  "/opt/evolution-master",
			Usage:  "directory that evolution-master is installed to",
			EnvVar: "EVOLUTION_MASTER_PATH",
		},
		cli.StringFlag{
			Name:   "http_proxy, p",
			Value:  "",
			Usage:  "proxy string for access to remote files",
			EnvVar: "http_proxy,HTTP_PROXY,https_proxy,HTTPS_PROXY",
		},
	}
	evo.Commands = []cli.Command{
		{
			Name:   "bootstrap",
			Usage:  "let the evolution-master install itself and interrogate you.",
			Action: bootstrap,
		},
		{
			Name:   "autoreclaim",
			Usage:  "instruct the evolution-master to reclaim its own disk space and remove itself from the system.",
			Action: autoreclaim,
		},
		{
			Name:   "splice",
			Usage:  "combine the gene instructions with the current machine state.",
			Action: splice,
		},
	}
	evo.Action = runEvolutionMaster
	evo.Run(os.Args)
}
