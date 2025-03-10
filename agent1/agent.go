package agent1

import (
	"fmt"
	"github.com/behavioral-ai/core/messaging"
	"github.com/behavioral-ai/domain/collective"
	"github.com/behavioral-ai/domain/common"
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
	running  bool
	uri      string
	origin   common.Origin
	duration time.Duration

	emissary   *messaging.Channel
	master     *messaging.Channel
	resolver   collective.Resolution
	dispatcher messaging.Dispatcher
}

func serviceAgentUri(origin common.Origin) string {
	return fmt.Sprintf("%v%v#%v", Name, strconv.Itoa(version), origin)
}

// New - create a new agent1 agent
func New(origin common.Origin, resolver collective.Resolution, dispatcher messaging.Dispatcher) messaging.Agent {
	return newOp(origin, resolver, dispatcher, 0)
}

func newOp(origin common.Origin, resolver collective.Resolution, dispatcher messaging.Dispatcher, d time.Duration) *agentT {
	r := new(agentT)
	r.origin = origin
	r.uri = serviceAgentUri(origin)
	if d <= 0 {
		r.duration = minDuration
	} else {
		r.duration = d
	}
	if resolver == nil {
		r.resolver = collective.Resolver
	} else {
		r.resolver = resolver
	}
	r.emissary = messaging.NewEmissaryChannel()
	r.master = messaging.NewMasterChannel()
	r.dispatcher = dispatcher
	return r
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
	go emissaryAttend(a, timeseries1.Observations)
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

func (a *agentT) dispatch(channel any, event string) {
	messaging.Dispatch(a, a.dispatcher, channel, event)
}

func (a *agentT) emissaryFinalize() {
	a.emissary.Close()
}

func (a *agentT) masterFinalize() {
	a.master.Close()
}
