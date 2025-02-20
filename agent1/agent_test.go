package agent1

import (
	"fmt"
	"github.com/behavioral-ai/domain/common"
)

func ExampleNewAgent() {
	a := New(common.Origin{Region: "us-central", Zone: "zone1", SubZone: "sub-zone", Host: "www.host.com"}, nil, nil)

	fmt.Printf("test: NewAgent() -> [%v]\n", a)

	//Output:
	//test: NewAgent() -> [uri:agent1:us-central..]

}
