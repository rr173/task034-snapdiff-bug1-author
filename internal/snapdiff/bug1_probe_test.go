package snapdiff

import (
	"encoding/json"
	"testing"
)

func TestBug1AdditionOnlyKeepsEmptyRemovedSlice(t *testing.T) {
	a, err := NewSnapshot(nil)
	if err != nil {
		t.Fatal(err)
	}
	b, err := NewSnapshot([]FileInput{{Path: "new.txt", Content: "new"}})
	if err != nil {
		t.Fatal(err)
	}
	report := Diff(a, b)
	if report.Removed == nil {
		t.Fatal("addition-only diff must expose removed as an empty slice")
	}
	payload, err := json.Marshal(report)
	if err != nil {
		t.Fatal(err)
	}
	if string(payload) == "" {
		t.Fatal("diff report must be JSON encodable")
	}
}
