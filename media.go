package gogrammy

import (
	"bytes"
	"log"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

type PhotoParams struct {
	Caption     string
	ParseMode   string
	Spoiler     bool
	ReplyMarkup [][]models.InlineKeyboardButton
}

func (c *Context) SendPhoto(chatID any, url string, params *PhotoParams) (*models.Message, error) {

	p := &bot.SendPhotoParams{
		ChatID: chatID,
		Photo:  &models.InputFileString{Data: url},
	}

	if params != nil {
		p.Caption = params.Caption
		p.HasSpoiler = params.Spoiler
		p.ParseMode = models.ParseMode(params.ParseMode)
		if params.ReplyMarkup != nil {
			p.ReplyMarkup = &models.InlineKeyboardMarkup{
				InlineKeyboard: params.ReplyMarkup,
			}
		}
	}

	msg, err := c.Bot.SendPhoto(c.Ctx, p)
	if err != nil {
		log.Println("SendPhoto error:", err)
		return nil, err
	}
	return msg, nil
}

func (c *Context) SendAudio(chatID any, filename, title, artist string, data, coverData []byte) (*models.Message, error) {

	p := &bot.SendAudioParams{
		ChatID:    chatID,
		Title:     title,
		Performer: artist,
		Thumbnail: &models.InputFileUpload{
			Filename: "cover.jpg",
			Data:     bytes.NewReader(coverData),
		},
		Audio: &models.InputFileUpload{
			Filename: filename,
			Data:     bytes.NewReader(data),
		},
	}

	msg, err := c.Bot.SendAudio(c.Ctx, p)
	if err != nil {
		log.Println("SendAudio error:", err)
		return nil, err
	}
	return msg, nil
}
