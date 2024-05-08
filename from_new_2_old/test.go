package main

import "fmt"

type channel interface {
	Notify()
}

type webIm struct {
}

type imsdk struct {
}

type wxpub struct {
}

func (w wxpub) Notify() {
	fmt.Println("wxpub notify")
}

func (w *webIm) Notify() {
	fmt.Println("webim notify")
}

func (i *imsdk) Notify() {
	fmt.Println("imsdk notify")
}

func main() {
	var channels []channel
	channels = append(channels, &webIm{})
	channels = append(channels, &imsdk{})
	channels = append(channels, &wxpub{})
	for _, c := range channels {
		c.Notify()
		fmt.Println("-----------------")
	}
}
