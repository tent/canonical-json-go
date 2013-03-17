package cjson

import (
	"testing"
)

func TestCanonicalization(t *testing.T) {
	input := struct {
		C map[string]interface{} `json:"c"`
		A string                 `json:"a"`
		D []int                  `json:"d"`
		B int                    `json:"b"`
	}{map[string]interface{}{"b": "b", "a": "\n\r", "c": "\"\\<>"}, "a", []int{1, 2, 3}, 1}
	expected := `{"a":"a","b":1,"c":{"a":"` + "\n\r" + `","b":"b","c":"\"\\<>"},"d":[1,2,3]}`

	output, err := Marshal(input)
	if err != nil {
		t.Errorf("got err = %v, want nil", err)
	}
	if expected != string(output) {
		t.Errorf("got %s, want %s", string(output), expected)
	}
}

func TestFloatError(t *testing.T) {
	input := struct{ A float64 }{1.0}

	_, err := Marshal(input)
	if err == nil {
		t.Errorf("want float error, got nil")
	}
}
