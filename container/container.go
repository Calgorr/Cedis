package container

import "github.com/Calgorr/Cedis/database"

type Container struct {
	databases       []*database.Cache
	currentDatabase *database.Cache
}

func NewContainer() *Container {
	return &Container{databases: nil, currentDatabase: nil}
}

func (c *Container) AddDatabase(id int) {
	db := database.NewCache(id)
	c.databases = append(c.databases, db)
}
