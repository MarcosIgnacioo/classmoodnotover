package arraylist

import (
	"fmt"
	"reflect"
)

type ArrayList struct {
	ArrayList []interface{}
	Length    uint
	Capacity  uint
	Type      reflect.Type
}

func (al *ArrayList) String() string {
	var s string
	for i := 0; i < int(al.Length); i++ {
		if al.ArrayList[i] != nil {
			s += fmt.Sprintf("%v ", al.ArrayList[i])
		}
	}
	return s
}

func NewArrayList(c uint) ArrayList {
	return ArrayList{
		ArrayList: make([]interface{}, c),
		Length:    0,
		Capacity:  c,
	}
}

func (arrayList *ArrayList) Enqueue(item interface{}) interface{} {

	if arrayList.Type == nil {
		arrayList.Type = reflect.TypeOf(item)
	}

	if arrayList.Type != reflect.TypeOf(item) {
		return nil
	}

	if arrayList.Length+1 > arrayList.Capacity {
		newArray := make([]interface{}, arrayList.Capacity*2)
		newArray = appendGamer(newArray, arrayList.ArrayList)
		arrayList.ArrayList = newArray
	}

	arrayList.ArrayList[arrayList.Length] = item
	arrayList.Length++
	return item
}

func (arrayList *ArrayList) Push(item interface{}) interface{} {
	if arrayList.Type == nil {
		arrayList.Type = reflect.TypeOf(item)
	}

	if arrayList.Type != reflect.TypeOf(item) {
		return nil
	}

	if arrayList.Length+1 > arrayList.Capacity {
		arrayList.Capacity = arrayList.Capacity * 2
		newArray := make([]interface{}, arrayList.Capacity)
		newArray = appendGamer(newArray, arrayList.ArrayList)
		arrayList.ArrayList = newArray
	}

	if arrayList.Length > 0 {
		for i := arrayList.Length; i > 0; i-- {
			arrayList.ArrayList[i] = arrayList.ArrayList[i-1]
		}
	}
	arrayList.ArrayList[0] = item
	arrayList.Length++
	return item
}

func (arrayList *ArrayList) Pop() {
	arrayList.ArrayList[arrayList.Length-1] = nil
	if arrayList.Length > 0 {
		arrayList.Length--
	}
}

func (arrayList *ArrayList) Dequeue() interface{} {

	var dequeued interface{}
	if arrayList.Length > 0 {
		dequeued = arrayList.ArrayList[0]
		for i := 0; i < int(arrayList.Length)-1; i++ {
			arrayList.ArrayList[i] = arrayList.ArrayList[i+1]
		}
		arrayList.ArrayList[arrayList.Length-1] = nil
	}

	if arrayList.Length > 0 {
		arrayList.Length--

	}
	return dequeued
}

func appendGamer(array []interface{}, secondArray []interface{}) []interface{} {
	for i := 0; i < len(secondArray); i++ {
		array[i] = secondArray[i]
	}
	return array
}
