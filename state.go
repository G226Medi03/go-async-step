package asyncstep

// State is the execution state of a step.
type State string

const (
	Pending   State = "pending"
	Running   State = "running"
	Succeeded State = "succeeded"
	Failed    State = "failed"
	Skipped   State = "skipped"
)
