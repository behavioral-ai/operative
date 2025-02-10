package knowledge1

import (
	"errors"
	"fmt"
	"github.com/behavioral-ai/core/aspect"
	"github.com/behavioral-ai/core/messaging"
	"github.com/behavioral-ai/domain/common"
	"math/rand"
	"reflect"
)

const (
	ContentTypeObservation = "application/observation"
)

type Observation struct {
	Origin   common.Origin `json:"origin"`
	Latency  int           `json:"latency"`  // Milliseconds for the 95th percentile
	Gradient int           `json:"gradient"` // Rate of change
	RPS      int           `json:"rps"`      // Requests per second
}

func GetObservation(h messaging.Notifier, agentId string, msg *messaging.Message) (Observation, *aspect.Status) {
	if !msg.IsContentType(ContentTypeObservation) {
		return Observation{}, aspect.StatusNotFound()
	}
	if p, ok := msg.Body.(Observation); ok {
		return p, aspect.StatusOK()
	}
	status := observationTypeErrorStatus(agentId, msg.Body)
	h.Notify(status)
	return Observation{}, status
}

func observationTypeErrorStatus(agentId string, t any) *aspect.Status {
	err := errors.New(fmt.Sprintf("error: observation type:%v is invalid for agent:%v", reflect.TypeOf(t), agentId))
	return aspect.NewStatusError(aspect.StatusInvalidArgument, err)
}

func newObservation() Observation {
	var o Observation

	minN := 10
	maxN := 3500
	o.Latency = rand.Intn(maxN-minN+1) + minN

	minN = 0
	maxN = 100
	o.Gradient = rand.Intn(maxN-minN+1) + minN
	return o

}
