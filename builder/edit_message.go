package builder

import (
	"context"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

type EditTextBuilder struct {
	bot         *bot.Bot
	userID      int64
	messageID   int
	text        string
	parseMode   models.ParseMode
	replyMarkup models.ReplyMarkup
}

func NewEditTextBuilder(b *bot.Bot, userID int64, messageID int, text string) *EditTextBuilder {
	return &EditTextBuilder{bot: b, userID: userID, messageID: messageID, text: text}
}

func (b *EditTextBuilder) ParseMode(mode models.ParseMode) *EditTextBuilder {
	b.parseMode = mode
	return b
}

func (b *EditTextBuilder) ReplyMarkup(kb models.ReplyMarkup) *EditTextBuilder {
	b.replyMarkup = kb
	return b
}

func (b *EditTextBuilder) Do(ctx context.Context) (*models.Message, error) {
	return b.bot.EditMessageText(ctx, &bot.EditMessageTextParams{
		ChatID:      b.userID,
		MessageID:   b.messageID,
		Text:        b.text,
		ParseMode:   b.parseMode,
		ReplyMarkup: b.replyMarkup,
	})
}
