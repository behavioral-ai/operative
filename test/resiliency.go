package test

import (
	"github.com/behavioral-ai/collective/content"
	"github.com/behavioral-ai/collective/test"
	"github.com/behavioral-ai/collective/testrsc"
	"github.com/behavioral-ai/core/iox"
	"github.com/behavioral-ai/core/messaging"
	"github.com/behavioral-ai/operative/urn"
)

func NewResiliencyResolver() (content.Resolution, *messaging.Status) {
	r := content.NewEphemeralResolver()
	status := loadResiliencyContent(r)
	if status.OK() {
		status = test.LoadProfile(r)
	}
	return r, status
}

func loadResiliencyContent(r content.Resolution) *messaging.Status {
	buf, err := iox.ReadFile(testrsc.ResiliencyInterpret1)
	if err != nil {
		return messaging.NewStatusError(messaging.StatusIOError, err, "")
	}
	status := r.PutValue(urn.ResiliencyInterpret, "author", buf, 1)
	if !status.OK() {
		return status
	}
	buf, err = iox.ReadFile(testrsc.ResiliencyThreshold1)
	if err != nil {
		return messaging.NewStatusError(messaging.StatusIOError, err, "")
	}
	return r.PutValue(urn.ResiliencyThreshold, "author", buf, 1)
}
