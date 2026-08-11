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

// =================== EDIT CAPTION ===========================

type EditCaptionBuilder struct {
	bot         *bot.Bot
	userID      int64
	messageID   int
	caption     string
	parseMode   models.ParseMode
	replyMarkup models.ReplyMarkup
}

func NewEditCaptionBuilder(b *bot.Bot, userID int64, messageID int, caption string) *EditCaptionBuilder {
	return &EditCaptionBuilder{bot: b, userID: userID, messageID: messageID, caption: caption}
}

func (b *EditCaptionBuilder) ParseMode(mode models.ParseMode) *EditCaptionBuilder {
	b.parseMode = mode
	return b
}

func (b *EditCaptionBuilder) ReplyMarkup(kb models.ReplyMarkup) *EditCaptionBuilder {
	b.replyMarkup = kb
	return b
}

func (b *EditCaptionBuilder) Do(ctx context.Context) (*models.Message, error) {
	return b.bot.EditMessageCaption(ctx, &bot.EditMessageCaptionParams{
		ChatID:      b.userID,
		MessageID:   b.messageID,
		Caption:     b.caption,
		ParseMode:   b.parseMode,
		ReplyMarkup: b.replyMarkup,
	})
}

// =================== EDIT MEDIA ==========================

type EditMediaBuilder struct {
	bot         *bot.Bot
	userID      int64
	messageID   int
	media       models.InputMedia
	replyMarkup models.ReplyMarkup
	err         error
}

func NewEditMediaBuilder(b *bot.Bot, userID int64, messageID int) *EditMediaBuilder {
	return &EditMediaBuilder{bot: b, userID: userID, messageID: messageID}
}

// ------------------- PHOTO ---------------------

func (b *EditMediaBuilder) PhotoFileID(id string, caption string) *EditMediaBuilder {
	b.media = &models.InputMediaPhoto{Media: id, Caption: caption}
	return b
}

func (b *EditMediaBuilder) PhotoURL(url string, caption string) *EditMediaBuilder {
	b.media = &models.InputMediaPhoto{Media: url, Caption: caption}
	return b
}

func (b *EditMediaBuilder) PhotoFromPath(path string, caption string) *EditMediaBuilder {
	data, err := os.ReadFile(path)
	if err != nil {
		b.err = err
		return b
	}
	filename := filepath.Base(path)
	b.media = &models.InputMediaPhoto{
		Media:           "attach://" + filename,
		Caption:         caption,
		MediaAttachment: bytes.NewReader(data),
	}
	return b
}

// ------------------- VIDEO ---------------------

func (b *EditMediaBuilder) VideoFileID(id string, caption string) *EditMediaBuilder {
	b.media = &models.InputMediaVideo{Media: id, Caption: caption}
	return b
}

func (b *EditMediaBuilder) VideoURL(url string, caption string) *EditMediaBuilder {
	b.media = &models.InputMediaVideo{Media: url, Caption: caption}
	return b
}

func (b *EditMediaBuilder) VideoFromPath(path string, caption string) *EditMediaBuilder {
	data, err := os.ReadFile(path)
	if err != nil {
		b.err = err
		return b
	}
	filename := filepath.Base(path)
	b.media = &models.InputMediaVideo{
		Media:           "attach://" + filename,
		Caption:         caption,
		MediaAttachment: bytes.NewReader(data),
	}
	return b
}

// ------------------- DOCUMENT ---------------------

func (b *EditMediaBuilder) DocumentFileID(id string, caption string) *EditMediaBuilder {
	b.media = &models.InputMediaDocument{Media: id, Caption: caption}
	return b
}

func (b *EditMediaBuilder) DocumentURL(url string, caption string) *EditMediaBuilder {
	b.media = &models.InputMediaDocument{Media: url, Caption: caption}
	return b
}

func (b *EditMediaBuilder) DocumentFromPath(path string, caption string) *EditMediaBuilder {
	data, err := os.ReadFile(path)
	if err != nil {
		b.err = err
		return b
	}
	filename := filepath.Base(path)
	b.media = &models.InputMediaDocument{
		Media:           "attach://" + filename,
		Caption:         caption,
		MediaAttachment: bytes.NewReader(data),
	}
	return b
}

// ------------------- AUDIO ---------------------

func (b *EditMediaBuilder) AudioFileID(id string, caption string) *EditMediaBuilder {
	b.media = &models.InputMediaAudio{Media: id, Caption: caption}
	return b
}

func (b *EditMediaBuilder) AudioURL(url string, caption string) *EditMediaBuilder {
	b.media = &models.InputMediaAudio{Media: url, Caption: caption}
	return b
}

func (b *EditMediaBuilder) AudioFromPath(path string, caption string) *EditMediaBuilder {
	data, err := os.ReadFile(path)
	if err != nil {
		b.err = err
		return b
	}
	filename := filepath.Base(path)
	b.media = &models.InputMediaAudio{
		Media:           "attach://" + filename,
		Caption:         caption,
		MediaAttachment: bytes.NewReader(data),
	}
	return b
}

func (b *EditMediaBuilder) ReplyMarkup(kb models.ReplyMarkup) *EditMediaBuilder {
	b.replyMarkup = kb
	return b
}

func (b *EditMediaBuilder) Do(ctx context.Context) (*models.Message, error) {
	if b.err != nil {
		return nil, b.err
	}
	if b.media == nil {
		return nil, fmt.Errorf("media is not set: call Photo, Video or similar")
	}
	return b.bot.EditMessageMedia(ctx, &bot.EditMessageMediaParams{
		ChatID:      b.userID,
		MessageID:   b.messageID,
		Media:       b.media,
		ReplyMarkup: b.replyMarkup,
	})
}
