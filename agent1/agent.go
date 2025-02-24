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
	Name            = "resiliency:agent/operative/agent"
	defaultDuration = time.Second * 10
	testDuration    = time.Second * 5
)

type service struct {
	running  bool
	uri      string
	origin   common.Origin
	duration time.Duration

	emissary   *messaging.Channel
	master     *messaging.Channel
	notifier   messaging.NotifyFunc
	dispatcher messaging.Dispatcher
}

func serviceAgentUri(origin common.Origin) string {
	return fmt.Sprintf("%v%v#%v", Name, strconv.Itoa(version), origin)
}

// New - create a new agent1 agent
func New(origin common.Origin, notifier messaging.NotifyFunc, dispatcher messaging.Dispatcher) messaging.Agent {
	return newOp(origin, notifier, dispatcher)
}

func newOp(origin common.Origin, notifier messaging.NotifyFunc, dispatcher messaging.Dispatcher) *service {
	r := new(service)
	r.origin = origin
	r.uri = serviceAgentUri(origin)
	r.duration = defaultDuration

	r.emissary = messaging.NewEmissaryChannel(true)
	r.master = messaging.NewMasterChannel(true)
	r.notifier = notifier
	if r.notifier == nil {
		r.notifier = collective.Resolver.Notify
	}
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
	default:
		s.emissary.C <- m
	}
}

// Run - run the agent
func (s *service) Run() {
	if s.running {
		return
	}
	go masterAttend(s, collective.Append, collective.Resolver)
	go emissaryAttend(s, timeseries1.Observe)
	s.running = true
}

// Shutdown - shutdown the agent
func (s *service) Shutdown() {
	if !s.running {
		return
	}
	s.running = false
	s.emissary.Enable()
	s.emissary.C <- messaging.Shutdown
	s.master.Enable()
	s.master.C <- messaging.Shutdown
}

func (s *service) notify(e messaging.Event) {
	s.notifier(e)
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
