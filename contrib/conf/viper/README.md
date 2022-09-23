# file config
> parsing the configuration file by [viper](https://github.com/spf13/viper)

## Usage
```go
package main

import (
	"fmt"

	"github.com/segmentfault/pacman/contrib/conf/viper"
)

func main() {
	c := &AllConfig{}
	config, err := viper.NewWithPath("/config/path")
	if err != nil {
		panic(err)
	}
	if err = config.Parse(&c); err != nil {
		panic(err)
	}
	fmt.Println(c.Server.HTTP)
}
```
