package main

import (
	"fmt"

	"github.com/Calgorr/Cedis/cli"
	"github.com/Calgorr/Cedis/container"
)

func main() {
	parser := cli.NewParser(container.NewContainer())
	fmt.Println(parser.StartProgrammingLoop())
}
