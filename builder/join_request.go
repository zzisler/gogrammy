package builder

import (
	"context"

	"github.com/go-telegram/bot"
)

type ApproveJoinBuilder struct {
	bot    *bot.Bot
	chatID int64
	userID int64
}

func NewApproveJoinBuilder(b *bot.Bot, chatID, userID int64) *ApproveJoinBuilder {
	return &ApproveJoinBuilder{bot: b, chatID: chatID, userID: userID}
}

func (b *ApproveJoinBuilder) Do(ctx context.Context) (bool, error) {
	return b.bot.ApproveChatJoinRequest(ctx, &bot.ApproveChatJoinRequestParams{
		ChatID: b.chatID,
		UserID: b.userID,
	})
}

type DeclineJoinBuilder struct {
	bot    *bot.Bot
	chatID int64
	userID int64
}

func NewDeclineJoinBuilder(b *bot.Bot, chatID, userID int64) *DeclineJoinBuilder {
	return &DeclineJoinBuilder{bot: b, chatID: chatID, userID: userID}
}

func (b *DeclineJoinBuilder) Do(ctx context.Context) (bool, error) {
	return b.bot.DeclineChatJoinRequest(ctx, &bot.DeclineChatJoinRequestParams{
		ChatID: b.chatID,
		UserID: b.userID,
	})
}
