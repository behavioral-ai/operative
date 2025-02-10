package knowledge1

import "github.com/behavioral-ai/core/aspect"

const (
	PkgPath = "github/behavioral-ai/ingress/knowledge1"
)

// IRetrieval - retrieval interface, a nod to Linus Torvalds and plain C
type IRetrieval struct {
	QueryThreshold func(h aspect.ErrorHandler) (Threshold, *aspect.Status)
	//SearchThings func(h core.ErrorHandler, tags Uri, created string) ([]SearchResponse, *core.Status)

	//QueryRelations  func(h core.ErrorHandler, t Triple) ([]Relation, *core.Status)
	//SearchRelations func(h core.ErrorHandler, tags Uri, created string) ([]SearchResponse, *core.Status)

	//QueryFrames  func(h core.ErrorHandler, name Urn) ([]Frame, *core.Status)
	//SearchFrames func(h core.ErrorHandler, tags Uri, created string) ([]SearchResponse, *core.Status)

	//Search func(h core.ErrorHandler, tags Uri, created string) ([]SearchResponse, *core.Status)
}

var Retrieval = func() *IRetrieval {
	return &IRetrieval{
		QueryThreshold: func(h aspect.ErrorHandler) (Threshold, *aspect.Status) {
			return Threshold{}, aspect.StatusOK()
		},
		/*
			SearchThings: func(h core.ErrorHandler, tags Uri, created string) ([]SearchResponse, *core.Status) {
				return nil, core.StatusOK()
			},
			QueryRelations: func(h core.ErrorHandler, t Triple) ([]Relation, *core.Status) {
				return nil, core.StatusOK()
			},
			SearchRelations: func(h core.ErrorHandler, tags Uri, created string) ([]SearchResponse, *core.Status) {
				return nil, core.StatusOK()
			},
			QueryFrames: func(h core.ErrorHandler, name Urn) ([]Frame, *core.Status) {
				return nil, core.StatusOK()
			},
			SearchFrames: func(h core.ErrorHandler, tags Uri, created string) ([]SearchResponse, *core.Status) {
				return nil, core.StatusOK()
			},
			Search: func(h core.ErrorHandler, tags Uri, created string) ([]SearchResponse, *core.Status) {
				return nil, core.StatusOK()
			},

		*/
	}
}()
