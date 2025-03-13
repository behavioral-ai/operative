package test

import (
	"fmt"
	"github.com/behavioral-ai/domain/metrics1"
	"github.com/behavioral-ai/operative/urn"
)

func ExampleNewResiliencyResolver() {
	r, status := NewResiliencyResolver()

	fmt.Printf("test: NewResiliencyResolver() -> [status:%v]\n", status)

	name := urn.ResiliencyThreshold
	buf, status1 := r.GetValue(name, 1)
	fmt.Printf("test: GetValue(\"%v\") -> [status:%v] [buf:%v]\n", name, status1, len(buf))

	name = urn.ResiliencyInterpret
	buf, status1 = r.GetValue(name, 1)
	fmt.Printf("test: GetValue(\"%v\") -> [status:%v] [buf:%v]\n", name, status1, len(buf))

	name = metrics1.ProfileName
	buf, status1 = r.GetValue(name, 1)
	fmt.Printf("test: GetValue(\"%v\") -> [status:%v] [buf:%v]\n", name, status1, len(buf))

	//Output:
	//test: NewResiliencyResolver() -> [status:OK]
	//test: GetValue("resiliency:type/operative/agent/threshold") -> [status:OK] [buf:303]
	//test: GetValue("resiliency:type/operative/agent/interpret") -> [status:OK] [buf:316]
	//test: GetValue("resiliency:type/domain/metrics/profile") -> [status:OK] [buf:1218]

}
