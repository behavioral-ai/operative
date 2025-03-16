package agent1

import (
	"github.com/behavioral-ai/core/messaging"
	"github.com/behavioral-ai/domain/common"
	"github.com/behavioral-ai/domain/content"
	"github.com/behavioral-ai/operative/test"
	"time"
)

func ExampleMaster() {
	ch := make(chan struct{})
	agent := newAgent(common.Origin{Region: common.WestRegion}, content.NewEphemeralResolver(), messaging.NewTraceDispatcher())

	go func() {
		go masterAttend(agent)
		agent.Message(messaging.NewMessage(messaging.Master, messaging.ObservationEvent))

		agent.Message(messaging.NewMessage(messaging.Master, messaging.PauseEvent))
		agent.Message(messaging.NewMessage(messaging.Master, messaging.ObservationEvent))
		agent.Message(messaging.NewMessage(messaging.Master, messaging.ResumeEvent))
		agent.Message(messaging.NewMessage(messaging.Master, messaging.ObservationEvent))

		agent.Shutdown() //Message(messaging.MasterShutdown)
		time.Sleep(testDuration)

		ch <- struct{}{}
	}()
	<-ch
	close(ch)
	//Output:
	//fail
}

func ExampleMaster_Observation() {
	ch := make(chan struct{})
	origin := common.Origin{Region: "us-west"}
	msg := messaging.NewMessage(messaging.Master, messaging.ObservationEvent)
	msg.SetContent(contentTypeObservation, observation{origin: origin, latency: 2350, gradient: 15})
	resolver, status := test.NewResiliencyResolver()
	if !status.OK() {
		messaging.Notify(status)
	}
	agent := newAgent(origin, resolver, messaging.NewTraceDispatcher())

	go func() {
		go masterAttend(agent)
		agent.Message(msg)
		time.Sleep(testDuration * 2)

		agent.Shutdown()
		time.Sleep(testDuration)

		ch <- struct{}{}
	}()
	<-ch
	close(ch)

	//Output:
	//fail

}
