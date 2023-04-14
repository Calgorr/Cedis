package cli

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/Calgorr/Cedis/container"
)

type Parser struct {
	redis  *container.Container
	reader *bufio.Reader
}

const (
	EX = "EX"
	PX = "PX"
)

func NewParser(c *container.Container) *Parser {
	return &Parser{redis: c, reader: bufio.NewReader(os.Stdin)}
}

func (p *Parser) StartProgrammingLoop() error {
	for {
		input, err := p.reader.ReadString('\n')
		if err != nil {
			return err
		}
		p.parse(strings.Split(strings.TrimSpace(input), " "))
	}
}

func (p *Parser) parse(input []string) error {
	cmd := input[0]

	switch cmd {
	case "set":
		var Duration string
		if p.redis.CurrentDatabase == nil {
			return errors.New("No Database selected")
		}
		key, value := input[1], input[2]
		if len(input) == 5 {
			if strings.Compare(EX, input[3]) == 0 {
				Duration = input[4] + "s"
			} else if strings.Compare(PX, input[3]) == 0 {
				Duration = input[4] + "ms"
			}
		}
		ttl, err := time.ParseDuration(Duration)
		if err != nil {
			ttl = 0
			p.redis.CurrentDatabase.Set(key, value, ttl)
		}
		p.redis.CurrentDatabase.Set(key, value, ttl)
	case "get":
		if p.redis.CurrentDatabase == nil {
			return errors.New("No Database selected")
		}
		key := input[1]
		value, ok := p.redis.CurrentDatabase.Get(key)
		if !ok {
			return errors.New("Key does not exist")
		}
		fmt.Println(value)
	case "del":
		if p.redis.CurrentDatabase == nil {
			return errors.New("No Database selected")
		}
		key := input[1]
		fmt.Println(p.redis.CurrentDatabase.Delete(key))
	case "keys":
		if p.redis.CurrentDatabase == nil {
			return errors.New("No Database selected")
		}
		pattern := input[1]
		keys, err := p.redis.CurrentDatabase.KeysMatchesPatern(pattern)
		if err != nil {
			return err
		}
		fmt.Println(keys)
	case "select":
		db, _ := strconv.Atoi(input[1])
		p.redis.CurrentDatabase = p.redis.GetDatabase(db)

	case "exit":
		os.Exit(0)
	default:
		return errors.New("invalid command")
	}
	return nil

}
