package agent1

import (
	"github.com/behavioral-ai/core/messaging"
	"github.com/behavioral-ai/domain/collective"
	"github.com/behavioral-ai/operative/frame1"
)

// master attention
func masterAttend(agent *service, resolver collective.Resolution) {
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
				_, status1 := reason(agent, o, resolver)
				// TODO : add action to data store
				if status1.OK() {

				}
			default:
			}
		default:
		}
	}
}

func reason(agent *service, o observation, resolver collective.Resolution) (frame1.Activity, *messaging.Status) {
	activity, status := frame1.Reason(o, resolver)
	if !status.OK() {
		agent.notify(status)
		return activity, status
	}
	resolver.AddActivity(agent, messaging.ObservationEvent, agent.master.Name(), activity.Desc)
	return activity, messaging.StatusOK()
}
