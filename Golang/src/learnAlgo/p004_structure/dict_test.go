package p004_structure

import (
	"testing"
)

func TestNewSequentialSearchMap(t *testing.T) {
	m := NewSequentialSearchMap()
	if m.first != nil {
		t.Error("First node is not nil !!")
	}
	if e := m.put("first", 1); (e != nil) || (m.first == nil) {
		t.Error("Failed when put first!")
	}
	m.put("second", 2)
	if v, e := m.get("first"); (e != nil) || (v != 1) {
		t.Error("Failed when get first!")
	}
	m.put("first", 100)
	if v, e := m.get("first"); (e != nil) || (v != 100) {
		t.Error("Failed when update first!")
	}
	if e := m.put(99, 99); (e != nil) || (m.first == nil) {
		t.Error("Failed when put diff type!")
	}
	if v, e := m.get(99); (e != nil) || (v != 99) {
		t.Error("Failed when get diff type!")
	}
}
