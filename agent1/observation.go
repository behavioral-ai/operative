package agent1

import (
	"fmt"
	"github.com/behavioral-ai/core/messaging"
	"github.com/behavioral-ai/domain/common"
	"net/http"
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

func (o observation) String() string {
	return fmt.Sprintf("latency: %v gradient: %v", o.Latency(), o.Gradient())
}

func getObservation(agent *service, msg *messaging.Message) (o observation, status *messaging.Status) {
	if msg.Body == nil {
		status = messaging.NewStatusError(http.StatusNoContent, nil, agent.Uri())
	} else {
		if p, ok := msg.Body.(observation); ok {
			return p, messaging.StatusOK()
		}
	}
	agent.resolver.Notify(status)
	return observation{}, status

}

/*
func observationTypeErrorStatus(t any) *messaging.Status {
	err := errors.New(fmt.Sprintf("error: observation type:%v is invalid for agent:%v", reflect.TypeOf(t), Name))
	return messaging.NewStatusError(messaging.StatusInvalidArgument, err, "", nil)
}


*/

/*
	if msg.ContentType() != contentTypeObservation {
		status := messaging.NewStatusError(messaging.StatusInvalidContent, errors.New("error: observation not found"), "", agent)
		agent.notify(status)
		return observation{}, status
	}
//status := observationTypeErrorStatus(msg.Body)
	//agent.notify(status)
*/
