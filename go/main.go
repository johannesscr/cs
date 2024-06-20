package main

import (
	"fmt"
	"net/http"
	"net/url"
)

type Service struct {
	URL     url.URL
	Values  url.Values
	Headers http.Header
}

func (s *Service) GetEnv() {
	fmt.Println("read env")
}

// here we cannot access the GetEnv method.. why not?? surely they are the same
// type?
// well in go, the idea of objects/classes OOP and inheritance does note exist.
//
// go is all about type and interface
//
// so instead of using OOP, let's do it with interfaces

// for a service to be mockable we need
// 1. to be able to set the host
// 2. to be able to set the scheme the rest does not really matter.

//type Service interface {
//	SetHost()
//	SetScheme()
//}

// now we can mock anything that is able to set the host and scheme

type MicroService Service

func (ms *MicroService) Set() {}

func main() {
	ms := &MicroService{}
	ms.Set()
}
