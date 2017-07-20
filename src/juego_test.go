package main

import (
	"log"
	"os"
	"testing"
)

func Test000(*testing.T) {
	log.Print("test")
}

func TestAuthenticate(t *testing.T) {
	var v string
	v = Authenticate()
	if v != string {
		t.Error("Expected a string when authenticatings")
	}
}

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}
