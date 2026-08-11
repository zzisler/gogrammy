package gogrammy

import (
	"context"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func (c *Client) Command(cmd string, handler func(*Context)) {
	c.handlers[cmd] = handler
}

func (c *Client) Start(ctx context.Context) error {
	c.Bot.RegisterHandlerMatchFunc(func(*models.Update) bool { return true }, c.dispatch)
	c.Bot.Start(ctx)
	return nil
}

func (c *Client) dispatch(ctx context.Context, _ *bot.Bot, update *models.Update) {
	if update.Message == nil {
		return
	}
	if handler, ok := c.handlers[update.Message.Text]; ok {
		handler(&Context{Client: c, Update: update, Ctx: ctx})
	}
}
