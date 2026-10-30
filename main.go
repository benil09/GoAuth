package main

import (
	"fmt"
	"AuthInGo/app"	
)

func main() {

		cfg := app.NewConfig()

		application := app.NewApplication(cfg)

	err := application.Run();
	
		if err != nil {
			fmt.Println(err)
		}
	

}
