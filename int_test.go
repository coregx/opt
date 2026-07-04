package opt

import (
	"encoding/json"
	"testing"
)

func TestIntFrom(t *testing.T) {
	i := IntFrom(42)
	if i.V != 42 || !i.Valid {
		t.Errorf("IntFrom(42): got {%v, %v}, want {42, true}", i.V, i.Valid)
	}

	zero := IntFrom(0)
	if !zero.Valid {
		t.Error("IntFrom(0) should be valid")
	}
}

func TestIntFromPtr(t *testing.T) {
	n := int64(42)
	i := IntFromPtr(&n)
	if i.V != 42 || !i.Valid {
		t.Errorf("IntFromPtr: got {%v, %v}", i.V, i.Valid)
	}

	null := IntFromPtr(nil)
	if null.Valid {
		t.Error("IntFromPtr(nil) should be invalid")
	}
}

func TestIntMarshalJSON(t *testing.T) {
	tests := []struct {
		name string
		val  Int
		want string
	}{
		{"positive", IntFrom(42), "42"},
		{"zero", IntFrom(0), "0"},
		{"negative", IntFrom(-1), "-1"},
		{"null", NewInt(0, false), "null"},
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
	tests := []struct {
		name    string
		input   string
		want    int64
		valid   bool
		wantErr bool
	}{
		{"number", "42", 42, true, false},
		{"zero", "0", 0, true, false},
		{"negative", "-1", -1, true, false},
		{"string number", `"42"`, 42, true, false},
		{"empty string", `""`, 0, false, false},
		{"null", "null", 0, false, false},
		{"bad string", `"abc"`, 0, false, true},
		{"bool", "true", 0, false, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var i Int
			err := json.Unmarshal([]byte(tt.input), &i)
			if (err != nil) != tt.wantErr {
				t.Fatalf("err = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr {
				if i.Valid != tt.valid || (i.Valid && i.V != tt.want) {
					t.Errorf("got {%v, %v}, want {%v, %v}", i.V, i.Valid, tt.want, tt.valid)
				}
			}
		})
	}
}

func TestIntMarshalText(t *testing.T) {
	i := IntFrom(42)
	data, err := i.MarshalText()
	if err != nil {
		t.Fatal(err)
	}
	if string(data) != "42" {
		t.Errorf("got %s, want 42", data)
	}

	null := NewInt(0, false)
	data, _ = null.MarshalText()
	if len(data) != 0 {
		t.Errorf("null MarshalText should return empty, got %s", data)
	}
}

func TestIntUnmarshalText(t *testing.T) {
	var i Int
	if err := i.UnmarshalText([]byte("42")); err != nil {
		t.Fatal(err)
	}
	if i.V != 42 || !i.Valid {
		t.Errorf("got {%v, %v}, want {42, true}", i.V, i.Valid)
	}

	var null Int
	if err := null.UnmarshalText([]byte("")); err != nil {
		t.Fatal(err)
	}
	if null.Valid {
		t.Error("empty text should produce invalid Int")
	}

	var nullText Int
	if err := nullText.UnmarshalText([]byte("null")); err != nil {
		t.Fatal(err)
	}
	if nullText.Valid {
		t.Error("\"null\" text should produce invalid Int")
	}
}

func TestIntEqual(t *testing.T) {
	if !IntFrom(42).Equal(IntFrom(42)) {
		t.Error("equal ints should be Equal")
	}
	if IntFrom(42).Equal(IntFrom(99)) {
		t.Error("different ints should not be Equal")
	}
	if !NewInt(0, false).Equal(NewInt(0, false)) {
		t.Error("two null ints should be Equal")
	}
}

func TestIntJSONRoundtrip(t *testing.T) {
	type item struct {
		Count Int `json:"count"`
		Price Int `json:"price"`
	}

	original := item{Count: IntFrom(5), Price: NewInt(0, false)}
	data, err := json.Marshal(original)
	if err != nil {
		t.Fatal(err)
	}

	want := `{"count":5,"price":null}`
	if string(data) != want {
		t.Errorf("marshal: got %s, want %s", data, want)
	}

	var decoded item
	if err := json.Unmarshal(data, &decoded); err != nil {
		t.Fatal(err)
	}
	if !decoded.Count.Equal(original.Count) || !decoded.Price.Equal(original.Price) {
		t.Error("roundtrip failed")
	}
}
