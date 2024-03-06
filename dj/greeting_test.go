package dj_test

import (
	"bytes"
	"testing"

	. "example.com/go-test/dj"
	"github.com/stretchr/testify/assert"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Chris")

	got := buffer.String()
	want := "Hello, Chris"

	assert.Equal(t, want, got)
}
