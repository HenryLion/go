package main

import (
	"fmt"
	"reflect"
	"time"
)

type User struct {
	name uint64
}

type User1 struct {
	name string
}

func print1(user any) {
	myname := reflect.ValueOf(user).Elem().FieldByName("name")
	fmt.Println(myname.Type().Name())
	if myname.Type().Name() == "string" {
		fmt.Println("name is: %s", myname.String())
	} else if myname.Type().Name() == "uint64" {
		fmt.Println("name is: %d", myname.Uint())
	}
}

func main() {
	p := fmt.Println
	p(time.Now().Weekday())

	p(time.Now())

	then := time.Date(1987, 9, 18, 20, 34, 58, 456789, time.UTC)
	p(then)

	p(then.Weekday())
	p(then.Nanosecond())
	p(then.Before(time.Now()))
	diff := time.Now().Sub(then)
	p(diff)
	p(diff.Hours())
	p(diff.Minutes())
	p(diff.Seconds())
	p(diff.Nanoseconds())

	p(then.Add(diff))
	p(then.Add(-diff))

	timeTemplate := "2006-01-02 15:04:05"
	tempstr := "2022-11-17 00:00:00"
	stamp, _ := time.ParseInLocation(timeTemplate, tempstr, time.Local)
	fmt.Println(stamp.Unix())

	timeTemplate1 := "2006-01-02"
	tempstr1 := "2022-11-17"
	stamp1, _ := time.ParseInLocation(timeTemplate1, tempstr1, time.Local)
	fmt.Println(stamp1.Unix())

	fmt.Println(time.Now().Format("2006-01-02"))

	_, err := time.ParseInLocation("2006-01-02 15:04:05", "", time.Local)
	if err != nil {
		fmt.Println(err)
	}

}
