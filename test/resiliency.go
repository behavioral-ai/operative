package test

import (
	"github.com/behavioral-ai/collective/content"
	"github.com/behavioral-ai/core/messaging"
	"github.com/behavioral-ai/operative/testrsc"
	"github.com/behavioral-ai/operative/urn"
	"net/url"
)

func NewResiliencyResolver() (content.Resolution, *messaging.Status) {
	r := content.NewEphemeralResolver()
	status := loadResiliencyContent(r)
	return r, status
}

func loadResiliencyContent(r content.Resolution) *messaging.Status {
	u, _ := url.Parse(testrsc.ResiliencyTrafficProfile1)
	status := r.PutValue(urn.ResiliencyTrafficProfile, "author", u, 1)
	if !status.OK() {
		return status
	}
	u, _ = url.Parse(testrsc.ResiliencyInterpret1)
	status = r.PutValue(urn.ResiliencyInterpret, "author", u, 1)
	if !status.OK() {
		return status
	}
	u, _ = url.Parse(testrsc.ResiliencyThreshold1)
	return r.PutValue(urn.ResiliencyThreshold, "author", u, 1)
}
