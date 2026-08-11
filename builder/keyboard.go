package builder

import "github.com/go-telegram/bot/models"

// ================== INLINE =====================

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

// ================== REPLY =====================

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
