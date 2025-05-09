package gh

import (
	"fmt"
	"os"
)

func ActionOutput(key, value string) {
	if githubOutput, ok := os.LookupEnv("GITHUB_OUTPUT"); ok {
		f, err := os.OpenFile(githubOutput, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			fmt.Fprintf(os.Stderr, "failed to open GITHUB_OUTPUT: %v\n", err)
			os.Exit(1)
			return
		}
		defer f.Close()
		fmt.Fprintf(f, "%s=%s\n", key, value)
	}
}
