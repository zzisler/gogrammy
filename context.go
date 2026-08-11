package gogrammy

import (
	"context"

	"github.com/go-telegram/bot/models"
)

type Context struct {
	*Client
	Ctx    context.Context
	Update *models.Update
}

func (c *Context) UserID() int64 {
	return c.Update.Message.From.ID
}
