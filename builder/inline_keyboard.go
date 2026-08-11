package builder

import "github.com/go-telegram/bot/models"

type InlineKeyboardBuilder struct {
	rows [][]models.InlineKeyboardButton
	cur  []models.InlineKeyboardButton
}

func NewInlineKeyboard() *InlineKeyboardBuilder {
	return &InlineKeyboardBuilder{}
}

func (k *InlineKeyboardBuilder) Text(text, callbackData string) *InlineKeyboardBuilder {
	k.cur = append(k.cur, models.InlineKeyboardButton{
		Text:         text,
		CallbackData: callbackData,
	})
	return k
}

func (k *InlineKeyboardBuilder) URL(text, url string) *InlineKeyboardBuilder {
	k.cur = append(k.cur, models.InlineKeyboardButton{
		Text: text,
		URL:  url,
	})
	return k
}

func (k *InlineKeyboardBuilder) Row() *InlineKeyboardBuilder {
	if len(k.cur) > 0 {
		k.rows = append(k.rows, k.cur)
		k.cur = nil
	}
	return k
}

func (k *InlineKeyboardBuilder) Build() models.InlineKeyboardMarkup {
	k.Row()
	return models.InlineKeyboardMarkup{InlineKeyboard: k.rows}
}
