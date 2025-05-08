package metadata

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/docker/cli/cli-plugins/metadata"
)

var (
	Version string = "0.0.0"
)

func init() {
	if len(os.Args) > 1 && os.Args[1] == "docker-cli-plugin-metadata" {
		metadata := metadata.Metadata{
			SchemaVersion:    "0.1.0",
			Vendor:           "Bloodmage Software",
			Version:          Version,
			ShortDescription: "A Docker Compose templating plugin",
			URL:              "https://github.com/bloodmagesoftware/docker-craft",
		}

		encoder := json.NewEncoder(os.Stdout)
		if err := encoder.Encode(metadata); err != nil {
			fmt.Fprintf(os.Stderr, "Error encoding metadata: %v\n", err)
			os.Exit(1)
		}
		os.Exit(0)
	}
}
