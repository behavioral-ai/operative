package testrsc

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/behavioral-ai/core/iox"
	"github.com/behavioral-ai/domain/collective"
	"github.com/behavioral-ai/operative/urn"
)

const (
	ResiliencyThreshold1 = "file:///f:/files/resiliency/threshold-1.json"
	ResiliencyInterpret1 = "file:///f:/files/resiliency/interpret-1.json"

	ResiliencyThreshold2 = "file:///f:/files/resiliency/threshold-2.json"
	ResiliencyInterpret2 = "file:///f:/files/resiliency/interpret-2.json"
)

// Resolver -
var Resolver = func() *collective.IResolver {
	return &collective.IResolver{
		Get: func(name string, version int) ([]byte, error) {
			return iox.ReadFile(fname(name, version))
		},
		GetRelated: func(name string, version int) ([]byte, error) {
			return nil, nil
		},
		Append: func(name string, content any, version int) error {
			var buf []byte
			if name == "" || content == nil || version <= 0 {
				return errors.New(fmt.Sprintf("error: invalid argument name %v content %v version %v", name, content, version))
			}
			switch ptr := content.(type) {
			case string:
				buf = []byte(ptr)
			case []byte:
				buf = ptr
			default:
				var err error

				buf, err = json.Marshal(ptr)
				if err != nil {
					return err
				}
			}
			return storeAppend(name, buf, version)
		},
	}
}()

// Get - generic typed get
func Get[T any](name string, version int, resolver *collective.IResolver) (T, error) {
	var t T

	body, status := resolver.Get(name, version)
	if status != nil {
		return t, status
	}
	switch ptr := any(&t).(type) {
	case *string:
		*ptr = string(body)
	case *[]byte:
		*ptr = body
	default:
		err := json.Unmarshal(body, ptr)
		if err != nil {
			return t, errors.New(fmt.Sprintf("error: JsonEncode %v", err))
		}
	}
	return t, nil
}

func fname(name string, version int) string {
	switch name {
	case urn.ResiliencyInterpret:
		if version == 1 {
			return ResiliencyInterpret1
		} else {
			return ResiliencyInterpret2
		}
	case urn.ResiliencyThreshold:
		if version == 1 {
			return ResiliencyThreshold1
		} else {
			return ResiliencyThreshold2
		}
	}
	return name
}
