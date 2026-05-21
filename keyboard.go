package gogrammy

import "github.com/go-telegram/bot/models"

func Row(buttons ...models.InlineKeyboardButton) []models.InlineKeyboardButton {
	return buttons
}

func Keyboard(rows ...[]models.InlineKeyboardButton) [][]models.InlineKeyboardButton {
	return rows
}

func Button(text string, data string) models.InlineKeyboardButton {
	return models.InlineKeyboardButton{
		Text:         text,
		CallbackData: data,
	}
}
