package util

import "testing"

func TestGenShortId2(t *testing.T) {
	id, err := GenShortId()
	if err != nil {
		t.Error("not pass")
	}
	t.Logf("pass and id is %s", id)
}
