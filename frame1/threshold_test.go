package frame1

import (
	"encoding/json"
	"fmt"
)

func ExampleThreshold() {
	t := threshold{
		Latency:    2000,
		Saturation: lookup{Low: 20, Medium: 40, High: 60},
		Gradient:   lookup{Low: 20, Medium: 50, High: 85},
	}

	buf, err := json.Marshal(t)
	fmt.Printf("test: Threshold() -> [%v] [%v]\n", err, string(buf))

	//Output:
	//fail
}
