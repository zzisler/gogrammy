package gogrammy

import (
	"context"

	"github.com/go-telegram/bot/models"
	"github.com/zzisler/gogrammy/builder"
)

type Context struct {
	*Client
	Ctx    context.Context
	Update *models.Update
}

func (c *Context) UserID() int64 {
	return c.Update.Message.From.ID
}

func (c *Context) AnswerCallback() *builder.AnswerCallbackBuilder {
	return builder.NewAnswerCallbackBuilder(c.Bot, c.Update.CallbackQuery.ID)
}

func (c *Context) SendChatAction(userID int64, action models.ChatAction) *builder.ChatActionBuilder {
	return builder.NewChatActionBuilder(c.Bot, userID, action)
}

func (c *Context) RemoveKeyboard() models.ReplyKeyboardRemove {
	return models.ReplyKeyboardRemove{RemoveKeyboard: true}
}
