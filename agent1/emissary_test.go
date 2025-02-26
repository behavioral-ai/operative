package agent1

import (
	"github.com/behavioral-ai/core/messaging"
	"github.com/behavioral-ai/domain/common"
	"github.com/behavioral-ai/operative/timeseries1"
	"time"
)

const (
	testDuration = time.Second * 5
)

func ExampleEmissary() {
	ch := make(chan struct{})
	agent := newOp(nil, common.Origin{Region: "us-west"}, messaging.Notify, messaging.NewTraceDispatcher(), testDuration)

	go func() {
		go emissaryAttend(agent, timeseries1.NewObservation(timeseries1.Entry{}, messaging.StatusNotFound()))
		agent.Message(messaging.NewMessage(messaging.EmissaryChannel, messaging.DataChangeEvent))
		time.Sleep(testDuration * 2)
		agent.Message(messaging.NewMessage(messaging.EmissaryChannel, messaging.PauseEvent))
		time.Sleep(testDuration * 2)
		agent.Message(messaging.NewMessage(messaging.EmissaryChannel, messaging.ResumeEvent))
		time.Sleep(testDuration * 2)
		agent.Shutdown() //Message(messaging.EmissaryShutdown)
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
	agent := newOp(nil, origin, messaging.Notify, messaging.NewTraceDispatcher(), testDuration)

	go func() {
		go emissaryAttend(agent, timeseries1.NewObservation(timeseries1.Entry{Origin: origin, Latency: 1500, Gradient: 15}, messaging.StatusOK()))
		time.Sleep(testDuration * 2)

		// Receive observation message
		msg := <-agent.master.C
		o, status := getObservation(agent, msg)
		status.AgentUri = agent.Uri()
		status.Msg = o.String()
		agent.notify(status)
		agent.Shutdown() //Message(messaging.EmissaryShutdown)
		time.Sleep(testDuration * 3)
		ch <- struct{}{}
	}()
	<-ch
	close(ch)

	//Output:
	//fail
}
