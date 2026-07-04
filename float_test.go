package opt

import (
	"encoding/json"
	"math"
	"testing"
)

func TestFloatFrom(t *testing.T) {
	f := FloatFrom(3.14)
	if f.V != 3.14 || !f.Valid {
		t.Errorf("FloatFrom(3.14): got {%v, %v}", f.V, f.Valid)
	}

	zero := FloatFrom(0)
	if !zero.Valid {
		t.Error("FloatFrom(0) should be valid")
	}
}

func TestFloatMarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		val     Float
		want    string
		wantErr bool
	}{
		{"positive", FloatFrom(3.14), "3.14", false},
		{"zero", FloatFrom(0), "0", false},
		{"negative", FloatFrom(-1.5), "-1.5", false},
		{"null", NewFloat(0, false), "null", false},
		{"inf", NewFloat(math.Inf(1), true), "", true},
		{"nan", NewFloat(math.NaN(), true), "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data, err := json.Marshal(tt.val)
			if (err != nil) != tt.wantErr {
				t.Fatalf("err = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr && string(data) != tt.want {
				t.Errorf("got %s, want %s", data, tt.want)
			}
		})
	}
}

func TestFloatUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    float64
		valid   bool
		wantErr bool
	}{
		{"number", "3.14", 3.14, true, false},
		{"integer", "42", 42, true, false},
		{"zero", "0", 0, true, false},
		{"string number", `"3.14"`, 3.14, true, false},
		{"empty string", `""`, 0, false, false},
		{"null", "null", 0, false, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var f Float
			err := json.Unmarshal([]byte(tt.input), &f)
			if (err != nil) != tt.wantErr {
				t.Fatalf("err = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr {
				if f.Valid != tt.valid || (f.Valid && f.V != tt.want) {
					t.Errorf("got {%v, %v}, want {%v, %v}", f.V, f.Valid, tt.want, tt.valid)
				}
			}
		})
	}
}

func TestFloatEqual(t *testing.T) {
	if !FloatFrom(3.14).Equal(FloatFrom(3.14)) {
		t.Error("equal floats should be Equal")
	}
	if FloatFrom(3.14).Equal(FloatFrom(2.71)) {
		t.Error("different floats should not be Equal")
	}
	if !NewFloat(0, false).Equal(NewFloat(0, false)) {
		t.Error("two null floats should be Equal")
	}
}

func TestFloatJSONRoundtrip(t *testing.T) {
	type measurement struct {
		Temp     Float `json:"temp"`
		Humidity Float `json:"humidity"`
	}

	original := measurement{Temp: FloatFrom(23.5), Humidity: NewFloat(0, false)}
	data, err := json.Marshal(original)
	if err != nil {
		t.Fatal(err)
	}

	want := `{"temp":23.5,"humidity":null}`
	if string(data) != want {
		t.Errorf("marshal: got %s, want %s", data, want)
	}

	var decoded measurement
	if err := json.Unmarshal(data, &decoded); err != nil {
		t.Fatal(err)
	}
	if !decoded.Temp.Equal(original.Temp) || !decoded.Humidity.Equal(original.Humidity) {
		t.Error("roundtrip failed")
	}
}
