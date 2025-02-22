package agent1

import (
	"github.com/behavioral-ai/core/messaging"
	"github.com/behavioral-ai/core/test"
	"github.com/behavioral-ai/domain/common"
	"github.com/behavioral-ai/operative/timeseries1"
	"time"
)

func ExampleEmissary() {
	ch := make(chan struct{})
	emissaryShutdown := messaging.NewMessage(messaging.EmissaryChannel, messaging.ShutdownEvent)
	dataChange := messaging.NewMessage(messaging.EmissaryChannel, messaging.DataChangeEvent)
	agent := newOp(common.Origin{Region: "us-west"}, test.Notify, messaging.NewTraceDispatcher())

	go func() {
		go emissaryAttend(agent, func() *timeseries1.Observation {
			return &timeseries1.Observation{
				Timeseries: func(origin common.Origin) (timeseries1.Entry, *messaging.Status) {
					return timeseries1.Entry{}, messaging.StatusNotFound()
				},
			}
		}())
		agent.Message(dataChange)
		time.Sleep(testDuration * 2)
		agent.Message(messaging.NewMessage(messaging.EmissaryChannel, messaging.PauseEvent))
		time.Sleep(testDuration * 2)
		agent.Message(messaging.NewMessage(messaging.EmissaryChannel, messaging.ResumeEvent))
		time.Sleep(testDuration * 2)
		agent.Message(emissaryShutdown)
		time.Sleep(testDuration)
		ch <- struct{}{}
	}()
	<-ch
	close(ch)

	//Output:
	//fail
}
