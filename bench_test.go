package opt

import (
	"encoding/json"
	"testing"
)

func BenchmarkIntMarshalJSON(b *testing.B) {
	v := IntFrom(123456)
	for b.Loop() {
		v.MarshalJSON()
	}
}

func BenchmarkIntUnmarshalJSON(b *testing.B) {
	data := []byte("123456")
	var v Int
	for b.Loop() {
		v.UnmarshalJSON(data)
	}
}

func BenchmarkStringMarshalJSON(b *testing.B) {
	v := StringFrom("hello world")
	for b.Loop() {
		v.MarshalJSON()
	}
}

func BenchmarkStringUnmarshalJSON(b *testing.B) {
	data := []byte(`"hello world"`)
	var v String
	for b.Loop() {
		v.UnmarshalJSON(data)
	}
}

func BenchmarkBoolMarshalJSON(b *testing.B) {
	v := BoolFrom(true)
	for b.Loop() {
		v.MarshalJSON()
	}
}

func BenchmarkBoolUnmarshalJSON(b *testing.B) {
	data := []byte("true")
	var v Bool
	for b.Loop() {
		v.UnmarshalJSON(data)
	}
}

func BenchmarkNullUnmarshalJSON(b *testing.B) {
	data := []byte("null")
	var v Int
	for b.Loop() {
		v.UnmarshalJSON(data)
	}
}

func BenchmarkValueMarshalJSON(b *testing.B) {
	v := From(42)
	for b.Loop() {
		v.MarshalJSON()
	}
}

func BenchmarkValueUnmarshalJSON(b *testing.B) {
	data := []byte("42")
	var v Value[int]
	for b.Loop() {
		v.UnmarshalJSON(data)
	}
}

func BenchmarkStructMarshalJSON(b *testing.B) {
	type payload struct {
		Name   String `json:"name"`
		Count  Int    `json:"count"`
		Score  Float  `json:"score"`
		Active Bool   `json:"active"`
	}
	v := payload{
		Name:   StringFrom("bench"),
		Count:  IntFrom(42),
		Score:  FloatFrom(3.14),
		Active: BoolFrom(true),
	}
	for b.Loop() {
		json.Marshal(v)
	}
}

func BenchmarkStructUnmarshalJSON(b *testing.B) {
	type payload struct {
		Name   String `json:"name"`
		Count  Int    `json:"count"`
		Score  Float  `json:"score"`
		Active Bool   `json:"active"`
	}
	data := []byte(`{"name":"bench","count":42,"score":3.14,"active":true}`)
	var v payload
	for b.Loop() {
		json.Unmarshal(data, &v)
	}
}
