package frame2

import (
	"github.com/behavioral-ai/core/messaging"
	"github.com/behavioral-ai/domain/collective"
)

const (
	PkgPath = "github/behavioral-ai/operative/frame2"
	version = 1
)

type Features struct {
	Action   int
	Latency  int
	Gradient int
}

type Observation interface {
	Metrics() []Features
}

// IFrame - frame interface
type IFrame struct {
	Reason func(o Observation, handler messaging.OpsAgent, resolver collective.IResolver)
}

var Frame = func() *IFrame {
	return &IFrame{
		Reason: func(o Observation, handler messaging.OpsAgent, resolver collective.IResolver) {
		},
	}
}()
