package container

import (
	"github.com/Calgorr/Cedis/database"
)

type Container struct {
	Databases       []*database.Cache
	CurrentDatabase *database.Cache
}

var container *Container

func NewContainer() *Container {
	if container == nil {
		container = &Container{Databases: nil, CurrentDatabase: nil}
		return container
	}
	return container
}

func (c *Container) AddDatabase(id int) {
	db := database.NewCache(id)
	c.Databases = append(c.Databases, db)
}

func (c *Container) GetDatabase(id int) *database.Cache {
	if len(c.Databases) >= id {
		c.AddDatabase(id)

	}
	return c.Databases[id]
}
