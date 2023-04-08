package cli

import (
	"bufio"
	"os"
	"strings"

	"github.com/Calgorr/Cedis/container"
)

type Parser struct {
	redis  *container.Container
	reader *bufio.Reader
}

func NewParser(c *container.Container) *Parser {
	return &Parser{redis: c, reader: bufio.NewReader(os.Stdin)}
}

func (p *Parser) StartProgrammingLoop() error {
	for {
		input, err := p.reader.ReadString('\n')
		if err != nil {
			return err
		}
		p.parse(strings.TrimSpace(input))
	}
	return nil
}

func (p *Parser) parse(input string) error {

}
