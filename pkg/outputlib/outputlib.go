package outputlib

import (
	logUtil "github.com/containers-ai/alameda/pkg/utils/log"
	"github.com/sheerun/queue"
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

type OutputLibrary interface {
	Write() error
	LoadConfig(config string, scope *logUtil.Scope) error
	SetAgentQueue(agentQueue *queue.Queue)
}