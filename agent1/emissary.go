package agent1

import (
	"github.com/behavioral-ai/core/messaging"
	"github.com/behavioral-ai/domain/guidance"
	"github.com/behavioral-ai/domain/timeseries1"
	"github.com/behavioral-ai/operative/knowledge1"
)

// emissary attention
func emissaryAttend(agent *service, observe *timeseries1.Observation) {
	paused := false
	comms := agent.emissary
	comms.dispatch(agent, messaging.StartupEvent)
	ticker := messaging.NewPrimaryTicker(agent.duration)

	ticker.Start(-1)
	for {
		select {
		case <-ticker.C():
			if !paused {
				e, status := observe.Timeseries(agent.handler, agent.origin)
				if status.OK() {
					m := messaging.NewControlMessage(messaging.MasterChannel, agent.Uri(), messaging.ObservationEvent)
					m.SetContent(knowledge1.ContentTypeObservation, knowledge1.Observation{
						Latency:  e.Latency,
						Gradient: e.Gradient})
					agent.Message(m)
					comms.dispatch(agent, messaging.ObservationEvent)
				}
			}
		default:
		}
		select {
		case msg := <-comms.channel().C:
			comms.setup(agent, msg.Event())
			switch msg.Event() {
			case messaging.PauseEvent:
				paused = true
			case messaging.ResumeEvent:
				paused = false
			case messaging.ShutdownEvent:
				ticker.Stop()
				comms.finalize()
				comms.dispatch(agent, msg.Event())
				return
			case messaging.DataChangeEvent:
				if p := guidance.GetCalendar(agent.handler, agent.Uri(), msg); p != nil {
					//comms.dispatch(agent, msg.Event())
				}
			default:
				agent.handler.Notify(messaging.EventErrorStatus(agent.Uri(), msg))
			}
			comms.dispatch(agent, msg.Event())
		default:
		}
	}
}
