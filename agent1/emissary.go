package agent1

import (
	"github.com/behavioral-ai/core/messaging"
	"github.com/behavioral-ai/operative/timeseries1"
)

// emissary attention
func emissaryAttend(agent *service, observe *timeseries1.Observation) {
	paused := false
	ticker := messaging.NewPrimaryTicker(agent.duration)
	ticker.Start(-1)
	agent.dispatch(agent.emissary, messaging.StartupEvent)

	for {
		select {
		case <-ticker.C():
			agent.dispatch(ticker, messaging.ObservationEvent)
			if !paused {
				e, status := observe.Timeseries(agent.origin)
				if status.OK() {
					m := messaging.NewControlMessage(messaging.MasterChannel, agent.Uri(), messaging.ObservationEvent)
					m.SetContent(contentTypeObservation, observation{origin: e.Origin, latency: e.Latency, gradient: e.Gradient})
					agent.Message(m)
				} else {
					agent.notify(status)
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
