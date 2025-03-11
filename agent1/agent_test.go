package agent1

import (
	"fmt"
	"github.com/behavioral-ai/core/messaging"
	"github.com/behavioral-ai/core/messaging/messagingtest"
	"github.com/behavioral-ai/domain/collective"
	"github.com/behavioral-ai/domain/common"
	"github.com/behavioral-ai/domain/timeseries1"
	"time"
)

func ExampleNewAgent() {
	a := New(common.Origin{Region: "us-central", Zone: "c-zone-a", SubZone: "sub-zone", Host: "www.host.com"}, collective.NewEphemeralResolver(), nil)

	fmt.Printf("test: NewAgent() -> [%v]\n", a)
	fmt.Printf("test: NewAgent() -> [%v]\n", a.Name())

	//Output:
	//test: NewAgent() -> [resiliency:agent/operative1#us-central.c-zone-a.sub-zone.www.host.com]
	//test: NewAgent() -> [resiliency:agent/operative]

}

func ExampleEphemeralAgent() {
	ch := make(chan struct{})
	dispatcher := messaging.NewTraceDispatcher()
	origin := common.Origin{Region: common.WestRegion, Zone: common.WestZoneA}
	s := messagingtest.NewTestSpanner(testDuration, testDuration)
	resolver, status := createResolver()
	if !status.OK() {
		messaging.Notify(status)
	}
	agent := newAgent(origin, resolver, dispatcher)

	go func() {
		go masterAttend(agent)
		go emissaryAttend(agent, timeseries1.Observations, s)
		//	go emissaryAttend(agent, timeseries1.NewObservation(timeseries1.Observation{Origin: origin, Latency: 1500, Gradient: 15}, messaging.StatusOK()))
		time.Sleep(testDuration * 5)

		agent.Shutdown()
		time.Sleep(testDuration * 2)
		ch <- struct{}{}
	}()
	<-ch
	close(ch)

	//Output:
	//fail
}

func ExampleAgent() {
	ch := make(chan struct{})
	collective.Startup(nil, nil, "")
	//dispatcher := messaging.NewTraceDispatcher()
	s := messagingtest.NewTestSpanner(testDuration, testDuration*2)
	origin := common.Origin{Region: common.WestRegion, Zone: common.WestZoneA}
	agent := newAgent(origin, nil, nil)

	go func() {
		go masterAttend(agent)
		go emissaryAttend(agent, timeseries1.Observations, s)
		time.Sleep(testDuration * 5)
		agent.Shutdown()
		time.Sleep(testDuration * 2)
		ch <- struct{}{}
	}()
	<-ch
	close(ch)

	//Output:
	//fail
}
