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

type DocumentBuilder struct {
	bot         *bot.Bot
	userID      int64
	document    models.InputFile
	caption     string
	parseMode   models.ParseMode
	replyMarkup models.ReplyMarkup
	replyTo     int
	err         error
}

func NewDocumentBuilder(b *bot.Bot, userID int64) *DocumentBuilder {
	return &DocumentBuilder{bot: b, userID: userID}
}

func (b *DocumentBuilder) FileID(id string) *DocumentBuilder {
	b.document = &models.InputFileString{Data: id}
	return b
}

func (b *DocumentBuilder) FileURL(url string) *DocumentBuilder {
	b.document = &models.InputFileString{Data: url}
	return b
}

func (b *DocumentBuilder) FilePath(path string) *DocumentBuilder {
	data, err := os.ReadFile(path)
	if err != nil {
		b.err = err
		return b
	}
	b.document = &models.InputFileUpload{Filename: filepath.Base(path), Data: bytes.NewReader(data)}
	return b
}

func (b *DocumentBuilder) Caption(caption string) *DocumentBuilder {
	b.caption = caption
	return b
}

func (b *DocumentBuilder) ParseMode(mode models.ParseMode) *DocumentBuilder {
	b.parseMode = mode
	return b
}

func (b *DocumentBuilder) ReplyMarkup(kb models.ReplyMarkup) *DocumentBuilder {
	b.replyMarkup = kb
	return b
}

func (b *DocumentBuilder) ReplyTo(messageID int) *DocumentBuilder {
	b.replyTo = messageID
	return b
}

func (b *DocumentBuilder) Do(ctx context.Context) (*models.Message, error) {
	if b.err != nil {
		return nil, b.err
	}
	if b.document == nil {
		return nil, fmt.Errorf("document source is not set: call FileID, URL or FromPath")
	}

	params := &bot.SendDocumentParams{
		ChatID:      b.userID,
		Document:    b.document,
		Caption:     b.caption,
		ParseMode:   b.parseMode,
		ReplyMarkup: b.replyMarkup,
	}
	if b.replyTo != 0 {
		params.ReplyParameters = &models.ReplyParameters{MessageID: b.replyTo}
	}

	return b.bot.SendDocument(ctx, params)
}
