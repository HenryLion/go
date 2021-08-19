package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func echo1() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}

func echo2() {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}

func exercise1_1() {
	s, sep := "", ""
	for _, arg := range os.Args[:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}

func exercise1_2() {
	for index, arg := range os.Args[:] {
		fmt.Println(index, arg)
	}
}

type Person struct {
	Name string `json:"name"`
	Age  uint32 `json:"age"`
	Sex  string `json:"sex"`
}

type class_one struct {
	Person1   Person `json:"person1"`
	ClassName string `json:"classname"`
}

type ShortPerson struct {
	Name string `json:"name"`
}

type Part_class struct {
	PersonName ShortPerson `json:"person1"`
	ClassName  string      `json:"classname"`
}

func main() {
	echo1()
	echo2()
	fmt.Println(strings.Join(os.Args[1:], " "))
	fmt.Println(os.Args[1:])
	fmt.Println("exercise below =====")
	exercise1_1()
	exercise1_2()

	person1 := Person{Name: "jk", Age: 23, Sex: "male"}
	fmt.Println(person1) // 原始变量
	byte_person, _ := json.Marshal(person1)
	fmt.Println(byte_person) // 序列化为[]byte
	str_person := string(byte_person)
	fmt.Println(str_person)         // []byte转化为string
	fmt.Println([]byte(str_person)) // string转化为[]byte

	var person_unmarshal Person
	json.Unmarshal([]byte(str_person), &person_unmarshal) //从byte反序列化为原始变量
	fmt.Println("new ----")
	fmt.Println(person_unmarshal)
	fmt.Println(person_unmarshal)

	map_unmarshal := make(map[string]interface{})
	json.Unmarshal(byte_person, &map_unmarshal) //从byte反序列化为map

	fmt.Println(map_unmarshal["name"].(string)) // 访问map中key
	fmt.Println(map_unmarshal["age"])
	fmt.Println(map_unmarshal["sex"].(string))

	fmt.Println("-------------------------")

	class1 := class_one{Person1: person1, ClassName: "huojianban"}
	fmt.Println(class1)
	byte_class, _ := json.Marshal(class1)
	fmt.Println(byte_class)
	str_class := string(byte_class)
	fmt.Println(str_class)

	var part_class Part_class
	json.Unmarshal(byte_class, &part_class)
	fmt.Println(part_class)
}
