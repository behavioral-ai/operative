package agent1

import (
	"github.com/behavioral-ai/core/iox"
	"github.com/behavioral-ai/core/messaging"
	"github.com/behavioral-ai/domain/collective"
	"github.com/behavioral-ai/domain/common"
	"github.com/behavioral-ai/operative/testrsc"
	"github.com/behavioral-ai/operative/urn"
	"time"
)

func _ExampleMaster() {
	ch := make(chan struct{})
	agent := newOp(common.Origin{Region: "us-west"}, collective.NewEphemeralResolver(), messaging.NewTraceDispatcher(), 0)

	go func() {
		go masterAttend(agent)
		agent.Message(messaging.NewMessage(messaging.Master, messaging.ObservationEvent))

		agent.Message(messaging.NewMessage(messaging.Master, messaging.PauseEvent))
		agent.Message(messaging.NewMessage(messaging.Master, messaging.ObservationEvent))
		agent.Message(messaging.NewMessage(messaging.Master, messaging.ResumeEvent))
		agent.Message(messaging.NewMessage(messaging.Master, messaging.ObservationEvent))

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
	msg := messaging.NewMessage(messaging.Master, messaging.ObservationEvent)
	msg.SetContent(contentTypeObservation, observation{origin: origin, latency: 2350, gradient: 15})
	resolver, status := createResolver()
	if !status.OK() {
		messaging.Notify(status)
	}
	agent := newOp(origin, resolver, messaging.NewTraceDispatcher(), 0)

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

func createResolver() (collective.Resolution, *messaging.Status) {
	resolver := collective.NewEphemeralResolver()
	buf, err := iox.ReadFile(testrsc.ResiliencyInterpret1)
	if err != nil {
		return nil, messaging.NewStatusError(messaging.StatusIOError, err, "")
	}
	status := resolver.PutContent(urn.ResiliencyInterpret, "author", buf, 1)
	if !status.OK() {
		return nil, status
	}
	buf, err = iox.ReadFile(testrsc.ResiliencyThreshold1)
	if err != nil {
		return nil, messaging.NewStatusError(messaging.StatusIOError, err, "")
	}
	return resolver, resolver.PutContent(urn.ResiliencyThreshold, "author", buf, 1)
}
