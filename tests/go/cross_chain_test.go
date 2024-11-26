package main

import (
	"rampx/backend/db"
	"testing"
)

func TestCrossChainSwaps(*testing.T) {
	db.RunCrossChainFaker()
}
