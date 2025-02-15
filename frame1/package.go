package frame1

import (
	"github.com/behavioral-ai/core/messaging"
	"github.com/behavioral-ai/domain/collective"
	"github.com/behavioral-ai/operative/urn"
)

const (
	PkgPath = "github/behavioral-ai/operative/frame1"
	version = 1
)

type Observation interface {
	Gradient() int
	Latency() int
}

// IFrame - frame interface
type IFrame struct {
	Reason func(o Observation, handler messaging.OpsAgent, resolver collective.IResolver)
}

var Frame = func() *IFrame {
	return &IFrame{
		Reason: func(o Observation, handler messaging.OpsAgent, resolver collective.IResolver) {
			t, err := newThreshold(urn.ResiliencyThreshold, resolver)
		},
	}
}()
