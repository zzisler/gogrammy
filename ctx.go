package gogrammy

import (
	"log"

	"github.com/go-telegram/bot"
)

func (c *Context) UserID() int64 {
	return c.Update.Message.From.ID
}

func (c *Context) FirstName() string {
	return c.Update.Message.From.FirstName
}

func (c *Context) ChatID() int64 {
	return c.Update.Message.Chat.ID
}

func (c *Context) Text() string {
	return c.Update.Message.Text
}

func (c *Context) AnswerCallback(text ...string) (bool, error) {
	p := &bot.AnswerCallbackQueryParams{
		CallbackQueryID: c.Update.CallbackQuery.ID,
	}
	if len(text) > 0 {
		p.Text = text[0]
	}
	isCallback, err := c.Bot.AnswerCallbackQuery(c.Ctx, p)
	if err != nil {
		log.Println("AnswerCallback error:", err)
		return false, err
	}
	return isCallback, nil
}
