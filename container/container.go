package container

import "github.com/Calgorr/Cedis/database"

type Container struct {
	databases       []*database.Cache
	currentDatabase *database.Cache
}

func NewContainer() *Container {
	return &Container{databases: nil, currentDatabase: nil}
}
