package main

import (
	"HW1_http/controller/httpserver"
	"fmt"
)

func main() {
	hs, err := httpserver.NewHttpServer(":8080")
	if err != nil {
		fmt.Println("Error occured:", err)
		return
	}
	hs.Start()
}

// func main() {
// 	//r := dto.Record{
// 	//	Name:       "Иван",
// 	//	LastName:   "Иванов",
// 	//	MiddleName: "Иванович",
// 	//	Address:    "Москва",
// 	//	Phone:      "1234567890",
// 	//}
// 	//err := psg.SelectRecord(r)
// 	//if err != nil {
// 	//	fmt.Println(err)
// 	//}
// 	// if err != nil {
// 	// 	fmt.Println("Error:", err)
// 	// }
// 	// //err = p.RecordSave(r)
// 	// fmt.Println(p.RecordGet("1234567890"))
// 	// if err != nil {
// 	// 	fmt.Println("Error:", err)
// 	// }
// }
