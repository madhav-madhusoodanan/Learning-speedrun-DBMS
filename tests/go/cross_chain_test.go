package main

import (
	"rampx/backend/internal/db"
	"testing"
)

func TestCrossChainSwaps(*testing.T) {
	db.RunCrossChainFaker()
}
