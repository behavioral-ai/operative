package agent1

import (
	"github.com/behavioral-ai/core/messaging"
	"github.com/behavioral-ai/domain/timeseries1"
	"time"
)

// emissary attention
func emissaryAttend(agent *agentT, observe *timeseries1.Observer, duration time.Duration) {
	agent.dispatch(agent.emissary, messaging.StartupEvent)
	paused := false
	agent.reviseTicker(duration)

	for {
		select {
		case <-agent.ticker.C():
			if !paused {
				agent.dispatch(agent.ticker, messaging.ObservationEvent)
				e, status := observe.Timeseries(agent.origin)
				if status.OK() {
					m := messaging.NewMessage(messaging.Master, messaging.ObservationEvent)
					m.SetContent(contentTypeObservation, observation{origin: e.Origin, latency: e.Latency, gradient: e.Gradient})
					agent.Message(m)
				} else {
					status.AgentUri = agent.Uri()
					agent.resolver.Notify(status)
				}
			}
			agent.reviseTicker(duration)
		default:
		}
		select {
		case msg := <-agent.emissary.C:
			agent.dispatch(agent.emissary, msg.Event())
			switch msg.Event() {
			case messaging.PauseEvent:
				paused = true
			case messaging.ResumeEvent:
				paused = false
			case messaging.ShutdownEvent:
				agent.emissaryFinalize()
				return
			default:
			}
		default:
		}
	}
}
