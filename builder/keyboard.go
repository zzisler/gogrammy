package builder

import "github.com/go-telegram/bot/models"

type KeyboardBuilder struct {
	rows [][]models.InlineKeyboardButton
	cur  []models.InlineKeyboardButton
}

func NewKeyBoard() *KeyboardBuilder {
	return &KeyboardBuilder{}
}

func (k *KeyboardBuilder) Text(text, callbackData string) *KeyboardBuilder {
	k.cur = append(k.cur, models.InlineKeyboardButton{
		Text:         text,
		CallbackData: callbackData,
	})
	return k
}

func (k *KeyboardBuilder) URL(text, url string) *KeyboardBuilder {
	k.cur = append(k.cur, models.InlineKeyboardButton{
		Text: text,
		URL:  url,
	})
	return k
}

func (k *KeyboardBuilder) Row() *KeyboardBuilder {
	if len(k.cur) > 0 {
		k.rows = append(k.rows, k.cur)
		k.cur = nil
	}
	return k
}

func (k *KeyboardBuilder) Build() models.InlineKeyboardMarkup {
	k.Row()
	return models.InlineKeyboardMarkup{InlineKeyboard: k.rows}
}
