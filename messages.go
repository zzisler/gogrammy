package gogrammy

import (
	"log"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

type MsgParams struct {
	ParseMode      string
	Reply          bool
	ReplyMarkup    [][]models.InlineKeyboardButton
	DisablePreview bool
}

// ---------------------------SEND------------------------------

func (c *Context) Send(chatID any, text string, params *MsgParams) (*models.Message, error) {

	p := &bot.SendMessageParams{
		ChatID: chatID,
		Text:   text,
	}

	if c.Update.BusinessMessage != nil {
		p.BusinessConnectionID = c.Update.BusinessMessage.BusinessConnectionID
	}

	if params != nil {
		p.ParseMode = models.ParseMode(params.ParseMode)
		if params.ReplyMarkup != nil {
			p.ReplyMarkup = &models.InlineKeyboardMarkup{
				InlineKeyboard: params.ReplyMarkup,
			}
		}
		if params.Reply {

			if c.Update.BusinessMessage != nil {
				p.ReplyParameters = &models.ReplyParameters{
					MessageID: c.Update.BusinessMessage.ID,
				}
			} else {
				p.ReplyParameters = &models.ReplyParameters{
					MessageID: c.Update.Message.ID,
				}
			}
		}
		if params.DisablePreview {
			p.LinkPreviewOptions = &models.LinkPreviewOptions{
				IsDisabled: bot.True(),
			}
		}
	}

	msg, err := c.Bot.SendMessage(c.Ctx, p)
	if err != nil {
		log.Println("Send error:", err)
		return nil, err
	}
	return msg, nil
}

// ---------------------------EDIT------------------------------

func (c *Context) Edit(text string, params *MsgParams) (*models.Message, error) {

	var chatID any
	var messageID int

	if c.Update.CallbackQuery != nil {
		chatID = c.Update.CallbackQuery.Message.Message.Chat.ID
		messageID = c.Update.CallbackQuery.Message.Message.ID
	} else if c.Update.Message != nil {
		chatID = c.Update.Message.Chat.ID
		messageID = c.Update.Message.ID
	}

	p := &bot.EditMessageTextParams{
		ChatID:    chatID,
		MessageID: messageID,
		Text:      text,
	}

	if params != nil {
		p.ParseMode = models.ParseMode(params.ParseMode)
		if params.ReplyMarkup != nil {
			p.ReplyMarkup = &models.InlineKeyboardMarkup{
				InlineKeyboard: params.ReplyMarkup,
			}
		}
		if params.DisablePreview {
			p.LinkPreviewOptions = &models.LinkPreviewOptions{
				IsDisabled: bot.True(),
			}
		}
	}

	msg, err := c.Bot.EditMessageText(c.Ctx, p)
	if err != nil {
		log.Println("Edit error:", err)
		return nil, err
	}
	return msg, nil
}

func (c *Context) EditMessage(messageID int, text string, params *MsgParams) (*models.Message, error) {

	var chatID any

	if c.Update.CallbackQuery != nil {
		chatID = c.Update.CallbackQuery.Message.Message.Chat.ID
	} else if c.Update.Message != nil {
		chatID = c.Update.Message.Chat.ID
	}

	p := &bot.EditMessageTextParams{
		ChatID:    chatID,
		MessageID: messageID,
		Text:      text,
	}

	if params != nil {
		p.ParseMode = models.ParseMode(params.ParseMode)
		if params.ReplyMarkup != nil {
			p.ReplyMarkup = &models.InlineKeyboardMarkup{
				InlineKeyboard: params.ReplyMarkup,
			}
		}
		if params.DisablePreview {
			p.LinkPreviewOptions = &models.LinkPreviewOptions{
				IsDisabled: bot.True(),
			}
		}
	}

	msg, err := c.Bot.EditMessageText(c.Ctx, p)
	if err != nil {
		log.Println("Edit error:", err)
		return nil, err
	}
	return msg, nil
}

// ---------------------------DELETE------------------------------

func (c *Context) Delete() (bool, error) {

	var chatID any
	var messageID int

	if c.Update.CallbackQuery != nil {
		chatID = c.Update.CallbackQuery.Message.Message.Chat.ID
		messageID = c.Update.CallbackQuery.Message.Message.ID
	} else if c.Update.Message != nil {
		chatID = c.Update.Message.Chat.ID
		messageID = c.Update.Message.ID
	}

	p := &bot.DeleteMessageParams{
		ChatID:    chatID,
		MessageID: messageID,
	}

	isDelete, err := c.Bot.DeleteMessage(c.Ctx, p)
	if err != nil {
		log.Println("Delete error:", err)
		return false, err
	}
	return isDelete, nil
}

func (c *Context) DeleteMessage(messageID int) (bool, error) {

	var chatID any

	if c.Update.CallbackQuery != nil {
		chatID = c.Update.CallbackQuery.Message.Message.Chat.ID
	} else if c.Update.Message != nil {
		chatID = c.Update.Message.Chat.ID
	}

	p := &bot.DeleteMessageParams{
		ChatID:    chatID,
		MessageID: messageID,
	}

	isDelete, err := c.Bot.DeleteMessage(c.Ctx, p)
	if err != nil {
		log.Println("Delete error:", err)
		return false, err
	}
	return isDelete, nil
}

// ---------------------------FOWARD------------------------------

func (c *Context) Forward(fromChatID any, messageID int) (*models.Message, error) {

	var chatID any

	if c.Update.CallbackQuery != nil {
		chatID = c.Update.CallbackQuery.Message.Message.Chat.ID
	} else if c.Update.Message != nil {
		chatID = c.Update.Message.Chat.ID
	}

	p := &bot.ForwardMessageParams{
		ChatID:     chatID,
		FromChatID: fromChatID,
		MessageID:  messageID,
	}

	msg, err := c.Bot.ForwardMessage(c.Ctx, p)
	if err != nil {
		log.Println("Forward error:", err)
		return nil, err
	}
	return msg, nil
}

func (c *Context) Copy(fromChatID any, messageID int, params *MsgParams) (*models.MessageID, error) {

	var chatID any

	if c.Update.CallbackQuery != nil {
		chatID = c.Update.CallbackQuery.Message.Message.Chat.ID
	} else if c.Update.Message != nil {
		chatID = c.Update.Message.Chat.ID
	}

	p := &bot.CopyMessageParams{
		ChatID:     chatID,
		FromChatID: fromChatID,
		MessageID:  messageID,
	}

	if params != nil {
		p.ParseMode = models.ParseMode(params.ParseMode)
		if params.ReplyMarkup != nil {
			p.ReplyMarkup = &models.InlineKeyboardMarkup{
				InlineKeyboard: params.ReplyMarkup,
			}
		}
		if params.Reply {
			p.ReplyParameters = &models.ReplyParameters{
				MessageID: c.Update.Message.ID,
			}
		}
	}

	msg, err := c.Bot.CopyMessage(c.Ctx, p)
	if err != nil {
		log.Println("Copy error:", err)
		return nil, err
	}
	return msg, nil
}
