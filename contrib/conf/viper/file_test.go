package viper

import (
	"fmt"
	"testing"
)

type AllConf struct {
	Base struct {
		WebPort string `json:"web_port" yaml:"web_port"`
	} `json:"base" yaml:"base"`
	Logger struct {
		Level string
	}
	Database []struct {
		Connection   string
		MaxIdleConns int
	}
}

func TestNewStaticConfigParser(t *testing.T) {
	configParser := NewStaticConfigParser("./testdata/config.yaml")
	c := &AllConf{}
	err := configParser.LoadAndSet(c)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(c)
	if c.Base.WebPort != "8080" || c.Logger.Level != "debug" {
		t.Error(fmt.Errorf("load config failed, config is %+v", c))
	}
}
