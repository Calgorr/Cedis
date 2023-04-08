package cli

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

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
		fmt.Println(p.parse(strings.TrimSpace(input)))
	}
}

func (p *Parser) parse(input string) error {
	cmd := strings.ToLower(strings.Split(input, " ")[0])

	switch cmd {
	case "set":
		if p.redis.CurrentDatabase == nil {
			return errors.New("No Database selected")
		}
		key, value, Duration := strings.Split(input, "")[0], strings.Split(input, "")[1], strings.Split(input, "")[2]
		ttl, _ := time.ParseDuration(Duration)
		p.redis.CurrentDatabase.Set(key, value, ttl)
	}
}
