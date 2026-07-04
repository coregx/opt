package opt_test

import (
	"encoding/json"
	"fmt"

	"github.com/coregx/opt"
)

func ExampleStringFrom() {
	s := opt.StringFrom("hello")
	data, _ := json.Marshal(s)
	fmt.Println(string(data))
	fmt.Println(s.Or("default"))

	null := opt.NewString("", false)
	data, _ = json.Marshal(null)
	fmt.Println(string(data))
	fmt.Println(null.Or("default"))
	// Output:
	// "hello"
	// hello
	// null
	// default
}

func ExampleIntFrom() {
	i := opt.IntFrom(42)
	data, _ := json.Marshal(i)
	fmt.Println(string(data))
	fmt.Println(i.OrZero())

	zero := opt.IntFrom(0)
	fmt.Println(zero.IsZero()) // false — 0 is valid, not null
	// Output:
	// 42
	// 42
	// false
}

func ExampleField_patchAPI() {
	type PatchUser struct {
		Name  opt.Field[string] `json:"name,omitzero"`
		Email opt.Field[string] `json:"email,omitzero"`
		Age   opt.Field[int]    `json:"age,omitzero"`
	}

	input := `{"name":"John","email":null}`
	var patch PatchUser
	json.Unmarshal([]byte(input), &patch)

	fmt.Println("name absent:", patch.Name.IsAbsent())
	fmt.Println("name value:", patch.Name.Or(""))
	fmt.Println("email null:", patch.Email.IsNull())
	fmt.Println("age absent:", patch.Age.IsAbsent())
	// Output:
	// name absent: false
	// name value: John
	// email null: true
	// age absent: true
}

func ExampleMap() {
	name := opt.From("John")
	length := opt.Map(name, func(s string) int { return len(s) })
	fmt.Println(length.OrZero())

	null := opt.New("", false)
	result := opt.Map(null, func(s string) int { return len(s) })
	fmt.Println(result.IsZero())
	// Output:
	// 4
	// true
}

func ExampleValue_structJSON() {
	type User struct {
		Name  opt.String `json:"name"`
		Age   opt.Int    `json:"age"`
		Score opt.Float  `json:"score,omitzero"`
	}

	user := User{
		Name: opt.StringFrom("Alice"),
		Age:  opt.NewInt(0, false),
	}

	data, _ := json.Marshal(user)
	fmt.Println(string(data))
	// Output:
	// {"name":"Alice","age":null}
}
