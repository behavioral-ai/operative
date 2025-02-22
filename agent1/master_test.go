package agent1

import (
	"fmt"
	"github.com/behavioral-ai/core/messaging"
	"github.com/behavioral-ai/core/test"
	"github.com/behavioral-ai/domain/collective"
	"github.com/behavioral-ai/domain/common"
)

var (
	masterShutdown = messaging.NewControlMessage(messaging.MasterChannel, "", messaging.ShutdownEvent)
	observationMsg = messaging.NewControlMessage(messaging.MasterChannel, "", messaging.ObservationEvent)
)

func ExampleMaster() {
	ch := make(chan struct{})
	agent := newOp(common.Origin{Region: "us-west"}, test.Notify, messaging.NewTraceDispatcher())

	go func() {
		go masterAttend(agent, collective.Append, collective.Resolver)
		agent.Message(observationMsg)
		agent.Message(masterShutdown)
		fmt.Printf("test: masterAttend() -> [finalized:%v]\n", true)
		ch <- struct{}{}
	}()
	<-ch
	close(ch)

	//Output:
	//test: masterAttend() -> [finalized:true]

}
