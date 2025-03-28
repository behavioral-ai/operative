package frame1

import (
	"encoding/json"
	"fmt"
)

func _ExampleInterpret() {
	t := interpret{LowLow: 0, LowMedium: 0, MediumMedium: 10, MediumHigh: 20, HighHigh: 30}
	buf, err := json.Marshal(t)
	fmt.Printf("test: interpret() -> [%v] [%v]\n", err, string(buf))

	//Output:
	//fail
}

func _ExampleNewInterpret() {
	i, err := newInterpret(1, nil)

	fmt.Printf("test: newInterpret() -> [%v] [%v]\n", err, i)

	//Output:
	//fail
}

func ExampleInterpretLowSaturation() {
	t := interpret{LowLow: 0, LowMedium: 10, LowHigh: 20}

	i := impression{Saturation: low, Gradient: low}
	action := t.action(i)
	fmt.Printf("test: Action() -> [%v] [%v]\n", i, action)

	i = impression{Saturation: low, Gradient: medium}
	action = t.action(i)
	fmt.Printf("test: Action() -> [%v] [%v]\n", i, action)

	i = impression{Saturation: low, Gradient: high}
	action = t.action(i)
	fmt.Printf("test: Action() -> [%v] [%v]\n", i, action)

	//Output:
	//test: Action() -> [{low low}] [0]
	//test: Action() -> [{low medium}] [10]
	//test: Action() -> [{low high}] [20]

}

func ExampleInterpretMediumSaturation() {
	t := interpret{MediumLow: 10, MediumMedium: 25, MediumHigh: 35}

	i := impression{Saturation: medium, Gradient: low}
	action := t.action(i)
	fmt.Printf("test: Action() -> [%v] [%v]\n", i, action)

	i = impression{Saturation: medium, Gradient: medium}
	action = t.action(i)
	fmt.Printf("test: Action() -> [%v] [%v]\n", i, action)

	i = impression{Saturation: medium, Gradient: high}
	action = t.action(i)
	fmt.Printf("test: Action() -> [%v] [%v]\n", i, action)

	//Output:
	//test: Action() -> [{medium low}] [10]
	//test: Action() -> [{medium medium}] [25]
	//test: Action() -> [{medium high}] [35]

}

func ExampleInterpretHighSaturation() {
	t := interpret{HighLow: 15, HighMedium: 35, HighHigh: 50}

	i := impression{Saturation: high, Gradient: low}
	action := t.action(i)
	fmt.Printf("test: Action() -> [%v] [%v]\n", i, action)

	i = impression{Saturation: high, Gradient: medium}
	action = t.action(i)
	fmt.Printf("test: Action() -> [%v] [%v]\n", i, action)

	i = impression{Saturation: high, Gradient: high}
	action = t.action(i)
	fmt.Printf("test: Action() -> [%v] [%v]\n", i, action)

	//Output:
	//test: Action() -> [{high low}] [15]
	//test: Action() -> [{high medium}] [35]
	//test: Action() -> [{high high}] [50]

}
