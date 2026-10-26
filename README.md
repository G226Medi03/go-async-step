# go-async-step

Common asynchronous step state and metadata model for Go.

```go
type SolveMeta struct {
    CaseID string `json:"caseId"`
}

step := asyncstep.Step[SolveMeta]{
    ID:    1,
    Node:  "extract",
    Seq:   20,
    State: asyncstep.Running,
}
```

`Step` contains only execution-common fields. Application-specific values belong
in the generic `Meta` field.
