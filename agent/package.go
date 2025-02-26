package agent

import (
	"github.com/behavioral-ai/core/messaging"
	"github.com/behavioral-ai/domain/common"
	"github.com/behavioral-ai/operative/agent1"
)

func New(handler messaging.OpsAgent, origin common.Origin, dispatcher messaging.Dispatcher) messaging.Agent {
	return agent1.New(handler, origin, dispatcher)
}
