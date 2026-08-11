package builder

import (
	"context"

	"github.com/go-telegram/bot"
)

type AnswerCallbackBuilder struct {
	bot           *bot.Bot
	callbackQuery string
	text          string
	showAlert     bool
}

func NewAnswerCallbackBuilder(b *bot.Bot, callbackQueryID string) *AnswerCallbackBuilder {
	return &AnswerCallbackBuilder{bot: b, callbackQuery: callbackQueryID}
}

func (b *AnswerCallbackBuilder) Text(text string) *AnswerCallbackBuilder {
	b.text = text
	return b
}

func (b *AnswerCallbackBuilder) ShowAlert() *AnswerCallbackBuilder {
	b.showAlert = true
	return b
}

func (b *AnswerCallbackBuilder) Do(ctx context.Context) (bool, error) {
	return b.bot.AnswerCallbackQuery(ctx, &bot.AnswerCallbackQueryParams{
		CallbackQueryID: b.callbackQuery,
		Text:            b.text,
		ShowAlert:       b.showAlert,
	})
}
