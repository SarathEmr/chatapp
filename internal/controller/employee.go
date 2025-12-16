package controller

import (
	"chatapp/internal/model"
	"chatapp/internal/repo"
	"context"
	"log"
)

type controller interface {
	Add(context.Context, model.Employee) error
}

type Controller struct {
	Repo repo.DBOps
}

func (c Controller) Add(ctx context.Context, emp model.Employee) error {
	log.Println("controller ...")

	err := c.Repo.Insert(ctx, emp)
	log.Println("ctrler .. err", err)
	return nil
}
