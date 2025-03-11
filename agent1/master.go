package agent1

import (
	"github.com/behavioral-ai/core/messaging"
	"github.com/behavioral-ai/operative/frame1"
)

// master attention
func masterAttend(agent *agentT) {
	agent.dispatch(agent.master, messaging.StartupEvent)
	paused := false

	for {
		select {
		case msg := <-agent.master.C:
			agent.dispatch(agent.master, msg.Event())
			switch msg.Event() {
			case messaging.PauseEvent:
				paused = true
			case messaging.ResumeEvent:
				paused = false
			case messaging.ShutdownEvent:
				agent.masterFinalize()
				return
			case messaging.ObservationEvent:
				if paused {
					continue
				}
				o, status := getObservation(agent, msg)
				if !status.OK() {
					continue
				}
				// Process reasoning
				_, status1 := reason(agent, o)
				if !status1.OK() {
					status.SetAgent(agent.uri)
					agent.resolver.Notify(status1)
				} else {
					// TODO : add action to data store

				}
			default:
			}
		default:
		}
	}
}

func reason(agent *agentT, o observation) (frame1.Action, *messaging.Status) {
	action, status := frame1.Reason(o, agent.resolver)
	if !status.OK() {
		agent.resolver.Notify(status)
		return action, status
	}
	agent.resolver.AddActivity(agent, messaging.ObservationEvent, agent.master.Name(), action.Desc)
	return action, messaging.StatusOK()
}
