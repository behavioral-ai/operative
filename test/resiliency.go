package test

import (
	"github.com/behavioral-ai/core/iox"
	"github.com/behavioral-ai/core/messaging"
	"github.com/behavioral-ai/domain/collective"
	"github.com/behavioral-ai/domain/test"
	"github.com/behavioral-ai/domain/testrsc"
	"github.com/behavioral-ai/operative/urn"
)

func NewResiliencyResolver() (collective.Resolution, *messaging.Status) {
	r := collective.NewEphemeralResolver()
	status := loadResiliencyContent(r)
	if status.OK() {
		status = test.LoadProfile(r)
	}
	return r, status
}

func loadResiliencyContent(r collective.Resolution) *messaging.Status {
	buf, err := iox.ReadFile(testrsc.ResiliencyInterpret1)
	if err != nil {
		return messaging.NewStatusError(messaging.StatusIOError, err, "")
	}
	status := r.PutContent(urn.ResiliencyInterpret, "author", buf, 1)
	if !status.OK() {
		return status
	}
	buf, err = iox.ReadFile(testrsc.ResiliencyThreshold1)
	if err != nil {
		return messaging.NewStatusError(messaging.StatusIOError, err, "")
	}
	return r.PutContent(urn.ResiliencyThreshold, "author", buf, 1)
}
