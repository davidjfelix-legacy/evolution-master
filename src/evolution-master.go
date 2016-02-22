package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
	"net/http"
	"net/url"

	"github.com/hashicorp/hcl/hcl/ast"
	"github.com/codegangsta/cli"
	"github.com/hashicorp/hcl"
)

type Config struct {
	GenepoolConfigs []*GenepoolConfig
}


func Parse(r io.Reader) (*Config, error) {
	var buf bytes.Buffer
	if _, err := io.Copy(&buf, r); err != nil {
		return nil, err
	}

	root, err := hcl.Parse(buf.String())
	if err != nil {
		return nil, err
	}
	buf.Reset()

	list, ok := root.Node.(*ast.ObjectList)
	if !ok {
		return nil, fmt.Errorf("error parsing: root should be an object")
	}

	genepools := list.Filter("genepool")
	if len(genepools.Items) == 0 {
		return nil, fmt.Errorf("no 'genepool' stanza found")
	}

	var config Config
	config.GenepoolConfigs, err = ParseGenepoolConfigs(genepools)
	if err != nil {
		return nil, err
	}

	return &config, nil
}


func bootstrap(ctx *cli.Context) {
}

func splice(ctx *cli.Context) {
	if len(ctx.Args()) != 1 {
		// FIXME: log here
		return
	}

	broodPath := ctx.Args().First()
	var config *Config

	if strings.HasPrefix(broodPath, "https://") || strings.HasPrefix(broodPath, "http://") {
		// FIXME: recover error and log here
		config, _ = loadWebConfig(broodPath, ctx.String("http_proxy"))
	} else {
		// FIXME: recover error and log here
		config, _ = loadFileConfig(broodPath)
	}
	fmt.Printf("%s\n", config)
}

func loadFileConfig(path string) (*Config, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	config, err := Parse(file)
	if err != nil {
		return nil, err
	}

	return config, nil
}

func loadWebConfig(fileRawURL string, proxyRawURL string) (*Config, error) {
	proxyURL, err := url.Parse(proxyRawURL)
	if err != nil {
		return nil, err
	}

	transport := &http.Transport{
		Proxy: http.ProxyURL(proxyURL),
	}
	client := &http.Client{
		Transport: transport,
	}

	resp, err := client.Get(fileRawURL)
	if err != nil {
		return nil, err
	}
	if resp.Status != 200 {
		return nil, fmt.Errorf("could not fetch URL")
	}

	config, err := Parse(resp.Body)
	if err != nil {
		return nil, err
	}

	return config, nil
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
