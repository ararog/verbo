package verbo

import (
	"strings"
)

import "testing"

func TestIsBlank(t *testing.T) {
  if ! IsBlank(" ") {
    t.Errorf("Wrong result: %s", "should not go here")
  }
}
