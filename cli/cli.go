package cli

import (
	"bufio"
	"os"

	"github.com/Calgorr/Cedis/container"
)

type Parser struct {
	redis  *container.Container
	reader *bufio.Reader
}

func NewParser(c *container.Container) *Parser {
	return &Parser{redis: c, reader: bufio.NewReader(os.Stdin)}
}
