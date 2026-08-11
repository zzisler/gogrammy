package builder

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

type PhotoBuilder struct {
	bot         *bot.Bot
	userID      int64
	photo       models.InputFile
	caption     string
	parseMode   models.ParseMode
	replyMarkup models.ReplyMarkup
	replyTo     int
	err         error
}

func NewPhotoBuilder(b *bot.Bot, userID int64) *PhotoBuilder {
	return &PhotoBuilder{bot: b, userID: userID}
}

func (b *PhotoBuilder) FileID(id string) *PhotoBuilder {
	b.photo = &models.InputFileString{Data: id}
	return b
}

func (b *PhotoBuilder) FileURL(url string) *PhotoBuilder {
	b.photo = &models.InputFileString{Data: url}
	return b
}

func (b *PhotoBuilder) FilePath(path string) *PhotoBuilder {
	data, err := os.ReadFile(path)
	if err != nil {
		b.err = err
		return b
	}
	b.photo = &models.InputFileUpload{Filename: filepath.Base(path), Data: bytes.NewReader(data)}
	return b
}

func (b *PhotoBuilder) Caption(caption string) *PhotoBuilder {
	b.caption = caption
	return b
}

func (b *PhotoBuilder) ParseMode(mode models.ParseMode) *PhotoBuilder {
	b.parseMode = mode
	return b
}

func (b *PhotoBuilder) ReplyMarkup(kb models.ReplyMarkup) *PhotoBuilder {
	b.replyMarkup = kb
	return b
}

func (b *PhotoBuilder) ReplyTo(messageID int) *PhotoBuilder {
	b.replyTo = messageID
	return b
}

func (b *PhotoBuilder) Do(ctx context.Context) (*models.Message, error) {
	if b.err != nil {
		return nil, b.err
	}
	if b.photo == nil {
		return nil, fmt.Errorf("photo source is not set: call FileID, URL or FromPath")
	}

	params := &bot.SendPhotoParams{
		ChatID:      b.userID,
		Photo:       b.photo,
		Caption:     b.caption,
		ParseMode:   b.parseMode,
		ReplyMarkup: b.replyMarkup,
	}
	if b.replyTo != 0 {
		params.ReplyParameters = &models.ReplyParameters{MessageID: b.replyTo}
	}

	return b.bot.SendPhoto(ctx, params)
}
