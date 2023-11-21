package main

import (
	"HW1_http/controller/httpserver"
	"HW1_http/gates/psg"
	"fmt"
)

func main() {
	p, err := psg.NewPsg("postgres://127.0.0.1:5432/web-programming", "postgres", "Den")
	
	if err != nil{
		fmt.Println("Error occured:", err)
		return
	}
	hs := httpserver.NewHttpServer(":8080", p)
	err = hs.Start()
	if err != nil {
		fmt.Println("Error occured:", err)
		return
	}
}
