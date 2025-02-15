package testrsc

import (
	"fmt"
	"github.com/behavioral-ai/operative/urn"
)

func ExampleResolveFailure() {
	_, err := Resolver.Get("invalid urn", 1)
	if err != nil {
		fmt.Printf("test: Resolver() -> [err:%v]\n", err)
	}

	//Output:
	//test: Resolver() -> [err:open error: scheme is invalid []: The system cannot find the file specified.]

}

func ExampleResolveSuccess() {
	buf, err := Resolver.Get(urn.ResiliencyThreshold, 1)
	if err != nil {
		fmt.Printf("test: Resolver() -> [err:%v]\n", err)
	} else {

		fmt.Printf("test: Resolver() -> [%v]\n", string(buf))
	}

	buf, err = Resolver.Get(urn.ResiliencyThreshold, 2)
	if err != nil {
		fmt.Printf("test: Resolver() -> [err:%v]\n", err)
	} else {
		fmt.Printf("test: Resolver() -> [%v]\n", string(buf))
	}

	//Output:
	//test: Resolver() -> [err:<nil>] [[{10 40 80}]]

}
