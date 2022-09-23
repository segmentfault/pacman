package viper

import (
	"github.com/segmentfault/pacman/conf"
	"github.com/spf13/viper"
)

var _ conf.StaticConfig = (*file)(nil)

type file struct {
	path string
}

// NewStaticConfigParser create config parser
func NewStaticConfigParser(path string) *file {
	return &file{path: path}
}

func (f *file) LoadAndSet(conf interface{}) error {
	configVip := viper.New()
	configVip.SetConfigFile(f.path)

	// load config
	if err := configVip.ReadInConfig(); err != nil {
		return err
	}

	// set config to conf
	if err := configVip.Unmarshal(conf); err != nil {
		return err
	}
	return nil
}
