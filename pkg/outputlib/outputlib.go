package main

import (
	logUtil "github.com/containers-ai/alameda/pkg/utils/log"
)

type outputlib struct {

}

func (o outputlib) Write() error {
	return nil
}

func (o outputlib) LoadConfig(config string, scope *logUtil.Scope) error {
	return nil
}


var OutputLib outputlib