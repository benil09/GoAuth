package main

import (
	"fmt"
	"AuthInGo/app"
)

func main() {
	cfg := app.Config{
		Addr : ":8080",
	}

	application := app.Application{Config: cfg}

	err := application.Run(); 
	if err != nil {
		fmt.Println(err)
	}	
}

 