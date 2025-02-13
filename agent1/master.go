package agent1

import (
	"github.com/behavioral-ai/core/messaging"
	"github.com/behavioral-ai/operative/knowledge1"
)

// master attention
func masterAttend(agent *service) {
	paused := false
	comms := agent.master
	comms.dispatch(agent, messaging.StartupEvent)

	for {
		// message processing
		select {
		case msg := <-comms.channel().C:
			comms.setup(agent, msg.Event())
			switch msg.Event() {
			case messaging.PauseEvent:
				paused = true
			case messaging.ResumeEvent:
				paused = false
			case messaging.ShutdownEvent:
				comms.finalize()
				comms.dispatch(agent, msg.Event())
				return
			case messaging.ObservationEvent:
				if !paused {
					observe, status := knowledge1.GetObservation(agent.handler, agent.Uri(), msg)
					if status.OK() {
						if observe.Gradient > 10 {
						}
						/*
							inf := runInference(r, observe)
							if inf == nil {
								continue
							}
							action := newAction(inf)
							rateLimiting.Limit = action.Limit
							rateLimiting.Burst = action.Burst
							common1.AddRateLimitingExperience(r.handler, r.origin, inf, action, exp)


						*/
					}
				}
			default:
				agent.handler.Notify(messaging.EventError(agent.Uri(), msg))
			}
			comms.dispatch(agent, msg.Event())
		default:
		}
	}
}
