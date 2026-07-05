package opt

import (
	"encoding/json"
	"testing"
)

func TestFieldThreeStates(t *testing.T) {
	absent := Field[string]{}
	if !absent.IsAbsent() {
		t.Error("zero value should be Absent")
	}
	if absent.IsNull() || absent.IsValue() {
		t.Error("absent should not be Null or Value")
	}

	null := FieldNull[string]()
	if !null.IsNull() {
		t.Error("FieldNull should be Null")
	}
	if null.IsAbsent() || null.IsValue() {
		t.Error("null should not be Absent or Value")
	}

	value := FieldFrom("hello")
	if !value.IsValue() {
		t.Error("FieldFrom should be Value")
	}
	if value.IsAbsent() || value.IsNull() {
		t.Error("value should not be Absent or Null")
	}
}

func TestFieldPATCHSemantics(t *testing.T) {
	type PatchUser struct {
		Name  Field[string] `json:"name,omitzero"`
		Email Field[string] `json:"email,omitzero"`
		Age   Field[int]    `json:"age,omitzero"`
	}

	tests := []struct {
		name       string
		input      string
		nameState  string
		emailState string
		ageState   string
	}{
		{
			name:       "all absent",
			input:      `{}`,
			nameState:  "absent",
			emailState: "absent",
			ageState:   "absent",
		},
		{
			name:       "name set, email null, age absent",
			input:      `{"name":"John","email":null}`,
			nameState:  "value",
			emailState: "null",
			ageState:   "absent",
		},
		{
			name:       "all present",
			input:      `{"name":"John","email":"john@example.com","age":30}`,
			nameState:  "value",
			emailState: "value",
			ageState:   "value",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var u PatchUser
			if err := json.Unmarshal([]byte(tt.input), &u); err != nil {
				t.Fatal(err)
			}

			checkState := func(name string, f interface {
				IsAbsent() bool
				IsNull() bool
				IsValue() bool
			}, want string) {
				var got string
				switch {
				case f.IsAbsent():
					got = "absent"
				case f.IsNull():
					got = "null"
				case f.IsValue():
					got = "value"
				}
				if got != want {
					t.Errorf("%s: got %s, want %s", name, got, want)
				}
			}

			checkState("name", u.Name, tt.nameState)
			checkState("email", u.Email, tt.emailState)
			checkState("age", u.Age, tt.ageState)
		})
	}
}

func TestFieldMarshalJSON(t *testing.T) {
	type response struct {
		Name  Field[string] `json:"name,omitzero"`
		Email Field[string] `json:"email,omitzero"`
		Age   Field[int]    `json:"age,omitzero"`
	}

	r := response{
		Name:  FieldFrom("John"),
		Email: FieldNull[string](),
		// Age is absent (zero value)
	}

	data, err := json.Marshal(r)
	if err != nil {
		t.Fatal(err)
	}

	want := `{"name":"John","email":null}`
	if string(data) != want {
		t.Errorf("got %s, want %s", data, want)
	}
}

func TestFieldOr(t *testing.T) {
	v := FieldFrom("hello")
	if v.Or("default") != "hello" {
		t.Error("valid Or should return value")
	}

	null := FieldNull[string]()
	if null.Or("default") != "default" {
		t.Error("null Or should return fallback")
	}

	absent := Field[string]{}
	if absent.Or("default") != "default" {
		t.Error("absent Or should return fallback")
	}
}

func TestFieldPtr(t *testing.T) {
	v := FieldFrom(42)
	ptr := v.Ptr()
	if ptr == nil || *ptr != 42 {
		t.Error("Ptr should return pointer to value")
	}

	null := FieldNull[int]()
	if null.Ptr() != nil {
		t.Error("null Ptr should return nil")
	}

	absent := Field[int]{}
	if absent.Ptr() != nil {
		t.Error("absent Ptr should return nil")
	}
}

func TestFieldToOption(t *testing.T) {
	v := FieldFrom("hello")
	val := v.ToOption()
	if val.V != "hello" || !val.Valid {
		t.Error("ToOption from value should be valid")
	}

	null := FieldNull[string]()
	nullVal := null.ToOption()
	if nullVal.Valid {
		t.Error("ToOption from null should be invalid")
	}

	absent := Field[string]{}
	absentVal := absent.ToOption()
	if absentVal.Valid {
		t.Error("ToOption from absent should be invalid")
	}
}

func TestFieldIsZeroOmitzero(t *testing.T) {
	absent := Field[string]{}
	if !absent.IsZero() {
		t.Error("absent should be IsZero (omitted from JSON)")
	}

	null := FieldNull[string]()
	if null.IsZero() {
		t.Error("null should NOT be IsZero (should appear as null in JSON)")
	}

	v := FieldFrom("hello")
	if v.IsZero() {
		t.Error("value should NOT be IsZero")
	}
}

func TestFieldJSONRoundtrip(t *testing.T) {
	type patch struct {
		Name Field[string] `json:"name,omitzero"`
		Age  Field[int]    `json:"age,omitzero"`
	}

	original := patch{
		Name: FieldFrom("John"),
		Age:  FieldNull[int](),
	}

	data, err := json.Marshal(original)
	if err != nil {
		t.Fatal(err)
	}

	var decoded patch
	if err := json.Unmarshal(data, &decoded); err != nil {
		t.Fatal(err)
	}

	if !decoded.Name.IsValue() || decoded.Name.V != "John" {
		t.Error("Name should roundtrip as value")
	}
	if !decoded.Age.IsNull() {
		t.Error("Age should roundtrip as null")
	}
}
