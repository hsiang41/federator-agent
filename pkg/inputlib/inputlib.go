package inputlib

import (
	logUtil "github.com/containers-ai/alameda/pkg/utils/log"
	"github.com/sheerun/queue"
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

type InputLibrary interface {
	Gather() error
	LoadConfig(config string, scope *logUtil.Scope) error
	SetAgentQueue(agentQueue *queue.Queue)
}