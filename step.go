package asyncstep

import "time"

// Step is the common execution unit for asynchronous work.
//
// Meta holds application-specific data. Keep Step free of product-specific
// fields so the same model can be reused by other pipelines.
type Step[M any] struct {
	ID         int64      `json:"id"`
	Node       string     `json:"node"`
	Seq        int        `json:"seq"`
	State      State      `json:"state"`
	Attempt    int        `json:"attempt"`
	Error      string     `json:"error,omitempty"`
	StartedAt  time.Time  `json:"startedAt"`
	FinishedAt *time.Time `json:"finishedAt,omitempty"`
	Meta       M          `json:"meta,omitempty"`
}

// AnyStep is Step with untyped metadata.
type AnyStep = Step[any]
