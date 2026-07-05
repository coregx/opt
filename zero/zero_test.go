package zero

import (
	"encoding/json"
	"testing"
	"time"
)

// --- String ---

func TestStringFrom(t *testing.T) {
	s := StringFrom("hello")
	if s.V != "hello" || !s.Valid {
		t.Errorf("StringFrom(\"hello\"): got {%v, %v}", s.V, s.Valid)
	}

	empty := StringFrom("")
	if empty.Valid {
		t.Error("StringFrom(\"\") should be invalid (zero = null)")
	}
}

func TestStringMarshalJSON(t *testing.T) {
	tests := []struct {
		name string
		val  String
		want string
	}{
		{"valid", StringFrom("hello"), `"hello"`},
		{"null", NewString("", false), `""`},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data, err := json.Marshal(tt.val)
			if err != nil {
				t.Fatal(err)
			}
			if string(data) != tt.want {
				t.Errorf("got %s, want %s", data, tt.want)
			}
		})
	}
}

func TestStringUnmarshalJSON(t *testing.T) {
	var s String
	json.Unmarshal([]byte(`"hello"`), &s)
	if s.V != "hello" || !s.Valid {
		t.Error("should be valid hello")
	}

	var empty String
	json.Unmarshal([]byte(`""`), &empty)
	if empty.Valid {
		t.Error("empty string should be invalid")
	}

	var null String
	json.Unmarshal([]byte("null"), &null)
	if null.Valid {
		t.Error("null should be invalid")
	}
}

// --- Int ---

func TestIntFrom(t *testing.T) {
	i := IntFrom(42)
	if i.V != 42 || !i.Valid {
		t.Errorf("IntFrom(42): got {%v, %v}", i.V, i.Valid)
	}

	zero := IntFrom(0)
	if zero.Valid {
		t.Error("IntFrom(0) should be invalid (zero = null)")
	}
}

func TestIntMarshalJSON(t *testing.T) {
	tests := []struct {
		name string
		val  Int
		want string
	}{
		{"valid", IntFrom(42), "42"},
		{"null", NewInt(0, false), "0"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data, err := json.Marshal(tt.val)
			if err != nil {
				t.Fatal(err)
			}
			if string(data) != tt.want {
				t.Errorf("got %s, want %s", data, tt.want)
			}
		})
	}
}

func TestIntUnmarshalJSON(t *testing.T) {
	var i Int
	json.Unmarshal([]byte("42"), &i)
	if i.V != 42 || !i.Valid {
		t.Error("should be valid 42")
	}

	var zero Int
	json.Unmarshal([]byte("0"), &zero)
	if zero.Valid {
		t.Error("0 should be invalid (zero = null)")
	}

	var str Int
	json.Unmarshal([]byte(`"42"`), &str)
	if str.V != 42 || !str.Valid {
		t.Error("string 42 should be valid")
	}
}

// --- Bool ---

func TestBoolFrom(t *testing.T) {
	b := BoolFrom(true)
	if !b.V || !b.Valid {
		t.Error("BoolFrom(true) should be valid")
	}

	f := BoolFrom(false)
	if f.Valid {
		t.Error("BoolFrom(false) should be invalid (zero = null)")
	}
}

func TestBoolMarshalJSON(t *testing.T) {
	b := BoolFrom(true)
	data, _ := json.Marshal(b)
	if string(data) != "true" {
		t.Errorf("got %s, want true", data)
	}

	null := NewBool(false, false)
	data, _ = json.Marshal(null)
	if string(data) != "false" {
		t.Errorf("null marshal: got %s, want false", data)
	}
}

// --- Float ---

func TestFloatFrom(t *testing.T) {
	f := FloatFrom(3.14)
	if f.V != 3.14 || !f.Valid {
		t.Error("FloatFrom(3.14) should be valid")
	}

	zero := FloatFrom(0)
	if zero.Valid {
		t.Error("FloatFrom(0) should be invalid (zero = null)")
	}
}

func TestFloatMarshalJSON(t *testing.T) {
	f := FloatFrom(3.14)
	data, _ := json.Marshal(f)
	if string(data) != "3.14" {
		t.Errorf("got %s", data)
	}

	null := NewFloat(0, false)
	data, _ = json.Marshal(null)
	if string(data) != "0" {
		t.Errorf("null marshal: got %s, want 0", data)
	}
}

// --- Time ---

func TestTimeFrom(t *testing.T) {
	now := time.Now()
	tv := TimeFrom(now)
	if !tv.Valid {
		t.Error("TimeFrom(now) should be valid")
	}

	zero := TimeFrom(time.Time{})
	if zero.Valid {
		t.Error("TimeFrom(zero) should be invalid")
	}
}

func TestTimeMarshalJSON(t *testing.T) {
	ts := time.Date(2026, 7, 4, 12, 0, 0, 0, time.UTC)
	tv := TimeFrom(ts)
	data, _ := json.Marshal(tv)
	if string(data) != `"2026-07-04T12:00:00Z"` {
		t.Errorf("got %s", data)
	}
}

// --- Generic Option ---

func TestOptionFrom(t *testing.T) {
	v := From(42)
	if v.V != 42 || !v.Valid {
		t.Error("From(42) should be valid")
	}

	zero := From(0)
	if zero.Valid {
		t.Error("From(0) should be invalid (zero = null)")
	}
}

func TestOptionMarshalJSON(t *testing.T) {
	v := From(42)
	data, _ := json.Marshal(v)
	if string(data) != "42" {
		t.Errorf("got %s", data)
	}

	null := New(0, false)
	data, _ = json.Marshal(null)
	if string(data) != "0" {
		t.Errorf("null marshal: got %s, want 0", data)
	}
}

// --- Semantic difference test ---

func TestZeroVsOptSemantics(t *testing.T) {
	type zeroModel struct {
		Name  String `json:"name"`
		Count Int    `json:"count"`
	}

	input := `{"name":"","count":0}`
	var m zeroModel
	json.Unmarshal([]byte(input), &m)

	if m.Name.Valid {
		t.Error("zero/String: empty string should be invalid")
	}
	if m.Count.Valid {
		t.Error("zero/Int: 0 should be invalid")
	}

	data, _ := json.Marshal(m)
	if string(data) != `{"name":"","count":0}` {
		t.Errorf("zero marshal: got %s, want empty/zero values", data)
	}
}
