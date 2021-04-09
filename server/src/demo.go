package demo

import (
	"demo/pb"
	"demo/web"
	"fmt"
)

// Start ...
func Start() {

	fmt.Println(&pb.Player{
		ID:    1,
		Level: 2,
		Class: pb.Class_Mage,
	})

	web.Server(5678)
}
