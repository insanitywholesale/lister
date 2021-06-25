package main

import (
	"testing"
)

//Test shitpost func
func TestServiceName(t *testing.T) {
	expectedResult := "lister"
	actualResult := getServiceName()
	if expectedResult != actualResult {
		t.Fatal("service name is wrong")
	}
	t.Log("service name is:", actualResult)
}
