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
	case "get":
		if p.redis.CurrentDatabase == nil {
			return errors.New("No Database selected")
		}
		key := strings.Split(input, " ")[0]
		value, ok := p.redis.CurrentDatabase.Get(key)
		if !ok {
			return errors.New("Key does not exist")
		}
		fmt.Println(value)
	case "del":
		if p.redis.CurrentDatabase == nil {
			return errors.New("No Database selected")
		}
		key := strings.Split(input, " ")[0]
		fmt.Println(p.redis.CurrentDatabase.Delete(key))
	case "keys":
		if p.redis.CurrentDatabase == nil {
			return errors.New("No Database selected")
		}
		pattern := strings.Split(input, " ")[1]
		keys, err := p.redis.CurrentDatabase.KeysMatchesPatern(pattern)
		if err != nil {
			return err
		}
		fmt.Println(keys)
	case "select":
		if p.redis.CurrentDatabase == nil {
			return errors.New("No Database selected")
		}
		db, _ := strconv.Atoi(strings.Split(input, " ")[1])
		p.redis.CurrentDatabase = p.redis.GetAllDatabases()[db]

	case "exit":
		os.Exit(0)
	default:
		return errors.New("invalid command")
	}
	return nil

}
