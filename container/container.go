package container

import "github.com/Calgorr/Cedis/database"

type Container struct {
	databases       []*database.Cache
	currentDatabase *database.Cache
}

var container *Container

func NewContainer() *Container {
	if container == nil {
		container = &Container{databases: nil, currentDatabase: nil}
		return container
	}
	return container
}

func (c *Container) AddDatabase(id int) {
	db := database.NewCache(id)
	c.databases = append(c.databases, db)
}

func (c *Container) GetAllDatabases() []*database.Cache {
	return c.databases
}
