package main

import (
	"encoding/json"
	"log"
)

type S struct {
	A int
	B *int
	C float64
	d func() string
	e chan struct{}
}

func main() {
	s := S{
		A: 1,
		B: nil,
		C: 12.15,
		d: func() string {
			return "NowCoder"
		},
		e: make(chan struct{}),
	}

	_, err := json.Marshal(s)
	if err != nil {
		log.Printf("err occurred..")
		return
	}
	log.Printf("everything is ok.")
	return
}
