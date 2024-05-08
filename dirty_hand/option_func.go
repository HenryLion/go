package main

import "fmt"

type ServeInfo struct {
	host     string
	port     int
	protocal string // tcp, udp
}

type SetServeInfoFunc func(*ServeInfo)

func SetHost(host string) SetServeInfoFunc {
	return func(s *ServeInfo) {
		s.host = host
	}
}

func SetPort(port int) SetServeInfoFunc {
	return func(s *ServeInfo) {
		s.port = port
	}
}

func SetProtocal(protocal string) SetServeInfoFunc {
	return func(s *ServeInfo) {
		s.protocal = protocal
	}
}

func New(setFieldsFuncs ...SetServeInfoFunc) *ServeInfo {
	s := &ServeInfo{}
	fmt.Println(*s)
	for _, setFunc := range setFieldsFuncs {
		setFunc(s)
	}
	return s
}

func main() {
	s := New(SetHost("localhost"),
		SetPort(8080),
		SetProtocal("tcp"))
	fmt.Println(*s)
}
