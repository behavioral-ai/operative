package agent1

import (
	"fmt"
	"github.com/behavioral-ai/core/messaging"
	"github.com/behavioral-ai/domain/common"
	"github.com/behavioral-ai/operative/timeseries1"
	"time"
)

func ExampleNewAgent() {
	a := New(nil, common.Origin{Region: "us-central", Zone: "c-zone-a", SubZone: "sub-zone", Host: "www.host.com"}, nil)

	fmt.Printf("test: NewAgent() -> [%v]\n", a)
	fmt.Printf("test: NewAgent() -> [%v]\n", a.Name())

	//Output:
	//test: NewAgent() -> [resiliency:agent/operative1#us-central.c-zone-a.sub-zone.www.host.com]
	//test: NewAgent() -> [resiliency:agent/operative]

}

func ExampleAgent() {
	ch := make(chan struct{})
	origin := common.Origin{Region: "us-west"}
	agent := newOp(nil, origin, messaging.Notify, messaging.NewTraceDispatcher(), testDuration)
	resolver, status := createResolver()
	if !status.OK() {
		messaging.Notify(status)
	}

	go func() {
		go masterAttend(agent, resolver)
		go emissaryAttend(agent, timeseries1.NewObservation(timeseries1.Entry{Origin: origin, Latency: 1500, Gradient: 15}, messaging.StatusOK()))
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
