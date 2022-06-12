package simplefactory

import (
	simple "SimpleFactory/SimpleFactory/Simple"
	"testing"
)

func TestType1(t *testing.T) {
	api := simple.NewAPI(1)
	s := api.Say("Mars")
	if s != "Hi, Mars" {
		t.Fatal("Type1 test fail")
	}
}
func TestType2(t *testing.T) {
	api := simple.NewAPI(2)
	s := api.Say("Mars")
	if s != "helloAPI, Mars" {
		t.Fatal("Type2 test fail")
	}
}
