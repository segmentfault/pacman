module github.com/segmentfault/pacman/contrib/log/zap

go 1.18

require (
	github.com/lestrrat-go/file-rotatelogs v2.4.0+incompatible
	github.com/segmentfault/pacman v1.0.0
	go.uber.org/zap v1.21.0
)

require (
	github.com/jonboulle/clockwork v0.3.0 // indirect
	github.com/lestrrat-go/strftime v1.0.6 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/stretchr/testify v1.8.0 // indirect
	go.uber.org/atomic v1.9.0 // indirect
	go.uber.org/multierr v1.8.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)

replace github.com/segmentfault/pacman => ../../../
