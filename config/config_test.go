package config

import (
	"fmt"
	"testing"
)

func TestConfig(t *testing.T) {
	_, err := ReadOrCreateConfig("./config.json")
	if err != nil {
		fmt.Printf("error:%v", err)
		return
	}
}
