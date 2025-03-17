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
				o, status := getObservation(msg)
				if !status.OK() {
					agent.notify(status.SetAgent(agent.Uri()))
					continue
				}
				// Process reasoning
				action, status1 := reason(agent, o)
				if !status1.OK() {
					continue
				}
				agent.resolver.AddActivity(agent, messaging.ObservationEvent, agent.master.Name(), action.Desc)
				// TODO : add action to data store
			default:
			}
		default:
		}
	}
}

func reason(agent *agentT, o observation) (frame1.Action, *messaging.Status) {
	action, status := frame1.Reason(o, agent.resolver)
	if !status.OK() {
		if status.NotFound() {
			status.SetAgent(agent.Uri())
		}
		agent.notify(status)
		return action, status
	}
	return action, messaging.StatusOK()
}
