package timeseries1

import "github.com/behavioral-ai/domain/common"

type Entry struct {
	Origin   common.Origin `json:"origin"`
	Latency  int           `json:"latency"` // Milliseconds for the 95th percentile
	Gradient int           `json:"gradient"`
}

/*

func GetObservation(h messaging.Notifier, agentId string, msg *messaging.Message) (Entry, *core.Status) {
	if !msg.IsContentType(ContentTypeObservation) {
		return Entry{}, core.StatusNotFound()
	}
	if p, ok := msg.Body.(Entry); ok {
		return p, core.StatusOK()
	}
	status := observationTypeErrorStatus(agentId, msg.Body)
	h.Notify(status)
	return Entry{}, status
}

func observationTypeErrorStatus(agentId string, t any) *core.Status {
	err := errors.New(fmt.Sprintf("error: observation type:%v is invalid for agent:%v", reflect.TypeOf(t), agentId))
	return core.NewStatusError(core.StatusInvalidArgument, err)
}


*/
