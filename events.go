package gogrammy

import (
	"strings"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func (a *App) On(eventType string, h Handler) {
	var matchFunc bot.MatchFunc

	switch eventType {
	case "join_request":
		matchFunc = func(update *models.Update) bool {
			return update.ChatJoinRequest != nil
		}
	case "my_chat_member":
		matchFunc = func(update *models.Update) bool {
			return update.MyChatMember != nil
		}
	case "callback":
		matchFunc = func(update *models.Update) bool {
			return update.CallbackQuery != nil
		}
	case "message":
		matchFunc = func(update *models.Update) bool {
			return update.Message != nil
		}
	case "business_message":
		matchFunc = func(update *models.Update) bool {
			return update.BusinessMessage != nil
		}
	default:
		matchFunc = func(update *models.Update) bool {
			return true
		}
	}

	a.Bot.RegisterHandlerMatchFunc(matchFunc, a.handle(h))
}

func (a *App) OnCallback(prefix string, h Handler) {
	a.Bot.RegisterHandlerMatchFunc(func(update *models.Update) bool {
		return update.CallbackQuery != nil && strings.HasPrefix(update.CallbackQuery.Data, prefix)
	}, a.handle(h))
}
