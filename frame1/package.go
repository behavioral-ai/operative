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
			if err != nil {
				handler.Notify(err)
				return
			}
			i, err1 := newInterpret(urn.ResiliencyInterpret, resolver)
			if err1 != nil {
				handler.Notify(err1)
				return
			}
			activity, result := reason(o, t, i)
			err2 := resolver.Append(urn.ResiliencyActivity, activity, version)
			if err2 != nil {
				handler.Notify(err2)
				return
			}
			// Do we want to trace on error??
			handler.Trace(handler, messaging.MasterChannel, messaging.ObservationEvent, result)
		},
	}
}()
