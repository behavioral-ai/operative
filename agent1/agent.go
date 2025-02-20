package agent1

import (
	"fmt"
	"github.com/behavioral-ai/core/messaging"
	"github.com/behavioral-ai/domain/collective"
	"github.com/behavioral-ai/domain/common"
	"github.com/behavioral-ai/operative/timeseries1"
	"time"
)

const (
	ClassUrn        = "resiliency:agent/operative/agent1"
	defaultDuration = time.Second * 10
)

type service struct {
	running  bool
	uri      string
	name     string
	origin   common.Origin
	duration time.Duration

	notifier   messaging.NotifyFunc
	emissary   *messaging.Channel
	master     *messaging.Channel
	dispatcher messaging.Dispatcher
}

func serviceAgentUri(origin common.Origin) string {
	return fmt.Sprintf("%v#%v", ClassUrn, origin)
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

	r.notifier = notifier
	r.emissary = messaging.NewEmissaryChannel(true) // //)newEmmissaryComms(global, emissary)
	r.master = messaging.NewMasterChannel(true)
	r.dispatcher = dispatcher
	return r
}

// String - identity
func (s *service) String() string { return s.Uri() }

// Uri - agent identifier
func (s *service) Uri() string { return s.uri }

// Name - agent urn
func (s *service) Name() string { return s.name }

// Notify - status notifications
func (s *service) Notify(status *messaging.Status) {
	if s.notifier != nil {
		s.notifier(status)
	}
}

// Message - message the agent
func (s *service) Message(m *messaging.Message) {
	if m == nil {
		return
	}
	switch m.To() {
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
	msg := messaging.NewControlMessage(s.Uri(), s.Uri(), messaging.ShutdownEvent)
	s.emissary.Enable()
	s.emissary.C <- msg
	s.master.Enable()
	s.master.C <- msg
}

func (s *service) IsFinalized() bool {
	return s.emissary.IsFinalized() && s.master.IsFinalized()
}

func (s *service) emissaryFinalize() {
	s.emissary.Close()
}

func (s *service) masterFinalize() {
	s.master.Close()
}

func (s *service) dispatch(channel any, event string) {
	if s.dispatcher == nil || channel == nil {
		return
	}
	if ch, ok := channel.(*messaging.Channel); ok {
		s.dispatcher.Dispatch(s, ch.Name(), event)
		return
	}
	if t, ok := channel.(*messaging.Ticker); ok {
		s.dispatcher.Dispatch(s, t.Name(), event)
	}
}
