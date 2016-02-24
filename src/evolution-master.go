package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
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
	// Set an environment variable so it can remove itself later
	/*file, err := os.OpenFile("/etc/environment", os.O_APPEND | os.O_CREATE | os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalf("Could not open /etc/environment: %s; try running as root", err.Error())
	}
	_, err = file.WriteString("EVOLUTION_MASTER_PATH=" + ctx.String("root-dir"))
	if err != nil {
		defer file.Close()
		log.Fatalf("Could not write to /etc/environment: %s", err.Error())
	}
	file.Close()
	*/
	// Build the opt directory structure
	// Clone
}

func splice(ctx *cli.Context) {
	// Get the Brood file path from the first command argument
	if len(ctx.Args()) != 1 {
		log.Fatal("Could not parse brood path. Form is: evolution-master splice [PATH | URL]")
	}
	broodPath := ctx.Args().First()

	var config *Config
	var err error

	// Load the config based on whether it's a path or a URL
	if strings.HasPrefix(broodPath, "https://") || strings.HasPrefix(broodPath, "http://") {
		config, err = loadWebConfig(broodPath, ctx.String("http_proxy"))
	} else {
		config, err = loadFileConfig(broodPath)
	}

	if err != nil {
		log.Fatalf("Could not get brood file: %s", err.Error())
	}

	// Walk through the config and work with the genepools
	for _, genepool := range config.GenepoolConfigs {
		// FIXME: handle error
		repoURL, _ := url.Parse(genepool.GitRepositoryURL)
		ctx.String("root-dir") + "/" + repoURL.Host + repoURL.Path
		// Walk through each gene and build its tree
	}
}

// Get the top level brood file
// for each genepool directive
// // get the genepool if it is not gotten
// // check out the commit/whatever
// // place the genes in a directory if that does not exist
// // get the brood file for the genes
// // recursively get the genepools if not gotten
// // recursively check out the commit/whatever
// // recursively place the genes in a directory if that does not exist
// // run the genes

func fetchGenepoolAndGetDependencies() ([]*GenepoolConfig, error) {
	return nil, nil
}

func setUpTopLevelGenes() {
	// Place top level genes into a runable dir
}

func runTopLevelGenes(config *Config) {
	// Run the top level genes
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
	// Utilize a proxy for the http client if there is one
	var client *http.Client
	if proxyRawURL != "" {
		proxyURL, err := url.Parse(proxyRawURL)
		if err != nil {
			log.Printf("Failed to parse proxy URL: %s", err.Error())
			return nil, err
		}
		transport := &http.Transport{Proxy: http.ProxyURL(proxyURL)}
		client = &http.Client{Transport: transport}

	} else {
		client = &http.Client{}
	}

	// Fetch the file at the given URL
	resp, err := client.Get(fileRawURL)
	if err != nil {
		log.Printf("Failed to get URL: %s", err.Error())
		return nil, err
	}
	if resp.Status != "200" {
		return nil, fmt.Errorf("could not fetch URL")
	}

	// Parse the file to get config values
	config, err := Parse(resp.Body)
	if err != nil {
		log.Printf("Failed to parse body: %s", err.Error())
		return nil, err
	}

	return config, nil
}

func autoreclaim(ctx *cli.Context) {
	rootDir := ctx.String("root-dir")
	err := os.RemoveAll(rootDir)
	if err != nil {
		log.Fatalf("Unable to autoreclaim %s", rootDir)
	}
}

func runEvolutionMaster(ctx *cli.Context) {
	/* Scratchpad
	rm -rf /opt/evolution-master
	mkdir -p /opt/evolution-master/
	mkdir /opt/evolution-master/genepools
	mkdir /opt/evolution-master/sequences
	chown -r splicer:splicer /opt/evolution-master
	chmod -r 774 /opt/evolution-master

	cd /opt/evolution-master/genepools
	-- for genepool in genepools
		git clone genepool.url genepool.name
		cd genepool.name
		git checkout genepool.commit
		cd genes
		-- for gene in genepool.genes
			cp -R gene.name /opt/evolution-master/sequences/genepool.name/genes/gene.name
		cd /opt/evolution-master/sequences/genepool.name/genes
		-- add all these genes to a queue
		-- while queue not empty
			-- pop gene from front of queue
			-- check for broodfile
			-- maybe fetch and maybe checkout git repos for genepools in broodfile
			-- move genes from repo
			-- push genes to back of list
		-- apply genes
	 */
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
