
package main

import "testing"

func TestSqrt(t *testing.T) {
	sqrt := Raiz(0.0001, 1000000000)
	if sqrt != "9999999.825159" {
		t.Errorf("Failed. %s", sqrt)
	}
}