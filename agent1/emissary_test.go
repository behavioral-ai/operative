package agent1

import (
	"github.com/behavioral-ai/core/messaging"
	"github.com/behavioral-ai/domain/collective"
	"github.com/behavioral-ai/domain/common"
	"github.com/behavioral-ai/domain/timeseries1"
	"time"
)

const (
	testDuration = time.Second * 5
)

func ExampleEmissary() {
	ch := make(chan struct{})
	agent := newOp(common.Origin{Region: "us-west"}, collective.NewEphemeralResolver(), messaging.NewTraceDispatcher(), testDuration)

	go func() {
		go emissaryAttend(agent, timeseries1.Observations)
		agent.Message(messaging.NewMessage(messaging.Emissary, messaging.DataChangeEvent))
		time.Sleep(testDuration * 2)
		agent.Message(messaging.NewMessage(messaging.Emissary, messaging.PauseEvent))
		time.Sleep(testDuration * 2)
		agent.Message(messaging.NewMessage(messaging.Emissary, messaging.ResumeEvent))
		time.Sleep(testDuration * 2)
		agent.Shutdown()
		time.Sleep(testDuration * 2)
		ch <- struct{}{}
	}()
	<-ch
	close(ch)

	//Output:
	//fail
}

func ExampleEmissary_Observation() {
	ch := make(chan struct{})
	origin := common.Origin{Region: "us-west"}
	agent := newOp(origin, collective.NewEphemeralResolver(), messaging.NewTraceDispatcher(), testDuration)

	go func() {
		go emissaryAttend(agent, timeseries1.NewObservation(timeseries1.Observation{Origin: origin, Latency: 1500, Gradient: 15}, messaging.StatusOK()))
		time.Sleep(testDuration * 2)

		// Receive observation message
		msg := <-agent.master.C
		o, status := getObservation(agent, msg)
		status.AgentUri = agent.Uri()
		status.Msg = o.String()
		agent.resolver.Notify(status)
		agent.Shutdown()
		time.Sleep(testDuration * 3)
		ch <- struct{}{}
	}()
	<-ch
	close(ch)

	//Output:
	//fail
}
