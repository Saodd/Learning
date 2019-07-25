package tools

import (
	"reflect"
	"testing"
)

func TestGen_ints_list(t *testing.T) {
	a, b := Gen_ints_list(5), Gen_ints_list(5)
	if reflect.DeepEqual(a, b) {
		t.Error(a, b)
	}
}
