package main

import (
	"fmt"
	"flag"
	"github.com/cloudfoundry/cli/plugin"
	"github.com/ArthurHlt/plugin-cf-manifest-generator/mg"
	"os"
)

/*
*	This is the struct implementing the interface defined by the core CLI. It can
*	be found at  "github.com/cloudfoundry/cli/plugin/plugin.go"
*
 */
type BasicPlugin struct{}

func (c *BasicPlugin) Run(cliConnection plugin.CliConnection, args []string) {
	// Ensure that we called the command basic-plugin-command
	if args[0] != "manifest-generator" {
		return
	}
	flagSet := flag.NewFlagSet("echo", flag.ExitOnError)
	manifestPath := flagSet.String("name", mg.DEFAULT_MANIFEST, "set name of the manifest file (can be a path)")
	flagSet.StringVar(manifestPath, "n", mg.DEFAULT_MANIFEST, "set name of the manifest file (can be a path)")
	err := flagSet.Parse(args[1:])
	checkError(err)
	generator := mg.NewGenerator(*manifestPath, cliConnection)
	err = generator.Generate()
	checkError(err)
	fmt.Println("Manifest generated in " + generator.ManifestPath)

}

func (c *BasicPlugin) GetMetadata() plugin.PluginMetadata {
	return plugin.PluginMetadata{
		Name: "manifest-generator",
		Version: plugin.VersionType{
			Major: 1,
			Minor: 0,
			Build: 0,
		},
		MinCliVersion: plugin.VersionType{
			Major: 6,
			Minor: 7,
			Build: 0,
		},
		Commands: []plugin.Command{
			plugin.Command{
				Name:     "manifest-generator",
				HelpText: "Help you to generate a manifest from 0",

				// UsageDetails is optional
				// It is used to show help of usage of each command
				UsageDetails: plugin.Usage{
					Usage: "manifest-generator\n   manifest-generator -n <name of manifest file>\n   manifest-generator --name=<name of manifest file>",
				},
			},
		},
	}
}
func checkError(err error) {
	if err != nil {
		fmt.Println(fmt.Sprintf("error: %v", err))
		os.Exit(1)
	}
}
func main() {
	plugin.Start(new(BasicPlugin))
}
