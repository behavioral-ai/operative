package agent1

import (
	"errors"
	"fmt"
	"github.com/behavioral-ai/core/messaging"
	"github.com/behavioral-ai/domain/common"
	"reflect"
)

const (
	contentTypeObservation = "application/observation"
)

type observation struct {
	origin   common.Origin `json:"origin"`
	latency  int           `json:"latency"`  // Milliseconds for the 95th percentile
	gradient int           `json:"gradient"` // Rate of change
}

func (o observation) Gradient() int {
	return o.gradient
}

func (o observation) Latency() int {
	return o.latency
}

func getObservation(agent *service, msg *messaging.Message) (observation, *messaging.Status) {
	if msg.ContentType() != contentTypeObservation {
		return observation{}, agent.notify(messaging.NewStatusError(messaging.StatusInvalidContent, errors.New("error: observation not found")))
	}
	if p, ok := msg.Body.(observation); ok {
		return p, messaging.StatusOK()
	}
	return observation{}, agent.notify(observationTypeErrorStatus(msg.Body))
}

func observationTypeErrorStatus(t any) *messaging.Status {
	err := errors.New(fmt.Sprintf("error: observation type:%v is invalid for agent:%v", reflect.TypeOf(t), Name))
	return messaging.NewStatusError(messaging.StatusInvalidArgument, err)
}
