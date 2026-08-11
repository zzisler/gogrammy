package builder

import (
	"context"

	"github.com/go-telegram/bot"
)

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
