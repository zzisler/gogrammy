package gogrammy

import (
	"context"
	"strings"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

type Handler func(*Context)

func (c *Client) handle(h Handler) bot.HandlerFunc {
	return func(ctx context.Context, b *bot.Bot, update *models.Update) {
		h(&Context{Client: c, Update: update, Ctx: ctx})
	}
}

func (c *Client) On(eventType string, h Handler) {
	var matchFunc bot.MatchFunc
	switch eventType {
	case "message":
		matchFunc = func(update *models.Update) bool { return update.Message != nil }
	case "callback":
		matchFunc = func(update *models.Update) bool { return update.CallbackQuery != nil }
	case "join_request":
		matchFunc = func(update *models.Update) bool { return update.ChatJoinRequest != nil }
	case "my_chat_member":
		matchFunc = func(update *models.Update) bool { return update.MyChatMember != nil }
	default:
		matchFunc = func(update *models.Update) bool { return true }
	}
	c.Bot.RegisterHandlerMatchFunc(matchFunc, c.handle(h))
}

func (c *Client) OnCallback(prefix string, h Handler) {
	c.Bot.RegisterHandlerMatchFunc(func(update *models.Update) bool {
		return update.CallbackQuery != nil && strings.HasPrefix(update.CallbackQuery.Data, prefix)
	}, c.handle(h))
}

func (c *Client) Command(cmd string, h Handler) {
	c.Bot.RegisterHandlerMatchFunc(func(update *models.Update) bool {
		return update.Message != nil && update.Message.Text == cmd
	}, c.handle(h))
}

func (c *Client) Start(ctx context.Context) {
	c.Bot.Start(ctx)
}
