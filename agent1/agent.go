package agent1

import (
	"fmt"
	"github.com/behavioral-ai/core/messaging"
	"github.com/behavioral-ai/domain/collective"
	"github.com/behavioral-ai/domain/common"
	"github.com/behavioral-ai/domain/metrics1"
	"github.com/behavioral-ai/domain/timeseries1"
	"strconv"
	"time"
)

const (
	Name        = "resiliency:agent/operative"
	minDuration = time.Second * 10
	maxDuration = time.Second * 15
)

type agentT struct {
	running bool
	uri     string
	traffic string
	origin  common.Origin

	ticker     *messaging.Ticker
	emissary   *messaging.Channel
	master     *messaging.Channel
	resolver   collective.Resolution
	dispatcher messaging.Dispatcher
}

func agentUri(origin common.Origin) string {
	return fmt.Sprintf("%v%v#%v", Name, strconv.Itoa(version), origin)
}

// New - create a new agent1 agent
func New(origin common.Origin, resolver collective.Resolution, dispatcher messaging.Dispatcher) messaging.Agent {
	return newAgent(origin, resolver, dispatcher)
}

func newAgent(origin common.Origin, resolver collective.Resolution, dispatcher messaging.Dispatcher) *agentT {
	a := new(agentT)
	a.origin = origin
	a.uri = agentUri(origin)

	if resolver == nil {
		a.resolver = collective.Resolver
	} else {
		a.resolver = resolver
	}
	a.ticker = messaging.NewTicker(messaging.Emissary, maxDuration)
	a.emissary = messaging.NewEmissaryChannel()
	a.master = messaging.NewMasterChannel()
	a.dispatcher = dispatcher
	return a
}

// String - identity
func (a *agentT) String() string { return a.Uri() }

// Uri - agent identifier
func (a *agentT) Uri() string { return a.uri }

// Name - agent urn
func (a *agentT) Name() string { return Name }

// Message - message the agent
func (a *agentT) Message(m *messaging.Message) {
	if m == nil {
		return
	}
	switch m.Channel() {
	case messaging.Emissary:
		a.emissary.C <- m
	case messaging.Master:
		a.master.C <- m
	case messaging.Control:
		a.emissary.C <- m
		a.master.C <- m
	default:
		a.emissary.C <- m
	}
}

// Run - run the agent
func (a *agentT) Run() {
	if a.running {
		return
	}
	go masterAttend(a)
	go emissaryAttend(a, timeseries1.Observations, 0)
	a.running = true
}

// Shutdown - shutdown the agent
func (a *agentT) Shutdown() {
	if !a.emissary.IsClosed() {
		a.emissary.C <- messaging.Shutdown
	}
	if !a.master.IsClosed() {
		a.master.C <- messaging.Shutdown
	}
}

func (a *agentT) reviseTicker(duration time.Duration) {
	if duration != 0 {
		a.ticker.Start(duration)
		return
	}
	p, status := collective.Resolve[metrics1.TrafficProfile](metrics1.ProfileName, 1, collective.Resolver)
	if !status.OK() {
		a.resolver.Notify(status)
		return
	}
	traffic := p.Now()
	if traffic == metrics1.TrafficMedium || traffic == a.traffic {
		return
	}
	if traffic == metrics1.TrafficLow {
		a.ticker.Start(maxDuration)
	} else {
		a.ticker.Start(minDuration)
	}
	a.traffic = traffic
}

func (a *agentT) dispatch(channel any, event string) {
	messaging.Dispatch(a, a.dispatcher, channel, event)
}

func (a *agentT) emissaryFinalize() {
	a.emissary.Close()
	a.ticker.Stop()
}

func (a *agentT) masterFinalize() {
	a.master.Close()
}
