package agent1

import (
	"fmt"
	"github.com/behavioral-ai/core/messaging"
	"github.com/behavioral-ai/core/test"
	"github.com/behavioral-ai/domain/collective"
	"github.com/behavioral-ai/domain/common"
)

func ExampleMaster() {
	ch := make(chan struct{})
	agent := newOp(nil, common.Origin{Region: "us-west"}, test.Notify, messaging.NewTraceDispatcher())

	go func() {
		go masterAttend(agent, collective.Resolver)
		//agent.Message(messaging.NewMessage(messaging.MasterChannel, messaging.ObservationEvent))
		agent.Message(messaging.MasterShutdown)
		fmt.Printf("test: masterAttend() -> [finalized:%v]\n", true)
		ch <- struct{}{}
	}()
	<-ch
	close(ch)

	//Output:
	//test: masterAttend() -> [finalized:true]

}

func _ExampleMaster_Observation() {
	ch := make(chan struct{})
	agent := newOp(nil, common.Origin{Region: "us-west"}, test.Notify, messaging.NewTraceDispatcher())

	go func() {
		go masterAttend(agent, collective.NewEphemeralResolver("", nil))
		agent.Message(messaging.NewMessage(messaging.MasterChannel, messaging.ObservationEvent))
		agent.Message(messaging.MasterShutdown)
		fmt.Printf("test: masterAttend() -> [finalized:%v]\n", true)
		ch <- struct{}{}
	}()
	<-ch
	close(ch)

	//Output:
	//test: masterAttend() -> [finalized:true]

}
