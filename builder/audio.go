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

type AudioBuilder struct {
	bot         *bot.Bot
	userID      int64
	audio       models.InputFile
	thumbnail   models.InputFile
	caption     string
	parseMode   models.ParseMode
	replyMarkup models.ReplyMarkup
	replyTo     int
	duration    int
	performer   string
	title       string
	err         error
}

func NewAudioBuilder(b *bot.Bot, userID int64) *AudioBuilder {
	return &AudioBuilder{bot: b, userID: userID}
}

func (b *AudioBuilder) FileID(id string) *AudioBuilder {
	b.audio = &models.InputFileString{Data: id}
	return b
}

func (b *AudioBuilder) FileURL(url string) *AudioBuilder {
	b.audio = &models.InputFileString{Data: url}
	return b
}

func (b *AudioBuilder) FilePath(path string) *AudioBuilder {
	data, err := os.ReadFile(path)
	if err != nil {
		b.err = err
		return b
	}
	b.audio = &models.InputFileUpload{Filename: filepath.Base(path), Data: bytes.NewReader(data)}
	return b
}

func (b *AudioBuilder) ThumbnailFileID(id string) *AudioBuilder {
	b.thumbnail = &models.InputFileString{Data: id}
	return b
}

func (b *AudioBuilder) ThumbnailFileURL(url string) *AudioBuilder {
	b.thumbnail = &models.InputFileString{Data: url}
	return b
}

func (b *AudioBuilder) ThumbnailFilePath(path string) *AudioBuilder {
	data, err := os.ReadFile(path)
	if err != nil {
		b.err = err
		return b
	}
	b.thumbnail = &models.InputFileUpload{Filename: filepath.Base(path), Data: bytes.NewReader(data)}
	return b
}

func (b *AudioBuilder) Caption(caption string) *AudioBuilder {
	b.caption = caption
	return b
}

func (b *AudioBuilder) ParseMode(mode models.ParseMode) *AudioBuilder {
	b.parseMode = mode
	return b
}

func (b *AudioBuilder) ReplyMarkup(kb models.ReplyMarkup) *AudioBuilder {
	b.replyMarkup = kb
	return b
}

func (b *AudioBuilder) ReplyTo(messageID int) *AudioBuilder {
	b.replyTo = messageID
	return b
}

func (b *AudioBuilder) Duration(second int) *AudioBuilder {
	b.duration = second
	return b
}

func (b *AudioBuilder) Performer(performer string) *AudioBuilder {
	b.performer = performer
	return b
}

func (b *AudioBuilder) Title(title string) *AudioBuilder {
	b.title = title
	return b
}

func (b *AudioBuilder) Do(ctx context.Context) (*models.Message, error) {
	if b.err != nil {
		return nil, b.err
	}
	if b.audio == nil {
		return nil, fmt.Errorf("audio source is not set: call FileID, URL or FromPath")
	}

	params := &bot.SendAudioParams{
		ChatID:      b.userID,
		Audio:       b.audio,
		Thumbnail:   b.thumbnail,
		Caption:     b.caption,
		ParseMode:   b.parseMode,
		ReplyMarkup: b.replyMarkup,
		Duration:    b.duration,
		Performer:   b.performer,
		Title:       b.title,
	}
	if b.replyTo != 0 {
		params.ReplyParameters = &models.ReplyParameters{MessageID: b.replyTo}
	}

	return b.bot.SendAudio(ctx, params)
}
