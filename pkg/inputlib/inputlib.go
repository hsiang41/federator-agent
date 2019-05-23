package main

import (
	logUtil "github.com/containers-ai/alameda/pkg/utils/log"
)

type inputLib struct {

}

func (i inputLib) Gather() error {
	return nil
}

func (i inputLib) LoadConfig(config string, scope *logUtil.Scope) error {
	return nil
}

var InputLib inputLib