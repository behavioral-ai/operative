package agent1

import (
	"fmt"
	"github.com/behavioral-ai/core/messaging"
	"github.com/behavioral-ai/domain/collective"
	"github.com/behavioral-ai/domain/common"
	"github.com/behavioral-ai/operative/timeseries1"
	"strconv"
	"time"
)

const (
	Name            = "resiliency:agent/operative"
	defaultDuration = time.Second * 10
)

type service struct {
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

func newOp(origin common.Origin, resolver collective.Resolution, dispatcher messaging.Dispatcher, d time.Duration) *service {
	r := new(service)
	r.origin = origin
	r.uri = serviceAgentUri(origin)
	if d <= 0 {
		r.duration = defaultDuration
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
func (s *service) String() string { return s.Uri() }

// Uri - agent identifier
func (s *service) Uri() string { return s.uri }

// Name - agent urn
func (s *service) Name() string { return Name }

// Message - message the agent
func (s *service) Message(m *messaging.Message) {
	if m == nil {
		return
	}
	switch m.Channel() {
	case messaging.EmissaryChannel:
		s.emissary.C <- m
	case messaging.MasterChannel:
		s.master.C <- m
	case messaging.ControlChannel:
		s.emissary.C <- m
		s.master.C <- m
	default:
		s.emissary.C <- m
	}
}

// Run - run the agent
func (s *service) Run() {
	if s.running {
		return
	}
	go masterAttend(s)
	go emissaryAttend(s, timeseries1.Observe)
	s.running = true
}

// Shutdown - shutdown the agent
func (s *service) Shutdown() {
	if !s.emissary.IsClosed() {
		s.emissary.C <- messaging.Shutdown
	}
	if !s.master.IsClosed() {
		s.master.C <- messaging.Shutdown
	}
}

func (s *service) notify(e messaging.Event) {
	s.resolver.Notify(e)
}

func (s *service) dispatch(channel any, event string) {
	messaging.Dispatch(s, s.dispatcher, channel, event)
}

func (s *service) emissaryFinalize() {
	s.emissary.Close()
}

func (s *service) masterFinalize() {
	s.master.Close()
}
