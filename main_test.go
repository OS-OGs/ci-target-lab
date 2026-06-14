package main

import (
	"os"
	"testing"
)

func TestPrivilegedContextReached(t *testing.T) {
	t.Log("pr_controlled_test_executed=true")

	if os.Getenv("DUMMY_SECRET") != "" {
		t.Log("dummy_secret_visible_to_pr_controlled_test=true")
	} else {
		t.Log("dummy_secret_visible_to_pr_controlled_test=false")
	}
}
