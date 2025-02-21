package agent

import (
	"github.com/behavioral-ai/core/messaging"
	"github.com/behavioral-ai/domain/common"
	"github.com/behavioral-ai/operative/agent1"
)

func New(origin common.Origin, notifier messaging.NotifyFunc, dispatcher messaging.Dispatcher) messaging.Agent {
	return agent1.New(origin, notifier, dispatcher)
}
