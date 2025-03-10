package agent1

import (
	"github.com/behavioral-ai/core/messaging"
	"github.com/behavioral-ai/domain/timeseries1"
)

// emissary attention
func emissaryAttend(agent *agentT, observe *timeseries1.Observer) {
	agent.dispatch(agent.emissary, messaging.StartupEvent)
	paused := false
	ticker := messaging.NewTicker(messaging.Emissary, agent.duration)
	ticker.Start(-1)

	for {
		select {
		case <-ticker.C():
			if !paused {
				agent.dispatch(ticker, messaging.ObservationEvent)
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
				ticker.Stop()
				agent.emissaryFinalize()
				return
			default:
			}
		default:
		}
	}
}
