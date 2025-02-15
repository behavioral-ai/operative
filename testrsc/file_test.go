package testrsc

import (
	"encoding/json"
	"fmt"
)

type Lookup struct {
	Low    int `json:"low"`
	Medium int `json:"medium"`
	High   int `json:"high"`
}

func _ExampleLookup() {
	l := Lookup{Low: 10, Medium: 40, High: 80}
	buf, err := json.Marshal(l)
	fmt.Printf("test: Lookup() -> [%v] [%v]\n", err, string(buf))

	//Output:
	//fail
}
