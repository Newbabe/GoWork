package main

import (
	"fmt"
	"sort"
)

type ListMap []map[string]interface{}

func (l ListMap) Len() int {
	return len(l)
}

func (l ListMap) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

type MapComparator struct {
	key string
}

/*func NewMapComparator(key string) *MapComparator {
	return &MapComparator{key: key}
}*/

func (mc *MapComparator) Compare(o1, o2 interface{}) int {
	m1 := o1.(map[string]interface{})
	m2 := o2.(map[string]interface{})
	n1 := m1[mc.key].(int)
	n2 := m2[mc.key].(int)
	if n1 < n2 {
		return -1
	} else if n1 == n2 {
		return 0
	} else {
		return 1
	}
}

func (l ListMap) Less(i, j int) bool {
	mc := MapComparator{key: "age"}
	return mc.Compare(l[i], l[j]) < 0
}

func main() {
	frList := []map[string]interface{}{
		{"name": "Alice", "age": 25},
		{"name": "Bob", "age": 30},
		{"name": "Charlie", "age": 20},
	}

	sort.Sort(ListMap(frList))

	fmt.Println(frList)
}
