//go:build !integration

package main

import (
	"log"
	"testing"
)

func TestMain(t *testing.T) {
	log.Fatal("Fatal error on unit tests")
}
