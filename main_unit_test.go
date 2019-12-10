// +build unit

package main

import "testing"

func TestGetNum(t *testing.T) {
	actual := GetNum()

	if actual != 10 {
		t.Fatal("Oh Hoy")
	}
}
