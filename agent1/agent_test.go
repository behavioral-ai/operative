package agent1

import (
	"fmt"
	"github.com/behavioral-ai/collective/content"
	"github.com/behavioral-ai/core/messaging"
	"github.com/behavioral-ai/core/messaging/messagingtest"
	"github.com/behavioral-ai/domain/common"
	"github.com/behavioral-ai/domain/timeseries1"
	"github.com/behavioral-ai/operative/test"
	"time"
)

func ExampleNewAgent() {
	a := New(common.Origin{Region: "us-central", Zone: "c-zone-a", SubZone: "sub-zone", Host: "www.host.com"}, content.NewEphemeralResolver(), nil)

	fmt.Printf("test: NewAgent() -> [%v]\n", a)
	fmt.Printf("test: NewAgent() -> [%v]\n", a.Name())

	//Output:
	//test: NewAgent() -> [resiliency:agent/operative1#us-central.c-zone-a.sub-zone.www.host.com]
	//test: NewAgent() -> [resiliency:agent/operative]

}

func ExampleAgent_Ephemeral() {
	ch := make(chan struct{})
	dispatcher := messaging.NewTraceDispatcher()
	origin := common.Origin{Region: common.WestRegion, Zone: common.WestZoneA}
	s := messagingtest.NewTestSpanner(time.Second*2, testDuration)
	resolver, status := test.NewResiliencyResolver()
	if !status.OK() {
		messaging.Notify(status)
	}
	agent := newAgent(origin, resolver, dispatcher)

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

func ExampleAgent_NotFound() {
	ch := make(chan struct{})
	dispatcher := messaging.NewTraceDispatcher()
	origin := common.Origin{Region: common.WestRegion, Zone: common.WestZoneA}
	agent := newAgent(origin, nil, dispatcher)

	go func() {
		agent.Run()
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

func _ExampleAgent_Resolver() {
	ch := make(chan struct{})
	dispatcher := messaging.NewTraceDispatcher()
	origin := common.Origin{Region: common.WestRegion, Zone: common.WestZoneA}
	agent := newAgent(origin, nil, dispatcher)
	//test2.Startup()

	go func() {
		agent.Run()
		//go masterAttend(agent)
		//go emissaryAttend(agent, timeseries1.Observations, s)
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
