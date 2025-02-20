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

	notifier messaging.NotifyFunc
	emissary *communications
	master   *communications
}

func serviceAgentUri(origin common.Origin) string {
	return fmt.Sprintf("%v#%v", ClassUrn, origin)
}

// New - create a new agent1 agent
func New(origin common.Origin, notifier messaging.NotifyFunc, global messaging.Dispatcher) messaging.Agent {
	return newOp(origin, notifier, global, newMasterDispatcher(false), newEmissaryDispatcher(false))
}

func newOp(origin common.Origin, notifier messaging.NotifyFunc, global messaging.Dispatcher, master, emissary dispatcher) *service {
	r := new(service)
	r.origin = origin
	r.uri = serviceAgentUri(origin)
	r.duration = defaultDuration

	r.notifier = notifier
	r.emissary = newEmmissaryComms(global, emissary)
	r.master = newMasterComms(global, master)
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
		s.emissary.send(m)
	case messaging.MasterChannel:
		s.master.send(m)
	default:
		s.emissary.send(m)
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
	s.emissary.enable()
	s.emissary.send(msg)
	s.master.enable()
	s.master.send(msg)
}

func (s *service) IsFinalized() bool {
	return s.emissary.isFinalized() && s.master.isFinalized()
}
