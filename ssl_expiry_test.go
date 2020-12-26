package main

import "testing"
import "fmt"

func TestProbe(t *testing.T) {
	want := "foo"
	if got := fmt.Sprint("foo"); got != want {
		t.Errorf("wanted %s, got %s", want, got)
	}
}
