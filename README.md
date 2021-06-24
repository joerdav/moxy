# Moxy

A small proxy built for developing micro-frontends locally.

## Usage

To get running fast just run:

`docker run -p 80:80 joedavidson1802/moxy`

To specify your own defaults config file run:

`docker run -p 80:80 -v ./my_defaults.toml:/app/defaults.toml joedavidson1802/moxy`

## Config file defaults

``` toml
moxy_path = "/moxy" # The default path that the moxy admin portal will appear on
default_upstream = "http://localhost:8081" # The upstream traffic will default to

# A list of upstreams (order matters they will be evaluated top down)
[[upstream]]
path_prefix = "/apps/basket"
host = "http://localhost:8085"
[[upstream]]
path_prefix = "/apps/browse"
host = "http://localhost:8085"
[[upstream]]
path_prefix = "/apps/details"
host = "http://localhost:8085"
[[upstream]]
path_prefix = "/apps/random"
host = "http://localhost:8085"
```