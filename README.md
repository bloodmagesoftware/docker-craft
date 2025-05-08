# docker-craft

A Docker Compose templating plugin.

## Why?

> If I miss something here and this whole thing makes no sense, please open an issue and let me know.
> If you don't have my problems and don't need such a tool, I'm happy for you.
>
> @tsukinoko_kun

### Why not use profiles?

I don't want to specify the whole service multiple times.
This also is extreamly verbose and not very readable.

### Why not use overwrite files?

They are a pretty good solution for most use cases, but I want the whole config in one file and configurable names (auto-generating network names in CI for example).
I could generate an overwrite file in CI but that just makes the problem a bit smaller instead of solving it.

### Why not environment variables?

I want to be able to use the same config for multiple environments (dev, staging, prod) and be able to just run `docker compose up -d` to deploy. I don't want to care about the right variables being set when I do `docker compose up -d`.

### Why not use a templating language?

I want syntax highlighting and a compiler (or something similar) that can check for errors. Additionally, I don't want any scripts in my codebase that are unrelated to the project itself.

### Why Lua?

I like Lua.

## Usage

```shell
docker craft docker-compose.yaml.lua
```

generates `docker-compose.yaml` based on `docker-compose.yaml.lua`

| Flag | Description | Default value         |
| ---- | ----------- | --------------------- |
| `-o` | Output file | `docker-compose.yaml` |
| `-i` | Indentation | `2`                   |

## Example

Just return a table with the compose configuration. Everything needs to be explicitly set (no shorthands).

```lua
local compose = {
    services = {
        web = {
            image = "nginx:latest",
            ports = {
                {
                    Published = "80",
                    Target = 80,
                }
            },
        }
    },
    networks = {
        default = {
            external = true
        }
    }
}

-- use environment variables to conditionally add configuration
if os.getenv("ADD_VOLUME") then
    compose.services.web.volumes = {
        {
            source = "./html",
            target = "/var/www/html",
            type = "bind",
        }
    }
end

-- return the compose configuration
return compose
```

## Installation

Install the application via Go install:

```shell
go install github.com/bloodmagesoftware/docker-craft@latest
```

or Homebrew:

```shell
brew install bloodmagesoftware/tap/docker-craft
```

then symlink the binary to `~/.docker/cli-plugins/`: _(requires admin privileges on Windows)_

```shell
docker-craft link
```

## License

MIT
