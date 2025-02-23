package agent1

import (
	"fmt"
	"github.com/behavioral-ai/core/messaging"
	"github.com/behavioral-ai/core/test"
	"github.com/behavioral-ai/domain/common"
	"github.com/behavioral-ai/operative/timeseries1"
	"time"
)

func ExampleEmissary() {
	ch := make(chan struct{})
	agent := newOp(common.Origin{Region: "us-west"}, test.Notify, messaging.NewTraceDispatcher())

	go func() {
		go emissaryAttend(agent, timeseries1.NewObservation(timeseries1.Entry{}, messaging.StatusNotFound()))
		agent.Message(messaging.NewMessage(messaging.EmissaryChannel, messaging.DataChangeEvent))
		time.Sleep(testDuration * 2)
		agent.Message(messaging.NewMessage(messaging.EmissaryChannel, messaging.PauseEvent))
		time.Sleep(testDuration * 2)
		agent.Message(messaging.NewMessage(messaging.EmissaryChannel, messaging.ResumeEvent))
		time.Sleep(testDuration * 2)
		//agent.Message(emissaryShutdown)
		agent.Shutdown()
		time.Sleep(testDuration * 2)
		ch <- struct{}{}
	}()
	<-ch
	close(ch)

	//Output:
	//fail
}

func _ExampleEmissary_Observation() {
	ch := make(chan struct{})
	origin := common.Origin{Region: "us-west"}
	agent := newOp(origin, test.Notify, messaging.NewTraceDispatcher())

	go func() {
		go emissaryAttend(agent, timeseries1.NewObservation(timeseries1.Entry{Origin: origin, Latency: 1500, Gradient: 15}, messaging.StatusOK()))
		time.Sleep(testDuration * 2)
		agent.Message(messaging.NewMessage(messaging.MasterChannel, messaging.ShutdownEvent))
		time.Sleep(testDuration)
		ch <- struct{}{}
	}()
	<-ch
	close(ch)
	// Receive observation message
	msg := <-agent.master.C
	o, status := getObservation(agent, msg)
	fmt.Printf("test: getObservation() -> [status:%v] [%v]\n", status, o)

	//Output:
	//fail
}
