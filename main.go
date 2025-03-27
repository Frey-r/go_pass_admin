package main

import (
	"fmt"
	"passcript/internal/controllers"
)

func main() {
	fmt.Println("Hello World!")
	//controllers.Encoder()
	fmt.Print("isDb: ", controllers.Check_db())
	controllers.Get_db()

}
