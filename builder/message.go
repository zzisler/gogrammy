package builder

import (
	"context"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

// ================== SEND =====================

type TextBuilder struct {
	bot         *bot.Bot
	userID      int64
	text        string
	parseMode   models.ParseMode
	replyTo     int
	replyMarkup models.ReplyMarkup
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

func (b *TextBuilder) ReplyMarkup(kb models.ReplyMarkup) *TextBuilder {
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

// ================== EDIT =====================

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

// ===================== DELETE ====================

type DeleteBuilder struct {
	bot       *bot.Bot
	userID    int64
	messageID int
}

func NewDeleteBuilder(b *bot.Bot, userID int64, messageID int) *DeleteBuilder {
	return &DeleteBuilder{bot: b, userID: userID, messageID: messageID}
}

func (b *DeleteBuilder) Do(ctx context.Context) (bool, error) {
	return b.bot.DeleteMessage(ctx, &bot.DeleteMessageParams{
		ChatID:    b.userID,
		MessageID: b.messageID,
	})
}
