package agent1

import (
	"github.com/behavioral-ai/core/messaging"
	"github.com/behavioral-ai/core/test"
	"github.com/behavioral-ai/domain/collective"
	"github.com/behavioral-ai/domain/common"
	"time"
)

func ExampleMaster() {
	ch := make(chan struct{})
	agent := newOp(nil, common.Origin{Region: "us-west"}, test.Notify, messaging.NewTraceDispatcher())

	go func() {
		go masterAttend(agent, collective.NewEphemeralResolver("", nil))
		agent.Message(messaging.NewMessage(messaging.MasterChannel, messaging.ObservationEvent))

		agent.Message(messaging.NewMessage(messaging.MasterChannel, messaging.PauseEvent))
		agent.Message(messaging.NewMessage(messaging.MasterChannel, messaging.ObservationEvent))
		agent.Message(messaging.NewMessage(messaging.MasterChannel, messaging.ResumeEvent))
		agent.Message(messaging.NewMessage(messaging.MasterChannel, messaging.ObservationEvent))

		agent.Message(messaging.MasterShutdown)
		time.Sleep(testDuration)

		ch <- struct{}{}
	}()
	<-ch
	close(ch)

}

func ExampleMaster_Observation() {
	ch := make(chan struct{})
	origin := common.Origin{Region: "us-west"}
	agent := newOp(nil, origin, test.Notify, messaging.NewTraceDispatcher())
	msg := messaging.NewMessage(messaging.MasterChannel, messaging.ObservationEvent)
	msg.SetContent(contentTypeObservation, observation{origin: origin, latency: 2350, gradient: 15})

	go func() {
		go masterAttend(agent, collective.NewEphemeralResolver("", nil))
		agent.Message(msg)

		agent.Message(messaging.MasterShutdown)
		time.Sleep(testDuration)

		ch <- struct{}{}
	}()
	<-ch
	close(ch)

	//Output:
	//test: masterAttend() -> [finalized:true]

}
