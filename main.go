package main

import (
	"fmt"

	"github.com/Calgorr/Cedis/cli"
	"github.com/Calgorr/Cedis/container"
)

func main() {
	parser := cli.NewParser(container.NewContainer())
	if err := parser.StartProgrammingLoop(); err != nil {
		fmt.Println(err)
	}
}
