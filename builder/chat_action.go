package builder

import (
	"context"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

type ChatActionBuilder struct {
	bot    *bot.Bot
	userID int64
	action models.ChatAction
}

func NewChatActionBuilder(b *bot.Bot, userID int64, action models.ChatAction) *ChatActionBuilder {
	return &ChatActionBuilder{bot: b, userID: userID, action: action}
}

func (b *ChatActionBuilder) Do(ctx context.Context) (bool, error) {
	return b.bot.SendChatAction(ctx, &bot.SendChatActionParams{
		ChatID: b.userID,
		Action: b.action,
	})
}
