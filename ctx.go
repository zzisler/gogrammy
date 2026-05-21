package gogrammy

import (
	"log"

	"github.com/go-telegram/bot"
)

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
