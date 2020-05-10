package main

import "testing"

func Test(t *testing.T) {
	mass := int64(100756)
	actual := fuel(mass) + fuelfuel(fuel(mass))
	expected := int64(50346)
	if actual != expected {
		t.Fatalf("expected %d, got %d", expected, actual)
	}
}
