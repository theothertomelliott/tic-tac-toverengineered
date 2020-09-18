package param_test

import (
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/theothertomelliott/tic-tac-toverengineered/pkg/param"
)

func TestParseString(t *testing.T) {
	var got string
	req := httptest.NewRequest("GET", "/?str=value", nil)
	err := param.Parse(req, "str", &got, param.ParseOptions{})
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, "value", got)
}

func TestParseQuotedString(t *testing.T) {
	var got string
	req := httptest.NewRequest("GET", "/?str=\"value\"", nil)
	err := param.Parse(req, "str", &got, param.ParseOptions{})
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, "value", got)
}

func TestParseInt(t *testing.T) {
	var got int
	req := httptest.NewRequest("GET", "/?num=123", nil)
	err := param.Parse(req, "num", &got, param.ParseOptions{})
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, int(123), got)
}

func TestParseInt32(t *testing.T) {
	var got int32
	req := httptest.NewRequest("GET", "/?num=123", nil)
	err := param.Parse(req, "num", &got, param.ParseOptions{Required: true})
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, int32(123), got)
}

func TestParseStruct(t *testing.T) {
	type T struct {
		A string
		B string
	}
	var got T
	req := httptest.NewRequest("GET", "/?t=%7B%22A%22%3A%20%22A%22%2C%20%22B%22%3A%20%22B%22%7D", nil)
	err := param.Parse(req, "t", &got, param.ParseOptions{})
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, T{A: "A", B: "B"}, got)
}

func TestDefaultValue(t *testing.T) {
	var got float64
	req := httptest.NewRequest("GET", "/?other=something", nil)
	err := param.Parse(req, "f", &got, param.ParseOptions{Default: float64(3.14)})
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, float64(3.14), got)
}

func TestRequiredValueMissing(t *testing.T) {
	var got float64
	req := httptest.NewRequest("GET", "/?other=something", nil)
	err := param.Parse(req, "f", &got, param.ParseOptions{Required: true})
	if err == nil {
		t.Error("Expected an error")
	}
}
