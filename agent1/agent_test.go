package agent1

import (
	"fmt"
	"github.com/behavioral-ai/domain/common"
)

func ExampleNewAgent() {
	a := New(nil, common.Origin{Region: "us-central", Zone: "c-zone-a", SubZone: "sub-zone", Host: "www.host.com"}, nil, nil)

	fmt.Printf("test: NewAgent() -> [%v]\n", a)
	fmt.Printf("test: NewAgent() -> [%v]\n", a.Name())

	//Output:
	//test: NewAgent() -> [resiliency:agent/operative/agent1#us-central.c-zone-a.sub-zone.www.host.com]
	//test: NewAgent() -> [resiliency:agent/operative/agent]

}
