module github.com/segmentfault/pacman/contrib/cache/memory

go 1.18

require (
	github.com/patrickmn/go-cache v2.1.0+incompatible
	github.com/segmentfault/pacman v1.0.0
	github.com/stretchr/testify v1.8.0
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/segmentfault/pacman => ../../../
