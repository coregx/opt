package opt

import (
	"encoding/json"
	"testing"
)

func TestByteFrom(t *testing.T) {
	b := ByteFrom(42)
	if b.V != 42 || !b.Valid {
		t.Errorf("ByteFrom(42): got {%v, %v}", b.V, b.Valid)
	}
}

func TestByteMarshalJSON(t *testing.T) {
	tests := []struct {
		name string
		val  Byte
		want string
	}{
		{"value", ByteFrom(42), "42"},
		{"zero", ByteFrom(0), "0"},
		{"null", NewByte(0, false), "null"},
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

func TestByteUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    byte
		valid   bool
		wantErr bool
	}{
		{"number", "42", 42, true, false},
		{"max", "255", 255, true, false},
		{"zero", "0", 0, true, false},
		{"null", "null", 0, false, false},
		{"overflow", "256", 0, false, true},
		{"negative", "-1", 0, false, true},
		{"large", "999", 0, false, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var b Byte
			err := json.Unmarshal([]byte(tt.input), &b)
			if (err != nil) != tt.wantErr {
				t.Fatalf("err = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr {
				if b.Valid != tt.valid || (b.Valid && b.V != tt.want) {
					t.Errorf("got {%v, %v}, want {%v, %v}", b.V, b.Valid, tt.want, tt.valid)
				}
			}
		})
	}
}

func TestByteEqual(t *testing.T) {
	if !ByteFrom(42).Equal(ByteFrom(42)) {
		t.Error("equal values should be Equal")
	}
	if ByteFrom(42).Equal(ByteFrom(99)) {
		t.Error("different values should not be Equal")
	}
}

func TestByteTextRoundtrip(t *testing.T) {
	b := ByteFrom(65) // 'A'
	data, err := b.MarshalText()
	if err != nil {
		t.Fatal(err)
	}
	if string(data) != "A" {
		t.Errorf("got %s, want A", data)
	}

	var parsed Byte
	if err := parsed.UnmarshalText(data); err != nil {
		t.Fatal(err)
	}
	if !parsed.Equal(b) {
		t.Error("text roundtrip failed")
	}
}
