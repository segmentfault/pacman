module github.com/segmentfault/pacman/contrib/cache/redis

go 1.18

require (
	github.com/go-redis/redis/v8 v8.11.5
	github.com/stretchr/testify v1.8.0
	github.com/segmentfault/pacman v1.0.0
)

require (
	github.com/cespare/xxhash/v2 v2.1.2 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/segmentfault/pacman => ../../../
