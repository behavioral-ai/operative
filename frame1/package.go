package frame1

import (
	"github.com/behavioral-ai/core/messaging"
	"github.com/behavioral-ai/domain/collective"
)

const (
	PkgPath = "github/behavioral-ai/operative/frame1"
)

const (
	version             = 1
	ResiliencyThreshold = "resiliency:thing/operative/frame/threshold"

	ResiliencyGradient   = "resiliency:thing/operative/agent/gradient/map"
	ResiliencySaturation = "resiliency:thing/operative/agent/saturation/map"
	ResiliencyAspect     = "resiliency:thing/operative/agent/aspect/map"

	ResiliencyAspectLow    = "resiliency:aspect/operative/agent/map#low"
	ResiliencyAspectMedium = "resiliency:aspect/operative/agent/map#medium"
	ResiliencyAspectHigh   = "resiliency:aspect/operative/agent/map#high"
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
			//return nil
		},
	}
}()
