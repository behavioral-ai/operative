package agent1

import (
	"github.com/behavioral-ai/core/messaging"
	"github.com/behavioral-ai/domain/collective"
	"github.com/behavioral-ai/operative/frame1"
)

// master attention
func masterAttend(agent *service, append collective.Appender, resolver collective.Resolution) {
	paused := false
	agent.dispatch(agent.master, messaging.StartupEvent)

	for {
		select {
		case msg := <-agent.master.C:
			switch msg.Event() {
			case messaging.PauseEvent:
				paused = true
			case messaging.ResumeEvent:
				paused = false
			case messaging.ShutdownEvent:
				agent.masterFinalize()
				agent.dispatch(agent.master, msg.Event())
				return
			case messaging.ObservationEvent:
				if !paused {
					o, status := getObservation(agent, msg)
					if status.OK() {
						reason(agent, o, append, resolver)
					}
				}
			default:
			}
			agent.dispatch(agent.master, msg.Event())
		default:
		}
	}
}

func reason(agent *service, o observation, append collective.Appender, resolver collective.Resolution) frame1.Activity {
	status, activity := frame1.Reason(o, resolver)
	if !status.OK() {
		agent.notify(status)
		return activity
	}
	append.Activity(agent, messaging.ObservationEvent, agent.master.Name(), activity.Desc)
	return activity
}
