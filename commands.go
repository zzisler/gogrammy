package gogrammy

import (
	"strings"

	"github.com/go-telegram/bot/models"
)

func (a *App) Command(cmd string, h Handler) {
	a.Bot.RegisterHandlerMatchFunc(func(update *models.Update) bool {
		if update.Message == nil {
			return false
		}
		text := update.Message.Text
		return text == cmd || strings.HasPrefix(text, cmd+"@"+a.Username)
	}, a.handle(h))
}

func (a *App) PrivateCommand(cmd string, h Handler) {
	a.Bot.RegisterHandlerMatchFunc(func(update *models.Update) bool {
		if update.Message == nil {
			return false
		}
		return update.Message.Chat.Type == "private" && update.Message.Text == cmd
	}, a.handle(h))
}

func (a *App) GroupCommand(cmd string, h Handler) {
	a.Bot.RegisterHandlerMatchFunc(func(update *models.Update) bool {
		if update.Message == nil {
			return false
		}
		text := update.Message.Text
		isGroup := update.Message.Chat.Type == "group" || update.Message.Chat.Type == "supergroup"
		return isGroup && (text == cmd || strings.HasPrefix(text, cmd+"@"+a.Username))
	}, a.handle(h))
}
