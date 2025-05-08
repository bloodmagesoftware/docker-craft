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

return compose
