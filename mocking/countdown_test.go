package mocking_test

import (
	"bytes"
	"testing"

	. "example.com/go-test/mocking"
	"github.com/stretchr/testify/assert"
)

func TestCountDown(t *testing.T) {
	buffer := &bytes.Buffer{}
	spySleepr := &SpySleeper{}
	CountDown(buffer, spySleepr)
	want := `3
2
1
Go!`
	got := buffer.String()
	assert.Equal(t, want, got)
	assert.True(t, spySleepr.Call == 3)
}

