package site

import (
	"testing"
)

func TestHello(t *testing.T) {
	if true == false {
		t.Fatalf(`my whole life is a lie`)
	}
}