package main
import (
	"fmt"
)


var m = make(map[string]int)
func k(list []string) string { return fmt.Sprintf("%q", list) }
func Add (list []string) {m[k(list)]++}
func Count(list []string) int { return m[k(list)] }

func main () {
	// make map1
	map1 := make(map[string]int)
	fmt.Println(len(map1))

	map1 = map[string]int {
		"alice": 31,
		"charlie": 34,
	}
	for i, v := range map1 {
		fmt.Println(i, v)
	}

	// make map2
	map2 := map[string]int{}
	if map2 == nil {
		fmt.Println("empty map2")
	} else {
		fmt.Println("map2 is not nil")
	}

	// map delete func
	delete (map2, "mantou")
	delete (map1, "alice")
	for i, v := range map1 {
		fmt.Println(i, v)
	}

	fmt.Println(map2["tita"]) // this return 0

	// random output
	map1["zhizhou"] = 33
	map1["some1"] = 35
	for i, v := range map1 {
		fmt.Println(i, v)
	}

	// get all keys
	fmt.Println("---------")
	var keys []string
	var vals []int
	for k,v := range map1 {
		keys = append(keys, k)
		vals = append(vals, v)
	}

	for _, v := range keys {
		fmt.Println(v)
	}

	for _, v := range vals {
		fmt.Println(v)
	}

	// compare with nil
	var map3 map[string]int
	fmt.Println(map3 == nil) // true
	// map3["coral"] = 21 // panic: assignment to entry in nil map

	map4 := map[string]int {}
	fmt.Println(map4 == nil) //fale
	map4["coral"] = 33 // OK

	map5 := make(map[string]int)
	map5["zheli"] = 22 // OK

	// key exist
	age, ok := map5["zheli"]
	if !ok {
		fmt.Println("nali is not a key in map5")
	} else {
		fmt.Println("nali value is ", age)
	}

	// key format: []string

	s1 := make([]string,3)
	s1 = append(s1,"one")
	Add(s1)
	fmt.Println(Count(s1))

	// map -- map
	var graph = make(map[string]map[string]bool)
	nnn := graph["nali"]
	if nnn == nil {
		fmt.Println("map map not exist")
	}

}