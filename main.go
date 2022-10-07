package main

import (
	connector "connector"
	"fmt"

	backend "ashish.com/backend/database"
	"example.com/ashishworkspace/main/frontend"
)

func main() {

	msg := "hello world\n"
	fmt.Printf("%v", msg)

	maintainer := Name
	fmt.Print(maintainer)

	frontend := frontend.React()
	fmt.Print(frontend)

	backend := backend.Database()
	fmt.Print(backend)

	connector := connector.Connector()
	fmt.Print(connector)
}
