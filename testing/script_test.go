package testing

import (
	"testing"

	"github.com/trisomy/gravity"
)

func TestCreateBuilder(t *testing.T) {
	_, err := gravity.NewProjectBuilder("../", "../build")

	if err != nil {
		t.Error(err)
	}
}
