package timeseries1

import (
	"github.com/behavioral-ai/core/messaging"
	"github.com/behavioral-ai/domain/common"
	"time"
)

const (
	PkgPath            = "github/behavioral-ai/operative/timeseries1"
	timeseriesDuration = time.Second * 2
)

// Observation - observation interface
type Observation struct {
	Timeseries func(origin common.Origin) (Entry, *messaging.Status)
}

var Observe = func() *Observation {
	return &Observation{
		Timeseries: func(origin common.Origin) (Entry, *messaging.Status) {
			//ctx, cancel := context.WithTimeout(context.Background(), timeseriesDuration)
			//defer cancel()
			//e, status := timeseries.Query(ctx, origin)
			//if !status.OK() && !status.NotFound() {
			//	h.Notify(status)
			//}
			return Entry{Gradient: 100, Latency: 55}, messaging.StatusOK()
		},
	}
}()
