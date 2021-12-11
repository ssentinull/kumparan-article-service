package utils_test

import (
	"testing"

	"github.com/ssentinull/kumparan-article-service/pkg/utils"
)

type dump struct {
	ID   int64
	Name string
}

func TestValidDump(t *testing.T) {
	test := dump{ID: 1, Name: "test name"}
	expect := `{"ID":1,"Name":"test name"}`
	if utils.Dump(test) != expect {
		t.Errorf("INCORRECT! Expected a %s", expect)
	}
}
