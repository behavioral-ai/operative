package frame1

const (
	PkgPath = "github/behavioral-ai/operative/frame1"
)

type Observation interface {
	Gradiant() int
	Latency() int
}

// IFrame - frame interface
type IFrame struct {
	Reason func(o Observation) error
}

var Frame = func() *IFrame {
	return &IFrame{
		Reason: func(o Observation) error {
			return nil
		},
	}
}()
