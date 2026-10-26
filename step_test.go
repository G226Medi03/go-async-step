package asyncstep

import (
	"encoding/json"
	"testing"
	"time"
)

func TestStepJSON(t *testing.T) {
	started := time.Date(2026, 1, 2, 3, 4, 5, 0, time.UTC)
	step := Step[map[string]string]{
		ID:        7,
		Node:      "extract",
		Seq:       20,
		State:     Running,
		Attempt:   2,
		StartedAt: started,
		Meta:      map[string]string{"owner": "case_1"},
	}
	encoded, err := json.Marshal(step)
	if err != nil {
		t.Fatalf("Marshal: %v", err)
	}
	var decoded Step[map[string]string]
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		t.Fatalf("Unmarshal: %v", err)
	}
	if decoded.State != Running || decoded.Meta["owner"] != "case_1" {
		t.Fatalf("decoded = %#v", decoded)
	}
}
