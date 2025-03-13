package test

import (
	"github.com/behavioral-ai/collective/content"
	"github.com/behavioral-ai/collective/test"
	"github.com/behavioral-ai/collective/testrsc"
	"github.com/behavioral-ai/core/messaging"
	"github.com/behavioral-ai/operative/urn"
	"net/url"
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
	//buf, err := iox.ReadFile(testrsc.ResiliencyInterpret1)
	//if err != nil {
	//	return messaging.NewStatusError(messaging.StatusIOError, err, "")
	//}
	u, _ := url.Parse(testrsc.ResiliencyInterpret1)
	status := r.PutValue(urn.ResiliencyInterpret, "author", u, 1)
	if !status.OK() {
		return status
	}
	//buf, err = iox.ReadFile(testrsc.ResiliencyThreshold1)
	//if err != nil {
	//	return messaging.NewStatusError(messaging.StatusIOError, err, "")
	//}
	u, _ = url.Parse(testrsc.ResiliencyThreshold1)
	return r.PutValue(urn.ResiliencyThreshold, "author", u, 1)
}
