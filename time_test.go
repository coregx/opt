package opt

import (
	"encoding/json"
	"testing"
	"time"
)

func TestTimeFrom(t *testing.T) {
	now := time.Now()
	tv := TimeFrom(now)
	if !tv.V.Equal(now) || !tv.Valid {
		t.Error("TimeFrom should produce valid Time")
	}
}

func TestTimeFromPtr(t *testing.T) {
	now := time.Now()
	tv := TimeFromPtr(&now)
	if !tv.V.Equal(now) || !tv.Valid {
		t.Error("TimeFromPtr should produce valid Time")
	}

	null := TimeFromPtr(nil)
	if null.Valid {
		t.Error("TimeFromPtr(nil) should be invalid")
	}
}

func TestTimeMarshalJSON(t *testing.T) {
	ts := time.Date(2026, 7, 4, 12, 0, 0, 0, time.UTC)
	tv := TimeFrom(ts)

	data, err := json.Marshal(tv)
	if err != nil {
		t.Fatal(err)
	}
	if string(data) != `"2026-07-04T12:00:00Z"` {
		t.Errorf("got %s, want \"2026-07-04T12:00:00Z\"", data)
	}

	null := NewTime(time.Time{}, false)
	data, err = json.Marshal(null)
	if err != nil {
		t.Fatal(err)
	}
	if string(data) != "null" {
		t.Errorf("null marshal: got %s, want null", data)
	}
}

func TestTimeUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		valid   bool
		wantErr bool
	}{
		{"valid", `"2026-07-04T12:00:00Z"`, true, false},
		{"null", "null", false, false},
		{"number", "42", false, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var tv Time
			err := json.Unmarshal([]byte(tt.input), &tv)
			if (err != nil) != tt.wantErr {
				t.Fatalf("err = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr && tv.Valid != tt.valid {
				t.Errorf("Valid = %v, want %v", tv.Valid, tt.valid)
			}
		})
	}
}

func TestTimeEqual(t *testing.T) {
	ts := time.Date(2026, 7, 4, 12, 0, 0, 0, time.UTC)
	a := TimeFrom(ts)
	b := TimeFrom(ts)
	if !a.Equal(b) {
		t.Error("same times should be Equal")
	}

	// Same instant, different timezone — Equal uses time.Equal (compares instants)
	c := TimeFrom(ts.In(time.FixedZone("UTC+5", 5*60*60)))
	if !a.Equal(c) {
		t.Error("Equal should match same instant in different timezone")
	}

	// ExactEqual distinguishes timezone
	if a.ExactEqual(c) {
		t.Error("ExactEqual should distinguish timezone")
	}

	null1 := NewTime(time.Time{}, false)
	null2 := NewTime(time.Time{}, false)
	if !null1.Equal(null2) {
		t.Error("two null times should be Equal")
	}
}

func TestTimeExactEqual(t *testing.T) {
	ts := time.Date(2026, 7, 4, 12, 0, 0, 0, time.UTC)
	a := TimeFrom(ts)
	b := TimeFrom(ts)
	if !a.ExactEqual(b) {
		t.Error("same times should be ExactEqual")
	}
}

func TestTimeMarshalText(t *testing.T) {
	ts := time.Date(2026, 7, 4, 12, 0, 0, 0, time.UTC)
	tv := TimeFrom(ts)

	data, err := tv.MarshalText()
	if err != nil {
		t.Fatal(err)
	}
	if string(data) != "2026-07-04T12:00:00Z" {
		t.Errorf("got %s", data)
	}

	null := NewTime(time.Time{}, false)
	data, _ = null.MarshalText()
	if len(data) != 0 {
		t.Errorf("null MarshalText should return empty, got %s", data)
	}
}

func TestTimeJSONRoundtrip(t *testing.T) {
	type event struct {
		Start Time `json:"start"`
		End   Time `json:"end"`
	}

	ts := time.Date(2026, 7, 4, 12, 0, 0, 0, time.UTC)
	original := event{Start: TimeFrom(ts), End: NewTime(time.Time{}, false)}

	data, err := json.Marshal(original)
	if err != nil {
		t.Fatal(err)
	}

	want := `{"start":"2026-07-04T12:00:00Z","end":null}`
	if string(data) != want {
		t.Errorf("marshal: got %s, want %s", data, want)
	}

	var decoded event
	if err := json.Unmarshal(data, &decoded); err != nil {
		t.Fatal(err)
	}
	if !decoded.Start.Equal(original.Start) || !decoded.End.Equal(original.End) {
		t.Error("roundtrip failed")
	}
}
