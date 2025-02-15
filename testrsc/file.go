package testrsc

import (
	"errors"
	"fmt"
	"github.com/behavioral-ai/core/iox"
	"sync"
)

func fileResolver(name string, version int) (buf []byte, err error) {
	buf, err = storeGet(name, version)
	if err == nil {
		return
	}
	return iox.ReadFile(fname(name, version))
}

// TODO: need to determine a partitioning scheme, append only
type entry struct {
	Name      string `json:"name"`    // Uuid
	Version   int    `json:"version"` // Semantic versioning MAJOR component only
	Content   []byte `json:"content"`
	CreatedTS string `json:"created-ts"`
}

var (
	sm    sync.Mutex
	store []entry
)

func storeAppend(name string, body []byte, version int) error {
	if storeExists(name, version) {
		return errors.New("error: bad request")
	}
	sm.Lock()
	defer sm.Unlock()
	store = append(store, entry{Name: name, Content: body, Version: version, CreatedTS: "2024-02-12"})
	return nil
}

func storeExists(name string, version int) bool {
	sm.Lock()
	defer sm.Unlock()
	for _, item := range store {
		if item.Name == name && item.Version == version {
			return true
		}
	}
	return false
}

func storeGet(name string, version int) ([]byte, error) {
	sm.Lock()
	defer sm.Unlock()
	for _, item := range store {
		if item.Name == name && item.Version == version {
			return item.Content, nil
		}
	}
	return nil, errors.New(fmt.Sprintf("error: name %v not found", name))
}
