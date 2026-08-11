package builder

import "github.com/go-telegram/bot/models"

type ReplyKeyboardBuilder struct {
	rows        [][]models.KeyboardButton
	cur         []models.KeyboardButton
	resize      bool
	oneTime     bool
	placeholder string
}

func NewReplyKeyboard() *ReplyKeyboardBuilder {
	return &ReplyKeyboardBuilder{}
}

func (k *ReplyKeyboardBuilder) Text(text string) *ReplyKeyboardBuilder {
	k.cur = append(k.cur, models.KeyboardButton{Text: text})
	return k
}

func (k *ReplyKeyboardBuilder) Row() *ReplyKeyboardBuilder {
	if len(k.cur) > 0 {
		k.rows = append(k.rows, k.cur)
		k.cur = nil
	}
	return k
}

func (k *ReplyKeyboardBuilder) Resize() *ReplyKeyboardBuilder {
	k.resize = true
	return k
}

func (k *ReplyKeyboardBuilder) OneTime() *ReplyKeyboardBuilder {
	k.oneTime = true
	return k
}

func (k *ReplyKeyboardBuilder) Build() models.ReplyKeyboardMarkup {
	k.Row()
	return models.ReplyKeyboardMarkup{
		Keyboard:        k.rows,
		ResizeKeyboard:  k.resize,
		OneTimeKeyboard: k.oneTime,
	}
}
