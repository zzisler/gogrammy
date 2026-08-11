package builder

import (
	"context"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

type TextBuilder struct {
	bot         *bot.Bot
	userID      int64
	text        string
	parseMode   models.ParseMode
	replyTo     int
	replyMarkup models.InlineKeyboardMarkup
}

func NewTextBuilder(b *bot.Bot, userID int64, text string) *TextBuilder {
	return &TextBuilder{bot: b, userID: userID, text: text}
}

func (b *TextBuilder) ParseMode(mode models.ParseMode) *TextBuilder {
	b.parseMode = mode
	return b
}

func (b *TextBuilder) ReplyTo(messageID int) *TextBuilder {
	b.replyTo = messageID
	return b
}

func (b *TextBuilder) ReplyMarkup(kb models.InlineKeyboardMarkup) *TextBuilder {
	b.replyMarkup = kb
	return b
}

func (b *TextBuilder) Do(ctx context.Context) (*models.Message, error) {
	return b.bot.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:      b.userID,
		Text:        b.text,
		ParseMode:   b.parseMode,
		ReplyMarkup: b.replyMarkup,
		ReplyParameters: &models.ReplyParameters{
			MessageID: b.replyTo,
		},
	})
}
