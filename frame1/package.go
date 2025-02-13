package frame1

import (
	"context"
	"fmt"
	"github.com/behavioral-ai/domain/common"
)

const (
	PkgPath = "github/behavioral-ai/ingress/frame1"
)

func Observe(ctx context.Context, origin common.Origin) (Observation, error) {
	return newObservation(), nil
}

func AddAction(ctx context.Context, origin common.Origin, action int) error {
	fmt.Printf("Action: %v %v\n", origin, action)
	return nil
}

func AddInference(ctx context.Context, origin common.Origin, f *Frame) error {
	if f != nil {
		addFrame(f)
		fmt.Printf("Frame: %v %v\n", origin, f.Observe)
	}
	return nil
}

func Inference(o Observation) *Frame {
	f := NewFrame(o)
	comprehension(f)
	reasoning(f)
	return f
}

func AddFeedback(ctx context.Context) {
	updateFeedback()
}
