package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"reflect"

	"maps"

	"github.com/bloodmage-software/docker-craft/gh"
	"github.com/bloodmage-software/docker-craft/link"
	"github.com/bloodmage-software/docker-craft/metadata"
	_ "github.com/bloodmage-software/docker-craft/metadata"
	"github.com/bloodmage-software/docker-craft/ttl"
	composetypes "github.com/compose-spec/compose-go/v2/types"
	"github.com/goccy/go-yaml"
	"github.com/yuin/gluamapper"
	lua "github.com/yuin/gopher-lua"
)

var motd = "docker-craft " + metadata.Version + "\n\n" +
	`Generate: docker craft docker-compose.lua
Or: docker-craft docker-compose.lua

Install as Docker CLI extension: docker-craft link

Initialize template: docker craft init
Or: docker-craft init
`

func main() {
	switch len(os.Args) {
	case 0, 1:
		fmt.Print(motd)
		os.Exit(0)
	case 2:
		switch os.Args[1] {
		case "link":
			link.Link()
			return
		case "init":
			// no args left
			cmdInit(nil)
			return
		default:
			fmt.Print(motd)
			os.Exit(1)
		}
	default:
		switch os.Args[1] {
		case "link":
			link.Link()
			return
		case "init":
			cmdInit(os.Args[2:])
		case "craft":
			if len(os.Args) > 2 && os.Args[2] == "init" {
				cmdInit(os.Args[3:])
				return
			}
			cmdMain(os.Args[2:])
			return
		default:
			cmdMain(os.Args[1:])
		}
	}
}

func cmdInit(args []string) {
	out := flag.String("o", "docker-compose.lua", "Output file")
	if err := flag.CommandLine.Parse(args); err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing flags: %s\n", err)
		os.Exit(1)
	}

	fmt.Println("initializing docker-craft template " + *out)
	f, err := os.Create(*out)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating file: %s\n", err)
		os.Exit(1)
		return
	}
	defer f.Close()
	configLuaTypeSymbol := ttl.TypeToLua(reflect.TypeOf(composetypes.Config{}))
	fmt.Fprintln(f, ttl.String())
	fmt.Fprintf(f, "---@type %s\nlocal config = {\n}\n\nreturn config\n", configLuaTypeSymbol)
	return
}

func cmdMain(args []string) {
	var (
		out    = flag.String("o", "docker-compose.yaml", "Output file")
		indent = flag.Uint("i", 2, "Indent")
	)
	if err := flag.CommandLine.Parse(args); err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing flags: %s\n", err)
		os.Exit(1)
	}

	files := flag.Args()
	if len(files) == 0 {
		fmt.Fprintf(os.Stderr, "No files specified\n")
		os.Exit(1)
		return
	}

	config := composetypes.Config{
		Filename: *out,
	}

	l := lua.NewState()
	if l == nil {
		os.Exit(1)
		return
	}
	defer l.Close()

	m := gluamapper.NewMapper(gluamapper.Option{
		NameFunc:    func(s string) string { return s },
		ErrorUnused: true,
		TagName:     "yaml",
	})

	for _, file := range files {
		if tbl, err := doFile(l, file); err != nil {
			fmt.Fprintf(os.Stderr, "Error executing file '%s': %s\n", file, err)
			os.Exit(1)
			return
		} else {
			c := composetypes.Config{}
			if err := m.Map(tbl, &config); err != nil {
				fmt.Fprintf(os.Stderr, "Error mapping table: %s\n", err)
				os.Exit(1)
				return
			}
			config = mergeConfigs(config, c)
		}
	}

	_ = os.MkdirAll(filepath.Dir(*out), 0755)
	f, err := os.Create(*out)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating file: %s\n", err)
		os.Exit(1)
		return
	}
	defer f.Close()
	fmt.Fprintf(f, "# yaml-language-server: $schema=https://raw.githubusercontent.com/compose-spec/compose-spec/master/schema/compose-spec.json\n---\n")

	ye := yaml.NewEncoder(f, yaml.Indent(int(*indent)))
	defer ye.Close()
	ye.Encode(config)

	gh.ActionOutput("docker-compose-file", *out)
}

func mergeMaps[K comparable, V any](a, b map[K]V) map[K]V {
	out := map[K]V{}
	maps.Copy(out, a)
	maps.Copy(out, b)
	return out
}

func mergeConfigs(a, b composetypes.Config) composetypes.Config {
	config := composetypes.Config{
		Name:       a.Name,
		Services:   mergeMaps(a.Services, b.Services),
		Networks:   mergeMaps(a.Networks, b.Networks),
		Volumes:    mergeMaps(a.Volumes, b.Volumes),
		Secrets:    mergeMaps(a.Secrets, b.Secrets),
		Configs:    mergeMaps(a.Configs, b.Configs),
		Extensions: mergeMaps(a.Extensions, b.Extensions),
		Include:    make([]composetypes.IncludeConfig, 0, len(a.Include)+len(b.Include)),
	}
	if len(b.Name) > 0 {
		config.Name = b.Name
	}
	config.Include = append(config.Include, a.Include...)
	config.Include = append(config.Include, b.Include...)
	return config
}

func doFile(L *lua.LState, file string) (*lua.LTable, error) {
	if err := L.DoFile(file); err != nil {
		fmt.Fprintf(os.Stderr, "Error executing file: %s\n", err)
		return nil, err
	}
	lv := L.Get(-1)
	if tbl, ok := lv.(*lua.LTable); ok {
		return tbl, nil
	} else {
		return nil, fmt.Errorf("expected table, got %s", lv.Type().String())
	}
}
