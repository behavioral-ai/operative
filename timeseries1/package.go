package timeseries1

import (
	"github.com/behavioral-ai/domain/common"
	"time"
)

const (
	PkgPath            = "github/behavioral-ai/domain/timeseries1"
	timeseriesDuration = time.Second * 2
)

// Observation - observation interface
type Observation struct {
	Timeseries func(origin common.Origin) (Entry, error)
}

var Observe = func() *Observation {
	return &Observation{
		Timeseries: func(origin common.Origin) (Entry, error) {
			//ctx, cancel := context.WithTimeout(context.Background(), timeseriesDuration)
			//defer cancel()
			//e, status := timeseries.Query(ctx, origin)
			//if !status.OK() && !status.NotFound() {
			//	h.Notify(status)
			//}
			return Entry{Gradient: 100, Latency: 55}, nil
		},
	}
}()
