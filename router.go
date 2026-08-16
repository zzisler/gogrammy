package gogrammy

import (
	"context"
	"strings"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

type Handler func(*Context)

func (c *Context) handle(h Handler) bot.HandlerFunc {
	return func(ctx context.Context, b *bot.Bot, update *models.Update) {
		h(&Context{Bot: b, Update: update, Ctx: ctx})
	}
}

func (c *Context) On(eventType string, h Handler) {
	var matchFunc bot.MatchFunc
	switch eventType {
	case "message":
		matchFunc = func(u *models.Update) bool { return u.Message != nil }
	case "edited_message":
		matchFunc = func(u *models.Update) bool { return u.EditedMessage != nil }
	case "channel_post":
		matchFunc = func(u *models.Update) bool { return u.ChannelPost != nil }
	case "edited_channel_post":
		matchFunc = func(u *models.Update) bool { return u.EditedChannelPost != nil }
	case "business_connection":
		matchFunc = func(u *models.Update) bool { return u.BusinessConnection != nil }
	case "business_message":
		matchFunc = func(u *models.Update) bool { return u.BusinessMessage != nil }
	case "edited_business_message":
		matchFunc = func(u *models.Update) bool { return u.EditedBusinessMessage != nil }
	case "deleted_business_message":
		matchFunc = func(u *models.Update) bool { return u.DeletedBusinessMessages != nil }
	case "message_reaction":
		matchFunc = func(u *models.Update) bool { return u.MessageReaction != nil }
	case "message_reaction_count":
		matchFunc = func(u *models.Update) bool { return u.MessageReactionCount != nil }
	case "managed_bot":
		matchFunc = func(u *models.Update) bool { return u.ManagedBot != nil }
	case "inline_query":
		matchFunc = func(u *models.Update) bool { return u.InlineQuery != nil }
	case "chosen_inline_result":
		matchFunc = func(u *models.Update) bool { return u.ChosenInlineResult != nil }
	case "callback":
		matchFunc = func(u *models.Update) bool { return u.CallbackQuery != nil }
	case "shipping_query":
		matchFunc = func(u *models.Update) bool { return u.ShippingQuery != nil }
	case "pre_checkout_query":
		matchFunc = func(u *models.Update) bool { return u.PreCheckoutQuery != nil }
	case "poll":
		matchFunc = func(u *models.Update) bool { return u.Poll != nil }
	case "poll_answer":
		matchFunc = func(u *models.Update) bool { return u.PollAnswer != nil }
	case "my_chat_member":
		matchFunc = func(u *models.Update) bool { return u.MyChatMember != nil }
	case "chat_member":
		matchFunc = func(u *models.Update) bool { return u.ChatMember != nil }
	case "join_request":
		matchFunc = func(u *models.Update) bool { return u.ChatJoinRequest != nil }
	case "chat_boost":
		matchFunc = func(u *models.Update) bool { return u.ChatBoost != nil }
	case "removed_chat_boost":
		matchFunc = func(u *models.Update) bool { return u.RemovedChatBoost != nil }
	default:
		matchFunc = func(u *models.Update) bool { return true }
	}
	c.Bot.RegisterHandlerMatchFunc(matchFunc, c.handle(h))
}

func (c *Context) OnCallback(prefix string, h Handler) {
	c.Bot.RegisterHandlerMatchFunc(func(update *models.Update) bool {
		return update.CallbackQuery != nil && strings.HasPrefix(update.CallbackQuery.Data, prefix)
	}, c.handle(h))
}

func (c *Context) Command(cmd string, h Handler) {
	c.Bot.RegisterHandlerMatchFunc(func(update *models.Update) bool {
		return update.Message != nil && update.Message.Text == cmd
	}, c.handle(h))
}

func (c *Context) Start(ctx context.Context) {
	c.Bot.Start(ctx)
}
