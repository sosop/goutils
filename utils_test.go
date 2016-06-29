package goutils

import "testing"

func TestGenerateUUID(t *testing.T) {
	uuid := GenerateUUID()
	if len(uuid) < 8 {
		t.Fatal("uuid is not correct!")
	}
	t.Log(uuid)
}
