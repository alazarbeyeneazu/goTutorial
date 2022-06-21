package main

import "testing"

func TestHello(t *testing.T) {
	want := "hello, world! "
	got := Hello("world")
	if want != got {
		t.Errorf("want %s  got %s", want, got)
	}
}
