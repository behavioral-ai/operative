package agent1

import (
	"fmt"
	"github.com/behavioral-ai/core/core"
)

func ExampleNewAgent() {
	a := New(core.Origin{Region: "us-central"}, nil, nil)

	fmt.Printf("test: NewAgent() -> [uri:%v]\n", a)

	//Output:
	//test: NewAgent() -> [uri:agent1:us-central..]

}
