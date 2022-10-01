package main

import (
	"fmt"
	"time"
)

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

}
