package dotenv

import (
	"testing"
)

func TestGet(t *testing.T) {
	Init()
	kv.set("TEST_1", "123")
	kv.set("TEST_2", "456")
	if Get("TEST_1") != "123" {
		t.Fail()
	}
	if Get("TEST_2") != "456" {
		t.Fail()
	}
	if Get("TEST_N") != "" {
		t.Fail()
	}
}
