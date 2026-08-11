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

type VideoBuilder struct {
	bot         *bot.Bot
	userID      int64
	video       models.InputFile
	caption     string
	parseMode   models.ParseMode
	replyMarkup models.ReplyMarkup
	replyTo     int
	duration    int
	width       int
	height      int
	err         error
}

func NewVideoBuilder(b *bot.Bot, userID int64) *VideoBuilder {
	return &VideoBuilder{bot: b, userID: userID}
}

func (b *VideoBuilder) FileID(id string) *VideoBuilder {
	b.video = &models.InputFileString{Data: id}
	return b
}

func (b *VideoBuilder) FileURL(url string) *VideoBuilder {
	b.video = &models.InputFileString{Data: url}
	return b
}

func (b *VideoBuilder) FilePath(path string) *VideoBuilder {
	data, err := os.ReadFile(path)
	if err != nil {
		b.err = err
		return b
	}
	b.video = &models.InputFileUpload{Filename: filepath.Base(path), Data: bytes.NewReader(data)}
	return b
}

func (b *VideoBuilder) Caption(caption string) *VideoBuilder {
	b.caption = caption
	return b
}

func (b *VideoBuilder) ParseMode(mode models.ParseMode) *VideoBuilder {
	b.parseMode = mode
	return b
}

func (b *VideoBuilder) ReplyMarkup(kb models.ReplyMarkup) *VideoBuilder {
	b.replyMarkup = kb
	return b
}

func (b *VideoBuilder) ReplyTo(messageID int) *VideoBuilder {
	b.replyTo = messageID
	return b
}

func (b *VideoBuilder) Duration(second int) *VideoBuilder {
	b.duration = second
	return b
}

func (b *VideoBuilder) Size(width, height int) *VideoBuilder {
	b.width = width
	b.height = height
	return b
}

func (b *VideoBuilder) Do(ctx context.Context) (*models.Message, error) {
	if b.err != nil {
		return nil, b.err
	}
	if b.video == nil {
		return nil, fmt.Errorf("video source is not set: call FileID, URL or FromPath")
	}

	params := &bot.SendVideoParams{
		ChatID:      b.userID,
		Video:       b.video,
		Caption:     b.caption,
		ParseMode:   b.parseMode,
		ReplyMarkup: b.replyMarkup,
		Duration:    b.duration,
		Width:       b.width,
		Height:      b.height,
	}
	if b.replyTo != 0 {
		params.ReplyParameters = &models.ReplyParameters{MessageID: b.replyTo}
	}

	return b.bot.SendVideo(ctx, params)
}
