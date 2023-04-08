package cli

import (
	"bufio"

	"github.com/Calgorr/Cedis/container"
)

type Parser struct {
	redis  *container.Container
	reader *bufio.Reader
}
