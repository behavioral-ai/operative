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
func (s *agentT) String() string { return s.Uri() }

// Uri - agent identifier
func (s *agentT) Uri() string { return s.uri }

// Name - agent urn
func (s *agentT) Name() string { return Name }

// Message - message the agent
func (s *agentT) Message(m *messaging.Message) {
	if m == nil {
		return
	}
	switch m.Channel() {
	case messaging.Emissary:
		s.emissary.C <- m
	case messaging.Master:
		s.master.C <- m
	case messaging.Control:
		s.emissary.C <- m
		s.master.C <- m
	default:
		s.emissary.C <- m
	}
}

// Run - run the agent
func (s *agentT) Run() {
	if s.running {
		return
	}
	go masterAttend(s)
	go emissaryAttend(s, timeseries1.Observations)
	s.running = true
}

// Shutdown - shutdown the agent
func (s *agentT) Shutdown() {
	if !s.emissary.IsClosed() {
		s.emissary.C <- messaging.Shutdown
	}
	if !s.master.IsClosed() {
		s.master.C <- messaging.Shutdown
	}
}

func (s *agentT) dispatch(channel any, event string) {
	messaging.Dispatch(s, s.dispatcher, channel, event)
}

func (s *agentT) emissaryFinalize() {
	s.emissary.Close()
}

func (s *agentT) masterFinalize() {
	s.master.Close()
}
