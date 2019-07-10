package p004_structure

import (
	"errors"
	"fmt"
)

type keyErrorString struct {
	key item
}

func (self keyErrorString) Error() string {
	return fmt.Sprintf("Key does not exist: %v\n", self.key)
}

// ----------------------------------------------------
//type MyDictStringInt struct {
//	keys []string
//	values []int
//}
//
////func NewMyDictStringInt() *MyDictStringInt{
////	return &MyDictStringInt{make([]string,0), make([]int,0)}
////}
//
//func (self *MyDictInt) Set(key string, value int) {
//
//}
//
//func (self *MyDictInt) Get(key string) (value int, err error) {
//	return 0, keyErrorString{key}
//}
//func (self *MyDictInt) Delete(key string) {
//
//}
//func (self *MyDictInt) Contains(key string) bool {
//	if _, err := self.Get(key); err == nil {
//		return true
//	}
//	return false
//}
//
//func (self *MyDictInt) Size() int {
//
//}
//func (self *MyDictInt) Keys() (keys []string) {
//	self.Set("A", none)
//}

// ----------------------------------------------------
type SequentialSearchMap struct {
	first *nodeOfMap
}
type item interface{}
type nodeOfMap struct {
	key  item
	val  item
	next *nodeOfMap
}

func NewSequentialSearchMap() *SequentialSearchMap {
	return &SequentialSearchMap{}
}

func (self *SequentialSearchMap) get(k item) (v item, e error) {
	defer func() {
		if r := recover(); r != nil {
			v = nil
			e = keyErrorString{k}
		}
	}()
	elem := self.first
	for elem != nil {
		if elem.key == k {
			return elem.val, nil
		} else {
			elem = elem.next
		}
	}
	return nil, keyErrorString{k}
}
func (self *SequentialSearchMap) put(k item, v item) (e error) {
	defer func() {
		if r := recover(); r != nil {
			e = errors.New("Error!")
		}
	}()
	elem := self.first
	for elem != nil {
		if elem.key == k {
			elem.val = v
			return nil
		} else {
			elem = elem.next
		}
	}
	self.first = &nodeOfMap{k, v, self.first}
	return nil
}
