package agent1

import (
	"fmt"
	"github.com/behavioral-ai/domain/common"
)

func ExampleNewAgent() {
	a := New(common.Origin{Region: "us-central"}, nil, nil)

	fmt.Printf("test: NewAgent() -> [uri:%v]\n", a)

	//Output:
	//test: NewAgent() -> [uri:agent1:us-central..]

}
