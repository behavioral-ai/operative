package frame2

import (
	"github.com/behavioral-ai/collective/content"
	"github.com/behavioral-ai/core/messaging"
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
	Reason func(o Observation, handler messaging.Agent, resolver content.Resolution)
}

var Frame = func() *IFrame {
	return &IFrame{
		Reason: func(o Observation, handler messaging.Agent, resolver content.Resolution) {
		},
	}
}()
