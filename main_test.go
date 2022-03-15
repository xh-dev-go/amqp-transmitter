package main

import "testing"

func TestValidate(t *testing.T) {
	err := validate("", "")
	if err == nil {
		t.Error("expecting error")
	} else if err != NoExchangeOrQueue {
		t.Error("No expected error")
	}
}
