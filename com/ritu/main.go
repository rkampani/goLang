package main

import (
	"fmt"

	"github.com/rituK/com/ritu/controller"

	"github.com/rituK/com/ritu/service"
)

var appName = "testapp-Here i go -is this really working-i have to clear this up before i say -YES"

func main() {
	fmt.Printf("Starting %v\n", appName)
	fmt.Printf("trying to access the package from different folder", controller.GetUsers)
	service.StartWebServer("8089")
}
